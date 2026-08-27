#!/usr/bin/env bash
# Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# This file is sourced by the deploy scripts; the variables it defines are read there.

# Shared handling of the signing key for the deploy scripts.
#
# A private key passed on a command line is readable by every local user for the
# lifetime of the process (/proc/<pid>/cmdline, ps) and is persisted in the operator's
# shell history, so neither these scripts nor the cast invocations they make accept one
# that way. The options below name a signer without carrying any secret material:
# a keystore path, a keystore account name, an interactive prompt, or a hardware wallet.

wallet_keystore=
wallet_account=
wallet_password_file=
wallet_interactive=false
wallet_ledger=false
wallet_trezor=false

# Number of arguments parse_wallet_arg consumed beyond the flag itself.
wallet_arg_shift=0

# The flags to hand to cast to select the signer. Set by resolve_cast_wallet_args.
cast_wallet_args=()

function printWalletUsage() {
    cat << USAGE
Signing options (choose one; none of these place key material on a command line):
    --keystore <path>                Encrypted keystore file (or directory) to sign with
    --account <name>                 Keystore account name in ~/.foundry/keystores
    --interactive                    Prompt for the private key on the terminal
    --ledger                         Sign with a Ledger hardware wallet
    --trezor                         Sign with a Trezor hardware wallet
    --password-file <path>           File holding the keystore password. Defaults to
                                     \$ETH_PASSWORD; cast prompts if neither is set.

To create a keystore from a raw private key, run this once and enter the key when
prompted (it is never passed as an argument):
    cast wallet import <name> --interactive
USAGE
}

# Handles one wallet-related option. Returns 0 if "$1" was recognized, 1 otherwise, and
# sets wallet_arg_shift to the number of extra arguments consumed.
# wallet_arg_shift is read by the sourcing script's argument loop, not here.
# shellcheck disable=SC2034
function parse_wallet_arg() {
    wallet_arg_shift=0
    case "$1" in
        --keystore)
            requireWalletValue "$1" "$2" && wallet_keystore=$2 && wallet_arg_shift=1 ;;
        --account)
            requireWalletValue "$1" "$2" && wallet_account=$2 && wallet_arg_shift=1 ;;
        --password-file)
            requireWalletValue "$1" "$2" && wallet_password_file=$2 && wallet_arg_shift=1 ;;
        --interactive)
            wallet_interactive=true ;;
        --ledger)
            wallet_ledger=true ;;
        --trezor)
            wallet_trezor=true ;;
        --private-key)
            cat >&2 << PRIVATE_KEY_REJECTED

--private-key is no longer accepted.

A key given on the command line is visible to every local user through the process
listing for as long as the deployment runs, and is written to your shell history.

Use --keystore, --account, --interactive, --ledger, or --trezor instead. To turn a raw
private key into a keystore, run this once and paste the key at the prompt:

    cast wallet import <name> --interactive

If the key you just passed is a real one, treat it as exposed on this host and rotate it.
PRIVATE_KEY_REJECTED
            exit 1 ;;
        *)
            return 1 ;;
    esac
    return 0
}

function requireWalletValue() {
    if [[ -z $2 || $2 == --* ]]; then
        echo "Missing value for $1" >&2
        exit 1
    fi
}

# True if the caller selected a signer.
function walletConfigured() {
    [[ -n $wallet_keystore || -n $wallet_account ]] && return 0
    [[ $wallet_interactive == true || $wallet_ledger == true || $wallet_trezor == true ]]
}

# Validates the selection and fills cast_wallet_args.
function resolve_cast_wallet_args() {
    local selected=0
    [[ -n $wallet_keystore ]] && selected=$((selected + 1))
    [[ -n $wallet_account ]] && selected=$((selected + 1))
    [[ $wallet_interactive == true ]] && selected=$((selected + 1))
    [[ $wallet_ledger == true ]] && selected=$((selected + 1))
    [[ $wallet_trezor == true ]] && selected=$((selected + 1))

    if [[ $selected -gt 1 ]]; then
        echo "Choose a single signing option." >&2
        printWalletUsage >&2
        exit 1
    fi

    cast_wallet_args=()
    if [[ -n $wallet_keystore ]]; then
        if [[ ! -e $wallet_keystore ]]; then
            echo "Keystore $wallet_keystore does not exist" >&2
            exit 1
        fi
        cast_wallet_args+=(--keystore "$wallet_keystore")
    elif [[ -n $wallet_account ]]; then
        cast_wallet_args+=(--account "$wallet_account")
    elif [[ $wallet_interactive == true ]]; then
        cast_wallet_args+=(--interactive)
    elif [[ $wallet_ledger == true ]]; then
        cast_wallet_args+=(--ledger)
    elif [[ $wallet_trezor == true ]]; then
        cast_wallet_args+=(--trezor)
    fi

    # The password is a path, not the password itself. Without it cast falls back to
    # $ETH_PASSWORD (also a path) or prompts on the terminal.
    if [[ -n $wallet_password_file ]]; then
        if [[ ! -f $wallet_password_file ]]; then
            echo "Password file $wallet_password_file does not exist" >&2
            exit 1
        fi
        cast_wallet_args+=(--password-file "$wallet_password_file")
    fi
}
