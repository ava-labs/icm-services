// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporter

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/evm/predicate"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	teleportermessenger "github.com/ava-labs/icm-services/abi-bindings/go/teleporter/TeleporterMessenger"
	gasUtils "github.com/ava-labs/icm-services/icm-contracts/utils/gas-utils"
	teleporterUtils "github.com/ava-labs/icm-services/icm-contracts/utils/teleporter-utils"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/mocks"
	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/icm-services/vms/evm"
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

	// Locate the boundary of the receipt-aware block gas limit check so it is tested from both
	// sides. The delivery gas limit is linear in the receipt count, so two samples give the fixed
	// cost and the per-receipt cost.
	_, gasNoReceipts := receiptHeavyUnsignedMessage(t, 0)
	_, gasOneReceipt := receiptHeavyUnsignedMessage(t, 1)
	perReceiptGas := gasOneReceipt - gasNoReceipts
	maxFittingReceipts := int((blockGasLimit - gasNoReceipts) / perReceiptGas)

	// The most receipts whose delivery still fits in a block must be relayed.
	receiptBoundaryWarpUnsignedMessage, minGasLimit := receiptHeavyUnsignedMessage(
		t,
		maxFittingReceipts,
	)
	require.LessOrEqual(t, minGasLimit, uint64(blockGasLimit))

	// One more receipt must be declined. The old check only looked at RequiredGasLimit and so
	// waved this through, after which the relayer broadcast an under-provisioned delivery and
	// paid for the out-of-gas revert.
	receiptHeavyWarpUnsignedMessage, minGasLimit := receiptHeavyUnsignedMessage(
		t,
		maxFittingReceipts+1,
	)
	require.Greater(t, minGasLimit, uint64(blockGasLimit))

	// A requiredGasLimit of 2^64 truncates to 0 if converted with Uint64() without a bounds check,
	// which would let it slip past the block gas limit guard.
	gasLimitOverflowTeleporterMessage := validTeleporterMessage
	gasLimitOverflowTeleporterMessage.RequiredGasLimit = new(big.Int).Lsh(big.NewInt(1), 64)
	gasLimitOverflowTeleporterMessageBytes, err := gasLimitOverflowTeleporterMessage.Pack()
	require.NoError(t, err)

	gasLimitOverflowAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		gasLimitOverflowTeleporterMessageBytes,
	)
	require.NoError(t, err)

	gasLimitOverflowWarpUnsignedMessage, err := warp.NewUnsignedMessage(
		0,
		sourceBlockchainID,
		gasLimitOverflowAddressedCall.Bytes(),
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
			name:                    "receipt count just fits in block gas limit",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     receiptBoundaryWarpUnsignedMessage,
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
			name:                    "receipt count exceeds block gas limit by one",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     receiptHeavyWarpUnsignedMessage,
			expectedResult:          false,
		},
		{
			name:                    "gas limit does not fit in uint64",
			destinationBlockchainID: destinationBlockchainID,
			warpUnsignedMessage:     gasLimitOverflowWarpUnsignedMessage,
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

// receiptHeavyUnsignedMessage builds a message shaped like one produced by Teleporter's
// permissionless sendSpecifiedReceipts: RequiredGasLimit of zero, no allowed-relayer restriction,
// and numReceipts receipts to mark. It also returns the minimum delivery gas limit that
// ShouldSendMessage will compute for it.
func receiptHeavyUnsignedMessage(t *testing.T, numReceipts int) (*warp.UnsignedMessage, uint64) {
	t.Helper()

	teleporterMessage := validTeleporterMessage
	teleporterMessage.RequiredGasLimit = big.NewInt(0)
	teleporterMessage.AllowedRelayerAddresses = nil
	teleporterMessage.Receipts = make([]teleportermessenger.TeleporterMessageReceipt, numReceipts)
	for i := range teleporterMessage.Receipts {
		teleporterMessage.Receipts[i] = teleportermessenger.TeleporterMessageReceipt{
			ReceivedMessageNonce: big.NewInt(int64(i + 1)),
			RelayerRewardAddress: validRelayerAddress,
		}
	}
	teleporterMessageBytes, err := teleporterMessage.Pack()
	require.NoError(t, err)

	addressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		teleporterMessageBytes,
	)
	require.NoError(t, err)

	unsignedMessage, err := warp.NewUnsignedMessage(0, ids.Empty, addressedCall.Bytes())
	require.NoError(t, err)

	minGasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(
		0,
		teleporterMessage.RequiredGasLimit,
		0,
		len(unsignedMessage.Payload),
		numReceipts,
	)
	require.NoError(t, err)

	return unsignedMessage, minGasLimit
}

// sendMessageTest is a handler for validTeleporterMessage, signed and ready to send, together with
// the mocks SendMessage talks to and what tests need to recognize its calls. The delivery
// decisions themselves are covered by the messages.Sender tests; these tests check that the
// handler hands the delivery the right transaction.
type sendMessageTest struct {
	handler           *messageHandler
	signedMessage     *warp.Message
	destinationClient *mock_vms.MockDestinationClient
	ethClient         *mock_evm.MockClient
	// gasLimit is the gas limit SendMessage computes for signedMessage.
	gasLimit uint64
	// callData is the receiveCrossChainMessage calldata SendMessage packs.
	callData []byte
	// messageReceivedInput is the calldata of the messageReceived query for the message.
	messageReceivedInput []byte
}

// newSendMessageTest builds the handler through the factory over a destination client whose only
// sender is validRelayerAddress, which the message allows.
func newSendMessageTest(t *testing.T, ctrl *gomock.Controller) *sendMessageTest {
	t.Helper()
	validMessageBytes, err := validTeleporterMessage.Pack()
	require.NoError(t, err)
	validAddressedCall, err := warpPayload.NewAddressedCall(
		messageProtocolAddress.Bytes(),
		validMessageBytes,
	)
	require.NoError(t, err)
	sourceBlockchainID := ids.Empty
	warpUnsignedMessage, err := warp.NewUnsignedMessage(0, sourceBlockchainID, validAddressedCall.Bytes())
	require.NoError(t, err)
	// The signature is verified on chain, not by the relayer, so an empty one will do.
	signedMessage, err := warp.NewMessage(warpUnsignedMessage, &warp.BitSetSignature{})
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
	callData, err := teleportermessenger.PackReceiveCrossChainMessage(
		0,
		common.HexToAddress(messageProtocolConfig.Settings["reward-address"].(string)),
	)
	require.NoError(t, err)

	destinationClient := mock_vms.NewMockDestinationClient(ctrl)
	ethClient := mock_evm.NewMockClient(ctrl)
	destinationClient.EXPECT().DestinationBlockchainID().Return(destinationBlockchainID).AnyTimes()
	destinationClient.EXPECT().BlockGasLimit().Return(uint64(15_000_000)).AnyTimes()
	destinationClient.EXPECT().SenderAddresses().Return([]common.Address{validRelayerAddress}).AnyTimes()
	destinationClient.EXPECT().Client().Return(ethClient).AnyTimes()

	factory, err := NewMessageHandlerFactory(messageProtocolAddress, messageProtocolConfig)
	require.NoError(t, err)
	handler, err := factory.NewMessageHandler(
		logging.NoLog{},
		sourceMessage(warpUnsignedMessage),
		destinationClient,
		nil,
		mocks.NewMockMetrics(ctrl),
		ids.Empty,
		0,
	)
	require.NoError(t, err)
	messageHandler, ok := handler.(*messageHandler)
	require.True(t, ok)

	numSigners, err := signedMessage.Signature.NumSigners()
	require.NoError(t, err)
	gasLimit, err := gasUtils.CalculateReceiveMessageGasLimit(
		numSigners,
		validTeleporterMessage.RequiredGasLimit,
		len(predicate.New(signedMessage.Bytes())),
		len(signedMessage.Payload),
		len(validTeleporterMessage.Receipts),
	)
	require.NoError(t, err)

	return &sendMessageTest{
		handler:              messageHandler,
		signedMessage:        signedMessage,
		destinationClient:    destinationClient,
		ethClient:            ethClient,
		gasLimit:             gasLimit,
		callData:             callData,
		messageReceivedInput: messageReceivedInput,
	}
}

// routeCalls answers the eth_call traffic SendMessage generates: the messageReceived query with
// [delivered], and the delivery simulation with [simulate], which also receives the simulated
// call for inspection.
func (s *sendMessageTest) routeCalls(
	t *testing.T,
	delivered bool,
	simulate func(call ethereum.CallMsg) ([]byte, error),
) {
	t.Helper()
	deliveredOutput, err := teleportermessenger.PackMessageReceivedOutput(delivered)
	require.NoError(t, err)
	s.ethClient.EXPECT().
		CallContract(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, call ethereum.CallMsg, _ *big.Int) ([]byte, error) {
			switch {
			case bytes.Equal(call.Data, s.messageReceivedInput):
				return deliveredOutput, nil
			case bytes.Equal(call.Data, s.callData):
				return simulate(call)
			default:
				t.Fatalf("unexpected eth_call to %v with data %x", call.To, call.Data)
				return nil, nil
			}
		}).
		AnyTimes()
}

// The handler must simulate, and then broadcast, the delivery of the signed message to the
// Teleporter contract: the relayer as sender, the computed gas limit, the packed
// receiveCrossChainMessage calldata and the signed message as the predicate access list.
func TestSendMessageSimulatesThenSends(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := newSendMessageTest(t, ctrl)

	var simulated ethereum.CallMsg
	s.routeCalls(t, false, func(call ethereum.CallMsg) ([]byte, error) {
		simulated = call
		return []byte{}, nil
	})
	expectedTxHash := common.HexToHash("0x01")
	s.destinationClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(
			_ logging.Logger,
			accessList types.AccessList,
			deliverers set.Set[common.Address],
			toAddress common.Address,
			gasLimit uint64,
			callData []byte,
		) (*types.Receipt, error) {
			require.Equal(t, utils.SignedWarpMessageToAccessList(s.signedMessage), accessList)
			require.Equal(t, set.Of(validRelayerAddress), deliverers)
			require.Equal(t, messageProtocolAddress, toAddress)
			require.Equal(t, s.gasLimit, gasLimit)
			require.Equal(t, s.callData, callData)
			return &types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: expectedTxHash}, nil
		}).
		Times(1)

	txHash, err := s.handler.SendMessage(context.Background(), s.signedMessage)
	require.NoError(t, err)
	require.Equal(t, expectedTxHash, txHash)

	require.Equal(t, validRelayerAddress, simulated.From)
	require.Equal(t, messageProtocolAddress, *simulated.To)
	require.Equal(t, s.gasLimit, simulated.Gas)
	require.Equal(t, utils.SignedWarpMessageToAccessList(s.signedMessage), simulated.AccessList)
}

// A delivery that reverts in simulation, such as one carrying a receipt the messenger cannot
// account for, is abandoned as non-retryable without being broadcast: no SendTx expectation is
// registered, so the mock fails the test if the handler sends it anyway.
func TestSendMessageSimulatedRevertIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	s := newSendMessageTest(t, ctrl)

	s.routeCalls(t, false, func(ethereum.CallMsg) ([]byte, error) {
		return nil, &evm.RevertError{Reason: "arithmetic underflow or overflow"}
	})

	_, err := s.handler.SendMessage(context.Background(), s.signedMessage)
	require.ErrorIs(t, err, messages.ErrDeliveryReverts)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
	require.ErrorContains(t, err, "arithmetic underflow or overflow")
}
