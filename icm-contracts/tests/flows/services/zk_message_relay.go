// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	zkadapter "github.com/ava-labs/icm-services/abi-bindings/go/verifiers/ethereum/ZKAdapter"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// ZKMessageRelay runs the relayer with the local Ethereum network as a zk message source and an
// Avalanche L1 as its destination. It checks that the relayer skips messages the destination
// messenger would reject, without sending a transaction for them.
//
// The local Ethereum network has no beacon chain, and delivering a message requires beacon data.
// This test therefore only sends messages that the relayer skips before it fetches any.
//
// TODO: serve synthetic beacon data for the local Ethereum network and set an anchor with the
// adapter's manualTransition, then assert the relayer delivers a message, does not resend one
// already delivered, and still delivers after a skipped message.
//
// Flow:
//  1. Deploy a ZKAdapter on Ethereum and the L1 at the same address (Nick's method).
//  2. Deploy TeleporterMessengerV2 bound to that adapter on both chains (same address).
//  3. Start the relayer configured with Ethereum as a source with a "zk" message-contracts entry,
//     and the L1 as the destination.
//  4. Emit a message through the Ethereum adapter whose origin is not the messenger, and assert
//     the relayer skips it: no transaction, height checkpointed, process still up.
//  5. Send a TeleporterV2 message from Ethereum that only another relayer may deliver, and assert
//     the relayer skips it the same way.
func ZKMessageRelay(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	log.Info("Starting ZKMessageRelay e2e test")

	l1Info := avalancheNetwork.GetL1Infos()[0]
	l1BlockchainID := l1Info.BlockchainID
	ethInfo := ethereumNetwork.EthereumTestInfo()
	ethBlockchainID := ethInfo.ChainID()

	fundedAddress, fundedKey := avalancheNetwork.GetFundedAccountInfo()
	ethFundedAddress, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()

	// Deploy the ZKAdapter on both chains, then the TeleporterMessengerV2 bound
	// to it. Nick's method yields identical addresses across chains for both contracts.
	//
	// The adapter's starting epoch must not begin in the future on either chain.
	beaconGenesisTime := uint64(time.Now().Add(-time.Hour).Unix())
	adapterAddr := utils.DeployZKAdapter(ctx, ethInfo, ethFundedKey, beaconGenesisTime, fundedAddress)
	adapterAddrL1 := utils.DeployZKAdapter(ctx, &l1Info, fundedKey, beaconGenesisTime, fundedAddress)
	Expect(adapterAddrL1).Should(Equal(adapterAddr))

	// The initializer address is a constructor argument, so the same key must be used on both
	// chains for the contracts to land at the same address. Fund it on Ethereum so it can send
	// the initialize tx.
	ethereumNetwork.FundAccount(ctx, fundedAddress, new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10)))
	ethTeleporterAddr := utils.DeployTeleporterV2(ctx, ethInfo, adapterAddr, ethFundedKey, fundedKey)
	l1TeleporterAddr := utils.DeployTeleporterV2(ctx, &l1Info, adapterAddr, fundedKey, fundedKey)
	Expect(l1TeleporterAddr).Should(Equal(ethTeleporterAddr))

	log.Info("Deployed ZKAdapter and TeleporterMessengerV2",
		zap.String("adapterAddress", adapterAddr.Hex()),
		zap.String("teleporterAddress", ethTeleporterAddr.Hex()),
	)

	// Configure and start the relayer
	Expect(utils.ClearRelayerStorage()).Should(BeNil())

	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	relayerAddress := crypto.PubkeyToAddress(relayerKey.PublicKey)
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1Info}, fundedKey, relayerKey)

	relayerConfig := createZKMessageRelayConfig(
		log, teleporter, l1Info, fundedAddress, relayerKey, ethereumNetwork,
		adapterAddr, ethTeleporterAddr, beaconGenesisTime, ethBlockchainID,
	)
	Expect(relayerConfig.Validate()).Should(BeNil())
	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	log.Info("Starting relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(ctx, log, relayerConfigPath, relayerConfig)
	defer relayerCleanup()

	startupCtx, startupCancel := context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	jsonDB, err := database.NewJSONFileStorage(
		logging.NoLog{},
		relayerConfig.StorageLocation,
		database.GetConfigRelayerIDs(&relayerConfig),
	)
	Expect(err).Should(BeNil())
	relayerID := database.CalculateRelayerID(
		adapterAddr,
		ethBlockchainID,
		l1BlockchainID,
		database.AllAllowedAddress,
		database.AllAllowedAddress,
	)

	// expectSkipped asserts that the relayer checkpoints past the message in [sourceHeight]
	// without broadcasting a transaction for it, and keeps running.
	expectSkipped := func(sourceHeight uint64, nonceBefore uint64) {
		Eventually(func() (uint64, error) {
			return utils.RelayerCheckpoint(jsonDB, relayerID)
		}, 120*time.Second, time.Second).Should(BeNumerically(">=", sourceHeight),
			"the relayer did not checkpoint past the skipped message")

		nonce, err := l1Info.EthClient.NonceAt(ctx, relayerAddress, nil)
		Expect(err).Should(BeNil())
		Expect(nonce).Should(Equal(nonceBefore),
			"the relayer broadcast a transaction for a message it should have skipped")

		utils.ExpectRelayerHealthy(ctx, relayerConfig.APIPort)
	}

	ethOpts, err := bind.NewKeyedTransactorWithChainID(ethFundedKey, ethInfo.EVMChainID)
	Expect(err).Should(BeNil())

	// A message whose origin is not the messenger is skipped
	//
	// The adapter lets any account emit a message that names the account itself as the origin
	// messenger. The relayer must recognize it is not from the configured messenger and skip it.
	nonceBefore, err := l1Info.EthClient.NonceAt(ctx, relayerAddress, nil)
	Expect(err).Should(BeNil())

	log.Info("Sending a message from Ethereum whose origin is not the messenger")
	ethAdapter, err := zkadapter.NewZKAdapter(adapterAddr, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())
	forgedTx, err := ethAdapter.SendMessage(ethOpts, zkadapter.TeleporterMessageV2{
		MessageNonce:            big.NewInt(1),
		OriginSenderAddress:     ethFundedAddress,
		OriginTeleporterAddress: ethFundedAddress,
		DestinationBlockchainID: l1BlockchainID,
		DestinationAddress:      fundedAddress,
		RequiredGasLimit:        big.NewInt(100_000),
		AllowedRelayerAddresses: []common.Address{},
		Receipts:                []zkadapter.TeleporterMessageReceipt{},
		Message:                 []byte("Not from the messenger"),
	})
	Expect(err).Should(BeNil())
	forgedReceipt := utils.WaitForTransactionSuccess(ctx, ethereumNetwork.EthClient, forgedTx.Hash())
	expectSkipped(forgedReceipt.BlockNumber.Uint64(), nonceBefore)

	// A message only another relayer may deliver is skipped
	nonceBefore, err = l1Info.EthClient.NonceAt(ctx, relayerAddress, nil)
	Expect(err).Should(BeNil())

	log.Info("Sending a message from Ethereum that only another relayer may deliver")
	ethTeleporter, err := teleportermessengerv2.NewTeleporterMessengerV2(ethTeleporterAddr, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())
	otherRelayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	restrictedTx, err := ethTeleporter.SendCrossChainMessage(ethOpts, teleportermessengerv2.TeleporterMessageInput{
		DestinationBlockchainID: l1BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessengerv2.TeleporterFeeInfo{
			FeeTokenAddress: common.Address{},
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit:        big.NewInt(100_000),
		AllowedRelayerAddresses: []common.Address{crypto.PubkeyToAddress(otherRelayerKey.PublicKey)},
		Message:                 []byte("For another relayer"),
	})
	Expect(err).Should(BeNil())
	restrictedReceipt := utils.WaitForTransactionSuccess(ctx, ethereumNetwork.EthClient, restrictedTx.Hash())
	expectSkipped(restrictedReceipt.BlockNumber.Uint64(), nonceBefore)

	log.Info("ZKMessageRelay e2e test PASSED: relayer skipped undeliverable Ethereum -> L1 messages")
}

// createZKMessageRelayConfig builds a relayer config whose only source is the local Ethereum
// network, with a zk message-contracts entry for [adapterAddr], and whose destination is the L1.
func createZKMessageRelayConfig(
	log logging.Logger,
	teleporter utils.TeleporterTestInfo,
	l1Info testinfo.L1TestInfo,
	fundedAddress common.Address,
	relayerKey *ecdsa.PrivateKey,
	ethereumNetwork *network.LocalEthereumNetwork,
	adapterAddr common.Address,
	teleporterAddr common.Address,
	beaconGenesisTime uint64,
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

	// Ethereum is not an Avalanche chain, so its subnet ID is unused. A zk-only source has no
	// validators for the relayer to track.
	baseConfig.SourceBlockchains = []*relayercfg.SourceBlockchain{
		{
			SubnetID:     ids.Empty.String(),
			BlockchainID: ethBlockchainID.String(),
			RPCEndpoint:  config.APIConfig{BaseURL: ethereumNetwork.BaseURL},
			WSEndpoint:   config.APIConfig{BaseURL: ethereumNetwork.WSURL},
			MessageContracts: map[string]relayercfg.MessageProtocolConfig{
				adapterAddr.Hex(): {
					MessageFormat: relayercfg.ZK.String(),
					// Numbers are float64, as the relayer reads them back from the JSON config.
					Settings: map[string]interface{}{
						"reward-address":      fundedAddress.Hex(),
						"adapter-address":     adapterAddr.Hex(),
						"teleporter-address":  teleporterAddr.Hex(),
						"beacon-genesis-time": float64(beaconGenesisTime),
						"seconds-per-slot":    float64(12),
						// No beacon node runs locally. The messages in this test are skipped before
						// the relayer requests beacon data, so the endpoint is never called.
						// TODO: point this at the synthetic beacon API once the test delivers
						// messages (see the TODO on ZKMessageRelay).
						"beacon-api-url": "http://127.0.0.1:1",
					},
				},
			},
			SupportedDestinations: []*relayercfg.SupportedDestination{
				{BlockchainID: l1Info.BlockchainID.String()},
			},
		},
	}

	return baseConfig
}
