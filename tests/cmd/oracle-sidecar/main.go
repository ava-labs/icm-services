// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// oracle-sidecar is a minimal HTTP sidecar for E2E testing. It accepts every
// /verify request unconditionally, standing in for a real Solana RPC sidecar.
// Usage: oracle-sidecar --port 9900
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

type verifyResponse struct {
	Error string `json:"error,omitempty"`
}

func main() {
	port := flag.Uint("port", 9900, "port to listen on")
	flag.Parse()

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(verifyResponse{}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	addr := fmt.Sprintf(":%d", *port)
	fmt.Fprintf(os.Stdout, "oracle-sidecar listening on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "oracle-sidecar: %v\n", err)
		os.Exit(1)
	}
}
