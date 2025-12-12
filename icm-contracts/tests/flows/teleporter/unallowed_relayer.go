package teleporter

import (
	"context"
	"math/big"

	teleportermessenger "github.com/ava-labs/icm-services/abi-bindings/go/teleporter/TeleporterMessenger"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

func UnallowedRelayer(
	ctx context.Context,
	network *localnetwork.LocalNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	l1AInfo := network.GetPrimaryNetworkInfo()
	l1BInfo, _ := network.GetTwoL1s()
	fundedAddress, fundedKey := network.GetFundedAccountInfo()

	//
	// Send a transaction to L1 A to issue an ICM Message from the Teleporter contract to L1 B
	// The Teleporter message includes an allowed relayer list that does NOT include the relayer
	//

	sendCrossChainMessageInput := teleportermessenger.TeleporterMessageInput{
		DestinationBlockchainID: l1BInfo.BlockchainID,
		DestinationAddress:      fundedAddress,
		FeeInfo: teleportermessenger.TeleporterFeeInfo{
			FeeTokenAddress: fundedAddress,
			Amount:          big.NewInt(0),
		},
		RequiredGasLimit: big.NewInt(1),
		AllowedRelayerAddresses: []common.Address{
			common.HexToAddress("0x0123456789012345678901234567890123456789"),
		},
		Message: []byte{1, 2, 3, 4},
	}

	log.Info(
		"Sending Teleporter transaction on source chain",
		zap.Stringer("destinationBlockchainID", l1BInfo.BlockchainID),
	)
	receipt, teleporterMessageID := utils.SendCrossChainMessageAndWaitForAcceptance(
		ctx, teleporter.TeleporterMessenger(l1AInfo), l1AInfo, l1BInfo, sendCrossChainMessageInput, fundedKey,
	)

	aggregator := network.GetSignatureAggregator()
	defer aggregator.Shutdown()

	//
	// Relay the message to the destination
	//
	teleporter.RelayTeleporterMessage(
		ctx,
		receipt,
		l1AInfo,
		l1BInfo,
		false,
		fundedKey,
		nil,
		aggregator,
	)

	//
	// Check Teleporter message was not received on the destination
	//
	delivered, err := teleporter.TeleporterMessenger(l1BInfo).MessageReceived(
		&bind.CallOpts{}, teleporterMessageID,
	)
	Expect(err).Should(BeNil())
	Expect(delivered).Should(BeFalse())
}
