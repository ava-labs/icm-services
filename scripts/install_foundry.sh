#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -euo pipefail

# foundryup uses XDG_CONFIG_HOME as the root of the install.
# This can vary for different environments, so it is set to $HOME for consistency.
export XDG_CONFIG_HOME="$HOME"

# The Foundry release to install. The three values below must be updated together.
# To bump the pin:
#   1. Set FOUNDRY_VERSION to the new release tag.
#   2. Set FOUNDRY_COMMIT to the commit that tag points to:
#        gh api repos/foundry-rs/foundry/git/ref/tags/<tag> --jq .object.sha
#   3. Set FOUNDRYUP_SHA256 to the checksum of the foundryup launcher at that commit:
#        curl -sSfL https://raw.githubusercontent.com/foundry-rs/foundry/<commit>/foundryup/foundryup | shasum -a 256
FOUNDRY_VERSION=v1.0.0
FOUNDRY_COMMIT=8692e926198056d0228c1e166b1b6c34a5bed66c
FOUNDRYUP_SHA256=980a7a4a7f6a453346191bbe5c03bb378a91c92b10573a86fd29ee6f4b7f5d35

FOUNDRY_DIR="${FOUNDRY_DIR:-$HOME/.foundry}"
FOUNDRY_BIN_DIR="$FOUNDRY_DIR/bin"

# Fetch the foundryup launcher directly from the pinned commit instead of running
# the upstream install script. The upstream install script hardcodes a download
# of foundryup from the master branch, so running it executes unpinned code.
# Everything else the upstream script does (creating directories, editing shell
# profiles) is either done below or intentionally left to the user.
#
# The launcher is pinned to a commit rather than a tag because tags can be moved,
# and its checksum is verified before it is made executable.
FOUNDRYUP_URL="https://raw.githubusercontent.com/foundry-rs/foundry/${FOUNDRY_COMMIT}/foundryup/foundryup"

sha256() {
  if command -v shasum >/dev/null 2>&1; then
    shasum -a 256 "$1" | awk '{print $1}'
  else
    sha256sum "$1" | awk '{print $1}'
  fi
}

echo "Installing foundryup ${FOUNDRY_VERSION} (${FOUNDRY_COMMIT}) to ${FOUNDRY_BIN_DIR}"
mkdir -p "$FOUNDRY_BIN_DIR"

FOUNDRYUP_TMP=$(mktemp)
trap 'rm -f "$FOUNDRYUP_TMP"' EXIT
curl --proto '=https' --tlsv1.2 -sSfL "$FOUNDRYUP_URL" -o "$FOUNDRYUP_TMP"

ACTUAL_SHA256=$(sha256 "$FOUNDRYUP_TMP")
if [[ "$ACTUAL_SHA256" != "$FOUNDRYUP_SHA256" ]]; then
  echo "error: foundryup checksum mismatch" >&2
  echo "  expected: ${FOUNDRYUP_SHA256}" >&2
  echo "  actual:   ${ACTUAL_SHA256}" >&2
  exit 1
fi

mv "$FOUNDRYUP_TMP" "$FOUNDRY_BIN_DIR/foundryup"
chmod +x "$FOUNDRY_BIN_DIR/foundryup"

# Remember whether the caller's PATH already covers the install directory so we
# can print a hint at the end. The upstream install script used to append the
# directory to the user's shell profile; this script does not modify profiles.
PATH_HAS_FOUNDRY=false
if [[ ":$PATH:" == *":${FOUNDRY_BIN_DIR}:"* ]]; then
  PATH_HAS_FOUNDRY=true
fi
export PATH="$PATH:$FOUNDRY_BIN_DIR"

foundryup --install "${FOUNDRY_VERSION}"

# Verify that foundryup actually installed the pinned version rather than
# silently falling back to a different release. forge reports the version
# without the leading "v", e.g. "forge Version: 1.0.0-v1.0.0".
EXPECTED_VERSION="${FOUNDRY_VERSION#v}"
INSTALLED_VERSION=$(forge --version | head -n 1)
if [[ "$INSTALLED_VERSION" != *"${EXPECTED_VERSION}"* ]]; then
  echo "error: expected forge version ${EXPECTED_VERSION}, but installed: ${INSTALLED_VERSION}" >&2
  exit 1
fi
echo "Installed ${INSTALLED_VERSION}"

if [[ "$PATH_HAS_FOUNDRY" == false ]]; then
  echo "Add ${FOUNDRY_BIN_DIR} to your PATH to use forge, cast, and anvil."
fi
