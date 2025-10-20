#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -o errexit
set -o nounset
set -o pipefail

RELAYER_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

source $RELAYER_PATH/scripts/versions.sh

go run github.com/golangci/golangci-lint/cmd/golangci-lint run --config=$RELAYER_PATH/.golangci.yml --build-tags=test ./... --timeout 5m

(cd proto && go run github.com/bufbuild/buf/cmd/buf lint)
