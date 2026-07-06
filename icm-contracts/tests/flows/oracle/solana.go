// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package oracle

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/mr-tron/base58"
	. "github.com/onsi/gomega"
)

// memoProgram is the Solana Memo Program v2, present on both mainnet and devnet.
const memoProgram = "MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr"

// solanaTxData holds the fields extracted from a Solana transaction that
// are needed to construct a matching OracleMessage.
type solanaTxData struct {
	txSigBytes []byte // raw 64-byte Ed25519 signature (justification for the sidecar)
	slot       uint64
	programID  string
	instrData  []byte
}

// solanaInstr mirrors the instruction fields returned by the Solana JSON-RPC
// getTransaction response.
type solanaInstr struct {
	ProgramIDIndex int    `json:"programIdIndex"`
	Data           string `json:"data"`
}

// fetchSolanaMemoTx discovers a recent Memo Program transaction from the given
// Solana RPC endpoint and extracts the fields needed for an OracleMessage.
func fetchSolanaMemoTx(ctx context.Context, rpcURL string) solanaTxData {
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

	// Step 1: find a recent Memo Program transaction.
	sigsRaw := post(map[string]any{
		"jsonrpc": "2.0", "id": 1,
		"method": "getSignaturesForAddress",
		"params": []any{memoProgram, map[string]any{"limit": 1}},
	})
	var sigsResp struct {
		Result []struct {
			Signature string `json:"signature"`
		} `json:"result"`
	}
	Expect(json.Unmarshal(sigsRaw, &sigsResp)).Should(BeNil())
	Expect(sigsResp.Result).ShouldNot(BeEmpty(), "no recent Memo Program transactions at SOLANA_RPC_URL")
	txSig := sigsResp.Result[0].Signature

	// Step 2: fetch the full transaction.
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
	Expect(txResp.Result).ShouldNot(BeNil(), "transaction not found for sig %s", txSig)

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

	// Collect all instructions: top-level first, then inner.
	// getSignaturesForAddress returns any tx that mentions a program (including CPI),
	// so the Memo call may appear only in inner instructions.
	allInstrs := make([]solanaInstr, 0, len(txResp.Result.Transaction.Message.Instructions))
	allInstrs = append(allInstrs, txResp.Result.Transaction.Message.Instructions...)
	for _, inner := range txResp.Result.Meta.InnerInstructions {
		allInstrs = append(allInstrs, inner.Instructions...)
	}

	var instrData []byte
	for _, instr := range allInstrs {
		if instr.ProgramIDIndex < 0 || instr.ProgramIDIndex >= len(keys) {
			continue
		}
		if keys[instr.ProgramIDIndex] != memoProgram {
			continue
		}
		data, err := base58.Decode(instr.Data)
		Expect(err).Should(BeNil())
		instrData = data
		break
	}
	Expect(instrData).ShouldNot(BeNil(), "could not find Memo instruction in transaction %s", txSig)

	sigBytes, err := base58.Decode(txSig)
	Expect(err).Should(BeNil())

	return solanaTxData{
		txSigBytes: sigBytes,
		slot:       txResp.Result.Slot,
		programID:  memoProgram,
		instrData:  instrData,
	}
}
