#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e
set -o pipefail

BASE_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

# Pass in the full name of the dependency.
# Parses go.mod for a matching entry and extracts the version number.
function getDepVersion() {
    grep -m1 "^\s*$1" $BASE_PATH/go.mod | cut -d ' ' -f2
}

function extract_commit() {
  local version=$1

  # Regex for a commit hash (assumed to be a 12+ character hex string)
  commit_hash_regex="-([0-9a-f]{12,})$"

  if [[ "$version" =~ $commit_hash_regex ]]; then
      # Extract the substring after the last '-'
      version=${BASH_REMATCH[1]}
  fi
  echo "$version"
}

# This needs to be exported to be picked up by the dockerfile.
export GO_VERSION=${GO_VERSION:-$(getDepVersion go)}

# ICM_SERVICES_VERSION is currently needed for the contracts E2E tests but is not a direct dependency
# since that would create a circular dependency. We should refactor the code until this is no longer the case.
# ICM_SERVICES_VERSION=${ICM_SERVICES_VERSION:-'signature-aggregator-v1.0.0-rc.0'}
ICM_SERVICES_VERSION=${ICM_SERVICES_VERSION:-'9564f00d296c7daeffbf26c4cc4866b3b6e98185'}

# Don't export them as they're used in the context of other calls
AVALANCHEGO_VERSION=${AVALANCHEGO_VERSION:-$(extract_commit "$(getDepVersion github.com/ava-labs/avalanchego)")}
SUBNET_EVM_VERSION=${SUBNET_EVM_VERSION:-$(extract_commit "$(getDepVersion github.com/ava-labs/subnet-evm)")}

# Extract the Solidity version from foundry.toml
SOLIDITY_VERSION=$(awk -F"'" '/^solc_version/ {print $2}' $BASE_PATH/foundry.toml)
EVM_VERSION=$(awk -F"'" '/^evm_version/ {print $2}' $BASE_PATH/foundry.toml)
