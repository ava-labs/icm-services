//! Standalone tester for the attestation circuit, driven by the lib's icm-services fixtures.
//!   cargo run --release -- execute          # fast: run guest, print public values + cycles
//!   cargo run --release -- prove            # STARK proof + self-verify
//!   cargo run --release -- prove --groth16  # on-chain proof (slow, heavy) + bytes for the contract
//!   cargo run --release -- vkey             # program verification key (bytes32)
use alloy_sol_types::SolValue;
use clap::Parser;
use merkle_sig_verification::{test_fixtures, PublicValues};
use sha2::{Digest, Sha256};
use sp1_sdk::{include_elf, HashableKey, ProverClient, SP1Stdin};

pub const ELF: &[u8] = include_elf!("zk-valset-program");

#[derive(Parser)]
enum Cmd {
    Execute,
    Prove {
        #[arg(long)]
        groth16: bool,
    },
    Vkey,
}

// Inputs in the exact order the guest reads them.
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
    let client = ProverClient::from_env();

    match Cmd::parse() {
        Cmd::Vkey => {
            let (_, vk) = client.setup(ELF);
            println!("{}", vk.bytes32());
        }
        Cmd::Execute => {
            // No proving — runs the guest in the executor. Seconds, not minutes.
            let (pv, report) = client.execute(ELF, &stdin).run().expect("execute failed");
            println!("executed in {} cycles", report.total_instruction_count());
            check(pv.as_slice());
        }
        Cmd::Prove { groth16 } => {
            let (pk, vk) = client.setup(ELF);
            let proof = if groth16 {
                client.prove(&pk, &stdin).groth16().run()
            } else {
                client.prove(&pk, &stdin).run()
            }
            .expect("prove failed");

            client.verify(&proof, &vk).expect("self-verify failed");
            println!("✓ proof generated and self-verified");
            check(proof.public_values.as_slice());

            if groth16 {
                println!("  publicValues = 0x{}", hex::encode(proof.public_values.as_slice()));
                println!("  proofBytes   = 0x{}", hex::encode(proof.bytes()));
            }
        }
    }
}
