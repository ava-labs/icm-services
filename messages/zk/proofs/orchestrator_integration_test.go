// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

package proofs

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/icm-services/messages/zk/proofs/beacon"
	"github.com/ava-labs/libevm/common"
	"github.com/stretchr/testify/require"
)

// The nightly integration test (in scripts/nightly_test.sh) sets these env
// variables. This includes the beacon endpoint to fetch from and the Ethereum
// fixture path that the nightly test just generated.
const (
	beaconAPIURLEnv = "BEACON_API_URL"
	fixturePathEnv  = "ETHEREUM_FIXTURE_PATH"
)

// The full pipeline, run against a real beacon node for the fixture's slots,
// must reproduce the fixture's proofs byte for byte. The fixture is generated
// by an independent implementation (lodestar), so this is the
// cross-implementation check on real data. It exercises the client, fork
// handling, lite state reduction, and proof building together.
func TestBuildExecutionProofForSlotsIntegration(t *testing.T) {
	// This test will skip in unit test CI runs since the env variables are not set.
	beaconAPIURL, fixturePath := os.Getenv(beaconAPIURLEnv), os.Getenv(fixturePathEnv)
	if beaconAPIURL == "" || fixturePath == "" {
		t.Skipf("%s and %s must be set", beaconAPIURLEnv, fixturePathEnv)
	}

	fixture := loadFixture(t, fixturePath)
	expected := fixture.ExecutionProof
	anchorBlockRoot := common.HexToHash(fixture.AnchorBeaconBlockRoot)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// Build execution proof.
	builder := NewProofBuilder(beacon.NewBeaconClient(beaconAPIURL))
	proof, err := builder.BuildExecutionProofForSlots(
		ctx, expected.AnchorSlot, expected.TargetSlot, anchorBlockRoot)
	require.NoError(t, err)

	// Check that proof produced by BuildExecutionProofForSlots matches the external party (lodestar)
	// generated proof.
	require.Equal(t, expected.AnchorSlot, proof.AnchorSlot)
	require.Equal(t, expected.TargetSlot, proof.TargetSlot)

	require.Equal(t, common.HexToHash(expected.AnchorBeaconStateRoot), common.Hash(proof.AnchorBeaconStateRoot))
	require.Equal(t, hexToHashes(expected.AnchorBeaconStateProof), proof.AnchorBeaconStateProof)

	require.Equal(t, common.HexToHash(expected.TargetBeaconStateRoot), common.Hash(proof.TargetBeaconStateRoot))
	require.Equal(t, hexToHashes(expected.TargetBeaconStateProof), proof.TargetBeaconStateProof)

	require.Equal(t, common.HexToHash(expected.TargetExecutionHeaderRoot), common.Hash(proof.TargetExecutionHeaderRoot))
	require.Equal(t, hexToHashes(expected.TargetExecutionHeaderProof), proof.TargetExecutionHeaderProof)

	require.Equal(t, common.HexToHash(expected.TargetReceiptsRoot), common.Hash(proof.TargetReceiptsRoot))
	require.Equal(t, hexToHashes(expected.TargetReceiptsProof), proof.TargetReceiptsProof)
}

// hexToHashes converts the fixture's hex sibling strings to the binding's
// bytes32[] representation.
func hexToHashes(hexes []string) [][32]byte {
	out := make([][32]byte, len(hexes))
	for i, h := range hexes {
		out[i] = common.HexToHash(h)
	}
	return out
}
