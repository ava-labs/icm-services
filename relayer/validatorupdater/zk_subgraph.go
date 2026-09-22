// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validatorupdater

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ava-labs/libevm/common"
)

// ConsensusProof is one Boundless consensus transition proof indexed by the
// subgraph. JournalData and Seal are decoded from the subgraph's hex strings.
type ConsensusProof struct {
	FinalizedSlot uint64
	JournalData   []byte
	Seal          []byte
}

// SubgraphClient queries the Boundless subgraph for consensus transition proofs.
type SubgraphClient struct {
	url    string
	client *http.Client
}

func NewSubgraphClient(url string) *SubgraphClient {
	return &SubgraphClient{url: url, client: &http.Client{Timeout: 30 * time.Second}}
}

// ConsensusProofsAfterSlot returns up to limit proofs with
// finalizedSlot > afterSlot, in ascending order. Ascending order matters
// because transitions are applied sequentially by the caller.
// Each journal's preState must match the contract's current state.
func (c *SubgraphClient) ConsensusProofsAfterSlot(
	ctx context.Context,
	afterSlot uint64,
	limit int,
) ([]*ConsensusProof, error) {
	query := fmt.Sprintf(`{
		signalEthereumProofs(
			first: %d, orderBy: finalizedSlot, orderDirection: asc,
			where: { finalizedSlot_gt: "%d" }
		) { finalizedSlot journalData seal }
	}`, limit, afterSlot)

	body, err := json.Marshal(map[string]string{"query": query})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("subgraph request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("subgraph returned status %d", resp.StatusCode)
	}

	// GraphQL returns HTTP 200 even for failed queries, reporting failures in
	// the top-level errors field. Without decoding it, a bad query looks like
	// an empty result.
	var result struct {
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors"`
		Data struct {
			SignalEthereumProofs []struct {
				FinalizedSlot uint64 `json:"finalizedSlot,string"`
				JournalData   string `json:"journalData"`
				Seal          string `json:"seal"`
			} `json:"signalEthereumProofs"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode subgraph response: %w", err)
	}
	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("subgraph query error: %s", result.Errors[0].Message)
	}

	proofs := make([]*ConsensusProof, 0, len(result.Data.SignalEthereumProofs))
	for _, p := range result.Data.SignalEthereumProofs {
		proofs = append(proofs, &ConsensusProof{
			FinalizedSlot: p.FinalizedSlot,
			JournalData:   common.FromHex(p.JournalData),
			Seal:          common.FromHex(p.Seal),
		})
	}
	return proofs, nil
}
