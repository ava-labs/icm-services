#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

SCRIPT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
REPO_ROOT=$(cd "$SCRIPT_DIR/.." && pwd)
ICM_MACROS_DIR="$REPO_ROOT/icm-macros"
INSTALL_DIR="$HOME/.local/bin"

cargo build --release --locked --manifest-path "$ICM_MACROS_DIR/Cargo.toml"

mkdir -p "$INSTALL_DIR"
cp "$ICM_MACROS_DIR/target/release/icm-macros" "$INSTALL_DIR/reforge"
chmod +x "$INSTALL_DIR/reforge"

export PATH="$INSTALL_DIR:$PATH"