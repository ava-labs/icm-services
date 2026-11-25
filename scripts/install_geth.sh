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
# SHA-256 of the release tarball. Must be updated together with GETH_VERSION.
GETH_SHA256=28c5f3806ef64d9d464f14ceefc3e9b5183b28f3fc5709e811931e8f028e2bba

GETH_DIR=$HOME/.geth
mkdir $GETH_DIR
export PATH=$PATH:$GETH_DIR

# Download, verify the checksum, and only then extract and install geth
wget https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-${GETH_VERSION}.tar.gz && \
    echo "${GETH_SHA256}  geth-linux-amd64-${GETH_VERSION}.tar.gz" | sha256sum -c - && \
    tar -xzf geth-linux-amd64-${GETH_VERSION}.tar.gz && \
    cp geth-linux-amd64-${GETH_VERSION}/geth $GETH_DIR && \
    rm -rf geth-linux-amd64-${GETH_VERSION}*

# Make the script executable
chmod +x ./ethereum/run_ethereum_network.sh

# Verify installation
geth version
