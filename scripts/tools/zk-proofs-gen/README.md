# zk-proofs-gen

SP1 zkVM proof generation for Avalanche validator-set Merkle attestations.

This workspace produces a single zero-knowledge proof that a quorum of a committed
validator set signed an ICM message. It is intended to be used by the Go-side e2e tests
for the `ZKValidatorSetRegistry` contract. It replaces the on-chain Merkle multi-inclusion
proof and aggregate BLS signature check (as done by `MerkleValidatorSetRegistry`) with one
SP1 Groth16 proof verified by `ZKValidatorSetRegistry`. The validator set stays anchored
on-chain by a Merkle root commitment; the signers, the multiproof, and the aggregate
signature are private inputs to the proof and never appear in calldata.

## Layout

```
zk-proofs-gen/
├── lib/          # merkle-sig-verification: the verification logic, shared by guest + host
├── sp1-program/  # the SP1 guest program (compiled to RISC-V, runs in the zkVM)
└── prover/       # host CLI: prove / prove-fixtures / vkey / inspect
```

- **`lib`** — parses a `ValidatorSetMerkleAttestation`, rebuilds the Merkle root, verifies
  the aggregate BLS12-381 signature, and checks the committed signing weight. The same logic
  runs in the zkVM and on the host. Defines `PublicValues`, the ABI-encoded struct the guest
  commits and the Solidity contract decodes. Also provides test-only helpers to sign and
  serialize new fixture attestations with the synthetic validator keys.
- **`sp1-program`** — the guest: reads the inputs, calls `lib::verify`, commits the
  ABI-encoded public values (`sourceBlockchainID`, `root`, `messageHash`, `signedWeight`).
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

For formatting and linting (`make fmt`, `make clippy`), also install the pinned nightly
toolchain and the formatting tools:

```bash
rustup toolchain install $(cat nightly-toolchain)
cargo install taplo-cli cargo-sort
```

## Usage

Run from the workspace root. The first build compiles the guest ELF via `prover/build.rs`,
so no separate build step is needed.

The proving backend is selected with the `SP1_PROVER` environment variable:
`mock` (fast fake proof, plumbing check), `cpu` (real Groth16, needs Docker + RAM), or
`network` (the Succinct prover network).

```bash
# Print the program verification key (bytes32) for the ZKValidatorSetRegistry constructor.
SP1_PROVER=cpu cargo run --release -p zk-valset-prover -- vkey

# Generate a Groth16 proof over the built-in test fixture and self-verify it.
#   --out PATH  write the {vkey, publicValues, proof, signedData} JSON fixture
SP1_PROVER=cpu cargo run --release -p zk-valset-prover -- prove-fixtures

# Generate a Groth16 proof over arbitrary inputs.
SP1_PROVER=cpu cargo run --release -p zk-valset-prover -- prove \
  --attestation <HEX> \
  --signed-data <HEX> \
  --source-blockchain-id <HEX32> \
  --root <HEX32> \
  --signed-weight <WEIGHT> \
  --out fixture.json

# Print the proof bytes and public values of a saved proof (debugging).
cargo run --release -p zk-valset-prover -- inspect --proof-file <PATH>
```

A Makefile wraps the common tasks (run from the workspace root):

```bash
make fmt          # rustfmt (nightly, grouped imports) + taplo + cargo-sort
make fmt-check    # CI-style check, no changes
make clippy       # lints with -D warnings
make test         # unit + integration + doc tests
make prove ARGS="prove-fixtures --out ../../../tests/testdata/zk-groth16-fixture.json"
make gen-fixture-inputs   # regenerate SIGNED_DATA_HEX / ATTESTATION_HEX
```

### Generating the on-chain proof fixture

Generate the Groth16 proof once and write the JSON fixture the Go e2e and forge tests load:

```bash
SP1_PROVER=cpu cargo run --release -p zk-valset-prover -- prove-fixtures \
  --out ../../../tests/testdata/zk-groth16-fixture.json
```

The `--out` path is relative to the current directory; from `scripts/tools/zk-proofs-gen`
that resolves to `icm-services/tests/testdata/`. The fixture contains:

```json
{ "vkey": "0x...", "publicValues": "0x...", "proof": "0x...", "signedData": "0x..." }
```

`signedData` is the exact unsigned warp message the proof commits to
(`sha256(signedData) == publicValues.messageHash`); the Go e2e derives the message it
submits to the contract from it. Commit this file. Tests load it directly — they never
re-prove. The `vkey` in the fixture is the same value the contract constructor needs;
regenerate the fixture whenever the guest program changes so the vkey and proof stay in
sync.

## Notes on proving

- `SP1_PROVER=mock`: runs in seconds on any machine; use it to validate the circuit and
  CLI plumbing before an expensive run. The resulting proof is not on-chain verifiable.
- `SP1_PROVER=cpu`: needs Docker running (the gnark wrap runs in a container) and ~64 GB
  RAM (e.g. `r6i.2xlarge`). First run downloads ~6 GB of circuit artifacts to `~/.sp1`.
- The committed fixture is regenerated rarely — only when the circuit or its inputs change.

## Test fixture

The proof is generated over the built-in test vectors (`lib::test_fixtures`): a synthetic
4-validator set with known keys, and an aggregate BLS signature by three of them over a
fixed `address(0)` (validator-set update) warp message. The message the proof commits to
(`messageHash`) must match the message `ZKValidatorSetRegistry.verifyICMMessage`
reconstructs, or on-chain verification will reject it.

To regenerate the fixture inputs over a new message, use the generator test in `lib`,
which signs with the synthetic keys, serializes the attestation, self-verifies the pair,
and prints the new `SIGNED_DATA_HEX` / `ATTESTATION_HEX` constants to paste into
`lib::test_fixtures`:

```bash
cargo test --release -p merkle-sig-verification gen_no_sender_fixtures -- --nocapture --ignored
```

THIS IS AN EXAMPLE OF UN-AUDITED CODE. DO NOT USE IN PRODUCTION.
