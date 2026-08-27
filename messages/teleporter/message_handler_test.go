// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter

import (
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleportermessenger "github.com/ava-labs/icm-services/abi-bindings/go/teleporter/TeleporterMessenger"
	gasUtils "github.com/ava-labs/icm-services/icm-contracts/utils/gas-utils"
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

	const blockGasLimit = 15_000_000
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

	// A message shaped like one produced by Teleporter's permissionless sendSpecifiedReceipts:
	// RequiredGasLimit of zero, no allowed-relayer restriction, and enough receipts that marking
	// them cannot fit in a block. The old check only looked at RequiredGasLimit and so waved
	// this through, after which the relayer broadcast an under-provisioned delivery and paid for
	// the out-of-gas revert.
	receiptHeavyTeleporterMessage := validTeleporterMessage
	receiptHeavyTeleporterMessage.RequiredGasLimit = big.NewInt(0)
	receiptHeavyTeleporterMessage.AllowedRelayerAddresses = nil
	numReceipts := blockGasLimit/int(gasUtils.MarkMessageReceiptGasCost) + 1
	receiptHeavyTeleporterMessage.Receipts = make(
		[]teleportermessenger.TeleporterMessageReceipt,
		numReceipts,
	)
	for i := range receiptHeavyTeleporterMessage.Receipts {
		receiptHeavyTeleporterMessage.Receipts[i] = teleportermessenger.TeleporterMessageReceipt{
			ReceivedMessageNonce: big.NewInt(int64(i + 1)),
			RelayerRewardAddress: common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567"),
		}
	}
	receiptHeavyTeleporterMessageBytes, err := receiptHeavyTeleporterMessage.Pack()
	require.NoError(t, err)

	receiptHeavyAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		receiptHeavyTeleporterMessageBytes,
	)
	require.NoError(t, err)

	receiptHeavyWarpUnsignedMessage, err := warp.NewUnsignedMessage(
		0,
		sourceBlockchainID,
		receiptHeavyAddressedCall.Bytes(),
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
		{
			name:                    "receipt count exceeds block gas limit",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     receiptHeavyWarpUnsignedMessage,
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
	mockClient.EXPECT().BlockGasLimit().Return(uint64(15_000_000)).AnyTimes()
	mockClient.EXPECT().
		Client().
		Return(mockEthClient).
		Times(1)

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

// A delivery that reverts having consumed its entire gas limit ran out of gas. That is
// deterministic, so the relayer must surface it as non-retryable rather than re-broadcasting and
// paying for the same reverted transaction again.
func TestSendMessageOutOfGasIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)

	validMessageBytes, err := validTeleporterMessage.Pack()
	require.NoError(t, err)

	validAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		validMessageBytes,
	)
	require.NoError(t, err)

	warpUnsignedMessage, err := warp.NewUnsignedMessage(0, ids.Empty, validAddressedCall.Bytes())
	require.NoError(t, err)

	signedMessage, err := warp.NewMessage(warpUnsignedMessage, &warp.BitSetSignature{})
	require.NoError(t, err)

	messageNotDelivered, err := teleportermessenger.PackMessageReceivedOutput(false)
	require.NoError(t, err)

	mockClient := mock_vms.NewMockDestinationClient(ctrl)
	mockClient.EXPECT().DestinationBlockchainID().Return(destinationBlockchainID).AnyTimes()
	mockClient.EXPECT().BlockGasLimit().Return(uint64(15_000_000)).AnyTimes()

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

	mockEthClient := mock_evm.NewMockClient(ctrl)
	mockClient.EXPECT().Client().Return(mockEthClient).Times(1)

	// Echo the gas limit the handler chose back as gas used: that is what an out-of-gas revert
	// looks like on chain.
	mockClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(
			_ logging.Logger,
			_ types.AccessList,
			_ set.Set[common.Address],
			_ common.Address,
			gasLimit uint64,
			_ []byte,
		) (*types.Receipt, error) {
			return &types.Receipt{
				Status:  types.ReceiptStatusFailed,
				GasUsed: gasLimit,
			}, nil
		}).Times(1)

	mockEthClient.EXPECT().
		CallContract(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(messageNotDelivered, nil).
		Times(1)

	_, err = handler.(*messageHandler).SendMessage(signedMessage)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
}

// A revert that left gas unspent is a logic failure, e.g. Warp verification against a validator
// set that has churned. A fresh signature may fix it, so it must stay retryable.
func TestSendMessageRevertWithGasLeftIsRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)

	validMessageBytes, err := validTeleporterMessage.Pack()
	require.NoError(t, err)
	validAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		validMessageBytes,
	)
	require.NoError(t, err)
	warpUnsignedMessage, err := warp.NewUnsignedMessage(0, ids.Empty, validAddressedCall.Bytes())
	require.NoError(t, err)
	signedMessage, err := warp.NewMessage(warpUnsignedMessage, &warp.BitSetSignature{})
	require.NoError(t, err)

	messageNotDelivered, err := teleportermessenger.PackMessageReceivedOutput(false)
	require.NoError(t, err)

	mockClient := mock_vms.NewMockDestinationClient(ctrl)
	mockClient.EXPECT().DestinationBlockchainID().Return(destinationBlockchainID).AnyTimes()
	mockClient.EXPECT().BlockGasLimit().Return(uint64(15_000_000)).AnyTimes()

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

	mockEthClient := mock_evm.NewMockClient(ctrl)
	mockClient.EXPECT().Client().Return(mockEthClient).Times(1)
	mockClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: 21_000}, nil).
		Times(1)
	mockEthClient.EXPECT().
		CallContract(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(messageNotDelivered, nil).
		Times(1)

	_, err = handler.(*messageHandler).SendMessage(signedMessage)
	require.Error(t, err)
	require.NotErrorIs(t, err, messages.ErrNonRetryable)
}
