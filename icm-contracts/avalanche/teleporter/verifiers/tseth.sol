// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.30;

struct Checkpoint {
    uint64 epoch;
    bytes32 root; // beacon block root for epoch boundary block
}

struct ConsensusState {
    Checkpoint currentJustifiedCheckpoint;
    Checkpoint finalizedCheckpoint;
}
