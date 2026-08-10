// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/mr-tron/base58"
	. "github.com/onsi/gomega"
)

// MemoProgram is the Solana Memo Program v2, present on both mainnet and devnet.
const MemoProgram = "MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr"

// SolanaTxData holds the fields extracted from a Solana transaction that
// are needed to construct a matching OracleMessage.
type SolanaTxData struct {
	TxSigBytes []byte // raw 64-byte Ed25519 signature (justification for the sidecar)
	Slot       uint64
	ProgramID  string
	InstrData  []byte
}

// solanaInstr mirrors the instruction fields returned by the Solana JSON-RPC
// getTransaction response.
type solanaInstr struct {
	ProgramIDIndex int    `json:"programIdIndex"`
	Data           string `json:"data"`
}

// FetchSolanaMemoTx discovers a recent Memo Program transaction from the given
// Solana RPC endpoint and extracts the fields needed for an OracleMessage.
func FetchSolanaMemoTx(ctx context.Context, rpcURL string) SolanaTxData {
	post := func(body any) []byte {
		b, err := json.Marshal(body)
		Expect(err).Should(BeNil())
		req, err := http.NewRequestWithContext(ctx, http.MethodPost, rpcURL, bytes.NewReader(b))
		Expect(err).Should(BeNil())
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		Expect(err).Should(BeNil())
		defer resp.Body.Close()
		out, err := io.ReadAll(resp.Body)
		Expect(err).Should(BeNil())
		return out
	}

	// Step 1: fetch a batch of recent transactions that mention the Memo program.
	// getSignaturesForAddress returns any tx where the address appears in any
	// role (invoked, plain account, LUT-resolved, etc.), not just ones that
	// actually invoke it as a program. We iterate through the batch until we
	// find one that contains a real Memo instruction.
	const candidateLimit = 20
	sigsRaw := post(map[string]any{
		"jsonrpc": "2.0", "id": 1,
		"method": "getSignaturesForAddress",
		"params": []any{MemoProgram, map[string]any{"limit": candidateLimit}},
	})
	var sigsResp struct {
		Result []struct {
			Signature string `json:"signature"`
		} `json:"result"`
	}
	Expect(json.Unmarshal(sigsRaw, &sigsResp)).Should(BeNil())
	Expect(sigsResp.Result).ShouldNot(BeEmpty(), "no recent Memo Program transactions at SOLANA_RPC_URL")

	// Step 2: for each candidate, fetch the tx and scan for a Memo instruction.
	// Return the first match.
	for _, cand := range sigsResp.Result {
		txSig := cand.Signature
		if data, slot, ok := tryExtractMemoInstruction(post, txSig); ok {
			sigBytes, err := base58.Decode(txSig)
			Expect(err).Should(BeNil())
			return SolanaTxData{
				TxSigBytes: sigBytes,
				Slot:       slot,
				ProgramID:  MemoProgram,
				InstrData:  data,
			}
		}
	}
	Expect(false).To(BeTrue(), "no Memo instruction found in the %d most recent Memo-tagged transactions", candidateLimit)
	return SolanaTxData{}
}

// tryExtractMemoInstruction fetches the transaction with the given signature and
// scans its top-level and inner instructions for one that invokes the Memo
// program with a non-empty payload. Returns (data, slot, true) on the first
// match, or (nil, 0, false) if the transaction contains no such Memo
// invocation. The transaction must exist; this helper does not retry lookup
// failures.
func tryExtractMemoInstruction(post func(any) []byte, txSig string) ([]byte, uint64, bool) {
	txRaw := post(map[string]any{
		"jsonrpc": "2.0", "id": 1,
		"method": "getTransaction",
		"params": []any{txSig, map[string]any{
			"encoding":                       "json",
			"maxSupportedTransactionVersion": 0,
		}},
	})
	var txResp struct {
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
	Expect(json.Unmarshal(txRaw, &txResp)).Should(BeNil())
	if txResp.Result == nil {
		return nil, 0, false
	}

	// For versioned (v0) transactions, programIdIndex refers to the combined account
	// list: static keys + loaded writable + loaded readonly. Programs resolved via
	// address lookup tables appear in meta.loadedAddresses, not in accountKeys, so
	// we must include both to correctly map an instruction's programIdIndex.
	keys := append(
		txResp.Result.Transaction.Message.AccountKeys,
		append(
			txResp.Result.Meta.LoadedAddresses.Writable,
			txResp.Result.Meta.LoadedAddresses.Readonly...,
		)...,
	)

	allInstrs := make([]solanaInstr, 0, len(txResp.Result.Transaction.Message.Instructions))
	allInstrs = append(allInstrs, txResp.Result.Transaction.Message.Instructions...)
	for _, inner := range txResp.Result.Meta.InnerInstructions {
		allInstrs = append(allInstrs, inner.Instructions...)
	}

	for _, instr := range allInstrs {
		if instr.ProgramIDIndex < 0 || instr.ProgramIDIndex >= len(keys) {
			continue
		}
		if keys[instr.ProgramIDIndex] != MemoProgram {
			continue
		}
		data, err := base58.Decode(instr.Data)
		Expect(err).Should(BeNil())
		// Skip memos with empty data so a match always carries a usable payload
		// and the (data, _, true) return is guaranteed non-empty.
		if len(data) == 0 {
			continue
		}
		return data, txResp.Result.Slot, true
	}
	return nil, 0, false
}
