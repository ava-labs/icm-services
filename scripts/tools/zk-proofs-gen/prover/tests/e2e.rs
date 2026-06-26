//! THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION. 

//! End-to-end test: feeds the icm-services fixtures through the SP1 guest via `execute`
//! (no proof generation) and asserts the committed public values are correct.
//! Fast — runs in CI with no prover network.
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
    let root = test_fixtures::expected_root();
    let total_weight = test_fixtures::TOTAL_WEIGHT;
    let signed_data_hash: [u8; 32] = Sha256::digest(&signed_data).into();

    let mut stdin = SP1Stdin::new();
    stdin.write(&attestation);
    stdin.write(&signed_data);
    stdin.write(&root);
    stdin.write(&total_weight);
    stdin.write(&signed_data_hash);

    let client = ProverClient::builder().cpu().build();
    let (public_values, _report) = client.execute(ELF, stdin).run().expect("execute failed");

    let pv = PublicValues::abi_decode(public_values.as_slice(), true).expect("decode public values");
    assert_eq!(pv.root.0, root, "root mismatch");
    assert_eq!(pv.totalWeight, total_weight, "total weight mismatch");
    assert_eq!(pv.messageHash.0, signed_data_hash, "message hash mismatch");
    assert!(pv.quorumReached, "signing weight should exceed 2/3 of total weight");
}
