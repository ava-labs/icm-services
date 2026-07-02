//! (c) 2026, Ava Labs, Inc. All rights reserved.
//! See the file LICENSE for licensing terms.
//! SPDX-License-Identifier: LicenseRef-Ecosystem
//!
//! SP1 guest: verifies a validator set Merkle attestation in-circuit and commits
//! ABI-encoded public values for the on-chain ZKValidatorSetRegistry.
//!
//! THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

#![no_main]
sp1_zkvm::entrypoint!(main);

use alloy_sol_types::SolValue;
use merkle_sig_verification::{verify, PublicValues, ValidatorSetMerkleAttestation};

pub fn main() {
    // Read inputs in the exact order the prover writes them.
    let attestation = sp1_zkvm::io::read::<ValidatorSetMerkleAttestation>();
    let message = sp1_zkvm::io::read::<Vec<u8>>();
    let source_blockchain_id = sp1_zkvm::io::read::<[u8; 32]>();
    let root = sp1_zkvm::io::read::<[u8; 32]>();
    let message_hash = sp1_zkvm::io::read::<[u8; 32]>();
    let signed_weight = sp1_zkvm::io::read::<u64>();

    // Reject (no proof) if the attestation does not verify.
    verify(&attestation, &message, &root, &message_hash, signed_weight)
        .expect("attestation verification failed");

    let pv = PublicValues {
        sourceBlockchainID: source_blockchain_id.into(),
        root: root.into(),
        messageHash: message_hash.into(),
        signedWeight: signed_weight,
    };
    sp1_zkvm::io::commit_slice(&pv.abi_encode());
}
