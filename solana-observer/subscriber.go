// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// LogEvent is emitted by the subscriber for each finalized transaction whose
// logs mention one of the watched programs. slot is included so downstream
// consumers can populate OracleMessage.SourceBlockHeight without a second RPC
// roundtrip if the notification carried it.
type LogEvent struct {
	TxSig   string
	Slot    uint64
	Program string
}

// runSubscriber opens one WebSocket connection per configured program and
// forwards finalized log events onto out. It runs until ctx is cancelled, and
// reconnects with a fixed backoff on any transport error.
//
// One goroutine per program. logsSubscribe's `mentions` filter accepts only a
// single pubkey per subscription, so we fan out.
func runSubscriber(
	ctx context.Context,
	log logging.Logger,
	wsURL string,
	commitment string,
	programs []string,
	out chan<- LogEvent,
) {
	for _, program := range programs {
		program := program
		go func() {
			for ctx.Err() == nil {
				if err := subscribeOnce(ctx, log, wsURL, commitment, program, out); err != nil && ctx.Err() == nil {
					log.Warn("subscription dropped, reconnecting",
						zap.String("program", program),
						zap.Error(err),
					)
					select {
					case <-ctx.Done():
					case <-time.After(2 * time.Second):
					}
				}
			}
		}()
	}
}

// subscribeOnce holds a single WebSocket subscription open until ctx is done or
// an error occurs. It returns nil only when ctx is cancelled.
func subscribeOnce(
	ctx context.Context,
	log logging.Logger,
	wsURL string,
	commitment string,
	program string,
	out chan<- LogEvent,
) error {
	dialer := websocket.Dialer{HandshakeTimeout: 10 * time.Second}
	conn, _, err := dialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		return fmt.Errorf("dial %s: %w", wsURL, err)
	}
	defer conn.Close()

	// Cancel the read loop if ctx dies.
	go func() {
		<-ctx.Done()
		_ = conn.Close()
	}()

	sub := map[string]any{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "logsSubscribe",
		"params": []any{
			map[string]any{"mentions": []string{program}},
			map[string]any{"commitment": commitment},
		},
	}
	if err := conn.WriteJSON(sub); err != nil {
		return fmt.Errorf("write subscription: %w", err)
	}

	log.Info("subscribed to Solana logs",
		zap.String("program", program),
		zap.String("commitment", commitment),
	)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			return fmt.Errorf("read: %w", err)
		}

		var envelope struct {
			Method string `json:"method"`
			Params struct {
				Result struct {
					Context struct {
						Slot uint64 `json:"slot"`
					} `json:"context"`
					Value struct {
						Signature string   `json:"signature"`
						Err       any      `json:"err"`
						Logs      []string `json:"logs"`
					} `json:"value"`
				} `json:"result"`
			} `json:"params"`
		}
		if err := json.Unmarshal(msg, &envelope); err != nil {
			// Skip non-notification frames (subscription-ack, keepalives).
			continue
		}
		if envelope.Method != "logsNotification" {
			continue
		}
		if envelope.Params.Result.Value.Err != nil {
			// Skip failed transactions.
			continue
		}
		if envelope.Params.Result.Value.Signature == "" {
			continue
		}
		select {
		case out <- LogEvent{
			TxSig:   envelope.Params.Result.Value.Signature,
			Slot:    envelope.Params.Result.Context.Slot,
			Program: program,
		}:
		case <-ctx.Done():
			return nil
		}
	}
}
