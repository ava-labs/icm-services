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
	teleportermessenger "github.com/ava-labs/icm-services/abi-bindings/go/teleporter/TeleporterMessenger"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/mocks"
	"github.com/ava-labs/icm-services/relayer/config"
	mock_evm "github.com/ava-labs/icm-services/vms/evm/mocks"
	mock_vms "github.com/ava-labs/icm-services/vms/mocks"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
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

// sourceMessage returns the source chain message that [unsignedMessage] was sent as by the
// Teleporter contract, as the relayer would pass it to the message handler factory.
func sourceMessage(unsignedMessage *warp.UnsignedMessage) *messages.SourceMessage {
	return &messages.SourceMessage{
		SourceBlockchainID: unsignedMessage.SourceChainID,
		ProtocolAddress:    messageProtocolAddress,
		Payload:            unsignedMessage.Bytes(),
	}
}

// warpMessageFor packs [teleporterMessage] the way the Teleporter contract does and wraps it in an
// unsigned Warp message sent from [sourceBlockchainID].
func warpMessageFor(
	t *testing.T,
	sourceBlockchainID ids.ID,
	teleporterMessage teleportermessenger.TeleporterMessage,
) *warp.UnsignedMessage {
	t.Helper()

	messageBytes, err := teleporterMessage.Pack()
	require.NoError(t, err)

	addressedCall, err := warpPayload.NewAddressedCall(messageProtocolAddress.Bytes(), messageBytes)
	require.NoError(t, err)

	unsignedMessage, err := warp.NewUnsignedMessage(0, sourceBlockchainID, addressedCall.Bytes())
	require.NoError(t, err)

	return unsignedMessage
}

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

	const blockGasLimit = 10_000_000

	gasLimitExceededTeleporterMessage := validTeleporterMessage
	gasLimitExceededTeleporterMessage.RequiredGasLimit = big.NewInt(blockGasLimit + 1)
	gasLimitExceededWarpUnsignedMessage := warpMessageFor(
		t,
		sourceBlockchainID,
		gasLimitExceededTeleporterMessage,
	)

	// A required gas limit exactly at the block gas limit still overflows a block once the
	// Teleporter and Warp verification overhead is added to the delivery transaction.
	gasLimitAtBlockLimitTeleporterMessage := validTeleporterMessage
	gasLimitAtBlockLimitTeleporterMessage.RequiredGasLimit = big.NewInt(blockGasLimit)
	gasLimitAtBlockLimitWarpUnsignedMessage := warpMessageFor(
		t,
		sourceBlockchainID,
		gasLimitAtBlockLimitTeleporterMessage,
	)

	// A required gas limit comfortably under the block gas limit still overflows a block once a
	// large payload is charged per byte for decoding and predicate verification.
	largePayloadTeleporterMessage := validTeleporterMessage
	largePayloadTeleporterMessage.RequiredGasLimit = big.NewInt(blockGasLimit - 500_000)
	largePayloadTeleporterMessage.Message = make([]byte, 20_000)
	largePayloadWarpUnsignedMessage := warpMessageFor(
		t,
		sourceBlockchainID,
		largePayloadTeleporterMessage,
	)

	// A required gas limit that does not fit in a uint64 must not be truncated into a passing value.
	overflowingGasLimitTeleporterMessage := validTeleporterMessage
	overflowingGasLimitTeleporterMessage.RequiredGasLimit = new(big.Int).Lsh(big.NewInt(1), 64)
	overflowingGasLimitWarpUnsignedMessage := warpMessageFor(
		t,
		sourceBlockchainID,
		overflowingGasLimitTeleporterMessage,
	)

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
		{
			name:                    "gas limit equal to block gas limit",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     gasLimitAtBlockLimitWarpUnsignedMessage,
			expectedResult:          false,
		},
		{
			name:                    "overhead of large payload exceeds gas limit",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     largePayloadWarpUnsignedMessage,
			expectedResult:          false,
		},
		{
			name:                    "gas limit does not fit in uint64",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     overflowingGasLimitWarpUnsignedMessage,
			expectedResult:          false,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// Set up mocks and the object under test
			ctrl := gomock.NewController(t)

			mockClient := mock_vms.NewMockDestinationClient(ctrl)

			factory, err := NewMessageHandlerFactory(
				messageProtocolAddress,
				messageProtocolConfig,
				nil,
			)
			require.NoError(t, err)
			mockClient.EXPECT().DestinationBlockchainID().Return(destinationBlockchainID).AnyTimes()
			handler, err := factory.NewMessageHandler(
				logging.NoLog{},
				sourceMessage(test.warpUnsignedMessage),
				mockClient,
				nil,
				mocks.NewMockMetrics(ctrl),
				ids.Empty,
				0,
			)
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
			result, err := handler.(*messageHandler).ShouldSendMessage()
			require.NoError(t, err)
			require.Equal(t, test.expectedResult, result)
		})
	}
}

func TestSendMessageAlreadyDelivered(t *testing.T) {
	// Set up test constants
	ctrl := gomock.NewController(t)

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
		messageProtocolAddress,
		messageProtocolConfig,
		nil,
	)
	require.NoError(t, err)
	mockClient.EXPECT().DestinationBlockchainID().Return(destinationBlockchainID).AnyTimes()
	handler, err := factory.NewMessageHandler(
		logging.NoLog{},
		sourceMessage(warpUnsignedMessage),
		mockClient,
		nil,
		mocks.NewMockMetrics(ctrl),
		ids.Empty,
		0,
	)
	require.NoError(t, err)

	mockEthClient := mock_evm.NewMockClient(ctrl)
	mockClient.EXPECT().
		Client().
		Return(mockEthClient).
		Times(1)

	mockClient.EXPECT().BlockGasLimit().Return(uint64(10_000_000)).AnyTimes()

	mockClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
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
	_, err = handler.(*messageHandler).SendMessage(signedMessage)
	require.NoError(t, err)
}

// TestSendMessageGasLimitExceedsBlockGasLimit checks that a message whose delivery transaction
// would not fit in a block is never handed to the destination client. Such a transaction is
// rejected deterministically by the destination's transaction pool, so retrying it would stall the
// relayer on the block containing the message.
func TestSendMessageGasLimitExceedsBlockGasLimit(t *testing.T) {
	ctrl := gomock.NewController(t)

	sourceBlockchainID := ids.Empty
	warpUnsignedMessage := warpMessageFor(t, sourceBlockchainID, validTeleporterMessage)

	signedMessage, err := warp.NewMessage(warpUnsignedMessage, &warp.BitSetSignature{})
	require.NoError(t, err)

	mockClient := mock_vms.NewMockDestinationClient(ctrl)
	mockClient.EXPECT().DestinationBlockchainID().Return(destinationBlockchainID).AnyTimes()
	// Below the fixed Teleporter overhead, so any delivery transaction exceeds it.
	mockClient.EXPECT().BlockGasLimit().Return(uint64(10_000)).AnyTimes()
	mockClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Times(0)

	factory, err := NewMessageHandlerFactory(messageProtocolAddress, messageProtocolConfig, nil)
	require.NoError(t, err)

	handler, err := factory.NewMessageHandler(
		logging.NoLog{},
		sourceMessage(warpUnsignedMessage),
		mockClient,
		nil,
		mocks.NewMockMetrics(ctrl),
		ids.Empty,
		0,
	)
	require.NoError(t, err)

	_, err = handler.(*messageHandler).SendMessage(signedMessage)
	require.ErrorIs(t, err, errUndeliverableGasLimit)
}
