// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Based on code from: https://github.com/boundless-xyz/boundless-transceiver
// Modifications: Builds a State Manager for the Ethereum beacon chain, including storing beacon state roots, 
// execution state roots, and receipt roots, on top of the infrastructure provided by the boundless-transceiver repository.

pragma solidity ^0.8.30;

struct Checkpoint {
    uint64 epoch;
    bytes32 root; // beacon block root for epoch boundary block
}

struct ConsensusState {
    Checkpoint currentJustifiedCheckpoint;
    Checkpoint finalizedCheckpoint;
}
