// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	oracleadapter "github.com/ava-labs/icm-services/abi-bindings/go/teleporterV2/OracleAdapter"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/icm-services/signature-aggregator/api"
	icmutils "github.com/ava-labs/icm-services/utils"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/ethclient"
	"go.uber.org/zap"
)

// oracleMsgABI mirrors the encoding at
// avalanchego/network/p2p/oracle/message.go and the E2E test flow. Duplicated
// here rather than imported so the observer does not require the sidecar
// branch to be published as a Go module.
var oracleMsgABI abi.Arguments

func init() {
	stringT, _ := abi.NewType("string", "", nil)
	addrT, _ := abi.NewType("address", "", nil)
	uint64T, _ := abi.NewType("uint64", "", nil)
	bytesT, _ := abi.NewType("bytes", "", nil)
	oracleMsgABI = abi.Arguments{
		{Type: stringT, Name: "sourceType"},
		{Type: stringT, Name: "sourceAddress"},
		{Type: addrT, Name: "destContract"},
		{Type: uint64T, Name: "sourceBlockHeight"},
		{Type: uint64T, Name: "nonce"},
		{Type: bytesT, Name: "payload"},
	}
}

// Relay owns the L1 delivery pipeline. Given a fetched Solana transaction, it
// constructs the OracleMessage, drives BLS aggregation against the running
// signature-aggregator, and submits the resulting warp message as the access
// list of a TeleporterMessengerV2 delivery call.
type Relay struct {
	log        logging.Logger
	cfg        *Config
	nonces     *NonceStore
	l1         *ethclient.Client
	l1ChainID  *big.Int
	blockchainID [32]byte
	fromAddr   common.Address
	fromKeyHex string
}

// NewRelay wires up L1 connectivity and the on-disk nonce store.
func NewRelay(ctx context.Context, log logging.Logger, cfg *Config) (*Relay, error) {
	l1, err := ethclient.DialContext(ctx, cfg.L1.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("dial l1: %w", err)
	}
	nonces, err := LoadNonces(cfg.NonceFile)
	if err != nil {
		return nil, fmt.Errorf("load nonce store: %w", err)
	}

	blockchainID, err := ids.FromString(cfg.L1.BlockchainID)
	if err != nil {
		return nil, fmt.Errorf("parse blockchain_id: %w", err)
	}

	// libevm's crypto.HexToECDSA requires no "0x" prefix.
	keyHex := strings.TrimPrefix(cfg.L1.DeliveryPrivateKeyHex, "0x")
	priv, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		return nil, fmt.Errorf("parse delivery private key: %w", err)
	}
	fromAddr := crypto.PubkeyToAddress(priv.PublicKey)

	return &Relay{
		log:          log,
		cfg:          cfg,
		nonces:       nonces,
		l1:           l1,
		l1ChainID:    new(big.Int).SetUint64(cfg.L1.ChainID),
		blockchainID: [32]byte(blockchainID),
		fromAddr:     fromAddr,
		fromKeyHex:   keyHex,
	}, nil
}

// Deliver runs the full pipeline for one Solana transaction. It is safe to call
// concurrently but callers should serialize per-source to keep nonces
// monotonic; the observer's default main loop is single-threaded.
func (r *Relay) Deliver(ctx context.Context, tx *SolanaTx) error {
	sourceType := "solana"
	nonce := r.nonces.Next(sourceType, tx.Program)

	oracleMsg := oracleadapter.OracleMessage{
		SourceType:        sourceType,
		SourceAddress:     tx.Program,
		DestContract:      r.cfg.L1.DestContract,
		SourceBlockHeight: new(big.Int).SetUint64(tx.Slot),
		Nonce:             new(big.Int).SetUint64(nonce),
		Payload:           tx.InstrData,
	}

	oraclePayload, err := oracleMsgABI.Pack(
		oracleMsg.SourceType,
		oracleMsg.SourceAddress,
		oracleMsg.DestContract,
		oracleMsg.SourceBlockHeight,
		oracleMsg.Nonce,
		oracleMsg.Payload,
	)
	if err != nil {
		return fmt.Errorf("abi pack: %w", err)
	}

	ac, err := payload.NewAddressedCall(nil, oraclePayload)
	if err != nil {
		return fmt.Errorf("new addressed call: %w", err)
	}
	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(r.cfg.L1.NetworkID, r.blockchainID, ac.Bytes())
	if err != nil {
		return fmt.Errorf("new unsigned message: %w", err)
	}

	signedMsg, err := r.aggregate(ctx, unsignedMsg, tx.SigBytes)
	if err != nil {
		return fmt.Errorf("aggregate: %w", err)
	}

	// Commit the nonce only after successful aggregation. If the L1 delivery
	// itself fails, we'd rather burn the nonce than let a retry reuse it
	// against an on-chain replay-protected mapping.
	if err := r.nonces.Save(); err != nil {
		r.log.Warn("failed to persist nonce", zap.Error(err))
	}

	if err := r.submitL1(ctx, signedMsg, oracleMsg); err != nil {
		return fmt.Errorf("submit l1: %w", err)
	}

	r.log.Info("oracle message delivered",
		zap.String("tx_sig", hex.EncodeToString(tx.SigBytes)),
		zap.String("program", tx.Program),
		zap.Uint64("slot", tx.Slot),
		zap.Uint64("nonce", nonce),
	)
	return nil
}

func (r *Relay) aggregate(
	ctx context.Context,
	unsigned *avalancheWarp.UnsignedMessage,
	justification []byte,
) (*avalancheWarp.Message, error) {
	body := api.AggregateSignatureRequest{
		Message:         "0x" + hex.EncodeToString(unsigned.Bytes()),
		Justification:   hex.EncodeToString(justification),
		SigningSubnetID: r.cfg.L1.SubnetID,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, r.cfg.Aggregator+api.OracleAPIPath, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("aggregator request: %w", err)
	}
	defer res.Body.Close()

	respBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("aggregator returned %s: %s", res.Status, string(respBytes))
	}

	var resp api.AggregateSignatureResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("parse aggregator response: %w", err)
	}
	decoded, err := hex.DecodeString(strings.TrimPrefix(resp.SignedMessage, "0x"))
	if err != nil {
		return nil, err
	}
	return avalancheWarp.ParseMessage(decoded)
}

func (r *Relay) submitL1(
	ctx context.Context,
	signed *avalancheWarp.Message,
	oracleMsg oracleadapter.OracleMessage,
) error {
	icmMsg, err := oracleadapter.BuildOracleICMMessage(
		0,
		oracleMsg,
		r.cfg.L1.TeleporterAddress,
		r.blockchainID,
		r.cfg.L1.NetworkID,
		new(big.Int).SetUint64(500_000),
	)
	if err != nil {
		return fmt.Errorf("build icm message: %w", err)
	}
	callData, err := teleportermessengerv2.PackReceiveCrossChainMessageV2(
		icmMsg.Message,
		r.blockchainID,
		icmMsg.Attestation,
		common.Address{},
	)
	if err != nil {
		return fmt.Errorf("pack receiveCrossChainMessage: %w", err)
	}

	gasFeeCap, gasTipCap, txNonce, err := r.txParams(ctx)
	if err != nil {
		return err
	}
	teleporter := r.cfg.L1.TeleporterAddress
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:    r.l1ChainID,
		Nonce:      txNonce,
		To:         &teleporter,
		Gas:        2_000_000,
		GasFeeCap:  gasFeeCap,
		GasTipCap:  gasTipCap,
		Value:      common.Big0,
		Data:       callData,
		AccessList: icmutils.SignedWarpMessageToAccessList(signed),
	})

	priv, err := crypto.HexToECDSA(r.fromKeyHex)
	if err != nil {
		return err
	}
	signed_tx, err := types.SignTx(tx, types.LatestSignerForChainID(r.l1ChainID), priv)
	if err != nil {
		return err
	}
	if err := r.l1.SendTransaction(ctx, signed_tx); err != nil {
		return fmt.Errorf("send tx: %w", err)
	}
	// Wait for inclusion with a bounded poll.
	deadline := time.Now().Add(60 * time.Second)
	for time.Now().Before(deadline) {
		receipt, err := r.l1.TransactionReceipt(ctx, signed_tx.Hash())
		if err == nil {
			if receipt.Status != types.ReceiptStatusSuccessful {
				return fmt.Errorf("delivery tx %s reverted", signed_tx.Hash().Hex())
			}
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(2 * time.Second):
		}
	}
	return fmt.Errorf("timed out waiting for delivery tx %s", signed_tx.Hash().Hex())
}

func (r *Relay) txParams(ctx context.Context) (*big.Int, *big.Int, uint64, error) {
	head, err := r.l1.BlockByNumber(ctx, nil)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("block by number: %w", err)
	}
	gasTipCap, err := r.l1.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("suggest gas tip cap: %w", err)
	}
	nonce, err := r.l1.NonceAt(ctx, r.fromAddr, nil)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("nonce at: %w", err)
	}
	baseFee := new(big.Int).Set(head.BaseFee())
	gasFeeCap := new(big.Int).Mul(baseFee, big.NewInt(2))
	gasFeeCap.Add(gasFeeCap, gasTipCap)
	return gasFeeCap, gasTipCap, nonce, nil
}
