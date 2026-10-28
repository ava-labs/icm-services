// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterV2

import (
	"context"
	"math/big"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	warpadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/WarpAdapter"
	"github.com/ava-labs/icm-services/database"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

// RevertingMessageAbandoned checks that the relayer abandons a TeleporterV2 message whose delivery
// reverts for a reason the message itself causes, instead of retrying it until the process exits
// and replaying it on every restart.
//
// The WarpAdapter lets any account emit a message that names the account itself as the origin
// messenger. The destination TeleporterMessengerV2 rejects such a message only after the Warp
// signature verifies, so it passes every relayer-side check and reverts on every delivery
// attempt. The relayer must see the revert in simulation, mine nothing, checkpoint past the
// block, keep serving both routes, and not replay the block after a restart.
func RevertingMessageAbandoned(
	ctx context.Context,
	log logging.Logger,
	network *localnetwork.LocalAvalancheNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	// The relayer keys its message contract entries, and so its checkpoints, by the adapter
	// address, which the universal deployer makes the same on every chain.
	warpAdapterAddress := teleporter.WarpAdapterAddress(l1AInfo.BlockchainID)
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	Expect(utils.ClearRelayerStorage()).Should(BeNil())

	//
	// Fund the relayer address on all subnets
	//
	log.Info("Funding relayer address on all subnets")
	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	relayerAddress := crypto.PubkeyToAddress(relayerKey.PublicKey)
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1AInfo, l1BInfo}, fundedKey, relayerKey)

	//
	// Set up and start the relayer
	//
	relayerConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]testinfo.L1TestInfo{l1AInfo, l1BInfo},
		[]testinfo.L1TestInfo{l1AInfo, l1BInfo},
		fundedAddress,
		relayerKey,
	)
	// The config needs to be validated in order to be passed to database.GetConfigRelayerIDs
	Expect(relayerConfig.Validate()).Should(BeNil())
	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	log.Info("Starting the relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(ctx, log, relayerConfigPath, relayerConfig)
	defer relayerCleanup()
	startupCtx, startupCancel := context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	log.Info("Relaying an ordinary message from L1 A to L1 B")
	utils.RelayBasicMessageV2(ctx, log, teleporter, l1AInfo, l1BInfo, fundedKey, fundedAddress)

	//
	// Emit a message through the adapter that names the sender as the origin messenger. The
	// destination messenger rejects it after the signature verifies, so only executing the
	// delivery reveals the revert.
	//
	abandonedBefore, err := utils.RevertingDeliveryCount(ctx, relayerConfig.MetricsPort)
	Expect(err).Should(BeNil())
	nonceBefore, err := l1BInfo.EthClient.NonceAt(ctx, relayerAddress, nil)
	Expect(err).Should(BeNil())

	log.Info("Sending a message from L1 A to L1 B that the destination messenger always rejects")
	adapter, err := warpadapter.NewWarpAdapter(warpAdapterAddress, l1AInfo.EthClient)
	Expect(err).Should(BeNil())
	opts, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := adapter.SendMessage(opts, warpadapter.TeleporterMessageV2{
		// A nonce the messenger never reaches, so the relayer does not take the forged message
		// for one the messenger has already received.
		MessageNonce:            new(big.Int).Lsh(big.NewInt(1), 128),
		OriginSenderAddress:     fundedAddress,
		OriginTeleporterAddress: fundedAddress,
		DestinationBlockchainID: l1BInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		RequiredGasLimit:        big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{},
		Receipts:                []warpadapter.TeleporterMessageReceipt{},
		Message:                 []byte{1, 2, 3, 4},
	})
	Expect(err).Should(BeNil())
	receipt := utils.WaitForTransactionSuccess(ctx, l1AInfo.EthClient, tx.Hash())

	utils.WaitForAbandonedDelivery(ctx, log, utils.AbandonedDelivery{
		RelayerConfig: relayerConfig,
		RelayerID: database.CalculateRelayerID(
			warpAdapterAddress,
			l1AInfo.BlockchainID,
			l1BInfo.BlockchainID,
			database.AllAllowedAddress,
			database.AllAllowedAddress,
		),
		SourceHeight:    receipt.BlockNumber.Uint64(),
		Destination:     l1BInfo.EthClient,
		RelayerAddress:  relayerAddress,
		NonceBefore:     nonceBefore,
		AbandonedBefore: abandonedBefore,
	})

	//
	// The relayer still serves both routes.
	//
	log.Info("Relaying ordinary messages in both directions after the abandoned delivery")
	utils.RelayBasicMessageV2(ctx, log, teleporter, l1AInfo, l1BInfo, fundedKey, fundedAddress)
	utils.RelayBasicMessageV2(ctx, log, teleporter, l1BInfo, l1AInfo, fundedKey, fundedAddress)

	//
	// A restart that processes missed blocks starts after the checkpoint, so the abandoned
	// message is never seen again.
	//
	log.Info("Restarting the relayer to check that the abandoned message is not replayed")
	relayerCleanup()
	relayerConfig.ProcessMissedBlocks = true
	relayerConfigPath = utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)
	relayerCleanup, readyChan = utils.RunRelayerExecutable(ctx, log, relayerConfigPath, relayerConfig)
	defer relayerCleanup()
	startupCtx, startupCancel = context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	Consistently(func() (float64, error) {
		return utils.RevertingDeliveryCount(ctx, relayerConfig.MetricsPort)
	}, 10*time.Second, time.Second).Should(BeZero(), "the restarted relayer replayed the abandoned message")
	utils.RelayBasicMessageV2(ctx, log, teleporter, l1AInfo, l1BInfo, fundedKey, fundedAddress)
}
