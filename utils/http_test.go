// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type decodeTestRequest struct {
	Value string `json:"value"`
}

func TestDecodeJSONBody(t *testing.T) {
	const limit = 64

	valid := `{"value":"abc"}`
	// JSON whitespace padding up to the limit, followed by an invalid token, must be read in full
	// and fail as undecodable rather than as too large.
	atLimit := strings.Repeat(" ", limit-len("not json")) + "not json"
	require.Len(t, atLimit, limit)

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantValue  string
	}{
		{name: "valid", body: valid, wantStatus: http.StatusOK, wantValue: "abc"},
		{name: "malformed", body: "not json", wantStatus: http.StatusBadRequest},
		{name: "malformed at the limit", body: atLimit, wantStatus: http.StatusBadRequest},
		{
			name:       "over the limit",
			body:       `{"value":"` + strings.Repeat("a", limit) + `"}`,
			wantStatus: http.StatusRequestEntityTooLarge,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			var dst decodeTestRequest
			status, err := DecodeJSONBody(httptest.NewRecorder(), r, limit, &dst)
			require.Equal(t, tt.wantStatus, status)
			if tt.wantStatus == http.StatusOK {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
			require.Equal(t, tt.wantValue, dst.Value)
		})
	}
}

// endlessReader yields 'a' forever and counts the bytes it has handed out.
type endlessReader struct {
	read int
}

func (e *endlessReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'a'
	}
	e.read += len(p)
	return len(p), nil
}

// A body that never ends, here a JSON string that is never closed, must be abandoned once the
// limit is reached instead of being buffered for as long as the sender keeps writing.
func TestDecodeJSONBodyStopsReadingAtLimit(t *testing.T) {
	const limit = 4096

	prefix := `{"value":"`
	endless := &endlessReader{}
	r := httptest.NewRequest(http.MethodPost, "/", io.MultiReader(strings.NewReader(prefix), endless))

	var dst decodeTestRequest
	status, err := DecodeJSONBody(httptest.NewRecorder(), r, limit, &dst)
	require.Equal(t, http.StatusRequestEntityTooLarge, status)
	var maxBytesErr *http.MaxBytesError
	require.ErrorAs(t, err, &maxBytesErr)
	// One byte past the limit is read to tell a body at the limit from one over it.
	require.LessOrEqual(t, len(prefix)+endless.read, limit+1)
}
