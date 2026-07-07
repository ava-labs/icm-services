// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// solana-observer subscribes to a Solana WebSocket endpoint, detects finalized
// transactions that invoke configured programs, and drives the oracle
// attestation pipeline to deliver each event as an OracleMessage on a
// destination Avalanche L1.
//
// Usage:
//
//	solana-observer --config-path /path/to/observer.json
//
// The config includes a `sidecar_config_path` field that points at the same
// sidecar config the validators consume; the observer reads the
// `verifiers.solana.allowed_programs` list from it so both the observer and
// the sidecar's SolanaVerifier agree on which programs are in scope.
package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("config-path", "", "path to observer config JSON")
	flag.Parse()
	if *configPath == "" {
		log.Fatal("--config-path is required")
	}

	cfg, err := Load(*configPath)
	if err != nil {
		log.Fatalf("failed to load observer config: %v", err)
	}

	programs, err := loadProgramAllowlist(cfg.SidecarConfigPath)
	if err != nil {
		log.Fatalf("failed to read program allowlist from sidecar config %s: %v", cfg.SidecarConfigPath, err)
	}
	if len(programs) == 0 {
		log.Fatalf("sidecar config %s has no solana programs to watch", cfg.SidecarConfigPath)
	}

	logger := logging.NewLogger("solana-observer", logging.NewWrappedCore(
		logging.Info,
		os.Stdout,
		logging.JSON.ConsoleEncoder(),
	))

	ctx, cancel := context.WithCancel(context.Background())
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		logger.Info("shutdown signal received")
		cancel()
	}()

	relay, err := NewRelay(ctx, logger, cfg)
	if err != nil {
		logger.Fatal("failed to init relay", zap.Error(err))
		os.Exit(1)
	}

	logger.Info("starting Solana observer",
		zap.Strings("programs", programs),
		zap.String("ws_url", cfg.Solana.WSURL),
		zap.String("aggregator", cfg.Aggregator),
	)

	events := make(chan LogEvent, 32)
	runSubscriber(ctx, logger, cfg.Solana.WSURL, cfg.Solana.Commitment, programs, events)

	for {
		select {
		case <-ctx.Done():
			logger.Info("observer stopped")
			return
		case ev := <-events:
			handleEvent(ctx, logger, cfg, relay, ev)
		}
	}
}

// handleEvent fetches the transaction body, extracts the instruction data, and
// hands it to the relay. A single failed event is logged and dropped —
// stopping the observer on a transient RPC error would be worse than skipping
// one demo message.
func handleEvent(ctx context.Context, logger logging.Logger, cfg *Config, relay *Relay, ev LogEvent) {
	// getTransaction sometimes lags a beat behind logsSubscribe finalization.
	// Retry a few times before giving up.
	var tx *SolanaTx
	var ok bool
	for attempt := 0; attempt < 5; attempt++ {
		var err error
		tx, ok, err = fetchTx(ctx, cfg.Solana.RPCURL, ev.TxSig, ev.Program)
		if err != nil {
			logger.Warn("fetchTx error", zap.String("sig", ev.TxSig), zap.Error(err))
			return
		}
		if ok {
			break
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):
		}
	}
	if !ok {
		logger.Warn("tx did not resolve to a matching instruction",
			zap.String("sig", ev.TxSig),
			zap.String("program", ev.Program),
		)
		return
	}
	if err := relay.Deliver(ctx, tx); err != nil {
		logger.Error("delivery failed",
			zap.String("sig", ev.TxSig),
			zap.Error(err),
		)
	}
}

// loadProgramAllowlist reads the sidecar config file, drills into
// verifiers.solana.allowed_programs, and returns the list. The observer treats
// the sidecar config as the single source of truth for which programs are in
// scope. Matches the shape defined by avalanchego/sidecar/config on the
// sidecar-verifier branch and the solanarpc.Config struct that consumes it.
func loadProgramAllowlist(path string) ([]string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var outer struct {
		Verifiers map[string]json.RawMessage `json:"verifiers"`
	}
	if err := json.Unmarshal(b, &outer); err != nil {
		return nil, err
	}
	solRaw, ok := outer.Verifiers["solana"]
	if !ok {
		return nil, errors.New(`sidecar config has no "solana" verifier entry`)
	}
	var sol struct {
		AllowedPrograms []string `json:"allowed_programs"`
	}
	if err := json.Unmarshal(solRaw, &sol); err != nil {
		return nil, err
	}
	return sol.AllowedPrograms, nil
}
