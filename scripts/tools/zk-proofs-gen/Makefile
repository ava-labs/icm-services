NIGHTLY      := +$(shell cat nightly-toolchain)

# sp1-program is a zkVM guest (#![no_main]) and cannot compile to the host target
ZKVM_EXCLUDE := --exclude zk-valset-program

# Clippy needs an additional exclusion beyond the zkVM guest:
#   zk-valset-prover — build.rs invokes sp1-build to compile the guest ELF, which
#                      breaks when clippy-driver is injected as a rustc wrapper
CLIPPY_EXCLUDE := $(ZKVM_EXCLUDE) \
	--exclude zk-valset-prover

# ── Formatting ───────────────────────────────────────────────────────────────

fmt:
	cargo $(NIGHTLY) fmt --all
	taplo fmt
	cargo sort --workspace --no-format

fmt-check:
	cargo $(NIGHTLY) fmt --all -- --check
	taplo fmt --check
	cargo sort --workspace --no-format --check

# ── Linting ──────────────────────────────────────────────────────────────────

clippy:
	cargo $(NIGHTLY) clippy --workspace $(CLIPPY_EXCLUDE) --all-targets --all-features -- -D warnings

# ── Build ─────────────────────────────────────────────────────────────────────

check:
	cargo check --workspace $(ZKVM_EXCLUDE) --all-targets --all-features

build:
	cargo build --workspace $(ZKVM_EXCLUDE)

build-release:
	cargo build --release -p zk-valset-prover

# ── Tests ─────────────────────────────────────────────────────────────────────

# Run all test suites
test: test-unit test-integration test-doc

# Unit tests: #[test] in library and binary source files
test-unit:
	cargo test --release --workspace $(ZKVM_EXCLUDE) --lib --bins

# Integration tests: tests/*.rs harness files (e.g. prover/tests/e2e.rs)
test-integration:
	cargo test --release --workspace $(ZKVM_EXCLUDE) --tests

# Doc tests: /// # Examples blocks
test-doc:
	cargo test --release --workspace $(ZKVM_EXCLUDE) --doc

# ── Fixtures ─────────────────────────────────────────────────────────────────

# Regenerate the fixture input constants (SIGNED_DATA_HEX / ATTESTATION_HEX):
# signs the address(0) message with the synthetic keys, self-verifies, and
# prints the new constants to paste into lib::test_fixtures.
gen-fixture-inputs:
	cargo test --release -p merkle-sig-verification gen_no_sender_fixtures -- --nocapture --ignored

# ── CLI ───────────────────────────────────────────────────────────────────────

# Run the prover CLI. Pass subcommand and flags via ARGS, e.g.:
#   make prove ARGS="prove-fixtures --out ../../../tests/testdata/zk-groth16-fixture.json"
#   make prove ARGS="vkey"
prove:
	cargo run --release -p zk-valset-prover -- $(ARGS)

.PHONY: fmt fmt-check clippy check build build-release test test-unit test-integration test-doc gen-fixture-inputs prove
