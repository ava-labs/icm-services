// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package teleporterv2

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	mock_evm "github.com/ava-labs/icm-services/vms/evm/mocks"
	mock_vms "github.com/ava-labs/icm-services/vms/mocks"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const testBlockGasLimit = 12_000_000

var (
	testTeleporterAddress = common.HexToAddress("0xd81545385803bCD83bd59f58Ba2d2c0562387F83")
	testRelayerAddress    = common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")
)

func newTestHandler(
	t *testing.T,
	destinationClient *mock_vms.MockDestinationClient,
	requiredGasLimit *big.Int,
) *messageHandler {
	t.Helper()

	unsignedMessage, err := warp.NewUnsignedMessage(0, ids.Empty, []byte{1, 2, 3, 4})
	require.NoError(t, err)

	return &messageHandler{
		logger: logging.NoLog{},
		teleporterMessage: &teleportermessengerv2.TeleporterMessageV2{
			MessageNonce:     big.NewInt(1),
			RequiredGasLimit: requiredGasLimit,
			Message:          []byte{1, 2, 3, 4},
		},
		unsignedMessage:   unsignedMessage,
		destinationClient: destinationClient,
		teleporterAddress: testTeleporterAddress,
	}
}

// TestShouldSendMessageGasLimitHeadroom checks that the RequiredGasLimit pre-check leaves room for
// the delivery overhead, so that a message at or just below the block gas limit is rejected before
// signature aggregation rather than delivered into a guaranteed revert.
func TestShouldSendMessageGasLimitHeadroom(t *testing.T) {
	testCases := []struct {
		name             string
		requiredGasLimit *big.Int
		// A message that clears the gas check goes on to query delivery status.
		expectMessengerCall bool
		expectedResult      bool
	}{
		{
			name:                "leaves room for the delivery overhead",
			requiredGasLimit:    big.NewInt(testBlockGasLimit - minDeliveryOverheadGas),
			expectMessengerCall: true,
			expectedResult:      true,
		},
		{
			name:             "equal to the block gas limit",
			requiredGasLimit: big.NewInt(testBlockGasLimit),
			expectedResult:   false,
		},
		{
			name:             "within the delivery overhead of the block gas limit",
			requiredGasLimit: big.NewInt(testBlockGasLimit - minDeliveryOverheadGas + 1),
			expectedResult:   false,
		},
		{
			name:             "exceeds the block gas limit",
			requiredGasLimit: big.NewInt(testBlockGasLimit + 1),
			expectedResult:   false,
		},
		{
			name:             "does not fit in uint64",
			requiredGasLimit: new(big.Int).Lsh(big.NewInt(1), 64),
			expectedResult:   false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			destinationClient := mock_vms.NewMockDestinationClient(ctrl)
			destinationClient.EXPECT().BlockGasLimit().Return(uint64(testBlockGasLimit)).AnyTimes()

			if test.expectMessengerCall {
				notDelivered, err := teleportermessengerv2.PackMessageReceivedOutput(false)
				require.NoError(t, err)

				ethClient := mock_evm.NewMockClient(ctrl)
				ethClient.EXPECT().
					CallContract(gomock.Any(), gomock.Any(), gomock.Any()).
					Return(notDelivered, nil)
				destinationClient.EXPECT().Client().Return(ethClient)
				destinationClient.EXPECT().
					SenderAddresses().
					Return([]common.Address{testRelayerAddress})
			}

			handler := newTestHandler(t, destinationClient, test.requiredGasLimit)
			shouldSend, err := handler.ShouldSendMessage()
			require.NoError(t, err)
			require.Equal(t, test.expectedResult, shouldSend)
		})
	}
}

// TestEstimateGasLimit checks that estimation never resolves to a gas limit it has not shown to be
// sufficient, since it is the only simulation before the delivery is signed and broadcast.
func TestEstimateGasLimit(t *testing.T) {
	testCases := []struct {
		name                  string
		estimated             uint64
		estimateErr           error
		expectedGasLimit      uint64
		expectedError         bool
		expectedUndeliverable bool
	}{
		{
			name:             "buffer applied below the block gas limit",
			estimated:        1_000_000,
			expectedGasLimit: 1_250_000,
		},
		{
			name:             "buffer capped at the block gas limit",
			estimated:        11_000_000,
			expectedGasLimit: testBlockGasLimit,
		},
		{
			name:                  "estimate exceeds the block gas limit",
			estimated:             testBlockGasLimit + 1,
			expectedError:         true,
			expectedUndeliverable: true,
		},
		{
			name:                  "estimation exhausts the search range",
			estimateErr:           errors.New("gas required exceeds allowance (12000000)"),
			expectedError:         true,
			expectedUndeliverable: true,
		},
		{
			name:                  "delivery reverts for insufficient gas",
			estimateErr:           errors.New("execution reverted: TeleporterMessenger: insufficient gas"),
			expectedError:         true,
			expectedUndeliverable: true,
		},
		{
			// Retried rather than skipped: verification can fail transiently if the registry's
			// committed validator set changes mid-delivery.
			name:          "delivery reverts for failed verification",
			estimateErr:   errors.New("execution reverted: TeleporterMessenger: message verification failed"),
			expectedError: true,
		},
		{
			name:          "transient rpc failure",
			estimateErr:   errors.New("connection refused"),
			expectedError: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			ethClient := mock_evm.NewMockClient(ctrl)
			ethClient.EXPECT().
				EstimateGas(gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, call ethereum.CallMsg) (uint64, error) {
					// The search must be bounded by the gas the delivery can carry.
					require.Equal(t, uint64(testBlockGasLimit), call.Gas)
					return test.estimated, test.estimateErr
				})

			destinationClient := mock_vms.NewMockDestinationClient(ctrl)
			destinationClient.EXPECT().Client().Return(ethClient)
			destinationClient.EXPECT().BlockGasLimit().Return(uint64(testBlockGasLimit)).AnyTimes()
			destinationClient.EXPECT().
				SenderAddresses().
				Return([]common.Address{testRelayerAddress}).
				AnyTimes()

			handler := newTestHandler(t, destinationClient, big.NewInt(1_000_000))
			gasLimit, err := handler.estimateGasLimit(context.Background(), []byte{1, 2, 3, 4})
			if !test.expectedError {
				require.NoError(t, err)
				require.Equal(t, test.expectedGasLimit, gasLimit)
				return
			}
			require.Error(t, err)
			require.Zero(t, gasLimit, "a failed estimation must not resolve to a gas limit")
			require.Equal(t, test.expectedUndeliverable, errors.Is(err, errUndeliverable))
		})
	}
}
