// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package zk

import (
	"maps"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

// The numeric settings arrive as JSON numbers, so they are float64 and must be whole and in
// range before they are converted. A slot duration that truncated to zero would become the
// divisor of every slot computation.
func TestConfigFromMapNumericSettings(t *testing.T) {
	testCases := []struct {
		name          string
		key           string
		value         any
		expectedError string
	}{
		{name: "valid seconds per slot", key: "seconds-per-slot", value: float64(12)},
		{name: "valid genesis time", key: "beacon-genesis-time", value: float64(1606824023)},
		{name: "zero genesis time", key: "beacon-genesis-time", value: float64(0)},
		{
			name:          "zero seconds per slot",
			key:           "seconds-per-slot",
			value:         float64(0),
			expectedError: "seconds-per-slot must be greater than 0",
		},
		{
			name:          "fractional seconds per slot below one",
			key:           "seconds-per-slot",
			value:         0.5,
			expectedError: "seconds-per-slot must be a non-negative integer",
		},
		{
			name:          "fractional seconds per slot above one",
			key:           "seconds-per-slot",
			value:         12.5,
			expectedError: "seconds-per-slot must be a non-negative integer",
		},
		{
			name:          "negative seconds per slot",
			key:           "seconds-per-slot",
			value:         float64(-12),
			expectedError: "seconds-per-slot must be a non-negative integer",
		},
		{
			name:          "seconds per slot is not a number",
			key:           "seconds-per-slot",
			value:         "12",
			expectedError: "seconds-per-slot not found",
		},
		{
			name:          "negative genesis time",
			key:           "beacon-genesis-time",
			value:         float64(-1),
			expectedError: "beacon-genesis-time must be a non-negative integer",
		},
		{
			name:          "fractional genesis time",
			key:           "beacon-genesis-time",
			value:         1606824023.5,
			expectedError: "beacon-genesis-time must be a non-negative integer",
		},
		{
			name:          "genesis time out of range",
			key:           "beacon-genesis-time",
			value:         float64(math.MaxUint64),
			expectedError: "beacon-genesis-time must be a non-negative integer",
		},
		{
			name:          "genesis time is not a number",
			key:           "beacon-genesis-time",
			value:         math.NaN(),
			expectedError: "beacon-genesis-time must be a non-negative integer",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			settings := maps.Clone(messageProtocolConfig.Settings)
			settings[test.key] = test.value

			cfg, err := ConfigFromMap(settings)
			if test.expectedError != "" {
				require.ErrorContains(t, err, test.expectedError)
				return
			}
			require.NoError(t, err)
			require.NotZero(t, cfg.SecondsPerSlot)
		})
	}
}
