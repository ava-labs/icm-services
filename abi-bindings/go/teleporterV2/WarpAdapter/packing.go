// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warpadapter

import (
	"fmt"
	"math"
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/libevm/common"
)

// PackReceiveCrossChainMessage packs a call to receiveCrossChainMessage for the WarpAdapter
// verification path. The attestation is the index of the signed Warp message in the
// transaction's predicate access list, ABI-encoded as the uint32 that WarpAdapter.sol decodes it
// as before passing it to getVerifiedWarpMessage.
func PackReceiveCrossChainMessage(
	teleporterMessage teleportermessengerv2.TeleporterMessageV2,
	sourceBlockChainID ids.ID,
	messageIndex uint64,
	relayerRewardAddress common.Address,
) ([]byte, error) {
	// WarpAdapter decodes the attestation as a uint32, so a larger index can never be delivered.
	// Reject it here rather than packing calldata that is guaranteed to revert on-chain.
	if messageIndex > math.MaxUint32 {
		return nil, fmt.Errorf(
			"warp message index %d exceeds the uint32 attestation WarpAdapter decodes", messageIndex,
		)
	}

	// Encode the full unsigned value. Converting through int64 would turn any index with the top
	// bit set into a negative number, and FillBytes would then write its absolute value.
	attestation := make([]byte, 32)
	new(big.Int).SetUint64(messageIndex).FillBytes(attestation)

	return teleportermessengerv2.PackReceiveCrossChainMessageV2(
		teleporterMessage,
		sourceBlockChainID,
		attestation,
		relayerRewardAddress,
	)
}
