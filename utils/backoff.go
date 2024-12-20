package utils

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/cenkalti/backoff/v4"
)

// WithMaxRetries uses an exponential backoff to run the operation until it
// succeeds or max elapsed time has been reached.
func WithMaxRetries(
	operation backoff.Operation,
	maxElapsedTime time.Duration,
	logger logging.Logger,
) error {
	expBackOff := backoff.NewExponentialBackOff(
		backoff.WithMaxElapsedTime(maxElapsedTime),
	)
	notify := func(err error, duration time.Duration) {
		logger.Warn("operation failed, retrying...")
	}
	err := backoff.RetryNotify(operation, expBackOff, notify)
	return err
}
