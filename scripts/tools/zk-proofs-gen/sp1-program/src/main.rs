//! SP1 guest: verifies a validator set Merkle attestation in-circuit and commits
//! ABI-encoded public values for the on-chain ZKValidatorSetRegistry.

//! THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION. 

#![no_main]
sp1_zkvm::entrypoint!(main);

use alloy_sol_types::SolValue;
use merkle_sig_verification::{verify, PublicValues, ValidatorSetMerkleAttestation};

pub fn main() {
    let attestation = sp1_zkvm::io::read::<ValidatorSetMerkleAttestation>();
    let signed_data = sp1_zkvm::io::read::<Vec<u8>>();
    let root = sp1_zkvm::io::read::<[u8; 32]>();
    let total_weight = sp1_zkvm::io::read::<u64>();
    let signed_data_hash = sp1_zkvm::io::read::<[u8; 32]>();

    let quorum_reached = verify(
        &attestation,
        &signed_data,
        &root,
        total_weight,
        &signed_data_hash,
    )
    .expect("attestation verification failed");

    let pv = PublicValues {
        root: root.into(),
        totalWeight: total_weight,
        quorumReached: quorum_reached,
        messageHash: signed_data_hash.into(),
    };
    sp1_zkvm::io::commit_slice(&pv.abi_encode());
}
