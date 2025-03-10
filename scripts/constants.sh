#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Use lower_case variables in the scripts and UPPER_CASE variables for override
# Use the constants.sh for env overrides

BASE_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

RELAYER_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd ../relayer && pwd
)

SIGNATURE_AGGREGATOR_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd ../signature-aggregator && pwd
)

# Where binaries go
relayer_path="$BASE_PATH/build/icm-relayer"
signature_aggregator_path="$BASE_PATH/build/signature-aggregator"

# Set the PATHS
GOPATH="$(go env GOPATH)"

ICM_CONTRACTS_PATH="$BASE_PATH"/tests/contracts/lib/icm-contracts
source $ICM_CONTRACTS_PATH/scripts/constants.sh

# Avalabs docker hub repo is avaplatform/icm-relayer.
# Here we default to the local image (icm-relayer) as to avoid unintentional pushes
# You should probably set it - export DOCKER_REPO='avaplatform/icm-relayer'
relayer_dockerhub_repo=${DOCKER_REPO:-"icm-relayer"}

# Current branch
current_branch=$(git symbolic-ref -q --short HEAD || git describe --tags --exact-match || true)

git_commit=${RELAYER_COMMIT:-$( git rev-list -1 HEAD )}

# Set the CGO flags to use the portable version of BLST
#
# We use "export" here instead of just setting a bash variable because we need
# to pass this flag to all child processes spawned by the shell.
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"
# While CGO_ENABLED doesn't need to be explicitly set, it produces a much more
# clear error due to the default value change in go1.20.
export CGO_ENABLED=1
