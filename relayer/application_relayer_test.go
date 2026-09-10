// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package relayer

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	mock_database "github.com/ava-labs/icm-services/database/mocks"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/messages/mocks"
	"github.com/ava-labs/icm-services/relayer/checkpoint"
	"github.com/ava-labs/libevm/common"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// A failure that will recur identically must be attempted once, not maxRetryCount times. Each
// retry of an under-provisioned delivery mines another reverted transaction and bills the
// relayer's hot wallet for the full gas limit, so retrying a deterministic failure multiplies the
// loss by maxRetryCount for no chance of success.
func TestProcessMessageDoesNotRetryNonRetryableError(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler := mocks.NewMockMessageHandler(ctrl)

	wrapped := fmt.Errorf("%w: delivery reverted out of gas", messages.ErrNonRetryable)
	handler.EXPECT().
		ProcessMessage().
		Return(common.Hash{}, wrapped).
		Times(1)

	relayer := &ApplicationRelayer{logger: logging.NoLog{}}
	_, err := relayer.ProcessMessage(handler)
	require.ErrorIs(t, err, messages.ErrNonRetryable)
}

// Failures that a fresh attempt may resolve, such as Warp verification against a validator set
// that has since churned, must still exhaust the retry budget.
func TestProcessMessageRetriesRetryableError(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler := mocks.NewMockMessageHandler(ctrl)

	handler.EXPECT().
		ProcessMessage().
		Return(common.Hash{}, errors.New("transaction failed with status: 0")).
		Times(maxRetryCount)

	relayer := &ApplicationRelayer{logger: logging.NoLog{}}
	_, err := relayer.ProcessMessage(handler)
	require.Error(t, err)
	require.NotErrorIs(t, err, messages.ErrNonRetryable)
}

// newTestCheckpointManager builds a real checkpoint manager over a mocked database so that
// ProcessHeight's StageCommittedHeight call exercises the real code path.
func newTestCheckpointManager(t *testing.T, ctrl *gomock.Controller) *checkpoint.CheckpointManager {
	t.Helper()
	db := mock_database.NewMockRelayerDatabase(ctrl)
	db.EXPECT().Get(gomock.Any(), gomock.Any()).Return([]byte("0"), nil).AnyTimes()
	db.EXPECT().Put(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	id := database.RelayerID{ID: common.BytesToHash([]byte("test-relayer"))}
	cm, err := checkpoint.NewCheckpointManager(
		logging.NoLog{},
		checkpoint.NewCheckpointManagerMetrics(prometheus.NewRegistry()),
		db,
		make(chan struct{}, 1),
		id,
		0,
	)
	require.NoError(t, err)
	return cm
}

func newTestApplicationRelayer(
	t *testing.T,
	ctrl *gomock.Controller,
	metrics messages.Metrics,
) *ApplicationRelayer {
	t.Helper()
	return &ApplicationRelayer{
		logger:                  logging.NoLog{},
		metrics:                 metrics,
		checkpointManager:       newTestCheckpointManager(t, ctrl),
		processMessageSemaphore: make(chan struct{}, 1),
	}
}

// A message that can never be delivered must not take the relayer down with it. Sending the error
// on to errChan makes the listener exit, and because the height is then never checkpointed every
// restart replays the same message - a crash loop that only an operator can break.
func TestProcessHeightSkipsNonRetryableMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler := mocks.NewMockMessageHandler(ctrl)
	handler.EXPECT().
		ProcessMessage().
		Return(common.Hash{}, fmt.Errorf("%w: used 5 of 5", messages.ErrDeliveryOutOfGas)).
		Times(1)

	// The message is gone for good, so it must be counted as abandoned rather than merely
	// failed, under a bounded reason label rather than the full error text.
	metrics := mocks.NewMockMetrics(ctrl)
	metrics.EXPECT().IncAbandonedRelayMessageCount("delivery out of gas").Times(1)

	relayer := newTestApplicationRelayer(t, ctrl, metrics)
	errChan := make(chan error, 1)
	relayer.ProcessHeight(1, []messages.MessageHandler{handler}, errChan)

	select {
	case err := <-errChan:
		t.Fatalf("non-retryable message should be skipped, but errored the listener: %v", err)
	default:
	}
}

// Failures that are not marked non-retryable keep the existing behaviour: they surface to the
// listener rather than being silently swallowed.
func TestProcessHeightSurfacesRetryableFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	handler := mocks.NewMockMessageHandler(ctrl)
	handler.EXPECT().
		ProcessMessage().
		Return(common.Hash{}, errors.New("some transient failure")).
		Times(maxRetryCount)

	relayer := newTestApplicationRelayer(t, ctrl, mocks.NewMockMetrics(ctrl))
	errChan := make(chan error, 1)
	relayer.ProcessHeight(1, []messages.MessageHandler{handler}, errChan)

	select {
	case err := <-errChan:
		require.Error(t, err)
	default:
		t.Fatal("retryable failure should have been surfaced to the listener")
	}
}
