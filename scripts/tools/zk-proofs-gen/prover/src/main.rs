//! (c) 2026, Ava Labs, Inc. All rights reserved.
//! See the file LICENSE for licensing terms.
//! SPDX-License-Identifier: LicenseRef-Ecosystem
//!
//! CLI for the validator-set attestation circuit. Set SP1_PROVER=mock|cpu|network to select
//! the proving backend. Generates a Groth16 proof over CLI-provided inputs (`prove`) or the
//! built-in icm-services test fixture (`prove-fixtures`), prints the program vkey (`vkey`),
//! and inspects a saved proof (`inspect`).
//!
//! THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

use std::{sync::Arc, time::Duration};

use alloy_sol_types::SolValue;
use clap::Parser;
use indicatif::{ProgressBar, ProgressStyle};
use merkle_sig_verification::{test_fixtures, PublicValues};
use sha2::{Digest, Sha256};
use sp1_sdk::{
    include_elf, utils, Elf, HashableKey, ProveRequest, Prover, ProverClient, ProvingKey,
    SP1ProofWithPublicValues, SP1Stdin,
};

pub const ELF: Elf = include_elf!("zk-valset-program");

#[derive(clap::Parser)]
#[command(name = "zk-valset-prover")]
struct Cli {
    #[command(subcommand)]
    command: Command,
}

#[derive(clap::Subcommand)]
enum Command {
    /// Generate a proof for a validator-set attestation from CLI-provided inputs.
    Prove {
        #[arg(
            long,
            value_name = "HEX",
            help = "Hex-encoded raw ValidatorSetMerkleAttestation bytes."
        )]
        attestation: String,
        #[arg(
            long,
            value_name = "HEX",
            help = "Hex-encoded signed data (the message attested to by validators)."
        )]
        signed_data: String,
        #[arg(long, value_name = "HEX", help = "Hex-encoded 32-byte source blockchain ID.")]
        source_blockchain_id: String,
        #[arg(
            long,
            value_name = "HEX",
            help = "Hex-encoded expected 32-byte Merkle root of the validator set."
        )]
        root: String,
        #[arg(
            long,
            value_name = "WEIGHT",
            help = "Total signing weight of the attestation's signers."
        )]
        signed_weight: u64,
        #[arg(
            long,
            value_name = "PATH",
            help = "Where to save the {vkey, publicValues, proof, signedData} JSON fixture."
        )]
        out: Option<String>,
    },
    /// Generate a proof using the built-in icm-services test fixtures.
    ProveFixtures {
        #[arg(long, value_name = "PATH", help = "Where to save the JSON fixture.")]
        out: Option<String>,
    },
    /// Print the SP1 program verification key hash (needed in the on-chain contract).
    Vkey,
    /// Print the public values and proof bytes of a saved fixture (for debugging).
    Inspect {
        #[arg(long, value_name = "PATH")]
        proof_file: String,
    },
}

// The witness inputs for one proof, plus the signed data recorded in the output fixture.
struct ProofInputs {
    attestation: Vec<u8>,
    signed_data: Vec<u8>,
    source_blockchain_id: [u8; 32],
    root: [u8; 32],
    signed_weight: u64,
}

impl ProofInputs {
    // Builds the guest's input witness. Write order must match the guest's io::read order exactly.
    fn stdin(&self) -> SP1Stdin {
        let signed_data_hash: [u8; 32] = Sha256::digest(&self.signed_data).into();
        let mut stdin = SP1Stdin::new();
        stdin.write(&self.attestation);
        stdin.write(&self.signed_data);
        stdin.write(&self.source_blockchain_id);
        stdin.write(&self.root);
        stdin.write(&signed_data_hash);
        stdin.write(&self.signed_weight);
        stdin
    }
}

// The built-in icm-services test fixture inputs.
fn fixture_inputs() -> ProofInputs {
    ProofInputs {
        attestation: hex::decode(test_fixtures::ATTESTATION_HEX).expect("attestation hex"),
        signed_data: hex::decode(test_fixtures::SIGNED_DATA_HEX).expect("signed_data hex"),
        source_blockchain_id: test_fixtures::SOURCE_BLOCKCHAIN_ID,
        root: test_fixtures::expected_root(),
        signed_weight: test_fixtures::SIGNING_WEIGHT,
    }
}

fn hex32(s: &str, what: &str) -> [u8; 32] {
    hex::decode(s)
        .unwrap_or_else(|_| panic!("--{what} must be valid hex"))
        .try_into()
        .unwrap_or_else(|_| panic!("--{what} must decode to exactly 32 bytes"))
}

// Writes the {vkey, publicValues, proof, signedData} fixture the Go e2e / forge test loads.
fn write_fixture(path: &str, vkey: String, proof: &SP1ProofWithPublicValues, signed_data: &[u8]) {
    let fixture = serde_json::json!({
        "vkey":         vkey,
        "publicValues": format!("0x{}", hex::encode(proof.public_values.as_slice())),
        "proof":        format!("0x{}", hex::encode(proof.bytes())),
        "signedData":   format!("0x{}", hex::encode(signed_data)),
    });
    if let Some(parent) = std::path::Path::new(path).parent() {
        std::fs::create_dir_all(parent).expect("create fixture dir");
    }
    std::fs::write(path, serde_json::to_string_pretty(&fixture).unwrap()).expect("write fixture");
    println!("wrote fixture to {path}");
}

fn print_public_values(pv_bytes: &[u8]) {
    let pv = PublicValues::abi_decode(pv_bytes, true).expect("decode public values");
    println!("  sourceBlockchainID = 0x{}", hex::encode(pv.sourceBlockchainID.0));
    println!("  root               = 0x{}", hex::encode(pv.root.0));
    println!("  messageHash        = 0x{}", hex::encode(pv.messageHash.0));
    println!("  signedWeight       = {}", pv.signedWeight);
}

// Sets up the program, generates a Groth16 proof (with a spinner), self-verifies it, prints the
// public values, and optionally writes the JSON fixture. The prover backend is selected by
// SP1_PROVER (mock|cpu|network) via ProverClient::from_env().
async fn prove_and_output(inputs: ProofInputs, out: Option<String>) {
    let client = ProverClient::from_env().await;
    let pk = client.setup(ELF).await.expect("setup failed");

    let spinner = Arc::new(ProgressBar::new_spinner());
    spinner.set_style(
        ProgressStyle::default_spinner()
            .tick_strings(&["⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"])
            .template("{spinner:.cyan} {msg}")
            .unwrap(),
    );
    spinner.set_message("Generating Groth16 proof...");
    let tick = tokio::spawn({
        let spinner = spinner.clone();
        async move {
            loop {
                spinner.tick();
                tokio::time::sleep(Duration::from_millis(80)).await;
            }
        }
    });

    let proof =
        client.prove(&pk, inputs.stdin()).groth16().await.expect("Groth16 proof generation failed");

    tick.abort();
    spinner.finish_with_message(format!(
        "Proof generated in {:.1}s.",
        spinner.elapsed().as_secs_f64()
    ));

    client.verify(&proof, pk.verifying_key(), None).expect("self-verify failed");
    println!("✓ proof generated and self-verified");
    print_public_values(proof.public_values.as_slice());

    if let Some(path) = out {
        write_fixture(&path, pk.verifying_key().bytes32(), &proof, &inputs.signed_data);
    }
}

#[tokio::main]
async fn main() {
    utils::setup_logger();

    match Cli::parse().command {
        Command::Vkey => {
            let client = ProverClient::from_env().await;
            let pk = client.setup(ELF).await.expect("setup failed");
            println!("{}", pk.verifying_key().bytes32());
        }
        Command::Inspect { proof_file } => {
            let proof = SP1ProofWithPublicValues::load(&proof_file).expect("load proof");
            println!("proof.bytes() hex: 0x{}", hex::encode(proof.bytes()));
            print_public_values(proof.public_values.as_slice());
        }
        Command::ProveFixtures { out } => {
            prove_and_output(fixture_inputs(), out).await;
        }
        Command::Prove {
            attestation,
            signed_data,
            source_blockchain_id,
            root,
            signed_weight,
            out,
        } => {
            let inputs = ProofInputs {
                attestation: hex::decode(attestation).expect("--attestation must be valid hex"),
                signed_data: hex::decode(signed_data).expect("--signed-data must be valid hex"),
                source_blockchain_id: hex32(&source_blockchain_id, "source-blockchain-id"),
                root: hex32(&root, "root"),
                signed_weight,
            };
            prove_and_output(inputs, out).await;
        }
    }
}
