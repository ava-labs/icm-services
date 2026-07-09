#!/usr/bin/env bash
# Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -o errexit
set -o nounset
set -o pipefail

BASE_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

binary_path="${1:-$BASE_PATH/build/solana-observer}"

echo "Building solana-observer at $binary_path"
go build -o "$binary_path" "$BASE_PATH/solana-observer/"
