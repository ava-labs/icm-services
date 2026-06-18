#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -eo pipefail

REPO_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $REPO_PATH/scripts/versions.sh

function solFormat() {
    # format solidity contracts
    echo "Formatting Solidity contracts..."
    reforge fmt --root $REPO_PATH $REPO_PATH/icm-contracts/**
}

function solFormatCheck() {
    # format solidity contracts
    echo "Checking formatting of Solidity contracts..."
    reforge fmt --check --root $REPO_PATH $REPO_PATH/icm-contracts/**
}

function solLinter() {
    # lint solidity contracts
    echo "Linting Solidity contracts..."
    cd $REPO_PATH

    # Expand macros to a temp dir
    EXPANDED_DIR=$(mktemp -d)
    trap "rm -rf '$EXPANDED_DIR'" EXIT

    for profile in default common ethereum; do
        current_file=""
        while IFS= read -r line; do
            if [[ "$line" =~ ^===\ (.+)\ ===$ ]]; then
                current_file="$EXPANDED_DIR/${BASH_REMATCH[1]}"
                mkdir -p "$(dirname "$current_file")"
                : > "$current_file"
            elif [[ -n "$current_file" ]]; then
                printf '%s\n' "$line" >> "$current_file"
            fi
        done < <(FOUNDRY_PROFILE="$profile" reforge --display '**/*.sol' build 2>/dev/null)
    done

    # solhint globs are relative to cwd, so cd into the expanded dir.
    # "solhint **/*.sol" runs differently than "solhint '**/*.sol'", where the latter checks sol files
    # in subdirectories. The former only checks sol files in the current directory and directories one level down.
    (cd "$EXPANDED_DIR" && solhint '**/*.sol' --config "$REPO_PATH/.solhint.json" --max-warnings 0)
}

function golangLinter() {
    echo "Linting Golang code..."
    cd $REPO_PATH
    go run github.com/golangci/golangci-lint/cmd/golangci-lint run --config=$REPO_PATH/.golangci.yml --build-tags=test ./... --timeout 5m --fix
    (cd proto && go run github.com/bufbuild/buf/cmd/buf lint)
}

function runAll() {
    solFormat
    solLinter
    golangLinter
}

function printHelp() {
    echo "Usage: ./scripts/lint.sh [OPTIONS]"
    echo "Lint/Format Teleporter Solidity contracts and E2E tests Golang code."
    echo "Pass no parameters to perform all checks"
    printUsage
}

function printUsage() {
    echo "Options:"
    echo "  -sfc, --sol-format-check            Check for proper formatted Solidity files. Exits with code 1 if not."
    echo "  -sf,  --sol-format                  Format Solidity contracts"
    echo "  -sl,  --sol-lint                    Run the Solidity linter"
    echo "  -gl,  --go-lint                     Run the Golang linter"
    echo "  -h,   --help                        Print this help message"
}

# if we have no args, perform all checks
if [ $# -eq 0 ]; then
    runAll
    exit 0
fi

while [ $# -gt 0 ]; do
    case "$1" in
        -sfc | --sol-format-check) 
            solFormatCheck ;;
        -sf  | --sol-format) 
            solFormat ;;
        -sl  | --sol-lint) 
            solLinter ;;
        -gl  | --go-lint) 
            golangLinter ;;
        -h   | --help) 
            printHelp ;;
        *) 
          echo "Invalid option: $1" && printHelp && exit 1;;
    esac
    shift
done


exit 0
