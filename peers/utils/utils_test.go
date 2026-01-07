// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"testing"

	"github.com/ava-labs/icm-services/config"
	"github.com/stretchr/testify/require"
)

func TestInitializeOptionsLength(t *testing.T) {
	tests := []struct {
		name           string
		apiConfig      *config.APIConfig
		expectedLength int
	}{
		{
			name: "two query params and two headers",
			apiConfig: &config.APIConfig{
				BaseURL: "http://localhost:9650",
				QueryParams: map[string]string{
					"param1": "value1",
					"param2": "value2",
				},
				HTTPHeaders: map[string]string{
					"header1": "value1",
					"header2": "value2",
				},
			},
			expectedLength: 4,
		},
		{
			name: "one query param",
			apiConfig: &config.APIConfig{
				BaseURL: "http://localhost:9650",
				QueryParams: map[string]string{
					"token": "value",
				},
			},
			expectedLength: 1,
		},
		{
			name: "empty config",
			apiConfig: &config.APIConfig{
				BaseURL: "http://localhost:9650",
			},
			expectedLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := tt.apiConfig.Options()
			require.Len(t, options, tt.expectedLength)
		})
	}
}

func TestInitializeOptionsCreatesCorrectTypes(t *testing.T) {
	apiConfig := &config.APIConfig{
		BaseURL: "http://localhost:9650",
		QueryParams: map[string]string{
			"token": "test-token",
		},
		HTTPHeaders: map[string]string{
			"Authorization": "Bearer xyz",
		},
	}

	options := apiConfig.Options()

	expectedCount := len(apiConfig.QueryParams) + len(apiConfig.HTTPHeaders)
	require.Len(t, options, expectedCount)

	for i, opt := range options {
		require.NotNil(t, opt, "option at index %d should not be nil", i)
	}
}
