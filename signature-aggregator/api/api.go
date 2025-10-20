// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	pchainapi "github.com/ava-labs/avalanchego/vms/platformvm/api"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/signature-aggregator/metrics"
	"github.com/ava-labs/icm-services/types"
	"github.com/ava-labs/icm-services/utils"
	"go.uber.org/zap"
)

const (
	APIPath                 = "/aggregate-signatures"
	DefaultQuorumPercentage = 67
)

// Defines a request interface for signature aggregation for a raw unsigned message.
type AggregateSignatureRequest struct {
	// Required: either Message or Justification must be provided.
	// hex-encoded message, optionally prefixed with "0x".
	Message string `json:"message"`
	// hex-encoded justification, optionally prefixed with "0x".
	Justification string `json:"justification"`
	// Optional hex or cb58 encoded signing subnet ID. If omitted will default to the subnetID of the source blockchain
	SigningSubnetID string `json:"signing-subnet-id"`
	// Optional. Integer from 0 to 100 representing the percentage of the weight of the signing Subnet that is required
	// to sign the message. Defaults to 67 if omitted.
	QuorumPercentage uint64 `json:"quorum-percentage"`
	// Optional. Integer from 0 to 100 representing the additional percentage of weight of the signing Subnet that
	// will be attempted to add to the signature. `QuorumPercentage`+`QuorumPercentageBuffer` must be less than or
	// equal to 100. Obtaining signatures from more validators can take a longer time, but signatures representing
	// a large percentage of the Subnet weight are less prone to become invalid due to validator weight changes.
	// Defaults to 0 if omitted.
	QuorumPercentageBuffer uint64 `json:"quorum-percentage-buffer"`
	// Optional P-Chain height for validator set selection. If 0 (default), validators at proposed height will be used.
	// If non-zero, validators at the specified P-Chain height will be used for signature aggregation.
	PChainHeight uint64 `json:"pchain-height"`
}

type AggregateSignatureResponse struct {
	// hex encoding of the signature
	SignedMessage string `json:"signed-message"`
}

type AggregateSignatureErrorResponse struct {
	Error string `json:"error"`
}

func HandleAggregateSignaturesByRawMsgRequest(
	logger logging.Logger,
	metrics *metrics.SignatureAggregatorMetrics,
	signatureAggregator *aggregator.SignatureAggregator,
) {
	http.Handle(
		APIPath,
		signatureAggregationAPIHandler(
			logger,
			metrics,
			signatureAggregator,
		),
	)
}

func writeJSONError(
	logger logging.Logger,
	w http.ResponseWriter,
	httpStatusCode int,
	errorMsg string,
) {
	resp, err := json.Marshal(
		AggregateSignatureErrorResponse{
			Error: errorMsg,
		},
	)
	if err != nil {
		msg := "Error marshalling JSON error response"
		logger.Error(msg, zap.Error(err))
		resp = []byte(msg)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)

	_, err = w.Write(resp)
	if err != nil {
		logger.Error("Error writing error response", zap.Error(err))
	}
}

func signatureAggregationAPIHandler(
	logger logging.Logger,
	metrics *metrics.SignatureAggregatorMetrics,
	aggregator *aggregator.SignatureAggregator,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics.AggregateSignaturesRequestCount.Inc()
		startTime := time.Now()

		var req AggregateSignatureRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			msg := "Could not decode request body"
			logger.Warn(msg, zap.Error(err))
			writeJSONError(logger, w, http.StatusBadRequest, msg)
			return
		}
		var decodedMessage []byte
		decodedMessage, err = hex.DecodeString(
			utils.SanitizeHexString(req.Message),
		)
		if err != nil {
			msg := "Could not decode message"
			logger.Warn(
				msg,
				zap.String("msg", req.Message),
				zap.Error(err),
			)
			writeJSONError(logger, w, http.StatusBadRequest, msg)
			return
		}
		message, err := types.UnpackWarpMessage(decodedMessage)
		if err != nil {
			msg := "Error unpacking warp message"
			logger.Warn(msg, zap.Error(err))
			writeJSONError(logger, w, http.StatusBadRequest, msg)
			return
		}

		justification, err := hex.DecodeString(
			utils.SanitizeHexString(req.Justification),
		)
		if err != nil {
			msg := "Could not decode justification"
			logger.Warn(
				msg,
				zap.String("justification", req.Justification),
				zap.Error(err),
			)
			writeJSONError(logger, w, http.StatusBadRequest, msg)
			return
		}

		if utils.IsEmptyOrZeroes(message.Bytes()) && utils.IsEmptyOrZeroes(justification) {
			writeJSONError(
				logger,
				w,
				http.StatusBadRequest,
				"Must provide either message or justification",
			)
			return
		}

		quorumPercentage := req.QuorumPercentage
		if quorumPercentage == 0 {
			quorumPercentage = DefaultQuorumPercentage
		} else if req.QuorumPercentage > 100 {
			msg := "Invalid quorum number"
			logger.Warn(msg, zap.Uint64("quorum-num", req.QuorumPercentage))
			writeJSONError(logger, w, http.StatusBadRequest, msg)
			return
		}

		if quorumPercentage+req.QuorumPercentageBuffer > 100 {
			msg := "Invalid quorum buffer number"
			logger.Warn(
				msg,
				zap.Uint64("quorum-buffer-num", req.QuorumPercentageBuffer),
			)
			writeJSONError(logger, w, http.StatusBadRequest, msg)
			return
		}

		var signingSubnetID ids.ID
		if req.SigningSubnetID != "" {
			signingSubnetID, err = utils.HexOrCB58ToID(
				req.SigningSubnetID,
			)
			if err != nil {
				msg := "Error parsing signing subnet ID"
				logger.Warn(
					msg,
					zap.Error(err),
					zap.String("input", req.SigningSubnetID),
				)
				writeJSONError(logger, w, http.StatusBadRequest, msg)
				return
			}
		}

		// Determine P-Chain height: use ProposedHeight (latest) if not specified
		pchainHeight := req.PChainHeight
		if pchainHeight == 0 {
			pchainHeight = pchainapi.ProposedHeight
			logger.Debug("Using ProposedHeight for current validators",
				zap.Uint64("pchainHeight", pchainHeight),
			)
		} else {
			logger.Debug("Using specified P-Chain height",
				zap.Uint64("pchainHeight", pchainHeight),
			)
		}

		ctx, cancel := context.WithTimeout(r.Context(), utils.DefaultCreateSignedMessageTimeout)
		defer cancel()

		signedMessage, err := aggregator.CreateSignedMessage(
			ctx,
			logger,
			message,
			justification,
			signingSubnetID,
			quorumPercentage,
			req.QuorumPercentageBuffer,
			false,
			pchainHeight, // ACP-181: Use determined P-Chain height for validator set selection
		)
		if err != nil {
			logger.Warn("Failed to aggregate signatures", zap.Error(err))
			msg := fmt.Errorf("failed to aggregate signatures. error: %w", err).Error()
			writeJSONError(logger, w, http.StatusInternalServerError, msg)
			return
		}
		resp, err := json.Marshal(
			AggregateSignatureResponse{
				SignedMessage: hex.EncodeToString(
					signedMessage.Bytes(),
				),
			},
		)

		if err != nil {
			msg := "Failed to marshal response"
			logger.Error(msg, zap.Error(err))
			writeJSONError(logger, w, http.StatusInternalServerError, msg)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(resp)
		if err != nil {
			logger.Error("Error writing response", zap.Error(err))
		}
		metrics.AggregateSignaturesLatencyMS.Set(
			float64(time.Since(startTime).Milliseconds()),
		)
	})
}
