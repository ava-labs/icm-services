#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -euo pipefail

# foundryup uses XDG_CONFIG_HOME as the root of the install.
# This can vary for different environments, so it is set to $HOME for consistency.
export XDG_CONFIG_HOME=$HOME

FOUNDRY_VERSION=v1.0.0
FOUNDRY_DIR="${FOUNDRY_DIR:-$HOME/.foundry}"
FOUNDRY_BIN_DIR="$FOUNDRY_DIR/bin"

# Fetch the foundryup launcher directly from the pinned tag instead of running
# the upstream install script. The upstream install script hardcodes a download
# of foundryup from the master branch, so running it executes unpinned code.
# Everything else the upstream script does (creating directories, editing shell
# profiles) is either done below or intentionally left to the user.
FOUNDRYUP_URL="https://raw.githubusercontent.com/foundry-rs/foundry/${FOUNDRY_VERSION}/foundryup/foundryup"

echo "Installing foundryup ${FOUNDRY_VERSION} to ${FOUNDRY_BIN_DIR}"
mkdir -p "$FOUNDRY_BIN_DIR"
curl --proto '=https' --tlsv1.2 -sSfL "$FOUNDRYUP_URL" -o "$FOUNDRY_BIN_DIR/foundryup"
chmod +x "$FOUNDRY_BIN_DIR/foundryup"

# Remember whether the caller's PATH already covers the install directory so we
# can print a hint at the end. The upstream install script used to append the
# directory to the user's shell profile; this script does not modify profiles.
PATH_HAS_FOUNDRY=false
if [[ ":$PATH:" == *":${FOUNDRY_BIN_DIR}:"* ]]; then
  PATH_HAS_FOUNDRY=true
fi
export PATH=$PATH:$FOUNDRY_BIN_DIR

foundryup --install "${FOUNDRY_VERSION}"

# Verify that foundryup actually installed the pinned version rather than
# silently falling back to a different release.
INSTALLED_VERSION=$(forge --version | head -n 1)
if [[ "$INSTALLED_VERSION" != *"${FOUNDRY_VERSION#v}"* ]]; then
  echo "error: expected forge ${FOUNDRY_VERSION}, but installed: ${INSTALLED_VERSION}" >&2
  exit 1
fi
echo "Installed ${INSTALLED_VERSION}"

if [[ "$PATH_HAS_FOUNDRY" == false ]]; then
  echo "Add ${FOUNDRY_BIN_DIR} to your PATH to use forge, cast, and anvil."
fi
