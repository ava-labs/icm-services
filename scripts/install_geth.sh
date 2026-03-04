#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

 # Install dependencies
 apt-get update && \
    apt-get install -y wget ca-certificates psmisc && \
    rm -rf /var/lib/apt/lists/*

# Set Geth version with commit hash
# The "12b4131f" is the commit hash for the official v1.16.1 release of geth.
# See: https://geth.ethereum.org/downloads/ and https://github.com/ethereum/go-ethereum/releases/tag/v1.16.1
GETH_VERSION=1.16.1-12b4131f

GETH_DIR=$HOME/.geth
mkdir $GETH_DIR
export PATH=$PATH:$GETH_DIR

# Download and install geth
wget https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-${GETH_VERSION}.tar.gz && \
    tar -xzf geth-linux-amd64-${GETH_VERSION}.tar.gz && \
    cp geth-linux-amd64-${GETH_VERSION}/geth $GETH_DIR && \
    rm -rf geth-linux-amd64-${GETH_VERSION}*

# Make the script executable
chmod +x ./ethereum/run_ethereum_network.sh

# Verify installation
geth version
