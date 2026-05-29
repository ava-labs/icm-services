#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

REPO_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $REPO_PATH/scripts/constants.sh
source $REPO_PATH/scripts/versions.sh

echo "Avalanche EVM Version: $AVALANCHE_EVM_VERSION"
echo "Avalanche Solidity Version: $AVALANCHE_SOLIDITY_VERSION"
echo "Common EVM Version: $COMMON_EVM_VERSION"
echo "Common Solidity Version: $COMMON_SOLIDITY_VERSION"
echo "Ethereum EVM Version: $ETHEREUM_EVM_VERSION"
echo "Ethereum Solidity Version: $ETHEREUM_SOLIDITY_VERSION"

DEFAULT_AVALANCHE_CONTRACT_LIST="TeleporterMessenger TeleporterRegistry WarpAdapter ExampleERC20 ExampleRewardCalculator TestMessenger ValidatorSetSig NativeTokenStakingManager ERC20TokenStakingManager
TokenHome TokenRemote ERC20TokenHome ERC20TokenHomeUpgradeable ERC20TokenRemote ERC20TokenRemoteUpgradeable NativeTokenHome NativeTokenHomeUpgradeable NativeTokenRemote NativeTokenRemoteUpgradeable
WrappedNativeToken MockERC20SendAndCallReceiver MockNativeSendAndCallReceiver ExampleERC20Decimals IStakingManager ACP99Manager ValidatorManager PoAManager BatchCrossChainMessenger INativeMinter
ECDSAVerifier ZKStateManager ZKAdapter RiscZeroGroth16Verifier"

DEFAULT_COMMON_CONTRACT_LIST="TeleporterMessengerV2 Adapter"

DEFAULT_ETHEREUM_CONTRACT_LIST="SubsetUpdater DiffUpdater MerkleValidatorSetRegistry"

PROXY_LIST="TransparentUpgradeableProxy ProxyAdmin"
ACCESS_LIST="OwnableUpgradeable"

EXTERNAL_LIBS="ValidatorMessages"

AVALANCHE_CONTRACT_LIST=
COMMON_CONTRACT_LIST=
ETHEREUM_CONTRACT_LIST=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -ac | --avalanche-contracts) AVALANCHE_CONTRACT_LIST=$2 ;;
        -cc | --common-contracts) COMMON_CONTRACT_LIST=$2 ;;
        -ec | --ethereum-contracts) ETHEREUM_CONTRACT_LIST=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/abi_bindings.sh [OPTIONS]"
    echo "Build contracts and generate Go bindings"
    echo ""
    echo "Options:"
    echo "  -ac, --avalanche-contracts contract1 contract2    Generate Go bindings for the contract. If empty, generate Go bindings for a default list of Avalanche contracts"
    echo "  -cc, --common-contracts contract1 contract2       Generate Go bindings for the contract. If empty, generate Go bindings for a default list of Common contracts"
    echo "  -ec, --ethereum-contracts contract1 contract      Generate Go bindings for the contract. If empty, generate Go bindings for a default list of Ethereum contracts"
    echo "  -h, --help                              Print this help message"
    exit 0
fi

if ! command -v reforge &> /dev/null; then
    echo "reforge not found. You can install by calling $REPO_PATH/scripts/build_reforge.sh" && exit 1
fi

if ! command -v solc &> /dev/null; then
    echo "solc not found. See https://docs.soliditylang.org/en/latest/installing-solidity.html for installation instructions" && exit 1
fi

# Get the version from solc output
solc_version_output=$(solc --version 2>&1)

# Extract the semver version from the output
extracted_version=$(solc --version 2>&1 | awk '/Version:/ {print $2}' | awk -F'+' '{print $1}')

# Check if the extracted version matches the expected version
if ! [[ "$extracted_version" == "$AVALANCHE_SOLIDITY_VERSION" ]]; then
    echo "Expected solc version $AVALANCHE_SOLIDITY_VERSION, but found $extracted_version. Please install the correct version." && exit 1
fi

# Install abigen
echo "Building libevm abigen"
go install github.com/ava-labs/libevm/cmd/abigen@$LIBEVM_VERSION

# Solc does not recursively expand remappings, so we must construct them manually
remappings=$(cat $REPO_PATH/remappings.txt)

# Recursively search for all remappings.txt files in the lib directory.
# For each file, prepend the remapping with the relative path to the file.
while read -r filepath; do
    relative_path="${filepath#$REPO_PATH/}"
    dir_path=$(dirname "$relative_path")
    echo $dir_path
  
    # Use sed to transform each line with the relative path if it matches the @token=remapping pattern,
    # so that each remapping is of the form @token=lib/path/to/remapping
    transformed_lines=$(sed -n "s|^\(@[^=]*=\)\(.*\)|\1$dir_path/\2|p" "$filepath")
    remappings+=" $transformed_lines "
done < <(find "$REPO_PATH/lib" -type f -name "remappings.txt" )

# Expand macros for a foundry profile into a target directory, maintaining the
# same relative path structure as the project root. All three profiles are
# written into the same dir so that cross-profile relative imports resolve.
function expand_profile_to_dir() {
    local profile="$1"
    local outdir="$2"
    local current_file=""
    while IFS= read -r line; do
        if [[ "$line" =~ ^===\ (.+)\ ===$ ]]; then
            current_file="$outdir/${BASH_REMATCH[1]}"
            mkdir -p "$(dirname "$current_file")"
            : > "$current_file"
        elif [[ -n "$current_file" ]]; then
            printf '%s\n' "$line" >> "$current_file"
        fi
    done < <(FOUNDRY_PROFILE="$profile" reforge --display '**/*.sol' build 2>/dev/null)
}

function convertToLower() {
    if [ "$(uname -m)" = 'arm64' ]; then
        echo $1 | perl -ne 'print lc'
    else
        echo $1 | sed -e 's/\(.*\)/\L\1/'
    fi
}

# Removes a matching string from a comma-separated list
remove_matching_string() {
    input_list="$1"
    match="$2"
    # Split the input list by commas
    IFS=',' read -ra elements <<< "$input_list"
    
    # Initialize an empty result array
    result=()

    # Iterate over each element
    for element in "${elements[@]}"; do
        # Check if the part after the colon matches the given string
        if [[ "${element#*:}" != "$match" ]]; then
        # If it doesn't match, add the element to the result array
        result+=("$element")
        fi
    done

    # Join the result array with commas and print
    (IFS=','; echo "${result[*]}")
}

function generate_bindings() {
    local evm_version="$1"
    local additional_flags="$2"
    # When non-empty, compile macro-expanded files from this root instead of
    # the original source tree.  Pass empty string for external libraries that
    # have no macro annotations.
    local source_root="$3"
    shift 3
    local contract_names=("$@")

    echo "EVM Version: $evm_version"
    echo "Solidity Version: $AVALANCHE_SOLIDITY_VERSION"
    echo "Additional flags: $additional_flags"

    # When compiling from the expanded dir, redirect project-source remappings
    # (those whose RHS does not start with lib/ or /) to paths inside source_root,
    # expressed as paths RELATIVE to $REPO_PATH (the solc CWD).  Keeping both the
    # input file and all remapping targets as relative paths ensures solc uses a
    # single consistent key for every file in its internal map, preventing the
    # "IWrappedNativeToken to IWrappedNativeToken" type-conflict that arises when
    # the same physical file is reached once via a relative path and once via an
    # absolute path derived from a remapping target.
    local source_root_rel="${source_root#$REPO_PATH/}"
    local effective_remappings="$remappings"
    if [[ -n "$source_root" ]]; then
        effective_remappings=""
        for remap in $remappings; do
            local rhs="${remap#*=}"
            if [[ "$rhs" == lib/* || "$rhs" == /* ]]; then
                effective_remappings+=" $remap"
            else
                effective_remappings+=" ${remap%%=*}=$source_root_rel/$rhs"
            fi
        done
    fi

    for contract_name in "${contract_names[@]}"
    do
        path=$(find . -name $contract_name.sol)
        dir=$(dirname $path)
        dir="${dir#./}"

        echo "Building $contract_name..."
        mkdir -p $REPO_PATH/out/$contract_name.sol

        combined_json=$REPO_PATH/out/$contract_name.sol/combined-output.json

        cwd=$(pwd)
        cd $REPO_PATH

        # Express the input file as a path relative to $REPO_PATH so that solc
        # records it under the same relative key it uses for all other project
        # files reached via the redirected remappings above.
        local sol_file
        if [[ -n "$source_root" ]]; then
            sol_file="$source_root_rel/${cwd#$REPO_PATH/}/$dir/$contract_name.sol"
        else
            sol_file="$cwd/$dir/$contract_name.sol"
        fi

        solc --optimize --evm-version $evm_version $additional_flags --metadata-hash none --combined-json abi,bin,metadata,ast,devdoc,userdoc --pretty-json $sol_file $effective_remappings > $combined_json
        cd $cwd

        # construct the exclude list
        contracts=$(jq -r '.contracts | keys | join(",")' $combined_json)

        # Filter out the contract we are generating bindings for
        filtered_contracts=$(remove_matching_string $contracts $contract_name)
        
        gen_path=$REPO_PATH/abi-bindings/go/$dir/$contract_name
        mkdir -p $gen_path
        echo "Generating Go bindings for $contract_name..."
        
        if [ -z "$filtered_contracts" ]; then
            echo "No external libraries found"
            $GOPATH/bin/abigen --pkg $(convertToLower $contract_name) \
                            --combined-json $combined_json \
                            --type $contract_name \
                            --out $gen_path/$contract_name.go
        else
            # Filter out external libraries
            for lib in $EXTERNAL_LIBS; do
                filtered_contracts=$(remove_matching_string $filtered_contracts $lib)
            done

            $GOPATH/bin/abigen --pkg $(convertToLower $contract_name) \
                            --combined-json $combined_json \
                            --type $contract_name \
                            --out $gen_path/$contract_name.go \
                            --exc $filtered_contracts
        fi
        
        echo "Done generating Go bindings for $contract_name."
        echo ""
    done
}

# If AVALANCHE_CONTRACT_LIST is empty, use DEFAULT_AVALANCHE_CONTRACT_LIST
if [[ -z "${CONTRACT_LIST}" ]]; then
    AVALANCHE_CONTRACT_LIST=($DEFAULT_AVALANCHE_CONTRACT_LIST)
fi

# If COMMON_CONTRACT_LIST is empty, use DEFAULT_COMMON_CONTRACT_LIST
if [[ -z "${CONTRACT_LIST}" ]]; then
    COMMON_CONTRACT_LIST=($DEFAULT_COMMON_CONTRACT_LIST)
fi

# If ETHEREUM_CONTRACT_LIST is empty, use DEFAULT_ETHEREUM_CONTRACT_LIST
if [[ -z "${ETHEREUM_CONTRACT_LIST}" ]]; then
    ETHEREUM_CONTRACT_LIST=($DEFAULT_ETHEREUM_CONTRACT_LIST)
fi

# Expand macros into a fixed directory so that the absolute paths embedded by
# solc in its metadata output are stable across runs, keeping generated Go
# bindings deterministic.  A random mktemp path would change every invocation
# and propagate into the metadata field of the combined JSON that abigen embeds
# in the generated file.
EXPANDED_DIR="$REPO_PATH/out/macro-expanded"
rm -rf "$EXPANDED_DIR"
mkdir -p "$EXPANDED_DIR"
trap "rm -rf '$EXPANDED_DIR'" EXIT
cd $REPO_PATH
for profile in default common ethereum; do
    expand_profile_to_dir "$profile" "$EXPANDED_DIR"
done

contract_names=(${AVALANCHE_CONTRACT_LIST[@]})
cd $AVALANCHE_ICM_PATH
generate_bindings "$AVALANCHE_EVM_VERSION" "" "$EXPANDED_DIR" "${contract_names[@]}"

contract_names=(${COMMON_CONTRACT_LIST[@]})
cd $COMMON_ICM_PATH
generate_bindings "$COMMON_EVM_VERSION" "" "$EXPANDED_DIR" "${contract_names[@]}"

contract_names=(${ETHEREUM_CONTRACT_LIST[@]})
cd $ETHEREUM_ICM_PATH
generate_bindings "$ETHEREUM_EVM_VERSION" "--via-ir" "$EXPANDED_DIR" "${contract_names[@]}"

contract_names=($PROXY_LIST)
cd $REPO_PATH/lib/openzeppelin-contracts-upgradeable/lib/openzeppelin-contracts/contracts/proxy/transparent
generate_bindings "$AVALANCHE_EVM_VERSION" "" "" "${contract_names[@]}"

contract_names=($ACCESS_LIST)
cd $REPO_PATH/lib/openzeppelin-contracts-upgradeable/contracts/access
generate_bindings "$AVALANCHE_EVM_VERSION" "" "" "${contract_names[@]}"

exit 0
