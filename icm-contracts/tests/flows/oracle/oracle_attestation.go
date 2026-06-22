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

	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/signature-aggregator/api"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

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
//  2. Submit to /oracle/aggregate-signatures with a dummy justification
//  3. Confirm a valid signed warp message is returned
//
// Requires:
//   - Validators from the boraplusplus/sidecar-verifier branch with
//     oracle.endpoint wired in their chain config
//   - An oracle sidecar (or mock) reachable at that endpoint
func OracleAttestation(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
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

	// Build an ABI-encoded OracleMessage payload.
	oraclePayload, err := oracleMsgABI.Pack(
		"solana",                   // sourceType
		"4oracle1testaddr",         // sourceAddress (dummy Solana address)
		common.Address{1, 2, 3},   // destContract
		uint64(100),               // sourceBlockHeight
		uint64(1),                 // nonce
		[]byte("e2e-test-payload"), // payload
	)
	Expect(err).Should(BeNil())

	networkID := avalancheNetwork.GetNetworkID()
	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		l1Info.BlockchainID,
		oraclePayload,
	)
	Expect(err).Should(BeNil())

	justification := []byte("dummy-solana-tx-signature")

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
