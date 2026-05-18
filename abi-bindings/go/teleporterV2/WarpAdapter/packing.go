// (c) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warpadapter

import (
	"math/big"

	"github.com/ava-labs/avalanchego/ids"
	teleportermessengerv2 "github.com/ava-labs/icm-services/abi-bindings/go/TeleporterMessengerV2"
	"github.com/ava-labs/libevm/common"
)

// This is specific to the warp adapter
func PackReceiveCrossChainMessage(
	teleporterMessage teleportermessengerv2.TeleporterMessageV2,
	sourceBlockChainID ids.ID,
	messageIndex uint64,
	relayerRewardAddress common.Address,
) ([]byte, error) {
	attestation := make([]byte, 32)
	big.NewInt(int64(messageIndex)).FillBytes(attestation)

	return teleportermessengerv2.PackReceiveCrossChainMessageV2(
		teleporterMessage,
		sourceBlockChainID,
		attestation,
		relayerRewardAddress,
	)
}
