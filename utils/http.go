// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

// DecodeJSONBody decodes the JSON body of [r] into [dst], reading at most [maxBytes] of it. A
// larger body fails before it is buffered in full, so a request cannot make the server allocate
// memory in proportion to a size chosen by its sender.
//
// If decoding fails, the returned status is the one to respond with:
// http.StatusRequestEntityTooLarge for a body over the limit, http.StatusBadRequest otherwise.
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, maxBytes int64, dst any) (int, error) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	err := json.NewDecoder(r.Body).Decode(dst)
	if err == nil {
		return http.StatusOK, nil
	}
	var maxBytesErr *http.MaxBytesError
	if errors.As(err, &maxBytesErr) {
		return http.StatusRequestEntityTooLarge, err
	}
	return http.StatusBadRequest, err
}
