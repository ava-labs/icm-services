// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"cmp"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"slices"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contracts/warp"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	predicateutils "github.com/ava-labs/avalanchego/vms/evm/predicate"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

var (
	ZeroAddress = common.Address{}

	// Errors
	ErrNilInput = errors.New("nil input")
	ErrTooLarge = errors.New("exceeds uint256 maximum value")
	// Generic private key parsing error used to obfuscate the actual error
	ErrInvalidPrivateKeyHex = errors.New("invalid account private key hex string")
)

const (
	DefaultRPCTimeout = 5 * time.Second

	// Re-exposing DefaultAppRequestTimeout for use by message creators to set deadlines
	DefaultAppRequestTimeout = constants.DefaultNetworkMaximumTimeout

	// Maximum amount of time to spend waiting for a connection to a quorum of validators for
	// a given subnetID.
	ConnectToValidatorsTimeout = 30 * time.Second

	// We make at most 3 RPC calls, one app request, and one round of retries to connect to validators.
	DefaultCreateSignedMessageTimeout = 3*DefaultRPCTimeout + DefaultAppRequestTimeout + ConnectToValidatorsTimeout
)

//
// ICM Utils
//

// CheckStakeWeightExceedsThreshold returns true if the accumulated signature weight is at
// least [quorumNum]/[quorumDen] of [totalWeight].
func CheckStakeWeightExceedsThreshold(
	accumulatedSignatureWeight *big.Int,
	totalWeight uint64,
	quorumNumerator uint64,
) bool {
	if accumulatedSignatureWeight == nil {
		return false
	}
	return accumulatedSignatureWeight.Cmp(RequiredSignatureWeight(totalWeight, quorumNumerator)) >= 0
}

// RequiredSignatureWeight returns the minimum signature weight that meets the
// [quorumNumerator]/WarpQuorumDenominator threshold of [totalWeight], i.e.
// ceil(totalWeight * quorumNumerator / WarpQuorumDenominator).
func RequiredSignatureWeight(totalWeight, quorumNumerator uint64) *big.Int {
	quorumDen := new(big.Int).SetUint64(warp.WarpQuorumDenominator)
	required := new(big.Int).Mul(
		new(big.Int).SetUint64(totalWeight),
		new(big.Int).SetUint64(quorumNumerator),
	)
	// ceil(required / quorumDen) == (required + quorumDen - 1) / quorumDen
	required.Add(required, new(big.Int).Sub(quorumDen, big.NewInt(1)))
	return required.Quo(required, quorumDen)
}

func SignedWarpMessageToAccessList(signedMessage *avalancheWarp.Message) types.AccessList {
	// Construct the actual transaction to broadcast on the destination chain
	// Create predicate from the signed warp message
	predicate := predicateutils.New(signedMessage.Bytes())

	// Create access list with the predicate for the warp precompile
	return types.AccessList{
		{
			Address:     warp.ContractAddress,
			StorageKeys: predicate,
		},
	}
}

// CalculateQuorumPercentageBuffer calculates the quorum percentage buffer based on the required quorum percentage
// and the desired quorum percentage buffer.
func CalculateQuorumPercentageBuffer(
	requiredQuorumPercentage uint64,
	desiredQuorumPercentageBuffer uint64,
) uint64 {
	if requiredQuorumPercentage >= 100 {
		return 0
	}
	if requiredQuorumPercentage+desiredQuorumPercentageBuffer > 100 {
		return 100 - requiredQuorumPercentage
	}
	return desiredQuorumPercentageBuffer
}

// SortByWeightDescending sorts [items] in place by descending weight, using [weight]
// to extract each item's weight. The sort is stable, so items of equal weight retain
// their original relative order.
func SortByWeightDescending[T any](items []T, weight func(T) uint64) {
	slices.SortStableFunc(items, func(a, b T) int {
		return cmp.Compare(weight(b), weight(a))
	})
}

//
// Generic Utils
//

func PrivateKeyToString(key *ecdsa.PrivateKey) string {
	// Use FillBytes so leading zeroes are not stripped.
	return hex.EncodeToString(key.D.FillBytes(make([]byte, 32)))
}

// SanitizeHexString removes the "0x" prefix from a hex string if it exists.
// Otherwise, returns the original string.
func SanitizeHexString(hex string) string {
	return strings.TrimPrefix(hex, "0x")
}

// StripFromString strips the input string starting from the first occurrence of the substring.
func StripFromString(input, substring string) string {
	index := strings.Index(input, substring)
	if index == -1 {
		// Substring not found, return the original string
		return input
	}

	// Strip the string starting from the found substring
	strippedString := input[:index]

	return strippedString
}

// Converts a '0x'-prefixed hex string or cb58-encoded string to an ID.
// Input length validation is handled by the ids package.
func HexOrCB58ToID(s string) (ids.ID, error) {
	if strings.HasPrefix(s, "0x") {
		bytes, err := hex.DecodeString(SanitizeHexString(s))
		if err != nil {
			return ids.ID{}, err
		}
		return ids.ToID(bytes)
	}
	return ids.FromString(s)
}

// IsEmptyOrZeroes returns true if the byte slice is empty or all zeroes
func IsEmptyOrZeroes(bytes []byte) bool {
	for _, b := range bytes {
		if b != 0 {
			return false
		}
	}
	return true
}
