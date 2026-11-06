// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/database"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/metrics"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethclient"
	. "github.com/onsi/gomega"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/common/model"
	"go.uber.org/zap"
)

// abandonedRelayMessageMetric is the relayer's abandoned-message counter as served on its metrics
// endpoint: the application relayer registry's "app" prefix followed by the counter's name.
const abandonedRelayMessageMetric = "app_abandoned_relay_message_count"

// AbandonedRelayMessageCount scrapes the relayer's metrics endpoint and returns the number of
// messages it has abandoned for [failureReason], summed over every route it serves.
func AbandonedRelayMessageCount(
	ctx context.Context,
	metricsPort uint16,
	failureReason string,
) (float64, error) {
	url := fmt.Sprintf("http://localhost:%d%s", metricsPort, metrics.MetricsPath)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("metrics endpoint returned status %d", resp.StatusCode)
	}

	parser := expfmt.NewTextParser(model.UTF8Validation)
	families, err := parser.TextToMetricFamilies(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to parse relayer metrics: %w", err)
	}
	family, ok := families[abandonedRelayMessageMetric]
	if !ok {
		// The counter is only exported once a message has been abandoned.
		return 0, nil
	}

	var count float64
	for _, metric := range family.GetMetric() {
		for _, label := range metric.GetLabel() {
			if label.GetName() == "failure_reason" && label.GetValue() == failureReason {
				count += metric.GetCounter().GetValue()
			}
		}
	}
	return count, nil
}

// RevertingDeliveryCount returns how many messages the relayer has abandoned because their
// delivery reverts in simulation.
func RevertingDeliveryCount(ctx context.Context, metricsPort uint16) (float64, error) {
	return AbandonedRelayMessageCount(
		ctx, metricsPort, messages.NonRetryableReason(messages.ErrDeliveryReverts),
	)
}

// RelayerCheckpoint returns the block height the relayer has committed for [relayerID]. It reads
// the relayer's JSON database from disk, so it can be polled while the relayer is running.
func RelayerCheckpoint(db database.RelayerDatabase, relayerID common.Hash) (uint64, error) {
	data, err := db.Get(relayerID, database.LatestProcessedBlockKey)
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(string(data), 10, 64)
}

// ExpectRelayerHealthy asserts that the relayer's health endpoint answers, which it only does
// while the process is running rather than having exited on an error.
func ExpectRelayerHealthy(ctx context.Context, apiPort uint16) {
	url := fmt.Sprintf("http://localhost:%d/health", apiPort)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	Expect(err).Should(BeNil())
	resp, err := http.DefaultClient.Do(req)
	Expect(err).Should(BeNil(), "relayer health endpoint unreachable")
	defer resp.Body.Close()
	Expect(resp.StatusCode).Should(Equal(http.StatusOK))
}

// AbandonedDelivery is a message whose delivery reverts for a reason the message itself causes.
// The relayer must give such a message up without broadcasting a transaction, instead of retrying
// it and exiting.
type AbandonedDelivery struct {
	// RelayerConfig is the running relayer's config. It names the metrics and API ports and the
	// checkpoint database.
	RelayerConfig relayercfg.Config
	// RelayerID is the checkpoint key of the route the message travels.
	RelayerID common.Hash
	// SourceHeight is the source-chain block that holds the message.
	SourceHeight uint64
	// Destination is the chain the relayer would have delivered the message to.
	Destination *ethclient.Client
	// RelayerAddress is the relayer's sender address on the destination.
	RelayerAddress common.Address
	// NonceBefore is the relayer's nonce on the destination before the message was sent.
	NonceBefore uint64
	// AbandonedBefore is the relayer's reverting-delivery count before the message was sent.
	AbandonedBefore float64
}

// WaitForAbandonedDelivery asserts that the relayer abandons the message as a reverting delivery,
// mines nothing for it, checkpoints past its block and keeps running.
func WaitForAbandonedDelivery(ctx context.Context, log logging.Logger, delivery AbandonedDelivery) {
	log.Info("Waiting for the relayer to abandon the reverting delivery",
		zap.Uint64("sourceHeight", delivery.SourceHeight),
	)
	// Signature aggregation and, on the Merkle path, validator-set and commitment reads happen
	// before the delivery is simulated, so allow as long as an ordinary delivery is allowed.
	Eventually(func() (float64, error) {
		return RevertingDeliveryCount(ctx, delivery.RelayerConfig.MetricsPort)
	}, 120*time.Second, time.Second).Should(BeNumerically(">", delivery.AbandonedBefore))

	// Without the simulation gate, each of the five retries mined a reverted transaction. None
	// may be broadcast now, so the relayer's nonce on the destination must not have moved.
	nonce, err := delivery.Destination.NonceAt(ctx, delivery.RelayerAddress, nil)
	Expect(err).Should(BeNil())
	Expect(nonce).Should(Equal(delivery.NonceBefore),
		"the relayer broadcast a transaction for a delivery that reverts in simulation")

	// Abandoning the message lets the height commit, so a restart never replays the block.
	jsonDB, err := database.NewJSONFileStorage(
		logging.NoLog{},
		delivery.RelayerConfig.StorageLocation,
		database.GetConfigRelayerIDs(&delivery.RelayerConfig),
	)
	Expect(err).Should(BeNil())
	Eventually(func() (uint64, error) {
		return RelayerCheckpoint(jsonDB, delivery.RelayerID)
	}, 30*time.Second, time.Second).Should(BeNumerically(">=", delivery.SourceHeight),
		"the relayer did not checkpoint past the abandoned message")

	ExpectRelayerHealthy(ctx, delivery.RelayerConfig.APIPort)
}
