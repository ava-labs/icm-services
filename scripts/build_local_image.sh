#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -o errexit
set -o nounset
set -o pipefail

# Root of the repository
BASE_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)
# The constants and the build scripts run git and resolve some of their inputs relative to the
# repository root.
cd "$BASE_PATH"

# Load the constants. This sets RELAYER_PATH to the relayer directory, where the Dockerfile is,
# and relayer_path to the location of the built binary.
source "$BASE_PATH"/scripts/constants.sh

# WARNING: this will use the most recent commit even if there are un-committed changes present
full_commit_hash="$(git rev-parse HEAD)"
commit_hash="${full_commit_hash::8}"

"$BASE_PATH"/scripts/build.sh

image_tags=(-t "$relayer_dockerhub_repo:$commit_hash")
# Also tag the image with the current branch, if there is one. A Docker tag cannot contain a
# slash, which branch names commonly do.
branch_tag="${current_branch//\//-}"
if [[ -n "$branch_tag" ]]; then
    image_tags+=(-t "$relayer_dockerhub_repo:$branch_tag")
fi

# The Dockerfile copies the binary from the root of the build context, so the directory the
# binary was built into is used as the context.
echo "Building Docker Image with tags: ${image_tags[*]}"
docker build "${image_tags[@]}" -f "$RELAYER_PATH/Dockerfile" "$(dirname "$relayer_path")"
