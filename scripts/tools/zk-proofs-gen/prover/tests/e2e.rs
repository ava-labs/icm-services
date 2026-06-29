//! (c) 2026, Ava Labs, Inc. All rights reserved.
//! See the file LICENSE for licensing terms.
//! SPDX-License-Identifier: LicenseRef-Ecosystem
//!
//! End-to-end test: feeds the icm-services fixtures through the SP1 guest via `execute`
//! (no proof generation) and asserts the committed public values are correct.
//! Fast — runs in CI with no prover network.
//!
//! THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

use alloy_sol_types::SolValue;
use merkle_sig_verification::{test_fixtures, PublicValues};
use sha2::{Digest, Sha256};
use sp1_sdk::blocking::{Prover, ProverClient};
use sp1_sdk::{include_elf, Elf, SP1Stdin};

const ELF: Elf = include_elf!("zk-valset-program");

#[test]
fn test_e2e_merkle_attestation() {
    sp1_sdk::utils::setup_logger();

    let attestation = hex::decode(test_fixtures::ATTESTATION_HEX).unwrap();
    let signed_data = hex::decode(test_fixtures::SIGNED_DATA_HEX).unwrap();
    let source_blockchain_id = test_fixtures::SOURCE_BLOCKCHAIN_ID;
    let root = test_fixtures::expected_root();
    let signed_data_hash: [u8; 32] = Sha256::digest(&signed_data).into();
    let signed_weight = test_fixtures::SIGNING_WEIGHT;

    // Inputs in the exact order the guest reads them.
    let mut stdin = SP1Stdin::new();
    stdin.write(&attestation);
    stdin.write(&signed_data);
    stdin.write(&source_blockchain_id);
    stdin.write(&root);
    stdin.write(&signed_data_hash);
    stdin.write(&signed_weight);

    let client = ProverClient::builder().cpu().build();
    let (public_values, _report) = client.execute(ELF, stdin).run().expect("execute failed");

    let pv = PublicValues::abi_decode(public_values.as_slice(), true).expect("decode public values");
    assert_eq!(pv.sourceBlockchainID.0, source_blockchain_id, "source blockchain ID mismatch");
    assert_eq!(pv.root.0, root, "root mismatch");
    assert_eq!(pv.messageHash.0, signed_data_hash, "message hash mismatch");
    assert_eq!(pv.signedWeight, signed_weight, "signed weight mismatch");
}
