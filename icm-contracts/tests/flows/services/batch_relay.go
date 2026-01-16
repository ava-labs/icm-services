package tests

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/icm-services/icm-contracts/tests/interfaces"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
)

// Processes multiple Warp messages contained in the same block
func BatchRelay(
	ctx context.Context,
	log logging.Logger,
	network *network.LocalAvalancheNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	l1AInfo, l1BInfo := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()
	err := utils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	//
	// Deploy the batch messenger contracts
	//
	_, batchMessengerA := utils.DeployBatchCrossChainMessenger(
		ctx,
		fundedKey,
		teleporter,
		fundedAddress,
		l1AInfo,
	)
	batchMessengerAddressB, batchMessengerB := utils.DeployBatchCrossChainMessenger(
		ctx,
		fundedKey,
		teleporter,
		fundedAddress,
		l1BInfo,
	)

	//
	// Fund the relayer address on all subnets
	//

	log.Info("Funding relayer address on all subnets")
	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []interfaces.L1TestInfo{l1AInfo, l1BInfo}, fundedKey, relayerKey)

	//
	// Set up relayer config
	//
	relayerConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]interfaces.L1TestInfo{l1AInfo, l1BInfo},
		[]interfaces.L1TestInfo{l1AInfo, l1BInfo},
		fundedAddress,
		relayerKey,
	)

	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	log.Info("Starting the relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(
		ctx,
		log,
		relayerConfigPath,
		relayerConfig,
	)
	defer relayerCleanup()

	// Wait for relayer to start up
	log.Info("Waiting for the relayer to start up")
	startupCtx, startupCancel := context.WithTimeout(ctx, 15*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	//
	// Send a batch message from subnet A -> B
	//

	newHeadsDest := make(chan *types.Header, 10)
	sub, err := l1BInfo.WSClient.SubscribeNewHead(ctx, newHeadsDest)
	Expect(err).Should(BeNil())
	defer sub.Unsubscribe()

	numMessages := 40
	sentMessages := set.NewSet[string](numMessages)
	for i := 0; i < numMessages; i++ {
		sentMessages.Add(strconv.Itoa(i))
	}

	optsA, err := bind.NewKeyedTransactorWithChainID(fundedKey, l1AInfo.EVMChainID)
	Expect(err).Should(BeNil())
	tx, err := batchMessengerA.SendMessages(
		optsA,
		l1BInfo.BlockchainID,
		batchMessengerAddressB,
		common.Address{},
		big.NewInt(0),
		big.NewInt(int64(300000*numMessages)),
		sentMessages.List(),
	)
	Expect(err).Should(BeNil())

	utils.WaitForTransactionSuccess(ctx, l1AInfo.RPCClient, tx.Hash())

	// Wait for the message on the destination
	maxWait := 40
	currWait := 0
	log.Info("Waiting to receive all messages on destination...")
	for {
		receivedMessages, err := batchMessengerB.GetCurrentMessages(&bind.CallOpts{}, l1AInfo.BlockchainID)
		Expect(err).Should(BeNil())

		// Remove the received messages from the set of sent messages
		sentMessages.Remove(receivedMessages...)
		if sentMessages.Len() == 0 {
			break
		}
		currWait++
		if currWait == maxWait {
			Expect(false).Should(BeTrue(),
				fmt.Sprintf(
					"did not receive all sent messages in time. received %d/%d",
					numMessages-sentMessages.Len(),
					numMessages,
				),
			)
		}
		time.Sleep(1 * time.Second)
	}
}
