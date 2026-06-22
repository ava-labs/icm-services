// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"net/http"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
)

const OracleAPIPath = "/oracle/aggregate-signatures"

// HandleOracleAggregateSignatures registers the oracle attestation endpoint.
// It accepts the same request/response format as HandleAggregateSignaturesByRawMsgRequest
// but routes requests to handler ID 4, so validators dispatch to their
// OracleVerifier (which calls the sidecar) instead of the native warp handler.
//
// Request body: AggregateSignatureRequest with justification = hex-encoded
// Solana transaction signature or other source-chain lookup hint.
// Response body: AggregateSignatureResponse with the signed warp message.
func HandleOracleAggregateSignatures(
	logger logging.Logger,
	metricsInstance *metrics.SignatureAggregatorMetrics,
	oracleAggregator *aggregator.OracleSignatureAggregator,
) {
	http.Handle(
		OracleAPIPath,
		signatureAggregationAPIHandler(
			logger,
			metricsInstance,
			oracleAggregator.SignatureAggregator,
		),
	)
}
