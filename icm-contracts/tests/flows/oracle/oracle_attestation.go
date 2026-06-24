// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package oracle contains E2E test flows for oracle attestation.
// These tests require validators built from the boraplusplus/sidecar-verifier
// branch of avalanchego, with the oracle.endpoint chain config pointing to a
// running sidecar process.
package oracle

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/mr-tron/base58"

	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/signature-aggregator/api"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// memoProgram is the Solana Memo Program v2, present on both mainnet and devnet.
const memoProgram = "MemoSq4gqABAXKb96qnH8TysNcWxMyWCqXgDLGmfcHr"

// solanaTxData holds the fields extracted from a real Solana transaction that
// are needed to construct a matching OracleMessage.
type solanaTxData struct {
	txSigBytes []byte // raw 64-byte Ed25519 signature (justification for the sidecar)
	slot       uint64
	programID  string
	instrData  []byte
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
		Result []struct{ Signature string `json:"signature"` } `json:"result"`
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
					AccountKeys  []string `json:"accountKeys"`
					Instructions []struct {
						ProgramIDIndex int    `json:"programIdIndex"`
						Data           string `json:"data"`
					} `json:"instructions"`
				} `json:"message"`
			} `json:"transaction"`
		} `json:"result"`
	}
	Expect(json.Unmarshal(txRaw, &txResp)).Should(BeNil())
	Expect(txResp.Result).ShouldNot(BeNil(), "transaction not found for sig %s", txSig)

	keys := txResp.Result.Transaction.Message.AccountKeys
	var instrData []byte
	for _, instr := range txResp.Result.Transaction.Message.Instructions {
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

// oracleMsgABI encodes the oracle message payload that OracleVerifier expects.
// Layout mirrors OracleMessage in network/p2p/oracle/message.go on the
// boraplusplus/sidecar-verifier branch.
var oracleMsgABI abi.Arguments

func init() {
	stringT, _ := abi.NewType("string", "", nil)
	addrT, _ := abi.NewType("address", "", nil)
	uint64T, _ := abi.NewType("uint64", "", nil)
	bytesT, _ := abi.NewType("bytes", "", nil)

	oracleMsgABI = abi.Arguments{
		{Type: stringT, Name: "sourceType"},
		{Type: stringT, Name: "sourceAddress"},
		{Type: addrT, Name: "destContract"},
		{Type: uint64T, Name: "sourceBlockHeight"},
		{Type: uint64T, Name: "nonce"},
		{Type: bytesT, Name: "payload"},
	}
}

// OracleAttestation tests the full oracle attestation path:
//  1. Construct an OracleMessage (ABI-encoded warp payload)
//  2. Submit to /oracle/aggregate-signatures
//  3. Confirm a valid signed warp message is returned
//
// When solanaRPCURL is empty the flow uses the mock sidecar with dummy data.
// When solanaRPCURL is set it fetches a real Memo Program transaction from that
// endpoint and uses its slot/program/payload as the oracle payload, exercising
// the real solanarpc sidecar end-to-end.
func OracleAttestation(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	solanaRPCURL string,
) {
	l1Infos := avalancheNetwork.GetL1Infos()
	Expect(len(l1Infos)).Should(BeNumerically(">=", 1), "oracle suite needs at least one L1")
	l1Info := l1Infos[0]

	sigAggConfig := utils.CreateDefaultSignatureAggregatorConfig(
		log,
		[]testinfo.L1TestInfo{l1Info},
	)
	sigAggConfigPath := utils.WriteSignatureAggregatorConfig(
		log,
		sigAggConfig,
		"sig-agg-oracle-config.json",
	)
	log.Info("Starting the signature aggregator for oracle test",
		zap.String("configPath", sigAggConfigPath),
	)
	sigAggCancel, readyChan := utils.RunSignatureAggregatorExecutable(
		ctx,
		log,
		sigAggConfigPath,
		sigAggConfig,
	)
	defer sigAggCancel()

	startupCtx, startupCancel := context.WithTimeout(ctx, 20*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	// Choose oracle message source: real Solana tx or dummy data.
	var (
		sourceAddress string
		blockHeight   uint64
		msgPayload    []byte
		justification []byte
	)
	if solanaRPCURL != "" {
		log.Info("Fetching real Memo Program transaction from Solana", zap.String("rpc", solanaRPCURL))
		txData := fetchSolanaMemoTx(ctx, solanaRPCURL)
		sourceAddress = txData.programID
		blockHeight = txData.slot
		msgPayload = txData.instrData
		justification = txData.txSigBytes
		log.Info("Using real Solana transaction",
			zap.String("program", sourceAddress),
			zap.Uint64("slot", blockHeight),
			zap.Int("payloadBytes", len(msgPayload)),
		)
	} else {
		sourceAddress = "4oracle1testaddr"
		blockHeight = 100
		msgPayload = []byte("e2e-test-payload")
		justification = []byte("dummy-solana-tx-signature")
	}

	// Build an ABI-encoded OracleMessage payload.
	oraclePayload, err := oracleMsgABI.Pack(
		"solana",
		sourceAddress,
		common.Address{1, 2, 3},
		blockHeight,
		uint64(1), // nonce
		msgPayload,
	)
	Expect(err).Should(BeNil())

	networkID := avalancheNetwork.GetNetworkID()
	ac, err := payload.NewAddressedCall(nil, oraclePayload)
	Expect(err).Should(BeNil())

	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		l1Info.BlockchainID,
		ac.Bytes(),
	)
	Expect(err).Should(BeNil())

	reqBody := api.AggregateSignatureRequest{
		Message:         "0x" + hex.EncodeToString(unsignedMsg.Bytes()),
		Justification:   hex.EncodeToString(justification),
		SigningSubnetID: l1Info.SubnetID.String(),
	}

	client := http.Client{Timeout: 30 * time.Second}
	requestURL := fmt.Sprintf("http://localhost:%d%s", sigAggConfig.APIPort, api.OracleAPIPath)

	log.Info("Sending oracle attestation request",
		zap.String("url", requestURL),
		zap.Stringer("blockchainID", l1Info.BlockchainID),
		zap.Stringer("signingSubnetID", l1Info.SubnetID),
	)

	b, err := json.Marshal(reqBody)
	Expect(err).Should(BeNil())

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL, bytes.NewReader(b))
	Expect(err).Should(BeNil())
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	Expect(err).Should(BeNil())
	Expect(res.Status).Should(Equal("200 OK"))
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	Expect(err).Should(BeNil())

	var response api.AggregateSignatureResponse
	err = json.Unmarshal(body, &response)
	Expect(err).Should(BeNil())
	Expect(response.SignedMessage).ShouldNot(BeEmpty())

	decodedMsg, err := hex.DecodeString(response.SignedMessage)
	Expect(err).Should(BeNil())

	signedMsg, err := avalancheWarp.ParseMessage(decodedMsg)
	Expect(err).Should(BeNil())
	Expect(signedMsg.ID()).Should(Equal(unsignedMsg.ID()),
		"signed message ID must match the submitted unsigned message")

	log.Info("Oracle attestation succeeded",
		zap.Stringer("messageID", signedMsg.ID()),
	)
}
