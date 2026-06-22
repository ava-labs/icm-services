// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package aggregator

import (
	"time"

	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/peers/clients"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
)

// oracleHandlerID is the p2p handler ID for oracle attestation requests.
// Mirrors validator.SignatureRequestHandlerID in
// github.com/ava-labs/avalanchego/network/p2p/oracle/validator.
const oracleHandlerID uint64 = 4

// OracleSignatureAggregator collects BLS signatures for oracle attestation
// messages. It uses the same aggregation logic as SignatureAggregator but
// routes requests to handler ID 4 (oracle attestation) instead of handler
// ID 2 (native warp).
//
// The counterpart on each validator is oracle/validator.OracleVerifier,
// registered at the same handler ID. Callers pass the Solana (or other
// source-chain) lookup hint as the justification argument to
// CreateSignedMessage; validators forward it to their sidecar.
type OracleSignatureAggregator struct {
	*SignatureAggregator
}

// NewOracleSignatureAggregator returns an OracleSignatureAggregator backed
// by the given network and validator state. The constructor signature mirrors
// NewSignatureAggregator exactly so callers can swap one for the other.
func NewOracleSignatureAggregator(
	network *peers.AppRequestNetwork,
	messageCreator message.Creator,
	signatureCacheSize uint64,
	aggregatorMetrics *metrics.SignatureAggregatorMetrics,
	validatorClient clients.CanonicalValidatorState,
	signatureRequestTimeout time.Duration,
) (*OracleSignatureAggregator, error) {
	sa, err := newSignatureAggregatorWithHandlerID(
		network,
		messageCreator,
		signatureCacheSize,
		aggregatorMetrics,
		validatorClient,
		signatureRequestTimeout,
		oracleHandlerID,
	)
	if err != nil {
		return nil, err
	}
	return &OracleSignatureAggregator{sa}, nil
}
