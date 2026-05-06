// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"sort"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/logging"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	warppayload "github.com/ava-labs/avalanchego/vms/platformvm/warp/payload"
	diffupdater "github.com/ava-labs/icm-services/abi-bindings/go/DiffUpdater"
	"github.com/ava-labs/icm-services/config"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	testinfo "github.com/ava-labs/icm-services/icm-contracts/tests/test-info"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/peers/clients"
	relayercfg "github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/icm-services/relayer/validatorupdater"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/crypto"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

// DiffUpdaterReset tests the reset path of the DiffUpdater contract by
// manually constructing and submitting a reset payload after the relayer
// has registered the L1 set. The relayer's auto-detection of stale anchor
// (the trigger that wraps a sign failure as `errAnchorUnsignable` and
// falls back to `performResetUpdate`) cannot be reliably exercised in
// tmpnet because the live failure mode requires P-chain state pruning,
// which doesn't occur in local devnets. So we bypass the auto-trigger and
// drive the same on-chain reset that `performResetUpdate` would have
// driven, validating:
//
//  1. The signature aggregator can produce an L1-quorum signature on a
//     reset-shaped `ValidatorSetMetadata` payload (`prevHeight=0`,
//     `prevTimestamp=0`, every current validator listed as an addition).
//  2. The justification format produced by `BuildDiffJustification` is
//     accepted by AvalancheGo's WarpAPI verifier for the reset shape.
//  3. The DiffUpdater contract verifies the L1 signature against the
//     already-registered L1 set and applies the reset, replacing the
//     registered validator set wholesale.
//  4. The contract emits `ValidatorSetReset` (not `ValidatorSetUpdated`
//     or `ValidatorSetRegistered`).
//
// For coverage of the relayer's auto-detection logic itself, see the Go
// unit tests in `relayer/validatorupdater/`.
func DiffUpdaterReset(
	ctx context.Context,
	log logging.Logger,
	avalancheNetwork *network.LocalAvalancheNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
	teleporter utils.TeleporterTestInfo,
) {
	log.Info("Starting DiffUpdaterReset e2e test")

	l1Info := avalancheNetwork.GetL1Infos()[0]
	blockchainID := l1Info.BlockchainID
	networkID := avalancheNetwork.GetNetworkID()

	log.Info("Test configuration",
		zap.Stringer("blockchainID", blockchainID),
		zap.Stringer("subnetID", l1Info.SubnetID),
		zap.Uint32("networkID", networkID),
	)

	ethClient := ethereumNetwork.EthClient
	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()
	chainID := ethereumNetwork.ChainID
	fundedAddress, fundedKey := avalancheNetwork.GetFundedAccountInfo()

	// =========================================================================
	// Setup: Fetch primary network validators for P-chain bootstrap.
	// (Same shape as DiffUpdater test.)
	// =========================================================================
	primaryNetworkInfo := avalancheNetwork.GetPrimaryNetworkInfo()
	pChainClient := clients.NewCanonicalValidatorClient(&config.APIConfig{
		BaseURL: primaryNetworkInfo.NodeURIs[0],
	})
	pChainHeight, err := pChainClient.GetLatestHeight(ctx)
	Expect(err).Should(BeNil())

	pChainWarpSet, err := pChainClient.GetProposedValidators(ctx, ids.Empty)
	Expect(err).Should(BeNil())

	pChainValidators := make([]*validatorupdater.Validator, len(pChainWarpSet.Validators))
	for i, vdr := range pChainWarpSet.Validators {
		pChainValidators[i] = &validatorupdater.Validator{
			UncompressedPublicKeyBytes: [96]byte(vdr.PublicKey.Serialize()),
			Weight:                     vdr.Weight,
		}
	}
	sort.Slice(pChainValidators, func(i, j int) bool {
		return string(pChainValidators[i].UncompressedPublicKeyBytes[:]) <
			string(pChainValidators[j].UncompressedPublicKeyBytes[:])
	})

	var pChainID [32]byte // all zeros = PlatformChainID
	pChainTimestamp, err := pChainClient.GetBlockTimestampAtHeight(ctx, pChainHeight)
	Expect(err).Should(BeNil())

	bootstrapHeight := pChainHeight + 1
	bootstrapTimestamp := pChainTimestamp + 1

	pChainShardBytesList, pChainShardHashes, err := validatorupdater.ShardValidatorsAsDiff(
		pChainValidators,
		testShardSize,
		ids.ID(pChainID),
		pChainHeight,
		pChainTimestamp,
		bootstrapHeight,
		bootstrapTimestamp,
	)
	Expect(err).Should(BeNil())

	pChainShardHashesBytes := make([][32]byte, len(pChainShardHashes))
	for i, h := range pChainShardHashes {
		pChainShardHashesBytes[i] = h
	}

	// =========================================================================
	// Setup: Deploy DiffUpdater contract and bootstrap P-chain validators.
	// =========================================================================
	txOpts, err := bind.NewKeyedTransactorWithChainID(ethFundedKey, chainID)
	Expect(err).Should(BeNil())
	txOpts.GasLimit = diffUpdaterBootstrapGasLimit

	initialMetadata := diffupdater.ValidatorSetMetadata{
		AvalancheBlockchainID: pChainID,
		PChainHeight:          pChainHeight,
		PChainTimestamp:       pChainTimestamp,
		ShardHashes:           pChainShardHashesBytes,
	}
	contractAddr := utils.DeployDiffUpdaterWithMetadata(
		ctx,
		ethereumNetwork.EthereumTestInfo(),
		ethFundedKey,
		networkID,
		initialMetadata,
	)
	contract, err := diffupdater.NewDiffUpdater(contractAddr, ethClient)
	Expect(err).Should(BeNil())

	for i, shardBytes := range pChainShardBytesList {
		shard := diffupdater.ValidatorSetShard{
			ShardNumber:           uint64(i + 1),
			AvalancheBlockchainID: pChainID,
		}
		tx, err := contract.UpdateValidatorSet(txOpts, shard, shardBytes)
		Expect(err).Should(BeNil())
		receipt, err := bind.WaitMined(ctx, ethClient, tx)
		Expect(err).Should(BeNil())
		Expect(receipt.Status).Should(Equal(types.ReceiptStatusSuccessful),
			"updateValidatorSet shard %d failed", i+1)
	}

	callOpts := &bind.CallOpts{Context: ctx}
	pChainInitialized, err := contract.PChainInitialized(callOpts)
	Expect(err).Should(BeNil())
	Expect(pChainInitialized).Should(BeTrue())

	// =========================================================================
	// Setup: Configure and start the relayer.
	// =========================================================================
	err = utils.ClearRelayerStorage()
	Expect(err).Should(BeNil())

	relayerKey, err := crypto.GenerateKey()
	Expect(err).Should(BeNil())
	utils.FundRelayers(ctx, []testinfo.L1TestInfo{l1Info}, fundedKey, relayerKey)

	relayerConfig := createDiffUpdaterResetRelayerConfig(
		log,
		teleporter,
		l1Info,
		fundedAddress,
		relayerKey,
		ethereumNetwork,
		contractAddr.Hex(),
		blockchainID.String(),
		l1Info.SubnetID.String(),
	)
	relayerConfigPath := utils.WriteRelayerConfig(log, relayerConfig, utils.DefaultRelayerCfgFname)

	log.Info("Starting relayer")
	relayerCleanup, readyChan := utils.RunRelayerExecutable(
		ctx,
		log,
		relayerConfigPath,
		relayerConfig,
	)
	defer relayerCleanup()

	startupCtx, startupCancel := context.WithTimeout(ctx, 60*time.Second)
	defer startupCancel()
	utils.WaitForChannelClose(startupCtx, readyChan)

	// =========================================================================
	// Wait for the relayer to register the initial L1 validator set.
	// =========================================================================
	pollCtx, pollCancel := context.WithTimeout(ctx, 120*time.Second)
	defer pollCancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	var preResetVS diffupdater.ValidatorSet
	for done := false; !done; {
		select {
		case <-pollCtx.Done():
			Expect(pollCtx.Err()).ShouldNot(HaveOccurred(),
				"Timed out waiting for relayer to register validator set")
		case <-ticker.C:
			vs, qerr := contract.GetValidatorSet(callOpts, blockchainID)
			if qerr != nil {
				log.Warn("Failed to query on-chain validator set", zap.Error(qerr))
				continue
			}
			if vs.TotalWeight == 0 {
				continue
			}
			preResetVS = vs
			done = true
		}
	}
	log.Info("Initial L1 set registered by relayer",
		zap.Int("validatorCount", len(preResetVS.Validators)),
		zap.Uint64("totalWeight", preResetVS.TotalWeight),
		zap.Uint64("pChainHeight", preResetVS.PChainHeight),
		zap.Uint64("pChainTimestamp", preResetVS.PChainTimestamp),
	)

	// =========================================================================
	// Reset phase: stop the relayer to avoid races, then manually craft and
	// submit an L1-signed reset.
	// =========================================================================
	log.Info("Stopping relayer to manually drive reset")
	relayerCleanup()

	// Force the P-chain to advance strictly past the registered height so
	// the contract's `diff.currentHeight > registered.pChainHeight`
	// monotonicity check passes. In tmpnet the P-chain doesn't advance on
	// its own without P-chain activity, so we issue a no-op BaseTx
	// (self-transfer with empty outputs) and poll for the resulting height
	// bump.
	log.Info("Issuing P-chain BaseTx to advance height past registered anchor",
		zap.Uint64("registeredHeight", preResetVS.PChainHeight),
	)
	pChainWallet := avalancheNetwork.GetPChainWallet()
	_, err = pChainWallet.IssueBaseTx(nil)
	Expect(err).Should(BeNil(), "failed to issue P-chain BaseTx to advance height")

	advanceCtx, advanceCancel := context.WithTimeout(ctx, 90*time.Second)
	defer advanceCancel()
	advanceTicker := time.NewTicker(2 * time.Second)
	defer advanceTicker.Stop()
	var resetPChainHeight uint64
	for {
		latest, lerr := pChainClient.GetLatestHeight(advanceCtx)
		if lerr == nil && latest > preResetVS.PChainHeight {
			resetPChainHeight = latest
			break
		}
		log.Debug("Waiting for P-chain to advance past registered height",
			zap.Uint64("registeredHeight", preResetVS.PChainHeight),
			zap.Uint64("latestHeight", latest),
			zap.Error(lerr),
		)
		select {
		case <-advanceCtx.Done():
			Expect(advanceCtx.Err()).ShouldNot(HaveOccurred(),
				"timed out waiting for P-chain to advance past registered height %d",
				preResetVS.PChainHeight)
		case <-advanceTicker.C:
		}
	}
	resetPChainTimestamp, err := pChainClient.GetBlockTimestampAtHeight(ctx, resetPChainHeight)
	Expect(err).Should(BeNil())
	Expect(resetPChainTimestamp).Should(BeNumerically(">", preResetVS.PChainTimestamp))
	log.Info("P-chain advanced; proceeding with manual reset",
		zap.Uint64("registeredHeight", preResetVS.PChainHeight),
		zap.Uint64("resetPChainHeight", resetPChainHeight),
	)

	// Fetch the canonical L1 validator set at the reset target height (same
	// source the relayer's performResetUpdate uses).
	currentL1Validators := fetchSortedL1ValidatorsAtHeight(
		ctx, pChainClient, l1Info.SubnetID, resetPChainHeight,
	)
	Expect(len(currentL1Validators)).Should(BeNumerically(">", 0))

	// Build reset shards: prevHeight=0 / prevTimestamp=0 marks "empty prior
	// state" so the contract treats this as a wholesale replacement; every
	// current validator is listed as an addition.
	resetShardBytesList, resetShardHashes, err := validatorupdater.ShardValidatorsAsDiff(
		currentL1Validators,
		testShardSize,
		blockchainID,
		0, 0,
		resetPChainHeight,
		resetPChainTimestamp,
	)
	Expect(err).Should(BeNil())

	resetShardHashesBytes := make([][32]byte, len(resetShardHashes))
	for i, h := range resetShardHashes {
		resetShardHashesBytes[i] = h
	}

	// Build the metadata that the L1 quorum will sign.
	metadataMsg, err := validatorupdater.NewValidatorSetMetadata(
		blockchainID,
		resetPChainHeight,
		resetPChainTimestamp,
		resetShardHashes,
	)
	Expect(err).Should(BeNil())

	addressedCall, err := warppayload.NewAddressedCall(nil, metadataMsg.Bytes())
	Expect(err).Should(BeNil())

	unsignedMsg, err := avalancheWarp.NewUnsignedMessage(
		networkID,
		ids.ID(pChainID), // signed warp payloads claim P-chain as source
		addressedCall.Bytes(),
	)
	Expect(err).Should(BeNil())

	// Sign with the L1 quorum: the contract has the L1 set registered and
	// will verify against it. signingSubnet = l1Info.SubnetID forces the
	// aggregator to query L1 validators (not the primary network).
	aggregator := avalancheNetwork.GetSignatureAggregator()
	defer aggregator.Shutdown()
	justification := validatorupdater.BuildDiffJustification(testShardSize, 0, 0)

	log.Info("Signing reset metadata via L1 quorum",
		zap.Uint64("resetPChainHeight", resetPChainHeight),
		zap.Int("numValidators", len(currentL1Validators)),
		zap.Int("numShards", len(resetShardBytesList)),
	)
	signedMsg, err := aggregator.CreateSignedMessage(
		unsignedMsg,
		justification,
		l1Info.SubnetID,
		67,
		l1Info,
	)
	Expect(err).Should(BeNil(), "L1 quorum sign of reset metadata failed")

	icmMessage, err := signedMessageToICMMessage(signedMsg)
	Expect(err).Should(BeNil())

	// Submit the reset to the contract. Shard 0 -> RegisterValidatorSet,
	// remaining shards -> UpdateValidatorSet.
	log.Info("Submitting reset RegisterValidatorSet", zap.String("contract", contractAddr.Hex()))
	tx, err := contract.RegisterValidatorSet(txOpts, icmMessage, resetShardBytesList[0])
	Expect(err).Should(BeNil())
	registerReceipt, err := bind.WaitMined(ctx, ethClient, tx)
	Expect(err).Should(BeNil())
	Expect(registerReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful),
		"reset RegisterValidatorSet reverted: %s", tx.Hash().Hex())

	finalReceipt := registerReceipt
	for i := 1; i < len(resetShardBytesList); i++ {
		shard := diffupdater.ValidatorSetShard{
			ShardNumber:           uint64(i + 1),
			AvalancheBlockchainID: blockchainID,
		}
		log.Info("Submitting reset UpdateValidatorSet", zap.Int("shardNumber", i+1))
		shardTx, terr := contract.UpdateValidatorSet(txOpts, shard, resetShardBytesList[i])
		Expect(terr).Should(BeNil())
		shardReceipt, terr := bind.WaitMined(ctx, ethClient, shardTx)
		Expect(terr).Should(BeNil())
		Expect(shardReceipt.Status).Should(Equal(types.ReceiptStatusSuccessful),
			"reset UpdateValidatorSet shard %d reverted: %s", i+1, shardTx.Hash().Hex())
		finalReceipt = shardReceipt
	}

	// =========================================================================
	// Assertions: ValidatorSetReset event + replaced validator set.
	// =========================================================================
	resetEvent := findValidatorSetResetEvent(contract, finalReceipt, blockchainID)
	Expect(resetEvent).ShouldNot(BeNil(),
		"contract must emit ValidatorSetReset for blockchainID=%s", blockchainID)
	log.Info("Contract emitted ValidatorSetReset",
		zap.Uint64("blockNumber", resetEvent.Raw.BlockNumber),
		zap.String("txHash", resetEvent.Raw.TxHash.Hex()),
	)

	postResetVS, err := contract.GetValidatorSet(callOpts, blockchainID)
	Expect(err).Should(BeNil())
	Expect(postResetVS.PChainHeight).Should(Equal(resetPChainHeight),
		"on-chain pChainHeight must advance to the reset target")
	Expect(postResetVS.PChainTimestamp).Should(Equal(resetPChainTimestamp))
	Expect(len(postResetVS.Validators)).Should(Equal(len(currentL1Validators)),
		"validator set length must match the canonical L1 set at the reset target height")

	var expectedTotalWeight uint64
	for _, exp := range currentL1Validators {
		expectedTotalWeight += exp.Weight
	}
	Expect(postResetVS.TotalWeight).Should(Equal(expectedTotalWeight))
	for i, exp := range currentL1Validators {
		Expect(postResetVS.Validators[i].Weight).Should(Equal(exp.Weight),
			"validator %d: weight mismatch after reset", i)
		Expect(postResetVS.Validators[i].BlsPublicKey).Should(
			Equal(padUncompressedBLSPublicKey(exp.UncompressedPublicKeyBytes[:])),
			"validator %d: BLS public key mismatch after reset", i)
	}

	log.Info("DiffUpdaterReset e2e test PASSED")
}

// signedMessageToICMMessage converts a signed warp message to the
// `ICMMessage` shape the DiffUpdater contract expects, mirroring the
// relayer's internal `buildDiffICMMessage` helper.
func signedMessageToICMMessage(
	signedMsg *avalancheWarp.Message,
) (diffupdater.ICMMessage, error) {
	addressedCall, err := warppayload.ParseAddressedCall(signedMsg.UnsignedMessage.Payload)
	Expect(err).Should(BeNil())

	bitSetSig, ok := signedMsg.Signature.(*avalancheWarp.BitSetSignature)
	Expect(ok).Should(BeTrue(), "expected BitSetSignature, got %T", signedMsg.Signature)

	sig, err := bls.SignatureFromBytes(bitSetSig.Signature[:])
	Expect(err).Should(BeNil())
	uncompressedSig := sig.Serialize()

	attestation := make([]byte, 0, len(bitSetSig.Signers)+len(uncompressedSig))
	attestation = append(attestation, bitSetSig.Signers...)
	attestation = append(attestation, uncompressedSig...)

	return diffupdater.ICMMessage{
		RawMessage:         addressedCall.Payload,
		SourceNetworkID:    signedMsg.UnsignedMessage.NetworkID,
		SourceBlockchainID: signedMsg.UnsignedMessage.SourceChainID,
		Attestation:        attestation,
	}, nil
}

// findValidatorSetResetEvent scans the receipt for a ValidatorSetReset
// event matching the provided blockchainID. Returns nil if not found.
func findValidatorSetResetEvent(
	contract *diffupdater.DiffUpdater,
	receipt *types.Receipt,
	blockchainID ids.ID,
) *diffupdater.DiffUpdaterValidatorSetReset {
	iter, err := contract.FilterValidatorSetReset(
		&bind.FilterOpts{
			Start: receipt.BlockNumber.Uint64(),
			End:   ptrUint64(receipt.BlockNumber.Uint64()),
		},
		[][32]byte{blockchainID},
	)
	Expect(err).Should(BeNil())
	defer iter.Close()
	if !iter.Next() {
		return nil
	}
	return iter.Event
}

func ptrUint64(v uint64) *uint64 { return &v }

// createDiffUpdaterResetRelayerConfig produces a relayer config tuned for
// the reset test: short staleness window, default threshold semantics.
// The relayer is only used to drive the *initial* L1 registration; we
// stop it before manually submitting the reset.
func createDiffUpdaterResetRelayerConfig(
	log logging.Logger,
	teleporter utils.TeleporterTestInfo,
	l1Info testinfo.L1TestInfo,
	fundedAddress common.Address,
	relayerKey *ecdsa.PrivateKey,
	ethereumNetwork *network.LocalEthereumNetwork,
	contractAddress string,
	blockchainID string,
	subnetID string,
) relayercfg.Config {
	baseConfig := utils.CreateDefaultRelayerConfig(
		log,
		teleporter,
		[]testinfo.L1TestInfo{l1Info},
		[]testinfo.L1TestInfo{l1Info},
		fundedAddress,
		relayerKey,
	)

	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()

	// Distinct from the DiffUpdater test's APIPort to avoid conflicts when
	// both flows run in the same suite.
	baseConfig.APIPort = 8084

	baseConfig.ExternalEVMDestinations = []*relayercfg.ExternalEVMDestination{
		{
			RPCEndpoint:              ethereumNetwork.BaseURL,
			PrivateKey:               hex.EncodeToString(crypto.FromECDSA(ethFundedKey)),
			ContractAddress:          contractAddress,
			BlockchainID:             blockchainID,
			SubnetID:                 subnetID,
			ShardSize:                testShardSize,
			PollIntervalSeconds:      testPollIntervalSeconds,
			ContractType:             "diff",
			WeightChangeThresholdPct: thresholdWeightChangeThresholdPct,
			MaxUpdateIntervalSeconds: thresholdMaxUpdateIntervalSeconds,
		},
	}

	return baseConfig
}
