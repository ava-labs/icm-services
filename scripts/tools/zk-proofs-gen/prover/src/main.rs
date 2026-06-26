//! (c) 2026, Ava Labs, Inc. All rights reserved.
//! See the file LICENSE for licensing terms.
//! SPDX-License-Identifier: LicenseRef-Ecosystem

//! CLI for the validator-set attestation circuit: run the guest (execute), generate and
//! self-verify proofs (prove, optionally Groth16 + a JSON fixture via --out), and print the
//! program verification key (vkey). Runs against the icm-services test fixture.
//!
//! This is example, un-audited code. Do not use in production.
// 
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
        /// Write a {vkey, publicValues, proof} JSON fixture to this path (use with --groth16).
        #[arg(long)]
        out: Option<String>,
    },
    Vkey,
}


// Builds the guest's input witness from the icm-services test fixture.
//
// The five values are written into SP1Stdin in the exact order the guest reads
// them with io::read — the stream is positional, so this order must stay in sync
// with sp1-program/src/main.rs or inputs deserialize into the wrong variables.
// All five are private witness inputs; the guest commits the public values separately.
fn fixture_stdin() -> SP1Stdin {
    let attestation = hex::decode(test_fixtures::ATTESTATION_HEX).expect("attestation hex");
    let signed_data = hex::decode(test_fixtures::SIGNED_DATA_HEX).expect("signed_data hex");
    let root = test_fixtures::expected_root();
    let total_weight = test_fixtures::TOTAL_WEIGHT;
    let signed_data_hash: [u8; 32] = Sha256::digest(&signed_data).into();

    let mut stdin = SP1Stdin::new();
    stdin.write(&attestation); // Vec<u8>; parsed by ValidatorSetMerkleAttestation::deserialize
    stdin.write(&signed_data); // Vec<u8>
    stdin.write(&root); // [u8; 32]
    stdin.write(&total_weight); // u64
    stdin.write(&signed_data_hash); // [u8; 32]
    stdin
}

fn check(pv_bytes: &[u8]) {
    let pv = PublicValues::abi_decode(pv_bytes, true).expect("decode public values");
    println!("  root          = 0x{}", hex::encode(pv.root.0));
    println!("  totalWeight   = {}", pv.totalWeight);
    println!("  quorumReached = {}", pv.quorumReached);
    println!("  messageHash   = 0x{}", hex::encode(pv.messageHash.0));
    assert_eq!(pv.root.0, test_fixtures::expected_root(), "root mismatch");
    assert!(pv.quorumReached, "expected quorum");
    println!("  ✓ public values match the fixture");
}

fn main() {
    sp1_sdk::utils::setup_logger();
    let stdin = fixture_stdin();
    let client = ProverClient::builder().cpu().build();

    match Cmd::parse() {
        // Prints the guest program's verification key.
        Cmd::Vkey => {
            let pk = client.setup(ELF).expect("setup failed");
            println!("{}", pk.verifying_key().bytes32());
        }
        // Runs the guest program in the zkVM executor without proving. 
        Cmd::Execute => {
            let (pv, report) = client.execute(ELF, stdin).run().expect("execute failed");
            println!("executed in {} cycles", report.total_instruction_count());
            check(pv.as_slice());
        }
        // Generates a proof in the zkVM and verifies it. Optionally writes out a JSON fixture. 
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

            // Write out the static fixture
            if let Some(path) = out {
                if !groth16 {
                    eprintln!("warning: fixture written from a non-groth16 proof is not on-chain verifiable; pass --groth16");
                }
                let fixture = serde_json::json!({
                    "vkey":         pk.verifying_key().bytes32(),
                    "publicValues": public_values_hex,
                    "proof":        proof_hex,
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
