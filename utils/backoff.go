package utils

import (
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
)

// WithRetriesTimeout uses an exponential backoff to run the operation until it
// succeeds or timeout limit has been reached. It is the caller's responsibility
// to ensure {operation} returns. It is safe for {operation} to take longer than {timeout}.
func WithRetriesTimeout(
	logger logging.Logger,
	operation backoff.Operation,
	timeout time.Duration,
	logMessage string,
) error {
	expBackOff := backoff.NewExponentialBackOff(
		backoff.WithMaxElapsedTime(timeout),
	)
	notify := func(err error, duration time.Duration) {
		logger.Warn("operation failed, retrying...",
			zap.String("logMessage", logMessage),
			zap.Duration("retryIn", duration),
			zap.Error(err),
		)
	}
	return backoff.RetryNotify(operation, expBackOff, notify)
}
