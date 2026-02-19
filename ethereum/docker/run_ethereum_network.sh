#!/bin/bash

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
data_dir=~/.local_geth_network
node_pid=""

# Address corresponding to private key 764A4A322753120B4667A20B6309CB5EC754A22BDBCBD62398BE8F803B255337
TEST_ADDRESS="0x6288dAdf62B57dd9A4ddcd02F88A98d0eb6c2598"

function delete_old_files() {
    rm -rf "$data_dir"
}

_term() {
    local signal=$1
    echo "Received signal $signal, cleaning up..."

    # Kill the specific geth process, not all geth processes
    if [ -n "$node_pid" ]; then
        echo "Stopping geth process (PID: $node_pid)..."
        kill -TERM "$node_pid" 2>/dev/null || true

        # Wait for graceful shutdown with timeout
        local timeout=10
        local count=0
        while kill -0 "$node_pid" 2>/dev/null && [ $count -lt $timeout ]; do
            sleep 1
            count=$((count + 1))
        done

        # Force kill if still running
        if kill -0 "$node_pid" 2>/dev/null; then
            echo "Force killing geth process..."
            kill -KILL "$node_pid" 2>/dev/null || true
        fi
    fi

    delete_old_files

    # Exit with appropriate signal-based exit code
    # Convention: 128 + signal number
    if [ "$signal" = "INT" ]; then
        exit 130  # 128 + 2 (SIGINT)
    elif [ "$signal" = "TERM" ]; then
        exit 143  # 128 + 15 (SIGTERM)
    else
        exit 1
    fi
}

trap '_term INT' INT
trap '_term TERM' SIGTERM

# Delete old nodes
delete_old_files

# Initialize new node
mkdir -p "$data_dir"

# Start the node in dev mode
echo "Starting Geth node in dev mode..."
if ! command -v geth &> /dev/null; then
    echo "Error: geth command not found in PATH"
    exit 1
fi

geth --dev \
    --http \
    --http.addr "0.0.0.0" \
    --http.port "5050" \
    --http.vhosts "*" \
    --http.corsdomain "*" \
    --datadir "$data_dir" \
    --http.api "eth,net,web3,personal" \
    --dev.period 1 &
node_pid=$!

# Verify the process started
sleep 1
if ! kill -0 "$node_pid" 2>/dev/null; then
    echo "Error: geth process failed to start"
    exit 1
fi

# Wait for the node to start with timeout
echo "Waiting for geth to start..."
max_wait=60
elapsed=0
until curl -s -X POST -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}' http://127.0.0.1:5050 > /dev/null 2>&1
do
    if [ $elapsed -ge $max_wait ]; then
        echo "Error: Timeout waiting for geth RPC to become accessible"
        kill -TERM "$node_pid" 2>/dev/null || true
        exit 1
    fi
    echo "Waiting for geth RPC to be accessible... ($elapsed/$max_wait seconds)"
    sleep 1
    elapsed=$((elapsed + 1))
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

# Wait for the geth process to exit
# This will return the exit code of geth
wait $node_pid
exit_code=$?

echo "Geth process exited with code $exit_code"
delete_old_files
exit $exit_code
