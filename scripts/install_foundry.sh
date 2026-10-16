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
FOUNDRY_VERSION=v1.6.0-rc1
FOUNDRY_COMMIT=b272ce3366987406406e9eb1b82596653a3ad628
FOUNDRYUP_SHA256=9f21de4687101a3e12adbad69bcb002ca2d454e9fff2352e33ab53f57074ec61

FOUNDRY_DIR="${FOUNDRY_DIR:-$HOME/.foundry}"
FOUNDRY_BIN_DIR="$FOUNDRY_DIR/bin"

FOUNDRYUP_ATTEMPTS="${FOUNDRYUP_ATTEMPTS:-4}"
FOUNDRYUP_RETRY_DELAY="${FOUNDRYUP_RETRY_DELAY:-15}"

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

# foundryup verifies release attestations, and GitHub's download endpoints
# intermittently return 504 to CI runners, which aborts the install. Retry
# instead of passing --force, which would skip that verification.
for attempt in $(seq 1 "$FOUNDRYUP_ATTEMPTS"); do
  if foundryup --install "${FOUNDRY_VERSION}"; then
    break
  fi
  if [[ "$attempt" -eq "$FOUNDRYUP_ATTEMPTS" ]]; then
    echo "error: foundryup could not install ${FOUNDRY_VERSION} after ${FOUNDRYUP_ATTEMPTS} attempts" >&2
    exit 1
  fi
  echo "foundryup failed (attempt ${attempt}/${FOUNDRYUP_ATTEMPTS}), retrying in ${FOUNDRYUP_RETRY_DELAY}s"
  sleep "$FOUNDRYUP_RETRY_DELAY"
done

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
