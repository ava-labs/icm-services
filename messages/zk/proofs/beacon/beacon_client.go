// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package beacon

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ava-labs/libevm/common"
)

// supportedForkVersion is the consensus fork this package's parsers assume.
// Blocks decode as Electra types and states as Fulu, both of which the
// beacon API reports under the Fulu fork name.
const supportedForkVersion = "fulu"

// maxBeaconObjectBytes bounds response bodies; mainnet states are ~200MB.
const maxBeaconObjectBytes = 1 << 30

// BeaconClient fetches SSZ encoded beacon objects from a beacon node's REST
// API. Beacon block sizes are roughly 100KB to 1MB, while beacon block states
// are roughly 100MB on mainnet.
type BeaconClient struct {
	baseURL string
	client  *http.Client
}

func NewBeaconClient(baseURL string) *BeaconClient {
	return &BeaconClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Minute}, // allow generous timeout for large beacon state fetches
	}
}

// BlockRoot returns the beacon block root for the given slot.
// For reference, see fetchBeaconBlockRoot in /scripts/tools/fixture-gen/generate_fixture.mts
// for how to fetch a block root from a beacon node.
func (c *BeaconClient) BlockRoot(ctx context.Context, slot uint64) (common.Hash, error) {
	url := fmt.Sprintf("%s/eth/v1/beacon/blocks/%d/root", c.baseURL, slot)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return common.Hash{}, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return common.Hash{}, fmt.Errorf("beacon block root fetch for slot %d failed: %w", slot, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return common.Hash{}, fmt.Errorf("beacon block root fetch for slot %d: status %d", slot, resp.StatusCode)
	}
	var result struct {
		Data struct {
			Root string `json:"root"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return common.Hash{}, fmt.Errorf("failed to decode beacon block root response: %w", err)
	}
	return common.HexToHash(result.Data.Root), nil
}

// Block returns the SSZ encoded SignedBeaconBlock at the given slot.
func (c *BeaconClient) Block(ctx context.Context, slot uint64) ([]byte, error) {
	return c.fetchSSZBytes(ctx, fmt.Sprintf("%s/eth/v2/beacon/blocks/%d", c.baseURL, slot))
}

// State returns the SSZ encoded BeaconState at the given slot.
func (c *BeaconClient) State(ctx context.Context, slot uint64) ([]byte, error) {
	return c.fetchSSZBytes(ctx, fmt.Sprintf("%s/eth/v2/debug/beacon/states/%d", c.baseURL, slot))
}

// fetchSSZBytes requests the raw SSZ encoding for the beacon object being fetched at url,
// rather than the API's default JSON encoding.
func (c *BeaconClient) fetchSSZBytes(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/octet-stream")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("SSZ fetch failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("SSZ fetch %s: status %d", url, resp.StatusCode)
	}

	// Reject unsupported consensus forks.
	if fork := resp.Header.Get("Eth-Consensus-Version"); !strings.EqualFold(fork, supportedForkVersion) {
		return nil, fmt.Errorf("unsupported consensus fork %q for %s: expected %s", fork, url, supportedForkVersion)
	}

	// Check if the response body is too large to read into memory.
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxBeaconObjectBytes+1))
	if err != nil {
		return nil, fmt.Errorf("SSZ fetch %s: %w", url, err)
	}
	if len(body) > maxBeaconObjectBytes {
		return nil, fmt.Errorf("SSZ fetch %s: response exceeds %d bytes", url, maxBeaconObjectBytes)
	}
	return body, nil
}
