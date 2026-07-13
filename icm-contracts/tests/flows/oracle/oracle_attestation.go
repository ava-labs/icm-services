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
	"math/big"
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	mockoraclereceiver "github.com/ava-labs/icm-services/abi-bindings/go/mocks/MockOracleReceiver"
	oracleadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/OracleAdapter"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/signature-aggregator/api"
	icmutils "github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/onsi/ginkgo/v2"
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
	uint256T, _ := abi.NewType("uint256", "", nil)
	bytesT, _ := abi.NewType("bytes", "", nil)

	oracleMsgABI = abi.Arguments{
		{Type: stringT, Name: "sourceType"},
		{Type: stringT, Name: "sourceAddress"},
		{Type: addrT, Name: "destContract"},
		{Type: uint256T, Name: "sourceBlockHeight"},
		{Type: uint256T, Name: "nonce"},
		{Type: bytesT, Name: "payload"},
	}
}

// OracleAttestation tests the full oracle attestation path:
//  1. Deploy OracleAdapter and TeleporterMessengerV2 on the L1
//  2. Deploy MockOracleReceiver pointing to TeleporterMessengerV2
//  3. Construct an OracleMessage and request BLS aggregate signature
//  4. Deliver the signed message via TeleporterMessengerV2.receiveCrossChainMessage
//  5. Assert MockOracleReceiver received the expected payload
//
// When solanaRPCURL is empty the flow uses the mock sidecar with dummy data.
// When solanaRPCURL is set it fetches a real Memo Program transaction from that
// endpoint and uses its slot/program/payload as the oracle payload, exercising
// the real solanarpc sidecar end-to-end.
func OracleAttestation(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	l1Info testinfo.L1TestInfo,
	solanaRPCURL string,
) {
	ginkgo.By("Step 1: Deploy OracleAdapter")
	_, fundedKey := avalancheNetwork.GetFundedAccountInfo()
	deployOpts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())

	adapterAddress, adapterDeployTx, adapterContract, err := oracleadapter.DeployOracleAdapter(
		deployOpts, l1Info.EthClient, deployOpts.From,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, adapterDeployTx.Hash())
	log.Info("Deployed OracleAdapter", zap.Stringer("address", adapterAddress))

	ginkgo.By("Step 1b: Deploy TeleporterMessengerV2 with OracleAdapter as verifier")
	teleporterAddress := utils.DeployTeleporterV2(ctx, &l1Info, adapterAddress, fundedKey)
	teleporterContract, err := teleportermessengerv2.NewTeleporterMessengerV2(
		teleporterAddress, l1Info.EthClient,
	)
	Expect(err).Should(BeNil())
	log.Info("Deployed TeleporterMessengerV2", zap.Stringer("address", teleporterAddress))

	ginkgo.By("Step 1c: Deploy MockOracleReceiver pointing to TeleporterMessengerV2")
	mockAddress, mockDeployTx, mockContract, err := mockoraclereceiver.DeployMockOracleReceiver(
		deployOpts, l1Info.EthClient, teleporterAddress,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, mockDeployTx.Hash())
	log.Info("Deployed MockOracleReceiver", zap.Stringer("address", mockAddress))

	ginkgo.By("Step 2: Start signature aggregator")
	sigAggConfig := utils.CreateDefaultSignatureAggregatorConfig(
		log,
		[]testinfo.L1TestInfo{l1Info},
	)
	sigAggConfigPath := utils.WriteSignatureAggregatorConfig(
		log,
		sigAggConfig,
		"sig-agg-oracle-config.json",
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
		blockHeight   *big.Int
		msgPayload    []byte
		justification []byte
	)
	if solanaRPCURL != "" {
		ginkgo.By("Step 3: Fetch real Memo Program transaction from Solana devnet")
		txData := fetchSolanaMemoTx(ctx, solanaRPCURL)
		sourceAddress = txData.programID
		blockHeight = new(big.Int).SetUint64(txData.slot)
		msgPayload = txData.instrData
		justification = txData.txSigBytes
		log.Info("Using real Solana transaction",
			zap.String("program", sourceAddress),
			zap.Stringer("slot", blockHeight),
			zap.Int("payloadBytes", len(msgPayload)),
		)
	} else {
		ginkgo.By("Step 3: Using mock oracle data (no SOLANA_RPC_URL set)")
		sourceAddress = "4oracle1testaddr"
		blockHeight = big.NewInt(100)
		msgPayload = []byte("e2e-test-payload")
		justification = []byte("dummy-solana-tx-signature")
	}

	ginkgo.By("Step 4: Allowlist source on OracleAdapter")
	allowTx, err := adapterContract.SetAllowedSource(deployOpts, "solana", sourceAddress, true)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, allowTx.Hash())

	ginkgo.By("Step 5: Request BLS aggregate signature from validators")
	oraclePayload, err := oracleMsgABI.Pack(
		"solana",
		sourceAddress,
		mockAddress,
		blockHeight,
		big.NewInt(1), // nonce
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

	log.Info("BLS aggregation succeeded", zap.Stringer("messageID", signedMsg.ID()))

	thisChainID := [32]byte(l1Info.BlockchainID)
	fundedAddress := utils.PrivateKeyToAddress(fundedKey)

	buildICMMessage := func(msg oracleadapter.OracleMessage) (teleportermessengerv2.TeleporterICMMessage, error) {
		return oracleadapter.BuildOracleICMMessage(
			0,
			msg,
			teleporterAddress,
			thisChainID,
			networkID,
			new(big.Int).SetUint64(500_000),
		)
	}

	sendExpectRevert := func(msg oracleadapter.OracleMessage) {
		icmMsg, buildErr := buildICMMessage(msg)
		Expect(buildErr).Should(BeNil())
		data, packErr := teleportermessengerv2.PackReceiveCrossChainMessageV2(
			icmMsg.Message,
			l1Info.BlockchainID,
			icmMsg.Attestation,
			common.Address{},
		)
		Expect(packErr).Should(BeNil())
		gasFeeCap, gasTipCap, txNonce := utils.CalculateTxParams(ctx, l1Info.EthClient, fundedAddress)
		tx := types.NewTx(&types.DynamicFeeTx{
			ChainID:    l1Info.EVMChainID,
			Nonce:      txNonce,
			To:         &teleporterAddress,
			Gas:        500_000,
			GasFeeCap:  gasFeeCap,
			GasTipCap:  gasTipCap,
			Value:      common.Big0,
			Data:       data,
			AccessList: icmutils.SignedWarpMessageToAccessList(signedMsg),
		})
		tx = utils.SignTransaction(tx, fundedKey, l1Info.EVMChainID)
		utils.SendTransactionAndWaitForFailure(ctx, l1Info.EthClient, tx)
	}

	ginkgo.By("Sad path 1: delivery from non-allowlisted source is rejected (SourceNotAllowed)")
	removeTx, removeErr := adapterContract.SetAllowedSource(deployOpts, "solana", sourceAddress, false)
	Expect(removeErr).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, removeTx.Hash())
	sendExpectRevert(oracleadapter.OracleMessage{
		SourceType:        "solana",
		SourceAddress:     sourceAddress,
		DestContract:      mockAddress,
		SourceBlockHeight: blockHeight,
		Nonce:             big.NewInt(1),
		Payload:           msgPayload,
	})
	restoreTx, restoreErr := adapterContract.SetAllowedSource(deployOpts, "solana", sourceAddress, true)
	Expect(restoreErr).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, restoreTx.Hash())

	ginkgo.By("Step 6: Deliver the signed oracle message via TeleporterMessengerV2")
	icmMsg, err := buildICMMessage(oracleadapter.OracleMessage{
		SourceType:        "solana",
		SourceAddress:     sourceAddress,
		DestContract:      mockAddress,
		SourceBlockHeight: blockHeight,
		Nonce:             big.NewInt(1),
		Payload:           msgPayload,
	})
	Expect(err).Should(BeNil())

	callData, packErr := teleportermessengerv2.PackReceiveCrossChainMessageV2(
		icmMsg.Message,
		l1Info.BlockchainID,
		icmMsg.Attestation,
		common.Address{},
	)
	Expect(packErr).Should(BeNil())

	gasFeeCap, gasTipCap, txNonce := utils.CalculateTxParams(ctx, l1Info.EthClient, fundedAddress)
	deliveryTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    l1Info.EVMChainID,
		Nonce:      txNonce,
		To:         &teleporterAddress,
		Gas:        2_000_000, // requiredGasLimit(500K) + TeleporterV2/OracleAdapter overhead
		GasFeeCap:  gasFeeCap,
		GasTipCap:  gasTipCap,
		Value:      common.Big0,
		Data:       callData,
		AccessList: icmutils.SignedWarpMessageToAccessList(signedMsg),
	})
	deliveryTx = utils.SignTransaction(deliveryTx, fundedKey, l1Info.EVMChainID)
	utils.SendTransactionAndWaitForSuccess(ctx, l1Info.EthClient, deliveryTx)

	ginkgo.By("Step 7: Assert MockOracleReceiver recorded the expected payload")
	receiveCount, assertErr := mockContract.ReceiveCount(&bind.CallOpts{})
	Expect(assertErr).Should(BeNil())
	Expect(receiveCount).Should(Equal(big.NewInt(1)))

	lastPayload, assertErr := mockContract.LastPayload(&bind.CallOpts{})
	Expect(assertErr).Should(BeNil())
	Expect(lastPayload).Should(Equal(msgPayload))

	lastSourceAddr, assertErr := mockContract.LastSourceAddress(&bind.CallOpts{})
	Expect(assertErr).Should(BeNil())
	Expect(lastSourceAddr).Should(Equal(sourceAddress))

	lastSourceChainID, assertErr := mockContract.LastSourceChainID(&bind.CallOpts{})
	Expect(assertErr).Should(BeNil())
	Expect(lastSourceChainID).Should(Equal(thisChainID))

	ginkgo.By("Step 8: Assert Teleporter recorded the message as received")
	msgID, assertErr := teleporterContract.CalculateMessageID(
		&bind.CallOpts{},
		thisChainID,
		thisChainID,
		icmMsg.Message.MessageNonce,
	)
	Expect(assertErr).Should(BeNil())
	received, assertErr := teleporterContract.MessageReceived(&bind.CallOpts{}, msgID)
	Expect(assertErr).Should(BeNil())
	Expect(received).Should(BeTrue())

	ginkgo.By("Sad path 3: replay of already-delivered nonce is rejected (AlreadyProcessed)")
	sendExpectRevert(oracleadapter.OracleMessage{
		SourceType:        "solana",
		SourceAddress:     sourceAddress,
		DestContract:      mockAddress,
		SourceBlockHeight: blockHeight,
		Nonce:             big.NewInt(1),
		Payload:           msgPayload,
	})
}
