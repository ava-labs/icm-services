// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {IMessageVerifier} from "./ITeleporterMessengerV2.sol";
import {TeleporterICMMessage} from "./TeleporterMessageV2.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

// An overly complicated implementation of a configurable attestation verifier.
// It supports up to eight IMessageVerifiers. The power set of eight IMessageVerifiers
// is enumerable via a uint256. Thus a uint256 is used as a config to indicate
// which subsets are considered accepting.
//
// example: For 3 adapters with a 2 out of 3 acceptance scheme, the valid results
// are 011, 101, 110, and 111. These correspond to 3, 5, 6, and 7, so the config
// should be 11101000 in binary, i.e. 232 in decimal.
//
// N.B. if your configuration is not even, you've likely done something horribly
// wrong as it means you accept if all IMessageVerifiers fail.
library ConfigurableAttestation {
    // Returns the CONFIG value for an "at least m of n adapters must succeed" policy.
    function mOfNConfig(
        uint8 m,
        uint8 n
    ) internal pure returns (uint256 configuration) {
        require(n <= 8, "n must be <= 8");
        require(m <= n, "m must be <= n");
        require(m > 0, "m must be > 0");
        uint256 subsets = uint256(1) << n;
        for (uint256 r = 0; r < subsets;) {
            if (_popcount(uint8(r)) >= m) {
                configuration |= uint256(1) << r;
            }
            unchecked {
                ++r;
            }
        }
    }

    // Returns the CONFIG value for an arbitrary acceptance policy, specified as an explicit
    // list of accepting subsets. Each element of `subsets` is a bitset of adapter indices
    // that, when matched, should be considered a passing result.
    //
    // example: For a 2-of-3 policy, pass [3, 5, 6, 7] (i.e. 011, 101, 110, 111 in binary).
    function subsetsToConfig(
        uint8[] memory subsets
    ) internal pure returns (uint256 configuration) {
        for (uint256 i = 0; i < subsets.length;) {
            configuration |= uint256(1) << subsets[i];
            unchecked {
                ++i;
            }
        }
    }

    function _popcount(
        uint8 x
    ) private pure returns (uint8 count) {
        while (x != 0) {
            count += x & 1;
            x >>= 1;
        }
    }
}

contract ConfigurableAttestationVerifier is IMessageVerifier {
    uint256 public immutable config;
    IMessageVerifier[] private _verifiers;

    constructor(
        uint256 config_,
        address[] memory verifiers_
    ) {
        config = config_;
        require(verifiers_.length <= 8, "Cannot configure more than eight adapters");
        // the configuration should not include subsets for sets with cardinality greater
        // than 2^length
        require(
            config_ < uint256(1) << (uint256(1) << verifiers_.length), "Malformed configuration"
        );
        for (uint256 i = 0; i < verifiers_.length;) {
            _verifiers.push(IMessageVerifier(verifiers_[i]));
            unchecked {
                ++i;
            }
        }
    }

    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external returns (bool) {
        // compute the subset of accepting IMessageVerifiers as bitset of 8 elements
        uint8 result = 0;
        for (uint8 i = 0; i < _verifiers.length;) {
            if (_verifiers[i].verifyMessage(message)) {
                result |= uint8(1) << i;
            }
            unchecked {
                ++i;
            }
        }
        // convert this subset to its corresponding uint256 flag
        uint256 resultMask = uint256(1) << result;
        // check if this flag is present in the config
        return config & resultMask != 0;
    }
}
