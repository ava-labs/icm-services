#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

HELP=
LOG_LEVEL=
reuse_network_dir=
root_dir=
network_dir=
reuse_network=false
while [ $# -gt 0 ]; do
    case "$1" in
        --network-dir)
            if [[ $2 != --* ]]; then
                reuse_network_dir=$2
            else 
                echo "Invalid network directory $2" && printHelp && exit 1
            fi 
            shift;;
        -v | --verbose) LOG_LEVEL=debug ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/e2e_test.sh [OPTIONS]"
    echo "Run E2E tests for ICM Services."
    echo ""
    echo "Options:"
    echo "  --network-dir                     Path to the network directory to reuse. If not provided, a new network will be created."
    echo "                                    If the directory exists and is not empty, it will be reused."
    echo "                                    If the directory does not exist or is empty, a new network will be created at the specified path."
    echo "  -v, --verbose                     Enable debug logs"
    echo "  -h, --help                        Print this help message"
    exit 0
fi

BASE_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source "$BASE_PATH"/scripts/constants.sh
source "$BASE_PATH"/scripts/versions.sh

BASEDIR=${BASEDIR:-"$HOME/.teleporter-deps"}

if [ -n "$reuse_network_dir" ]; then
    if [ -d "$reuse_network_dir" ] && [ "$(ls -a "$reuse_network_dir")" ]; then
        network_dir=$reuse_network_dir
        reuse_network=true
        echo "Reuse network directory: $network_dir"
    else
        echo "Network directory $reuse_network_dir does not exist or is empty. Creating a new network at root $reuse_network_dir."
        mkdir -p "$reuse_network_dir"
        root_dir=$reuse_network_dir
    fi
fi

cwd=$(pwd)
# Install the avalanchego and subnet-evm binaries
rm -rf $BASEDIR/avalanchego
BASEDIR=$BASEDIR AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego ./scripts/install_avalanchego_release.sh
BASEDIR=$BASEDIR ./scripts/install_subnetevm_release.sh

# Install signature-aggregator to the location used by the tests
SIGNATURE_AGGREGATOR_PATH=$BASEDIR/icm-services/signature-aggregator
./scripts/build_signature_aggregator.sh $SIGNATURE_AGGREGATOR_PATH
echo "signature-aggregator Path: ${SIGNATURE_AGGREGATOR_PATH}"

cp ${BASEDIR}/subnet-evm/subnet-evm ${BASEDIR}/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy
echo "Copied ${BASEDIR}/subnet-evm/subnet-evm binary to ${BASEDIR}/avalanchego/plugins/"

export AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego
export AVALANCHEGO_PATH=$BASEDIR/avalanchego/avalanchego
export AVAGO_PLUGIN_DIR=$BASEDIR/avalanchego/plugins

go run github.com/onsi/ginkgo/v2/ginkgo build ./tests/
go build -v -o tests/cmd/decider/decider ./tests/cmd/decider/

# Run the tests
echo "Running e2e tests $RUN_E2E"
RUN_E2E=true LOG_LEVEL=${LOG_LEVEL} SIG_AGG_PATH=${SIG_AGG_PATH:-"$BASEDIR/icm-services/signature-aggregator"} ./tests/tests.test \
  --root-network-dir=${root_dir} \
  --reuse-network=${reuse_network} \
  --network-dir=${network_dir} \
  --ginkgo.vv \
  --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
  --ginkgo.focus=${GINKGO_FOCUS:-""} 

echo "e2e tests passed"
exit 0
