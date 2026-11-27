#!/usr/bin/env bash
# Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Stop at the first build that fails. Without this the exit status is that of the last build
# alone, so a failed relayer build goes unnoticed by callers.
set -o errexit
set -o nounset
set -o pipefail

# Root directory
root=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

"$root"/scripts/build_relayer.sh
"$root"/scripts/build_signature_aggregator.sh
