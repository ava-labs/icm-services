#!/bin/bash

# Exit on any errors
set -e

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
data_dir=~/.local_geth_network

# Address corresponding to private key 764A4A322753120B4667A20B6309CB5EC754A22BDBCBD62398BE8F803B255337
TEST_ADDRESS="0x6288dAdf62B57dd9A4ddcd02F88A98d0eb6c2598"

function delete_old_files() {
    rm -rf $data_dir
}

_term() { # Cleanup when Docker Compose sends a SIGTERM
    echo "Cleaning up..."
    killall geth 2>/dev/null || true
    delete_old_files
    exit 0
}

trap _term INT # Cleanup child processes and write log on Ctrl+C
trap _term SIGTERM # Cleanup when Docker Compose kills the container

# Delete old nodes
delete_old_files

# Initialize new node
mkdir -p $data_dir

# Start the node in dev mode
echo "Starting Geth node in dev mode..."
geth --dev \
    --http \
    --http.addr "0.0.0.0" \
    --http.port "5050" \
    --http.vhosts "*" \
    --http.corsdomain "*" \
    --datadir $data_dir \
    --http.api "eth,net,web3,personal" \
    --dev.period 1 &
node_pid=$!

# Wait for the node to start
echo "Waiting for geth to start..."
sleep 3
until curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' http://127.0.0.1:5050 > /dev/null 2>&1
do
    echo 'Waiting for geth RPC to be accessible...'
    sleep 1
done

echo "Geth is accessible. Funding test address..."

# Get the dev account (first account in eth_accounts)
DEV_ACCOUNT=$(curl -s -X POST -H "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":1}' \
    http://127.0.0.1:5050 | grep -o '"0x[a-fA-F0-9]*"' | head -1 | tr -d '"')

echo "Dev account: $DEV_ACCOUNT"

# Transfer funds from dev account to test address (10000 ETH)
TRANSFER_RESULT=$(curl -s -X POST -H "Content-Type: application/json" \
    --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_sendTransaction\",\"params\":[{\"from\":\"$DEV_ACCOUNT\",\"to\":\"$TEST_ADDRESS\",\"value\":\"0x21E19E0C9BAB2400000\"}],\"id\":1}" \
    http://127.0.0.1:5050)

echo "Transfer result: $TRANSFER_RESULT"

# Wait for the transaction to be mined
sleep 2

# Verify the balance
BALANCE=$(curl -s -X POST -H "Content-Type: application/json" \
    --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\",\"params\":[\"$TEST_ADDRESS\", \"latest\"],\"id\":1}" \
    http://127.0.0.1:5050)

echo "Test address balance: $BALANCE"

echo "Successfully started network and funded test address."

wait $node_pid
