package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warpPayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

// Both endpoints must reject a body larger than MaxRequestBodySize instead of buffering it, and
// must not reject a body that fits. Both cases fail before the message coordinator is used, so it
// can be nil.
func TestRelayAPIBodySizeLimit(t *testing.T) {
	mux := http.NewServeMux()
	HandleRelay(mux, logging.NoLog{}, nil)
	HandleRelayMessage(mux, logging.NoLog{}, nil)

	// One string field that alone exceeds the whole body limit, which is the shape of the memory
	// exhaustion attack on each endpoint.
	oversizedFields := map[string]string{
		RelayAPIPath:        "message-id",
		RelayMessageAPIPath: "unsigned-message-bytes",
	}
	for path, field := range oversizedFields {
		t.Run(path, func(t *testing.T) {
			oversized := `{"` + field + `":"` + strings.Repeat("a", MaxRequestBodySize) + `"}`
			req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(oversized))
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			require.Equal(t, http.StatusRequestEntityTooLarge, rec.Code)
			require.Contains(t, rec.Body.String(), "request body too large")

			// A body exactly at the limit is read in full: JSON whitespace padding up to the limit
			// followed by an invalid token must produce a decode error, not a size error.
			atLimit := strings.Repeat(" ", MaxRequestBodySize-len("not json")) + "not json"
			require.Len(t, atLimit, MaxRequestBodySize)
			req = httptest.NewRequest(http.MethodPost, path, strings.NewReader(atLimit))
			rec = httptest.NewRecorder()
			mux.ServeHTTP(rec, req)
			require.Equal(t, http.StatusBadRequest, rec.Code)
			require.Contains(t, rec.Body.String(), "could not decode request body")
		})
	}
}

// The body limit must leave room for the largest message that can be relayed: a Warp payload of
// the maximum size, in the larger of the two encodings /relay/message accepts.
func TestRelayMessageBodyLimitFitsLargestWarpMessage(t *testing.T) {
	unsignedMessage, err := avalancheWarp.NewUnsignedMessage(
		constants.UnitTestID,
		ids.GenerateTestID(),
		make([]byte, warpPayload.MaxMessageSize),
	)
	require.NoError(t, err)
	_, eventData, err := warp.PackSendWarpMessageEvent(
		common.Address{},
		common.Hash(unsignedMessage.ID()),
		unsignedMessage.Bytes(),
	)
	require.NoError(t, err)

	// The invalid source address stops the handler after it has decoded the body and unpacked the
	// Warp message, and before it uses the message coordinator.
	body, err := json.Marshal(ManualWarpMessageRequest{
		UnsignedMessageBytes: eventData,
		SourceAddress:        "not an address",
	})
	require.NoError(t, err)
	require.LessOrEqual(t, len(body), MaxRequestBodySize)

	mux := http.NewServeMux()
	HandleRelayMessage(mux, logging.NoLog{}, nil)
	req := httptest.NewRequest(http.MethodPost, RelayMessageAPIPath, bytes.NewReader(body))
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)
	require.Equal(t, http.StatusBadRequest, rec.Code)
	require.Contains(t, rec.Body.String(), "invalid source address")
}
