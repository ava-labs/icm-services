// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	merkleregistry "github.com/ava-labs/icm-services/abi-bindings/go/MerkleValidatorSetRegistry"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/peers/clients"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// MerkleMessageRelay verifies that the relayer process itself (not the test) delivers a
// TeleporterV2 application message from an Avalanche L1 to an external EVM chain (the local
// "Ethereum" network) using the Merkle verification path.
//
// Flow:
//  1. Deploy a MerkleValidatorSetRegistry (the TeleporterMessengerV2 adapter) with the P-chain
//     genesis commitment on both the L1 and Ethereum at the same address (universal deployer).
//  2. Deploy TeleporterMessengerV2 bound to that registry on both chains (same address).
//  3. Start the relayer configured with: a source L1 "teleporter-v2" message-contracts entry and an
//     external-EVM delivery destination (contract-type "merkle"). The same external destination
//     also runs the Merkle validator-set updater that registers/refreshes the L1 set on Ethereum.
//  4. Wait for the relayer to register the L1 validator set on Ethereum.
//  5. Send a TeleporterV2 message from the L1 to Ethereum and assert the relayer delivers it
//     (MessageReceived becomes true on the Ethereum TeleporterMessengerV2).
//  6. Emit a message through the L1 registry adapter that the Ethereum messenger always rejects,
//     and assert the relayer abandons it: no transaction, height checkpointed, process still up.
//  7. Send another TeleporterV2 message and assert the relayer still delivers it.
func MerkleMessageRelay(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	log.Info("Starting MerkleMessageRelay e2e test")

	l1Info := avalancheNetwork.GetL1Infos()[2]
	l1BlockchainID := l1Info.BlockchainID
	networkID := avalancheNetwork.GetNetworkID()
	ethInfo := ethereumNetwork.EthereumTestInfo()
	ethBlockchainID := ethInfo.ChainID()

	fundedAddress, fundedKey := avalancheNetwork.GetFundedAccountInfo()
	ethFundedAddress, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()

	// =========================================================================
	// Compute the P-chain genesis commitment for the registry constructor
	// =========================================================================
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	pChainClient := clients.NewCanonicalValidatorClient(&config.APIConfig{
		BaseURL: primaryNetworkInfo.NodeURIs[0],
	})
	pChainHeight, err := pChainClient.GetLatestHeight(ctx)
	Expect(err).Should(BeNil())
	pChainTimestamp, err := pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())
	pChainWarpSet, err := pChainClient.GetProposedValidators(ctx, ids.Empty)
	Expect(err).Should(BeNil())

	pChainValidators := make([]*validatorupdater.Validator, len(pChainWarpSet.Validators))
	for i, val := range pChainWarpSet.Validators {
		pChainValidators[i] = &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(val.PublicKey.Serialize()),
			Weight:                     val.Weight,
		}
	}
	utils.SortValidators(pChainValidators)
	pChainGenesisRoot := validatorupdater.BuildMerkleRoot(pChainValidators)
	var pChainTotalWeight uint64
	for _, v := range pChainValidators {
		pChainTotalWeight += v.Weight
	}

	// =========================================================================
	// Deploy the MerkleValidatorSetRegistry (adapter) on both chains, then the
	// TeleporterMessengerV2 bound to it. The universal deployer yields identical
	// addresses across chains for both contracts.
	// =========================================================================
	registryAddr := utils.DeployMerkleValidatorSetRegistry(
		ctx, ethInfo, ethFundedKey, networkID, constants.PlatformChainID,
		pChainGenesisRoot, pChainTotalWeight, pChainHeight, pChainTimestamp, true,
	)
	registryAddrL1 := utils.DeployMerkleValidatorSetRegistry(
		ctx, &l1Info, fundedKey, networkID, constants.PlatformChainID,
		pChainGenesisRoot, pChainTotalWeight, pChainHeight, pChainTimestamp, true,
	)
	Expect(registryAddrL1).Should(Equal(registryAddr))

	// The initializer address is a constructor argument, so the same key must be used on both
	// chains for the contracts to land at the same address. Fund it on Ethereum so it can send
	// the initialize tx.
	ethereumNetwork.FundAccount(ctx, fundedAddress, new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10)))
	ethTeleporterAddr := utils.DeployTeleporterV2(ctx, ethInfo, registryAddr, ethFundedKey, fundedKey)
	l1TeleporterAddr := utils.DeployTeleporterV2(ctx, &l1Info, registryAddr, fundedKey, fundedKey)
	Expect(l1TeleporterAddr).Should(Equal(ethTeleporterAddr))

	registry, err := merkleregistry.NewMerkleValidatorSetRegistry(registryAddr, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())

	log.Info("Deployed merkle registry and TeleporterMessengerV2",
		zap.String("registryAddress", registryAddr.Hex()),
		zap.String("teleporterAddress", l1TeleporterAddr.Hex()),
	)

	// =========================================================================
	// Configure and start the relayer
	// =========================================================================
	Expect(utils.ClearRelayerStorage()).Should(BeNil())

	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1Info}, fundedKey, relayerKey)

	// The message-delivery client must use a different Ethereum account than the validator-set
	// updater (ethFundedKey) so their nonces never collide. Generate a dedicated delivery key
	// and fund it from the network's funded account.
	deliveryKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	deliveryAddress := crypto.PubkeyToAddress(deliveryKey.PublicKey)
	ethereumNetwork.FundAccount(ctx, deliveryAddress, new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10)))

	relayerConfig := createMerkleMessageRelayConfig(
		log, teleporter, l1Info, fundedAddress, relayerKey, ethereumNetwork,
		registryAddr, l1TeleporterAddr, ethFundedKey, deliveryKey, ethBlockchainID,
	)
	Expect(relayerConfig.Validate()).Should(BeNil())
	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	log.Info("Starting relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(ctx, log, relayerConfigPath, relayerConfig)
	defer relayerCleanup()

	startupCtx, startupCancel := context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	// =========================================================================
	// Wait for the relayer's Merkle updater to register the L1 validator set on Ethereum
	// =========================================================================
	callOpts := &bind.CallOpts{Context: ctx}
	log.Info("Waiting for relayer to register L1 validator set on Ethereum")
	pollForCommitmentUpdate(
		ctx, log, registry, callOpts, l1BlockchainID, 120*time.Second,
		"Timed out waiting for relayer to register L1 validator set",
		func(cmt merkleregistry.ValidatorSetMerkleCommitment) bool {
			return cmt.TotalWeight != 0
		},
		func(cmt merkleregistry.ValidatorSetMerkleCommitment) {
			log.Info("L1 validator set registered on Ethereum",
				zap.Uint64("totalWeight", cmt.TotalWeight),
				zap.Uint64("pChainHeight", cmt.PChainHeight),
				zap.String("root", hex.EncodeToString(cmt.Root[:])),
			)

			// Assert the on-chain commitment matches the expected L1 validator set at the recorded
			// P-chain height, not merely that some non-empty set was registered.
			expectedValidators := fetchSortedL1ValidatorsAtHeight(
				ctx, pChainClient, l1Info.SubnetID, cmt.PChainHeight,
			)
			expectedRoot := validatorupdater.BuildMerkleRoot(expectedValidators)
			var expectedWeight uint64
			for _, v := range expectedValidators {
				expectedWeight += v.Weight
			}
			Expect(cmt.Root).Should(Equal(expectedRoot),
				"Merkle root should match the L1 validator set at the recorded P-chain height")
			Expect(cmt.TotalWeight).Should(Equal(expectedWeight),
				"Total weight should match the sum of L1 validator weights")
		},
	)

	// =========================================================================
	// Send a TeleporterV2 message from the L1 to Ethereum and assert the relayer delivers it
	// =========================================================================
	l1Teleporter, err := teleportermessengerv2.NewTeleporterMessengerV2(l1TeleporterAddr, l1Info.EthClient)
	Expect(err).Should(BeNil())
	ethTeleporter, err := teleportermessengerv2.NewTeleporterMessengerV2(ethTeleporterAddr, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())
	l1Opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1Info.EVMChainID)
	Expect(err).Should(BeNil())

	route := merkleMessageRoute{
		log:              log,
		l1Info:           l1Info,
		l1Opts:           l1Opts,
		fundedKey:        fundedKey,
		l1Teleporter:     l1Teleporter,
		l1TeleporterAddr: l1TeleporterAddr,
		ethTeleporter:    ethTeleporter,
		ethBlockchainID:  ethBlockchainID,
		ethFundedAddress: ethFundedAddress,
	}
	route.relayMessage(ctx, []byte("Hello Ethereum from the relayer!"))

	// =========================================================================
	// A message the Ethereum messenger always rejects is abandoned, not retried
	// =========================================================================
	// The registry adapter lets any account emit a message that names the account itself as the
	// origin messenger. The Ethereum messenger rejects it only after the signature verifies, so
	// the message passes every relayer-side check and the relayer's gas estimation is what reveals
	// the revert. The relayer must mine nothing, checkpoint past the block and keep running.
	abandonedBefore, err := utils.RevertingDeliveryCount(ctx, relayerConfig.MetricsPort)
	Expect(err).Should(BeNil())
	nonceBefore, err := ethereumNetwork.EthClient.NonceAt(ctx, deliveryAddress, nil)
	Expect(err).Should(BeNil())

	log.Info("Sending a message from the L1 to Ethereum that the destination messenger always rejects")
	l1Registry, err := merkleregistry.NewMerkleValidatorSetRegistry(registryAddr, l1Info.EthClient)
	Expect(err).Should(BeNil())
	forgedTx, err := l1Registry.SendMessage(l1Opts, merkleregistry.TeleporterMessageV2{
		// A nonce the messenger never reaches, so the relayer does not take the forged message
		// for one the messenger has already received.
		MessageNonce:            new(big.Int).Lsh(big.NewInt(1), 128),
		OriginSenderAddress:     fundedAddress,
		OriginTeleporterAddress: fundedAddress,
		DestinationBlockchainID: ethBlockchainID,
		DestinationAddress:      ethFundedAddress,
		RequiredGasLimit:        big.NewInt(100_000),
		AllowedRelayerAddresses: []common.Address{},
		Receipts:                []merkleregistry.TeleporterMessageReceipt{},
		Message:                 []byte("Not from the messenger"),
	})
	Expect(err).Should(BeNil())
	forgedReceipt := utils.WaitForTransactionSuccess(ctx, l1Info.EthClient, forgedTx.Hash())
	Expect(utils.IssueTxsToAdvanceChain(ctx, l1Info.EVMChainID, fundedKey, l1Info.EthClient, 5)).Should(BeNil())

	utils.WaitForAbandonedDelivery(ctx, log, utils.AbandonedDelivery{
		RelayerConfig: relayerConfig,
		RelayerID: database.CalculateRelayerID(
			registryAddr,
			l1BlockchainID,
			ethBlockchainID,
			database.AllAllowedAddress,
			database.AllAllowedAddress,
		),
		SourceHeight:    forgedReceipt.BlockNumber.Uint64(),
		Destination:     ethereumNetwork.EthClient,
		RelayerAddress:  deliveryAddress,
		NonceBefore:     nonceBefore,
		AbandonedBefore: abandonedBefore,
	})

	// =========================================================================
	// The route still delivers after the abandoned message
	// =========================================================================
	route.relayMessage(ctx, []byte("Hello again from the relayer!"))

	log.Info("MerkleMessageRelay e2e test PASSED: relayer delivered L1 -> Ethereum messages")
}

// createMerkleMessageRelayConfig builds a relayer config whose source L1 routes its TeleporterV2
// messages (emitted via the Merkle registry adapter) to an external EVM delivery destination, and
// which also runs the Merkle validator-set updater against the same registry.
func createMerkleMessageRelayConfig(
	log logging.Logger,
	teleporter utils.TeleporterTestInfo,
	l1Info testinfo.L1TestInfo,
	fundedAddress common.Address,
	relayerKey *ecdsa.PrivateKey,
	ethereumNetwork *network.LocalEthereumNetwork,
	registryAddr common.Address,
	teleporterAddr common.Address,
	ethFundedKey *ecdsa.PrivateKey,
	deliveryKey *ecdsa.PrivateKey,
	ethBlockchainID ids.ID,
) relayercfg.Config {
	baseConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]testinfo.L1TestInfo{l1Info},
		[]testinfo.L1TestInfo{l1Info},
		fundedAddress,
		relayerKey,
	)
	baseConfig.APIPort = 8086

	// Register the TeleporterV2 (Merkle) message protocol for the source L1, keyed by the registry
	// adapter address (the warp message origin sender for these messages).
	source := baseConfig.SourceBlockchains[0]
	source.MessageContracts[registryAddr.Hex()] = relayercfg.MessageProtocolConfig{
		MessageFormat: relayercfg.TELEPORTER_V2.String(),
		Settings: map[string]interface{}{
			"reward-address":     fundedAddress.Hex(),
			"registry-address":   registryAddr.Hex(),
			"teleporter-address": teleporterAddr.Hex(),
		},
	}
	// Only deliver to the external Ethereum destination in this test.
	source.SupportedDestinations = []*relayercfg.SupportedDestination{
		{BlockchainID: ethBlockchainID.String()},
	}

	baseConfig.ExternalEVMDestinations = []*relayercfg.ExternalEVMDestination{
		{
			RPCEndpoint:              ethereumNetwork.BaseURL,
			PrivateKey:               hex.EncodeToString(crypto.FromECDSA(ethFundedKey)),
			ContractAddress:          registryAddr.Hex(),
			BlockchainID:             l1Info.BlockchainID.String(),
			SubnetID:                 l1Info.SubnetID.String(),
			PollIntervalSeconds:      testPollIntervalSeconds,
			MaxUpdateIntervalSeconds: merkleMaxUpdateIntervalSeconds,

			// Message delivery (enabled by setting DestinationBlockchainID). The delivery
			// client uses a dedicated key so it does not share a sender account with the
			// validator-set updater (PrivateKey above).
			DestinationBlockchainID: ethBlockchainID.String(),
			TeleporterAddress:       teleporterAddr.Hex(),
			DeliveryPrivateKey:      hex.EncodeToString(crypto.FromECDSA(deliveryKey)),
		},
	}

	return baseConfig
}

// merkleMessageRoute is the L1 to Ethereum TeleporterV2 route the relayer under test serves.
type merkleMessageRoute struct {
	log              logging.Logger
	l1Info           testinfo.L1TestInfo
	l1Opts           *bind.TransactOpts
	fundedKey        *ecdsa.PrivateKey
	l1Teleporter     *teleportermessengerv2.TeleporterMessengerV2
	l1TeleporterAddr common.Address
	ethTeleporter    *teleportermessengerv2.TeleporterMessengerV2
	ethBlockchainID  ids.ID
	ethFundedAddress common.Address
}

// relayMessage sends [message] from the L1 and asserts that the relayer delivers it unchanged to
// the Ethereum TeleporterMessengerV2.
func (r *merkleMessageRoute) relayMessage(ctx context.Context, message []byte) {
	input := teleportermessengerv2.TeleporterMessageInput{
		DestinationBlockchainID: r.ethBlockchainID,
		DestinationAddress:      r.ethFundedAddress,
		FeeInfo: teleportermessengerv2.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(100_000),
		AllowedRelayerAddresses: []common.Address{},
		Message:                 message,
	}
	tx, err := r.l1Teleporter.SendCrossChainMessage(r.l1Opts, input)
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, r.l1Info.EthClient, tx.Hash())

	sendEvent, err := utils.GetEventFromLogs(receipt.Logs, r.l1Teleporter.ParseSendCrossChainMessage)
	Expect(err).Should(BeNil())
	Expect(sendEvent.DestinationBlockchainID[:]).Should(Equal(r.ethBlockchainID[:]))

	messageID, err := teleporterUtils.CalculateMessageID(
		r.l1TeleporterAddr, r.l1Info.BlockchainID, r.ethBlockchainID, sendEvent.Message.MessageNonce,
	)
	Expect(err).Should(BeNil())
	r.log.Info("Sent TeleporterV2 message from L1 to Ethereum", zap.Stringer("messageID", messageID))

	// Advance the source chain a few blocks so the relayer's WS subscription reliably surfaces the
	// block containing the message. subnet-evm only produces blocks when there are transactions, so
	// without follow-on activity the message block's new-head notification can be delayed.
	Expect(utils.IssueTxsToAdvanceChain(ctx, r.l1Info.EVMChainID, r.fundedKey, r.l1Info.EthClient, 5)).Should(BeNil())

	deliveryCtx, deliveryCancel := context.WithTimeout(ctx, 120*time.Second)
	defer deliveryCancel()
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	delivered := false
	for !delivered {
		select {
		case <-deliveryCtx.Done():
			Expect(false).Should(BeTrue(), "Timed out waiting for relayer to deliver message to Ethereum")
		case <-ticker.C:
			received, recvErr := r.ethTeleporter.MessageReceived(&bind.CallOpts{Context: ctx}, messageID)
			if recvErr != nil {
				r.log.Warn("Failed to query MessageReceived", zap.Error(recvErr))
				continue
			}
			if received {
				delivered = true
			}
		}
	}

	// Explicitly filter for the ReceiveCrossChainMessage log emitted by the relayer's delivery and
	// assert the delivered message matches what was sent from the L1.
	receiveIter, err := r.ethTeleporter.FilterReceiveCrossChainMessage(
		&bind.FilterOpts{Context: ctx}, [][32]byte{messageID}, nil, nil,
	)
	Expect(err).Should(BeNil())
	defer receiveIter.Close()
	Expect(receiveIter.Next()).Should(BeTrue(),
		"expected a ReceiveCrossChainMessage log for the delivered message")
	Expect(receiveIter.Event.Message.Message).Should(Equal(message))
}
