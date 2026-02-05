#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

REPO_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

function printHelp() {
    echo "Usage: ./scripts/e2e_test.sh [--components component]"
    echo ""
    printUsage
}

function printUsage() {
    cat << EOF
Arguments:
    --components component1,component2            Comma separated list of test suites to run. Valid components are:
                                                  $(echo $valid_components | tr ' ' '\n' | sort | tr '\n' ' ')
                                                  (default: all)
    --network-dir path                            Path to the network directory. 
                                                  If the directory does not exist or is empty, it will be used as the root network directory for a new network.
                                                  If the directory exists and is non-empty, the network will be reused.
                                                  If not set, a new network will be created at the default root network directory.
    --epoch-duration duration                     Set to override the default test epoch duration.
Options:
    --help                                        Print this help message
    -v | --verbose                                Enable debug logs
EOF
}

valid_components=$(ls -d $REPO_PATH/icm-contracts/tests/suites/*/ | xargs -n 1 basename)
components=
reuse_network_dir=
root_dir=
network_dir=
reuse_network=false
epoch_duration=
while [ $# -gt 0 ]; do
    case "$1" in
        --components)  
            if [[ $2 != --* ]]; then
                components=$2
            else 
                echo "Invalid components $2" && printHelp && exit 1
            fi 
            shift;;
        --network-dir)
            if [[ $2 != --* ]]; then
                reuse_network_dir=$2
            else 
                echo "Invalid network directory $2" && printHelp && exit 1
            fi 
            shift;;
        --epoch-duration)
            if [[ $2 != --* ]]; then
                epoch_duration=$2
            else 
                echo "Invalid epoch duration $2" && printHelp && exit 1
            fi 
            shift;;
        --help) 
            printHelp && exit 0 ;;
        -v | --verbose)
            LOG_LEVEL=debug ;;
        *) 
            echo "Invalid option: $1" && printHelp && exit 1;;
    esac
    shift
done

# Run all suites if no component is provided
if [ -z "$components" ]; then
    components=$valid_components
fi

# Exit if invalid component is provided
for component in $(echo $components | tr ',' ' '); do
    if [[ $valid_components != *$component* ]]; then
        echo "Invalid component $component" && exit 1
    fi
done

if [ -n "$reuse_network_dir" ]; then
    if [ -d "$reuse_network_dir" ] && [ "$(ls -A "$reuse_network_dir")" ]; then
        network_dir=$reuse_network_dir
        reuse_network=true
        echo "Reuse network directory: $network_dir"
    else
        echo "Network directory $reuse_network_dir does not exist or is empty. Creating a new network at root $reuse_network_dir."
        mkdir -p "$reuse_network_dir"
        root_dir=$reuse_network_dir
    fi
fi

if [ -n "$epoch_duration" ]; then
    export GRANITE_EPOCH_DURATION=$epoch_duration
    echo "GRANITE_EPOCH_DURATION: $GRANITE_EPOCH_DURATION"
fi

source "$REPO_PATH"/scripts/constants.sh
source "$REPO_PATH"/scripts/versions.sh

BASEDIR=${BASEDIR:-"$HOME/.teleporter-deps"}

cwd=$(pwd)
# Install the avalanchego and subnet-evm binaries
rm -rf $BASEDIR/avalanchego
BASEDIR=$BASEDIR AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego "${REPO_PATH}/scripts/install_avalanchego_release.sh"
BASEDIR=$BASEDIR "${REPO_PATH}/scripts/install_subnetevm_release.sh"

cp ${BASEDIR}/subnet-evm/subnet-evm ${BASEDIR}/avalanchego/plugins/srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy
echo "Copied ${BASEDIR}/subnet-evm/subnet-evm binary to ${BASEDIR}/avalanchego/plugins/"

export AVALANCHEGO_BUILD_PATH=$BASEDIR/avalanchego
export AVALANCHEGO_PATH=$AVALANCHEGO_BUILD_PATH/avalanchego
export AVAGO_PLUGIN_DIR=$AVALANCHEGO_BUILD_PATH/plugins
export PATH=$PATH:$HOME/.foundry/bin

# Install signature-aggregator binary
"$REPO_PATH"/scripts/build_signature_aggregator.sh

cd "$REPO_PATH"
forge build --skip test
FOUNDRY_PROFILE=common forge build --skip test
FOUNDRY_PROFILE=ethereum forge build --skip test

# Start Ethereum network if running services tests
ETHEREUM_PID=
function start_ethereum_network() {
    echo "Starting local Ethereum network..."
    "${REPO_PATH}/ethereum/docker/run_ethereum_network.sh" > /dev/null 2>&1 &
    ETHEREUM_PID=$!
    
    # Wait for Ethereum network to be ready
    echo "Waiting for Ethereum network to be ready..."
    for i in {1..30}; do
        if curl -s -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
            http://127.0.0.1:5050 > /dev/null 2>&1; then
            echo "Ethereum network is ready"
            return 0
        fi
        sleep 1
    done
    echo "Failed to start Ethereum network"
    return 1
}

function stop_ethereum_network() {
    if [ -n "$ETHEREUM_PID" ]; then
        echo "Stopping Ethereum network (PID: $ETHEREUM_PID)..."
        kill $ETHEREUM_PID 2>/dev/null || true
        wait $ETHEREUM_PID 2>/dev/null || true
        ETHEREUM_PID=
    fi
}

# Cleanup on exit
trap stop_ethereum_network EXIT

for component in $(echo $components | tr ',' ' '); do
    # Start Ethereum network for services tests
    if [[ "$component" == "services" ]] && [ -z "$ETHEREUM_PID" ]; then
        start_ethereum_network
    fi
    echo "Building e2e tests for $component"
    go run github.com/onsi/ginkgo/v2/ginkgo build ${REPO_PATH}/icm-contracts/tests/suites/$component

    echo "Running e2e tests for $component"

    RUN_E2E=true LOG_LEVEL=${LOG_LEVEL} SIG_AGG_PATH=${REPO_PATH}/build/signature-aggregator ./icm-contracts/tests/suites/$component/$component.test \
    --root-network-dir=${root_dir} \
    --reuse-network=${reuse_network} \
    --network-dir=${network_dir} \
    --ginkgo.vv \
    --ginkgo.label-filter=${GINKGO_LABEL_FILTER:-""} \
    --ginkgo.focus=${GINKGO_FOCUS:-""} \
    --ginkgo.trace

    echo "$component e2e tests passed"
    echo ""
done
exit 0
