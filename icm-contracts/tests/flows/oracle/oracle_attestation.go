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
	"crypto/ecdsa"
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
	signatureaggregatorcfg "github.com/ava-labs/icm-services/signature-aggregator/config"
	icmutils "github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

const (
	oracleSourceType       = "solana"
	oracleRequiredGasLimit = 500_000
	// oracleRequiredGasLimit plus TeleporterV2/OracleAdapter delivery overhead.
	oracleDeliveryGasLimit = 2_000_000
)

// oracleTestContracts bundles the contracts deployed for an oracle attestation
// flow together with the keys used to deploy and administer them.
type oracleTestContracts struct {
	fundedKey  *ecdsa.PrivateKey
	deployOpts *bind.TransactOpts

	adapterContract    *oracleadapter.OracleAdapter
	teleporterAddress  common.Address
	teleporterContract *teleportermessengerv2.TeleporterMessengerV2
	mockAddress        common.Address
	mockContract       *mockoraclereceiver.MockOracleReceiver
}

// OracleAttestation tests the full oracle attestation path:
//  1. Deploy OracleAdapter, TeleporterMessengerV2, and MockOracleReceiver on the L1
//  2. Construct an OracleMessage and request BLS aggregate signature
//  3. Deliver the signed message via TeleporterMessengerV2.receiveCrossChainMessage
//  4. Assert MockOracleReceiver received the expected payload
//  5. Verify non-allowlisted sources and replayed nonces are rejected
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
	contracts := deployOracleTestContracts(ctx, log, avalancheNetwork, l1Info)

	ginkgo.By("Step 4: Start signature aggregator")
	sigAggConfig, sigAggCancel := startOracleSignatureAggregator(ctx, log, l1Info)
	defer sigAggCancel()

	oracleMsg, justification := buildOracleMessage(ctx, log, solanaRPCURL, contracts.mockAddress)

	ginkgo.By("Step 6: Allowlist source on OracleAdapter")
	allowTx, err := contracts.adapterContract.SetAllowedSource(
		contracts.deployOpts, oracleSourceType, oracleMsg.SourceAddress, true,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, allowTx.Hash())

	ginkgo.By("Step 7: Request BLS aggregate signature from validators")
	networkID := avalancheNetwork.GetNetworkID()
	signedMsg := requestOracleAggregateSignature(
		ctx, log, sigAggConfig, networkID, l1Info, oracleMsg, justification,
	)

	thisChainID := [32]byte(l1Info.BlockchainID)
	icmMsg := utils.BuildOracleICMMessage(
		0, // warpIndex
		oracleMsg,
		contracts.teleporterAddress,
		thisChainID,
		networkID,
		big.NewInt(oracleRequiredGasLimit),
	)

	ginkgo.By("Step 8: Verify delivery from a non-allowlisted source is rejected (SourceNotAllowed)")
	removeTx, err := contracts.adapterContract.SetAllowedSource(
		contracts.deployOpts, oracleSourceType, oracleMsg.SourceAddress, false,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, removeTx.Hash())
	expectOracleDeliveryFailure(ctx, l1Info, contracts, signedMsg, icmMsg)
	restoreTx, err := contracts.adapterContract.SetAllowedSource(
		contracts.deployOpts, oracleSourceType, oracleMsg.SourceAddress, true,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, restoreTx.Hash())

	ginkgo.By("Step 9: Deliver the signed oracle message via TeleporterMessengerV2")
	deliveryTx := buildOracleDeliveryTx(
		ctx, l1Info, contracts, signedMsg, icmMsg, oracleDeliveryGasLimit,
	)
	utils.SendTransactionAndWaitForSuccess(ctx, l1Info.EthClient, deliveryTx)

	assertOracleDelivered(contracts, oracleMsg, thisChainID, icmMsg)

	ginkgo.By("Step 12: Verify replay of an already-delivered nonce is rejected (AlreadyProcessed)")
	expectOracleDeliveryFailure(ctx, l1Info, contracts, signedMsg, icmMsg)
}

// deployOracleTestContracts deploys OracleAdapter, TeleporterMessengerV2 (with
// the adapter as verifier), and MockOracleReceiver on the L1.
func deployOracleTestContracts(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	l1Info testinfo.L1TestInfo,
) oracleTestContracts {
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

	ginkgo.By("Step 2: Deploy TeleporterMessengerV2 with OracleAdapter as verifier")
	teleporterAddress := utils.DeployTeleporterV2(ctx, &l1Info, adapterAddress, fundedKey)
	teleporterContract, err := teleportermessengerv2.NewTeleporterMessengerV2(
		teleporterAddress, l1Info.EthClient,
	)
	Expect(err).Should(BeNil())
	log.Info("Deployed TeleporterMessengerV2", zap.Stringer("address", teleporterAddress))

	ginkgo.By("Step 3: Deploy MockOracleReceiver pointing to TeleporterMessengerV2")
	mockAddress, mockDeployTx, mockContract, err := mockoraclereceiver.DeployMockOracleReceiver(
		deployOpts, l1Info.EthClient, teleporterAddress,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, mockDeployTx.Hash())
	log.Info("Deployed MockOracleReceiver", zap.Stringer("address", mockAddress))

	return oracleTestContracts{
		fundedKey:          fundedKey,
		deployOpts:         deployOpts,
		adapterContract:    adapterContract,
		teleporterAddress:  teleporterAddress,
		teleporterContract: teleporterContract,
		mockAddress:        mockAddress,
		mockContract:       mockContract,
	}
}

// startOracleSignatureAggregator starts a signature-aggregator executable
// configured for the given L1 and blocks until its health check passes.
func startOracleSignatureAggregator(
	ctx context.Context,
	log logging.Logger,
	l1Info testinfo.L1TestInfo,
) (signatureaggregatorcfg.Config, context.CancelFunc) {
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

	startupCtx, startupCancel := context.WithTimeout(ctx, 20*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	return sigAggConfig, sigAggCancel
}

// buildOracleMessage chooses the oracle message source. With a Solana RPC URL it
// fetches a real Memo Program transaction; otherwise it uses mock data accepted
// unconditionally by the mock sidecar. Returns the oracle message and the
// justification bytes handed to the sidecar.
func buildOracleMessage(
	ctx context.Context,
	log logging.Logger,
	solanaRPCURL string,
	destContract common.Address,
) (oracleadapter.OracleMessage, []byte) {
	oracleMsg := oracleadapter.OracleMessage{
		SourceType:   oracleSourceType,
		DestContract: destContract,
		Nonce:        big.NewInt(1),
	}
	var justification []byte

	if solanaRPCURL != "" {
		ginkgo.By("Step 5: Fetch real Memo Program transaction from Solana devnet")
		txData := utils.FetchSolanaMemoTx(ctx, solanaRPCURL)
		oracleMsg.SourceAddress = txData.ProgramID
		oracleMsg.SourceBlockHeight = new(big.Int).SetUint64(txData.Slot)
		oracleMsg.Payload = txData.InstrData
		justification = txData.TxSigBytes
		log.Info("Using real Solana transaction",
			zap.String("program", oracleMsg.SourceAddress),
			zap.Stringer("slot", oracleMsg.SourceBlockHeight),
			zap.Int("payloadBytes", len(oracleMsg.Payload)),
		)
	} else {
		ginkgo.By("Step 5: Using mock oracle data (no SOLANA_RPC_URL set)")
		oracleMsg.SourceAddress = "4oracle1testaddr"
		oracleMsg.SourceBlockHeight = big.NewInt(100)
		oracleMsg.Payload = []byte("e2e-test-payload")
		justification = []byte("dummy-solana-tx-signature")
	}

	return oracleMsg, justification
}

// requestOracleAggregateSignature submits the oracle warp payload to the
// signature aggregator's oracle endpoint and returns the BLS-signed message.
func requestOracleAggregateSignature(
	ctx context.Context,
	log logging.Logger,
	sigAggConfig signatureaggregatorcfg.Config,
	networkID uint32,
	l1Info testinfo.L1TestInfo,
	oracleMsg oracleadapter.OracleMessage,
	justification []byte,
) *avalancheWarp.Message {
	oraclePayload := utils.PackOracleWarpPayload(oracleMsg)

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
	return signedMsg
}

// buildOracleDeliveryTx constructs a signed receiveCrossChainMessage transaction
// carrying the signed oracle warp message in its access list.
func buildOracleDeliveryTx(
	ctx context.Context,
	l1Info testinfo.L1TestInfo,
	contracts oracleTestContracts,
	signedMsg *avalancheWarp.Message,
	icmMsg teleportermessengerv2.TeleporterICMMessage,
	gasLimit uint64,
) *types.Transaction {
	callData, err := teleportermessengerv2.PackReceiveCrossChainMessageV2(
		icmMsg.Message,
		l1Info.BlockchainID,
		icmMsg.Attestation,
		common.Address{},
	)
	Expect(err).Should(BeNil())

	fundedAddress := utils.PrivateKeyToAddress(contracts.fundedKey)
	gasFeeCap, gasTipCap, txNonce := utils.CalculateTxParams(ctx, l1Info.EthClient, fundedAddress)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    l1Info.EVMChainID,
		Nonce:      txNonce,
		To:         &contracts.teleporterAddress,
		Gas:        gasLimit,
		GasFeeCap:  gasFeeCap,
		GasTipCap:  gasTipCap,
		Value:      common.Big0,
		Data:       callData,
		AccessList: icmutils.SignedWarpMessageToAccessList(signedMsg),
	})
	return utils.SignTransaction(tx, contracts.fundedKey, l1Info.EVMChainID)
}

// expectOracleDeliveryFailure delivers the message and asserts the transaction
// reverts (used for the SourceNotAllowed and AlreadyProcessed sad paths).
func expectOracleDeliveryFailure(
	ctx context.Context,
	l1Info testinfo.L1TestInfo,
	contracts oracleTestContracts,
	signedMsg *avalancheWarp.Message,
	icmMsg teleportermessengerv2.TeleporterICMMessage,
) {
	tx := buildOracleDeliveryTx(ctx, l1Info, contracts, signedMsg, icmMsg, oracleRequiredGasLimit)
	utils.SendTransactionAndWaitForFailure(ctx, l1Info.EthClient, tx)
}

// assertOracleDelivered asserts MockOracleReceiver recorded the delivered
// payload and Teleporter marked the message as received.
func assertOracleDelivered(
	contracts oracleTestContracts,
	oracleMsg oracleadapter.OracleMessage,
	thisChainID [32]byte,
	icmMsg teleportermessengerv2.TeleporterICMMessage,
) {
	ginkgo.By("Step 10: Assert MockOracleReceiver recorded the expected payload")
	receiveCount, err := contracts.mockContract.ReceiveCount(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(receiveCount).Should(Equal(big.NewInt(1)))

	lastPayload, err := contracts.mockContract.LastPayload(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(lastPayload).Should(Equal(oracleMsg.Payload))

	lastSourceAddr, err := contracts.mockContract.LastSourceAddress(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(lastSourceAddr).Should(Equal(oracleMsg.SourceAddress))

	lastSourceChainID, err := contracts.mockContract.LastSourceChainID(&bind.CallOpts{})
	Expect(err).Should(BeNil())
	Expect(lastSourceChainID).Should(Equal(thisChainID))

	ginkgo.By("Step 11: Assert Teleporter recorded the message as received")
	msgID, err := contracts.teleporterContract.CalculateMessageID(
		&bind.CallOpts{},
		thisChainID,
		thisChainID,
		icmMsg.Message.MessageNonce,
	)
	Expect(err).Should(BeNil())
	received, err := contracts.teleporterContract.MessageReceived(&bind.CallOpts{}, msgID)
	Expect(err).Should(BeNil())
	Expect(received).Should(BeTrue())
}
