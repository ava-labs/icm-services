//! (c) 2026, Ava Labs, Inc. All rights reserved.
//! See the file LICENSE for licensing terms.
//! SPDX-License-Identifier: LicenseRef-Ecosystem
//!
//! CLI for the validator-set attestation circuit: run the guest (execute), generate and
//! self-verify proofs (prove, optionally Groth16 + a JSON fixture via --out), and print the
//! program verification key (vkey). Runs against the icm-services test fixture.
//!
//! THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

use alloy_sol_types::SolValue;
use clap::Parser;
use merkle_sig_verification::{test_fixtures, PublicValues};
use sha2::{Digest, Sha256};
use sp1_sdk::blocking::{Prover, ProveRequest, ProverClient};
use sp1_sdk::{include_elf, Elf, HashableKey, ProvingKey, SP1Stdin};

pub const ELF: Elf = include_elf!("zk-valset-program");

#[derive(Parser)]
enum Cmd {
    Execute,
    Prove {
        #[arg(long)]
        groth16: bool,
        /// Write a {vkey, publicValues, proof, signedData} JSON fixture to this path (use with --groth16).
        #[arg(long)]
        out: Option<String>,
    },
    Vkey,
}

// Packs the icm-services test fixture into the guest's input witness and 
// returns it along with the signed data the proof commits to. The order 
// of items placed into stdin must match the guest's io::read order exactly. 
fn generate_fixture_stdin_from_test() -> (SP1Stdin, Vec<u8>) {
    let attestation = hex::decode(test_fixtures::ATTESTATION_HEX).expect("attestation hex");
    let signed_data = hex::decode(test_fixtures::SIGNED_DATA_HEX).expect("signed_data hex");
    let source_blockchain_id = test_fixtures::SOURCE_BLOCKCHAIN_ID;
    let root = test_fixtures::expected_root();
    let signed_data_hash: [u8; 32] = Sha256::digest(&signed_data).into();
    let signed_weight = test_fixtures::SIGNING_WEIGHT;

    let mut stdin = SP1Stdin::new();
    stdin.write(&attestation);
    stdin.write(&signed_data);
    stdin.write(&source_blockchain_id);
    stdin.write(&root); 
    stdin.write(&signed_data_hash); 
    stdin.write(&signed_weight); 
    (stdin, signed_data)
}

fn check(pv_bytes: &[u8]) {
    let pv = PublicValues::abi_decode(pv_bytes, true).expect("decode public values");
    println!("  sourceBlockchainID = 0x{}", hex::encode(pv.sourceBlockchainID.0));
    println!("  root               = 0x{}", hex::encode(pv.root.0));
    println!("  messageHash        = 0x{}", hex::encode(pv.messageHash.0));
    println!("  signedWeight       = {}", pv.signedWeight);
    assert_eq!(pv.root.0, test_fixtures::expected_root(), "root mismatch");
    assert_eq!(pv.signedWeight, test_fixtures::SIGNING_WEIGHT, "signed weight mismatch");
    println!("  ✓ public values match the fixture");
}

fn main() {
    sp1_sdk::utils::setup_logger();
    let (stdin, signed_data) = generate_fixture_stdin_from_test(); 
    let client = ProverClient::builder().cpu().build();

    match Cmd::parse() {
        Cmd::Vkey => {
            let pk = client.setup(ELF).expect("setup failed");
            println!("{}", pk.verifying_key().bytes32());
        }
        Cmd::Execute => {
            let (pv, report) = client.execute(ELF, stdin).run().expect("execute failed");
            println!("executed in {} cycles", report.total_instruction_count());
            check(pv.as_slice());
        }
        Cmd::Prove { groth16, out } => {
            let pk = client.setup(ELF).expect("setup failed");
            let proof = if groth16 {
                client.prove(&pk, stdin).groth16().run()
            } else {
                client.prove(&pk, stdin).run()
            }
            .expect("prove failed");

            client.verify(&proof, pk.verifying_key(), None).expect("self-verify failed");
            println!("✓ proof generated and self-verified");
            check(proof.public_values.as_slice());

            let public_values_hex = format!("0x{}", hex::encode(proof.public_values.as_slice()));
            let proof_hex = format!("0x{}", hex::encode(proof.bytes()));

            if groth16 {
                println!("  publicValues = {public_values_hex}");
                println!("  proofBytes   = {proof_hex}");
            }

            // Write the static fixture the Go e2e / forge test will load.
            if let Some(path) = out {
                if !groth16 {
                    eprintln!("warning: fixture written from a non-groth16 proof is not on-chain verifiable; pass --groth16");
                }
                let fixture = serde_json::json!({
                    "vkey":         pk.verifying_key().bytes32(),
                    "publicValues": public_values_hex,
                    "proof":        proof_hex,
                    "signedData":   format!("0x{}", hex::encode(signed_data)),
                });
                if let Some(parent) = std::path::Path::new(&path).parent() {
                    std::fs::create_dir_all(parent).expect("create fixture dir");
                }
                std::fs::write(&path, serde_json::to_string_pretty(&fixture).unwrap())
                    .expect("write fixture");
                println!("wrote fixture to {path}");
            }
        }
    }
}
