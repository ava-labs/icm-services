// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleportermessenger "github.com/ava-labs/icm-contracts/abi-bindings/go/teleporter/TeleporterMessenger"
	teleporterUtils "github.com/ava-labs/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/relayer/config"
	mock_evm "github.com/ava-labs/icm-services/vms/evm/mocks"
	mock_vms "github.com/ava-labs/icm-services/vms/mocks"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type CallContractChecker struct {
	input          []byte
	expectedResult []byte
	times          int
}

var (
	messageProtocolAddress = common.HexToAddress("0xd81545385803bCD83bd59f58Ba2d2c0562387F83")
	messageProtocolConfig  = config.MessageProtocolConfig{
		MessageFormat: config.TELEPORTER.String(),
		Settings: map[string]interface{}{
			"reward-address": "0x27aE10273D17Cd7e80de8580A51f476960626e5f",
		},
	}
	destinationBlockchainIDString = "S4mMqUXe7vHsGiRAma6bv3CKnyaLssyAxmQ2KvFpX1KEvfFCD"
	destinationBlockchainID       ids.ID
	validRelayerAddress           = common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")
	validTeleporterMessage        teleportermessenger.TeleporterMessage
)

func init() {
	var err error
	destinationBlockchainID, err = ids.FromString(destinationBlockchainIDString)
	if err != nil {
		panic(err)
	}

	validTeleporterMessage = teleportermessenger.TeleporterMessage{
		MessageNonce:            big.NewInt(1),
		OriginSenderAddress:     common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		DestinationBlockchainID: destinationBlockchainID,
		DestinationAddress:      common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		RequiredGasLimit:        big.NewInt(2),
		AllowedRelayerAddresses: []common.Address{
			validRelayerAddress,
		},
		Receipts: []teleportermessenger.TeleporterMessageReceipt{
			{
				ReceivedMessageNonce: big.NewInt(1),
				RelayerRewardAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
			},
		},
		Message: []byte{1, 2, 3, 4},
	}
}

func TestShouldSendMessage(t *testing.T) {
	// Define test constants
	validMessageBytes, err := validTeleporterMessage.Pack()
	require.NoError(t, err)

	validAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		validMessageBytes,
	)
	require.NoError(t, err)

	sourceBlockchainID := ids.Empty
	warpUnsignedMessage, err := warp.NewUnsignedMessage(
		0,
		sourceBlockchainID,
		validAddressedCall.Bytes(),
	)
	require.NoError(t, err)

	messageID, err := teleporterUtils.CalculateMessageID(
		messageProtocolAddress,
		sourceBlockchainID,
		destinationBlockchainID,
		validTeleporterMessage.MessageNonce,
	)
	require.NoError(t, err)

	messageReceivedInput, err := teleportermessenger.PackMessageReceived(messageID)
	require.NoError(t, err)

	messageNotDelivered, err := teleportermessenger.PackMessageReceivedOutput(false)
	require.NoError(t, err)

	messageDelivered, err := teleportermessenger.PackMessageReceivedOutput(true)
	require.NoError(t, err)

	invalidAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		validMessageBytes,
	)
	require.NoError(t, err)
	invalidWarpUnsignedMessage, err := warp.NewUnsignedMessage(
		0,
		sourceBlockchainID,
		append(invalidAddressedCall.Bytes(), []byte{1, 2, 3, 4}...),
	)
	require.NoError(t, err)

	const blockGasLimit = 10_000
	gasLimitExceededTeleporterMessage := validTeleporterMessage
	gasLimitExceededTeleporterMessage.RequiredGasLimit = big.NewInt(blockGasLimit + 1)
	gasLimitExceededTeleporterMessageBytes, err := gasLimitExceededTeleporterMessage.Pack()
	require.NoError(t, err)

	gasLimitExceededAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		gasLimitExceededTeleporterMessageBytes,
	)
	require.NoError(t, err)

	gasLimitExceededWarpUnsignedMessage, err := warp.NewUnsignedMessage(
		0,
		sourceBlockchainID,
		gasLimitExceededAddressedCall.Bytes(),
	)
	require.NoError(t, err)

	testCases := []struct {
		name                    string
		destinationBlockchainID ids.ID
		warpUnsignedMessage     *warp.UnsignedMessage
		senderAddressesResult   []common.Address
		senderAddressesTimes    int
		clientTimes             int
		messageReceivedCall     *CallContractChecker
		expectedParseError      bool
		expectedResult          bool
	}{
		{
			name:                    "valid message",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     warpUnsignedMessage,
			senderAddressesResult:   []common.Address{validRelayerAddress},
			senderAddressesTimes:    1,
			clientTimes:             1,
			messageReceivedCall: &CallContractChecker{
				input:          messageReceivedInput,
				expectedResult: messageNotDelivered,
				times:          1,
			},
			expectedResult: true,
		},
		{
			name:                    "invalid message",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     invalidWarpUnsignedMessage,
			expectedParseError:      true,
		},
		{
			name:                    "invalid destination chain id",
			destinationBlockchainID: ids.Empty,
			senderAddressesResult:   []common.Address{common.Address{}},
			senderAddressesTimes:    1,
			warpUnsignedMessage:     warpUnsignedMessage,
		},
		{
			name:                    "not allowed",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     warpUnsignedMessage,
			senderAddressesResult:   []common.Address{common.Address{}},
			senderAddressesTimes:    1,
			clientTimes:             0,
			expectedResult:          false,
		},
		{
			name:                    "message already delivered",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     warpUnsignedMessage,
			senderAddressesResult:   []common.Address{validRelayerAddress},
			senderAddressesTimes:    1,
			clientTimes:             1,
			messageReceivedCall: &CallContractChecker{
				input:          messageReceivedInput,
				expectedResult: messageDelivered,
				times:          1,
			},
			expectedResult: false,
		},
		{
			name:                    "gas limit exceeded",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     gasLimitExceededWarpUnsignedMessage,
			expectedResult:          false,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// Set up mocks and the object under test
			ctrl := gomock.NewController(t)
			logger := logging.NoLog{}

			mockClient := mock_vms.NewMockDestinationClient(ctrl)

			factory, err := NewMessageHandlerFactory(
				logger,
				messageProtocolAddress,
				messageProtocolConfig,
				nil,
			)
			require.NoError(t, err)
			mockClient.EXPECT().DestinationBlockchainID().Return(destinationBlockchainID).AnyTimes()
			messageHandler, err := factory.NewMessageHandler(test.warpUnsignedMessage, mockClient)
			if test.expectedParseError {
				// If we expect an error parsing the Warp message, we should not call ShouldSendMessage
				require.Error(t, err)
				return
			} else {
				require.NoError(t, err)
			}
			mockEthClient := mock_evm.NewMockClient(ctrl)
			mockClient.EXPECT().
				Client().
				Return(mockEthClient).
				Times(test.clientTimes)
			mockClient.EXPECT().
				SenderAddresses().
				Return(test.senderAddressesResult).
				Times(test.senderAddressesTimes)
			mockClient.EXPECT().BlockGasLimit().Return(uint64(blockGasLimit)).AnyTimes()
			if test.messageReceivedCall != nil {
				messageReceivedInput := ethereum.CallMsg{
					From: bind.CallOpts{}.From,
					To:   &messageProtocolAddress,
					Data: test.messageReceivedCall.input,
				}
				mockEthClient.EXPECT().
					CallContract(gomock.Any(), gomock.Eq(messageReceivedInput), gomock.Any()).
					Return(test.messageReceivedCall.expectedResult, nil).
					Times(test.messageReceivedCall.times)
			}

			// Call the method under test
			result, err := messageHandler.ShouldSendMessage()
			require.NoError(t, err)
			require.Equal(t, test.expectedResult, result)
		})
	}
}

func TestSendMessageAlreadyDelivered(t *testing.T) {
	// Set up test constants
	ctrl := gomock.NewController(t)
	logger := logging.NoLog{}

	validMessageBytes, err := validTeleporterMessage.Pack()
	require.NoError(t, err)

	validAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		validMessageBytes,
	)
	require.NoError(t, err)

	sourceBlockchainID := ids.Empty
	warpUnsignedMessage, err := warp.NewUnsignedMessage(
		0,
		sourceBlockchainID,
		validAddressedCall.Bytes(),
	)
	require.NoError(t, err)

	messageID, err := teleporterUtils.CalculateMessageID(
		messageProtocolAddress,
		sourceBlockchainID,
		destinationBlockchainID,
		validTeleporterMessage.MessageNonce,
	)
	require.NoError(t, err)

	messageReceivedCallData, err := teleportermessenger.PackMessageReceived(messageID)
	require.NoError(t, err)

	messageReceivedInput := ethereum.CallMsg{
		From: bind.CallOpts{}.From,
		To:   &messageProtocolAddress,
		Data: messageReceivedCallData,
	}

	messageDeliveredResult, err := teleportermessenger.PackMessageReceivedOutput(true)
	require.NoError(t, err)

	signedMessage, err := warp.NewMessage(
		warpUnsignedMessage,
		&warp.BitSetSignature{},
	)
	require.NoError(t, err)

	// Set up mocks and the object under test
	mockClient := mock_vms.NewMockDestinationClient(ctrl)

	factory, err := NewMessageHandlerFactory(
		logger,
		messageProtocolAddress,
		messageProtocolConfig,
		nil,
	)
	require.NoError(t, err)
	mockClient.EXPECT().DestinationBlockchainID().Return(destinationBlockchainID).AnyTimes()
	messageHandler, err := factory.NewMessageHandler(warpUnsignedMessage, mockClient)
	require.NoError(t, err)

	mockEthClient := mock_evm.NewMockClient(ctrl)
	mockClient.EXPECT().
		Client().
		Return(mockEthClient).
		Times(1)

	mockClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(
			&types.Receipt{
				Status: types.ReceiptStatusFailed,
			},
			nil,
		).Times(1)

	mockEthClient.EXPECT().
		CallContract(gomock.Any(), gomock.Eq(messageReceivedInput), gomock.Any()).
		Return(messageDeliveredResult, nil).
		Times(1)

	// Call the method under test
	_, err = messageHandler.SendMessage(signedMessage)
	require.NoError(t, err)
}
