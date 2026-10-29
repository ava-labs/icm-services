// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messages

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/icm-services/vms/evm"
	mock_evm "github.com/ava-labs/icm-services/vms/evm/mocks"
	mock_vms "github.com/ava-labs/icm-services/vms/mocks"
	ethereum "github.com/ava-labs/libevm"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

const (
	// testEstimatedGas is what the node estimates for the delivery; testGasLimit is that estimate
	// with the sender's safety margin, which is what the transaction carries.
	testEstimatedGas  = uint64(240_000)
	testGasLimit      = uint64(300_000)
	testBlockGasLimit = uint64(12_000_000)
)

var (
	testTo           = common.HexToAddress("0x253b2784c75e510dD0fF1da844684a1aC0aa5fcf")
	testRelayer      = common.HexToAddress("0x0123456789abcdef0123456789abcdef01234567")
	testOtherRelayer = common.HexToAddress("0x89abcdef0123456789abcdef0123456789abcdef")
	testCallData     = []byte{0xde, 0xad, 0xbe, 0xef, 0x01}
	testAccessList   = types.AccessList{
		{
			Address:     common.HexToAddress("0x0200000000000000000000000000000000000005"),
			StorageKeys: []common.Hash{{0x01}, {0x02}},
		},
	}
	testTxHash = common.HexToHash("0x01")
)

// rpcError mirrors the error libevm's RPC client returns for a JSON-RPC error object: the message
// plus whatever the node put in the "data" field, which is the revert output for a reverted call.
type rpcError struct {
	msg  string
	data any
}

func (e *rpcError) Error() string {
	return e.msg
}

func (e *rpcError) ErrorData() any {
	return e.data
}

// panicRevert is the error a node returns for a call that reverted with Panic(0x11), as checked
// arithmetic does on overflow.
func panicRevert(t *testing.T) error {
	t.Helper()
	uintType, err := abi.NewType("uint256", "", nil)
	require.NoError(t, err)
	packed, err := abi.Arguments{{Type: uintType}}.Pack(big.NewInt(0x11))
	require.NoError(t, err)
	data := append(crypto.Keccak256([]byte("Panic(uint256)"))[:4], packed...)
	return &rpcError{msg: "execution reverted", data: hexutil.Encode(data)}
}

// senderTest is a Sender over mocked clients. The destination client has two senders, of
// which only testRelayer is allowed to deliver the message.
type senderTest struct {
	sender            *Sender
	destinationClient *mock_vms.MockDestinationClient
	ethClient         *mock_evm.MockClient
	// delivered and deliveredErr are what the Delivered callback reports.
	delivered    bool
	deliveredErr error
}

func newSenderTest(t *testing.T, ctrl *gomock.Controller) *senderTest {
	t.Helper()
	destinationClient := mock_vms.NewMockDestinationClient(ctrl)
	ethClient := mock_evm.NewMockClient(ctrl)
	destinationClient.EXPECT().Client().Return(ethClient).AnyTimes()
	destinationClient.EXPECT().BlockGasLimit().Return(testBlockGasLimit).AnyTimes()
	destinationClient.EXPECT().
		SenderAddresses().
		Return([]common.Address{testOtherRelayer, testRelayer}).
		AnyTimes()

	st := &senderTest{destinationClient: destinationClient, ethClient: ethClient}
	st.sender = &Sender{
		Logger:            logging.NoLog{},
		DestinationClient: destinationClient,
		To:                testTo,
		AllowedRelayers:   []common.Address{testRelayer},
		Delivered: func() (bool, error) {
			return st.delivered, st.deliveredErr
		},
	}
	return st
}

// expectEstimation answers the one gas estimation with [estimated] and [err], and returns where
// the estimated call is recorded.
func (st *senderTest) expectEstimation(t *testing.T, estimated uint64, err error) *ethereum.CallMsg {
	t.Helper()
	estimatedCall := &ethereum.CallMsg{}
	st.ethClient.EXPECT().
		EstimateGas(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, call ethereum.CallMsg) (uint64, error) {
			*estimatedCall = call
			return estimated, err
		}).
		Times(1)
	return estimatedCall
}

// expectSendTx answers the one broadcast with the receipt [receipt] builds from the gas limit the
// sender chose.
func (st *senderTest) expectSendTx(t *testing.T, receipt func(gasLimit uint64) *types.Receipt) {
	t.Helper()
	st.destinationClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(
			_ logging.Logger,
			accessList types.AccessList,
			deliverers set.Set[common.Address],
			toAddress common.Address,
			gasLimit uint64,
			callData []byte,
		) (*types.Receipt, error) {
			require.Equal(t, testAccessList, accessList)
			require.Equal(t, set.Of(testRelayer), deliverers)
			require.Equal(t, testTo, toAddress)
			require.Equal(t, testGasLimit, gasLimit)
			require.Equal(t, testCallData, callData)
			return receipt(gasLimit), nil
		}).
		Times(1)
}

func successfulReceipt(uint64) *types.Receipt {
	return &types.Receipt{Status: types.ReceiptStatusSuccessful, TxHash: testTxHash}
}

func revertedReceipt(uint64) *types.Receipt {
	return &types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: 21_000, TxHash: testTxHash}
}

func send(st *senderTest) (common.Hash, error) {
	return st.sender.Send(context.Background(), testAccessList, testCallData)
}

// The estimation must execute the transaction that would be sent: same sender (an allowed one,
// not just the first), target, calldata and access list, bounded by the block gas limit, so that
// whatever it reports holds for the real delivery. The estimate comes back with the safety margin.
func TestSenderEstimateGasMirrorsTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	estimated := st.expectEstimation(t, testEstimatedGas, nil)

	gasLimit, err := st.sender.EstimateGas(context.Background(), testAccessList, testCallData)
	require.NoError(t, err)
	require.Equal(t, testGasLimit, gasLimit)
	require.Equal(t, testRelayer, estimated.From)
	require.Equal(t, testTo, *estimated.To)
	require.Equal(t, testBlockGasLimit, estimated.Gas)
	require.Equal(t, testCallData, estimated.Data)
	require.Equal(t, testAccessList, estimated.AccessList)
}

// The margin never pushes the gas limit past the block gas limit.
func TestSenderEstimateGasIsCappedAtBlockGasLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.expectEstimation(t, testBlockGasLimit-1, nil)

	gasLimit, err := st.sender.EstimateGas(context.Background(), testAccessList, testCallData)
	require.NoError(t, err)
	require.Equal(t, testBlockGasLimit, gasLimit)
}

// A delivery that reverts when the node executes it for the estimate, such as one carrying a
// receipt the messenger cannot account for, is abandoned as non-retryable without being
// broadcast: no SendTx expectation is registered, so the mock fails the test if it is sent
// anyway. The node reports the revert as an RPC error carrying the revert data, which is decoded.
func TestSenderSendAbandonsEstimationRevert(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.expectEstimation(t, 0, panicRevert(t))

	_, err := send(st)
	require.ErrorIs(t, err, ErrDeliveryReverts)
	require.ErrorIs(t, err, ErrNonRetryable)
	require.ErrorContains(t, err, "arithmetic underflow or overflow")
}

// A delivery that succeeds with no gas the node may try, bounded by the block gas limit, cannot
// be mined. The node reports that instead of an estimate, and the message is abandoned without
// anything being broadcast.
func TestSenderSendExceedingGasAllowanceIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.expectEstimation(t, 0, errors.New("gas required exceeds allowance (12000000)"))

	_, err := send(st)
	require.ErrorIs(t, err, ErrDeliveryExceedsBlockGasLimit)
	require.ErrorIs(t, err, ErrNonRetryable)
}

// A revert because another relayer delivered the message first is not a failure: there is
// nothing left to send.
func TestSenderSendEstimationRevertAlreadyDelivered(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.delivered = true
	st.expectEstimation(t, 0, &evm.RevertError{Reason: "TeleporterMessenger: message already received"})

	txHash, err := send(st)
	require.NoError(t, err)
	require.Equal(t, common.Hash{}, txHash)
}

// A revert is only abandoned once it is known that nobody else delivered the message.
func TestSenderSendEstimationRevertUnknownDeliveryIsRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.deliveredErr = errors.New("connection refused")
	st.expectEstimation(t, 0, &evm.RevertError{Reason: "arithmetic underflow or overflow"})

	_, err := send(st)
	require.Error(t, err)
	require.NotErrorIs(t, err, ErrNonRetryable)
}

// If the gas cannot be estimated at all, nothing is known about the delivery: it is neither
// abandoned nor broadcast blind, and the failure stays retryable.
func TestSenderSendEstimationFailureIsRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.expectEstimation(t, 0, errors.New("connection refused"))

	_, err := send(st)
	require.Error(t, err)
	require.NotErrorIs(t, err, ErrNonRetryable)
}

// A delivery that estimates successfully is broadcast with the buffered estimate.
func TestSenderSendBroadcastsBufferedEstimate(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.expectEstimation(t, testEstimatedGas, nil)
	st.expectSendTx(t, successfulReceipt)

	txHash, err := send(st)
	require.NoError(t, err)
	require.Equal(t, testTxHash, txHash)
}

// A broadcast that fails before a receipt is obtained says nothing about the message.
func TestSenderSendBroadcastFailureIsRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.expectEstimation(t, testEstimatedGas, nil)
	st.destinationClient.EXPECT().
		SendTx(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, errors.New("timed out waiting for transaction result")).
		Times(1)

	_, err := send(st)
	require.Error(t, err)
	require.NotErrorIs(t, err, ErrNonRetryable)
}

// A mined delivery that reverted having consumed its entire gas limit ran out of gas. A retry
// would fail identically and only pay for a second reverted transaction, so the message must be
// abandoned.
func TestSenderSendOutOfGasReceiptIsNonRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.expectEstimation(t, testEstimatedGas, nil)
	// Echo the gas limit back as gas used: that is what an out-of-gas revert looks like on chain.
	st.expectSendTx(t, func(gasLimit uint64) *types.Receipt {
		return &types.Receipt{Status: types.ReceiptStatusFailed, GasUsed: gasLimit}
	})

	_, err := send(st)
	require.ErrorIs(t, err, ErrDeliveryOutOfGas)
	require.ErrorIs(t, err, ErrNonRetryable)
}

// A mined revert that left gas unspent after a successful estimate is a verification failure,
// which a fresh signature may resolve, so it stays retryable.
func TestSenderSendRevertWithGasLeftIsRetryable(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.expectEstimation(t, testEstimatedGas, nil)
	st.expectSendTx(t, revertedReceipt)

	_, err := send(st)
	require.Error(t, err)
	require.NotErrorIs(t, err, ErrNonRetryable)
}

// A mined revert because another relayer delivered the message first counts as delivered.
func TestSenderSendRevertedReceiptAlreadyDelivered(t *testing.T) {
	ctrl := gomock.NewController(t)
	st := newSenderTest(t, ctrl)
	st.delivered = true
	st.expectEstimation(t, testEstimatedGas, nil)
	st.expectSendTx(t, revertedReceipt)

	txHash, err := send(st)
	require.NoError(t, err)
	require.Equal(t, testTxHash, txHash)
}

func TestSenderAddress(t *testing.T) {
	tests := []struct {
		name            string
		senders         []common.Address
		allowedRelayers []common.Address
		expected        common.Address
	}{
		{
			name:            "allowed sender is preferred over the first",
			senders:         []common.Address{testOtherRelayer, testRelayer},
			allowedRelayers: []common.Address{testRelayer},
			expected:        testRelayer,
		},
		{
			name:     "any sender when the message allows all relayers",
			senders:  []common.Address{testOtherRelayer, testRelayer},
			expected: testOtherRelayer,
		},
		{
			name:            "first sender when none is allowed",
			senders:         []common.Address{testOtherRelayer},
			allowedRelayers: []common.Address{testRelayer},
			expected:        testOtherRelayer,
		},
		{
			name:     "no senders",
			expected: common.Address{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			destinationClient := mock_vms.NewMockDestinationClient(ctrl)
			destinationClient.EXPECT().SenderAddresses().Return(tt.senders).AnyTimes()
			sender := &Sender{
				DestinationClient: destinationClient,
				AllowedRelayers:   tt.allowedRelayers,
			}
			require.Equal(t, tt.expected, sender.SenderAddress())
		})
	}
}
