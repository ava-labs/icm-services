#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# The macro snapshot tests only diff generated text against the checked-in
# fixtures, so a fixture can drift into code that does not compile without any
# test noticing. Compile them here to close that gap.

set -euo pipefail

REPO_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

if ! command -v solc &> /dev/null; then
    echo "solc not found. See https://docs.soliditylang.org/en/latest/installing-solidity.html" >&2
    exit 1
fi

# Generated unpack code keeps many locals alive at once, which overflows the
# stack on the legacy pipeline, so compile everything via IR.
failed=0
shopt -s nullglob
for file in "$REPO_PATH"/icm-macros/testing/*/expected/*.sol; do
    if output=$(solc --via-ir --bin "$file" 2>&1); then
        echo "ok      ${file#"$REPO_PATH"/}"
    else
        echo "FAILED  ${file#"$REPO_PATH"/}" >&2
        echo "$output" >&2
        failed=1
    fi
done

if [ "$failed" -ne 0 ]; then
    echo "error: at least one macro fixture does not compile" >&2
    exit 1
fi
