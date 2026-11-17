package api

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strings"

	"github.com/ava-labs/avalanchego/utils/logging"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/icm-services/messages"
	"github.com/ava-labs/icm-services/relayer"
	"github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/common"
	"go.uber.org/zap"
)

const (
	RelayAPIPath        = "/relay"
	RelayMessageAPIPath = RelayAPIPath + "/message"

	// MaxRequestBodySize bounds the size of a relayer API request body. The largest valid request
	// is a /relay/message one, which carries a base64 encoded unsigned Warp message. The message
	// handlers reject a Warp payload over warpPayload.MaxMessageSize, so twice that leaves room
	// for the base64 expansion (4/3), the framing around the payload and the rest of the JSON.
	MaxRequestBodySize = 2 * warpPayload.MaxMessageSize
)

func sanitizeLogValue(value string) string {
	value = strings.ReplaceAll(value, "\n", "")
	value = strings.ReplaceAll(value, "\r", "")
	return value
}

type RelayMessageRequest struct {
	// Required. cb58-encoded or "0x" prefixed hex-encoded source blockchain ID for the message
	BlockchainID string `json:"blockchain-id"`
	// Required. cb58-encoded or "0x" prefixed hex-encoded warp message ID
	MessageID string `json:"message-id"`
	// Required. Block number that the message was sent in
	BlockNum uint64 `json:"block-num"`
}

type RelayMessageResponse struct {
	// hex encoding of the transaction hash containing the processed message
	TransactionHash string `json:"transaction-hash"`
}

// Defines a manual warp message to be sent from the relayer through the API.
type ManualWarpMessageRequest struct {
	UnsignedMessageBytes []byte `json:"unsigned-message-bytes"`
	SourceAddress        string `json:"source-address"`
}

func HandleRelayMessage(mux *http.ServeMux, logger logging.Logger, messageCoordinator *relayer.MessageCoordinator) {
	mux.Handle(RelayMessageAPIPath, relayMessageAPIHandler(logger, messageCoordinator))
}

func HandleRelay(mux *http.ServeMux, logger logging.Logger, messageCoordinator *relayer.MessageCoordinator) {
	mux.Handle(RelayAPIPath, relayAPIHandler(logger, messageCoordinator))
}

// decodeRequest decodes the body of [r] into [req], reading at most MaxRequestBodySize of it. If
// that fails, it writes the error response and returns false.
func decodeRequest(logger logging.Logger, w http.ResponseWriter, r *http.Request, req any) bool {
	status, err := utils.DecodeJSONBody(w, r, MaxRequestBodySize, req)
	if err == nil {
		return true
	}
	msg := "could not decode request body"
	if status == http.StatusRequestEntityTooLarge {
		msg = "request body too large"
	}
	logger.Warn("Could not decode request body", zap.Error(err))
	http.Error(w, msg, status)
	return false
}

func relayMessageAPIHandler(logger logging.Logger, messageCoordinator *relayer.MessageCoordinator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req ManualWarpMessageRequest
		if !decodeRequest(logger, w, r, &req) {
			return
		}

		unsignedMessage, err := utils.UnpackWarpMessage(req.UnsignedMessageBytes)
		if err != nil {
			logger.Warn("Error unpacking warp message", zap.Error(err))
			http.Error(w, "error unpacking warp message", http.StatusBadRequest)
			return
		}

		if !common.IsHexAddress(req.SourceAddress) {
			logger.Warn("Invalid source address")
			http.Error(w, "invalid source address", http.StatusBadRequest)
			return
		}
		address := common.HexToAddress(req.SourceAddress)

		// The message protocol at [address] parses the payload itself, so pass along the
		// unsigned message bytes as provided in the request.
		sourceMessage := &messages.SourceMessage{
			SourceBlockchainID: unsignedMessage.SourceChainID,
			ProtocolAddress:    address,
			Payload:            req.UnsignedMessageBytes,
		}
		logger.Info(
			"Processing manual warp message",
			zap.Stringer("sourceAddress", address),
			zap.Stringer("messageID", unsignedMessage.ID()),
		)
		txHash, err := messageCoordinator.ProcessMessage(sourceMessage)
		if err != nil {
			logger.Error("Error processing message", zap.Error(err))
			http.Error(w, "error processing message", http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(
			RelayMessageResponse{
				TransactionHash: txHash.Hex(),
			},
		)
		if err != nil {
			logger.Error("Error marshaling response", zap.Error(err))
			http.Error(w, "error marshaling response", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(resp)
		if err != nil {
			logger.Error("Error writing response", zap.Error(err))
		}
	})
}

func relayAPIHandler(logger logging.Logger, messageCoordinator *relayer.MessageCoordinator) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req RelayMessageRequest
		if !decodeRequest(logger, w, r, &req) {
			return
		}

		blockchainID, err := utils.HexOrCB58ToID(req.BlockchainID)
		if err != nil {
			logger.Warn("Invalid blockchainID", zap.String("blockchainID", sanitizeLogValue(req.BlockchainID)), zap.Error(err))
			http.Error(w, "invalid blockchainID", http.StatusBadRequest)
			return
		}
		messageID, err := utils.HexOrCB58ToID(req.MessageID)
		if err != nil {
			logger.Warn("Invalid messageID", zap.String("messageID", sanitizeLogValue(req.MessageID)), zap.Error(err))
			http.Error(w, "invalid messageID", http.StatusBadRequest)
			return
		}

		txHash, err := messageCoordinator.ProcessMessageID(blockchainID, messageID, new(big.Int).SetUint64(req.BlockNum))
		if err != nil {
			logger.Error("Error processing message", zap.Error(err))
			http.Error(w, "error processing message", http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(
			RelayMessageResponse{
				TransactionHash: txHash.Hex(),
			},
		)
		if err != nil {
			logger.Error("Error marshalling response", zap.Error(err))
			http.Error(w, "error marshalling response", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(resp)
		if err != nil {
			logger.Error("Error writing response", zap.Error(err))
		}
	})
}
