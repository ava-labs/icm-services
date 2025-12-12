// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"fmt"
	"net/url"

	"github.com/ava-labs/avalanchego/utils/rpc"
)

// API configuration containing the base URL and query parameters
type APIConfig struct {
	BaseURL     string            `mapstructure:"base-url" json:"base-url"`
	QueryParams map[string]string `mapstructure:"query-parameters" json:"query-parameters"`
	HTTPHeaders map[string]string `mapstructure:"http-headers" json:"http-headers"`
}

func (c *APIConfig) Validate() error {
	if _, err := url.ParseRequestURI(c.BaseURL); err != nil {
		return fmt.Errorf("invalid base URL: %w", err)
	}
	return nil
}

// Options initializes and returns the rpc options for an API
func (c *APIConfig) Options() []rpc.Option {
	options := make([]rpc.Option, 0, len(c.QueryParams)+len(c.HTTPHeaders))
	for key, value := range c.QueryParams {
		options = append(options, rpc.WithQueryParam(key, value))
	}
	for key, value := range c.HTTPHeaders {
		options = append(options, rpc.WithHeader(key, value))
	}
	return options
}
