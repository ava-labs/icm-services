// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mr-tron/base58"
)

// memoProgram is the SPL Memo program ID. The Anchor escrow program hex-encodes
// the ABI payload into a Memo CPI so it's valid UTF-8; the relay hex-decodes it.
const memoProgram = "MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr"

// SolanaTx is the fully resolved view of a Solana transaction the observer
// needs to construct an OracleMessage. It is a strict subset of what the JSON-RPC
// getTransaction endpoint returns.
type SolanaTx struct {
	Slot      uint64
	Program   string
	InstrData []byte
	SigBytes  []byte
}

type solanaInstr struct {
	ProgramIDIndex int    `json:"programIdIndex"`
	Data           string `json:"data"`
}

// fetchTx pulls the full transaction body for the given signature and extracts
// the first instruction (top-level or CPI-inner) whose programIdIndex resolves
// to program. Returns (nil, false) if the transaction does not exist yet or
// does not invoke the given program.
//
// This is the same discovery logic used by the E2E flow test in
// icm-contracts/tests/flows/oracle/solana.go, factored for production use.
func fetchTx(ctx context.Context, rpcURL, txSig, program string) (*SolanaTx, bool, error) {
	body, err := json.Marshal(map[string]any{
		"jsonrpc": "2.0", "id": 1,
		"method": "getTransaction",
		"params": []any{txSig, map[string]any{
			"encoding":                       "json",
			"commitment":                     "finalized",
			"maxSupportedTransactionVersion": 0,
		}},
	})
	if err != nil {
		return nil, false, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, rpcURL, bytes.NewReader(body))
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, false, fmt.Errorf("getTransaction request: %w", err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}

	var out struct {
		Result *struct {
			Slot        uint64 `json:"slot"`
			Transaction struct {
				Message struct {
					AccountKeys  []string      `json:"accountKeys"`
					Instructions []solanaInstr `json:"instructions"`
				} `json:"message"`
			} `json:"transaction"`
			Meta struct {
				LoadedAddresses struct {
					Writable []string `json:"writable"`
					Readonly []string `json:"readonly"`
				} `json:"loadedAddresses"`
				InnerInstructions []struct {
					Index        int           `json:"index"`
					Instructions []solanaInstr `json:"instructions"`
				} `json:"innerInstructions"`
			} `json:"meta"`
		} `json:"result"`
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return nil, false, fmt.Errorf("parse getTransaction: %w", err)
	}
	if out.Result == nil {
		return nil, false, nil
	}

	// v0 transactions can resolve program addresses via lookup tables, so include
	// loadedAddresses in the key space that programIdIndex references.
	keys := append(
		out.Result.Transaction.Message.AccountKeys,
		append(
			out.Result.Meta.LoadedAddresses.Writable,
			out.Result.Meta.LoadedAddresses.Readonly...,
		)...,
	)

	allInstrs := make([]solanaInstr, 0, len(out.Result.Transaction.Message.Instructions))
	allInstrs = append(allInstrs, out.Result.Transaction.Message.Instructions...)
	for _, inner := range out.Result.Meta.InnerInstructions {
		allInstrs = append(allInstrs, inner.Instructions...)
	}

	// Confirm the subscription program is present in the transaction.
	found := false
	for _, instr := range allInstrs {
		if instr.ProgramIDIndex >= 0 && instr.ProgramIDIndex < len(keys) && keys[instr.ProgramIDIndex] == program {
			found = true
			break
		}
	}
	if !found {
		return nil, false, nil
	}

	// The ABI payload always lives in the Memo instruction. When subscribing to
	// the escrow program the Memo CPI is an inner instruction; when subscribing
	// directly to Memo it's the top-level instruction. Either way we want the
	// Memo instruction data.
	dataProgram := memoProgram
	if program == memoProgram {
		// Already subscribed directly to Memo — keep original behaviour.
		dataProgram = memoProgram
	}

	for _, instr := range allInstrs {
		if instr.ProgramIDIndex < 0 || instr.ProgramIDIndex >= len(keys) {
			continue
		}
		if keys[instr.ProgramIDIndex] != dataProgram {
			continue
		}
		data, err := base58.Decode(instr.Data)
		if err != nil {
			return nil, false, fmt.Errorf("base58 decode instruction data: %w", err)
		}
		sigBytes, err := base58.Decode(txSig)
		if err != nil {
			return nil, false, fmt.Errorf("base58 decode tx signature: %w", err)
		}
		return &SolanaTx{
			Slot:      out.Result.Slot,
			Program:   dataProgram, // always the Memo program — sidecar verifies Memo instruction
			InstrData: data,
			SigBytes:  sigBytes,
		}, true, nil
	}
	return nil, false, nil
}
