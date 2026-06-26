# zk-proofs-gen

SP1 zkVM proof generation for Avalanche validator-set Merkle attestations.

This workspace produces a single zero-knowledge proof that a quorum of a committed
validator set signed an ICM message. It is intended to be used by the Go-side e2e tests
for the`ZKValidatorSetRegistry` contract. It replaces the on-chain Merkle multi-inclusion
proof and aggregate BLS signature check (as done by `MerkleValidatorSetRegistry`) with one
SP1 Groth16 proof verified by `ZKValidatorSetRegistry`. The validator set stays anchored
on-chain by a Merkle root commitment; the signers, the multiproof, and the aggregate
signature are private inputs to the proof and never appear in calldata.

## Layout

```
zk-proofs-gen/
├── lib/          # merkle-sig-verification: the verification logic, shared by guest + host
├── sp1-program/  # the SP1 guest program (compiled to RISC-V, runs in the zkVM)
└── prover/       # host CLI: execute / prove / vkey
```

- **`lib`** — parses a `ValidatorSetMerkleAttestation`, rebuilds the Merkle root, verifies
  the aggregate BLS12-381 signature, and checks the stake-weighted quorum. The same logic
  runs in the zkVM and on the host. Defines `PublicValues`, the ABI-encoded struct the guest
  commits and the Solidity contract decodes.
- **`sp1-program`** — the guest: reads the inputs, calls `lib::verify`, commits the
  ABI-encoded public values (`root`, `totalWeight`, `quorumReached`, `messageHash`).
- **`prover`** — the host CLI that runs the guest and generates proofs.

## Prerequisites

```bash
# Install system dependencies 
sudo apt-get update && sudo apt-get install -y \
  build-essential pkg-config libssl-dev clang cmake git curl protobuf-compiler

# Install Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
. "$HOME/.cargo/env"

# Install SP1 toolchain 
curl -L https://sp1up.succinct.xyz | bash
sp1up

# Docker — required by SP1's CPU Groth16 backend (the gnark wrap runs in a container)
sudo apt-get install -y docker.io
sudo systemctl enable --now docker
sudo usermod -aG docker $USER   # then re-login (or `newgrp docker`)
```

## Usage

Run from the workspace root. The first build compiles the guest ELF via `prover/build.rs`,
so no separate build step is needed.

```bash
# Run the guest on the test fixture WITHOUT proving. Fast, low memory.
# Validates the circuit (leaf hashing, multiproof, BLS, quorum, public values).
cargo run --release -p zk-valset-prover -- execute

# Print the program verification key (bytes32) for the ZKValidatorSetRegistry constructor.
cargo run --release -p zk-valset-prover -- vkey

# Generate a proof and self-verify it.
#   (default)  STARK proof
#   --groth16  on-chain-verifiable Groth16 proof
#   --out PATH write a {vkey, publicValues, proof} JSON fixture (use with --groth16)
cargo run --release -p zk-valset-prover -- prove
```

### Generating the on-chain proof fixture

The on-chain proof is **Groth16**. Generate it once and write the JSON fixture the Go e2e
and forge tests load:

```bash
cargo run --release -p zk-valset-prover -- prove --groth16 \
  --out ../../../test/testdata/zk-groth16-fixture.json
```

The `--out` path is relative to the current directory; from `scripts/tools/zk-proofs-gen`
that resolves to `icm-services/test/testdata/`. The fixture contains:

```json
{ "vkey": "0x...", "publicValues": "0x...", "proof": "0x..." }
```

Commit this file. Tests load it directly — they never re-prove. The `vkey` in the fixture
is the same value the contract constructor needs; regenerate the fixture whenever the guest
program changes so the vkey and proof stay in sync.

## Notes on proving
 
- `execute` and STARK `prove`: a modest machine works (~16-32 GB RAM).
- `prove --groth16`: needs Docker running (the gnark wrap runs in a container) and ~64 GB
  RAM (e.g. `r6i.2xlarge`). First run downloads ~6 GB of circuit artifacts to `~/.sp1`.
- The committed fixture is regenerated rarely — only when the circuit or its inputs change.


## Test fixture

The proof is generated over the built-in icm-services test vectors (`lib::test_fixtures`).
The message the proof commits to (`messageHash`) must match the message the consuming e2e
test reconstructs, or on-chain verification will reject it.

THIS IS AN EXAMPLE OF UN-AUDITED CODE. DO NOT USE IN PRODUCTION. 
