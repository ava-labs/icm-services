package utils

import (
	"errors"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/stretchr/testify/require"
)

func TestWithMaxRetries(t *testing.T) {
	t.Run("NotEnoughRetry", func(t *testing.T) {
		retryable := newMockRetryableFn(3)
		err := WithMaxRetries(
			func() (err error) {
				_, err = retryable.Run()
				return err
			},
			// using default values: we want to run max 2 tries.
			624*time.Millisecond,
			logging.NoLog{},
		)
		require.Error(t, err)
	})
	t.Run("EnoughRetry", func(t *testing.T) {
		retryable := newMockRetryableFn(2)
		var res bool
		err := WithMaxRetries(
			func() (err error) {
				res, err = retryable.Run()
				return err
			},
			// using default values we want to run 3 tries.
			2000*time.Millisecond,
			logging.NoLog{},
		)
		require.NoError(t, err)
		require.True(t, res)
	})
}

type mockRetryableFn struct {
	counter uint64
	trigger uint64
}

func newMockRetryableFn(trigger uint64) mockRetryableFn {
	return mockRetryableFn{
		counter: 0,
		trigger: trigger,
	}
}

func (m *mockRetryableFn) Run() (bool, error) {
	if m.counter >= m.trigger {
		return true, nil
	}
	m.counter++
	return false, errors.New("error")
}
