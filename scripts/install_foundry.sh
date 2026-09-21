#!/usr/bin/env bash
# Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

# The foundry install script uses XDG_CONFIG_HOME as the root of the install.
# This can vary for different environments, so it is set to $HOME for consistency.
export XDG_CONFIG_HOME=$HOME

FOUNDRYUP_ATTEMPTS=${FOUNDRYUP_ATTEMPTS:-4}
FOUNDRYUP_RETRY_DELAY=${FOUNDRYUP_RETRY_DELAY:-15}

FOUNDRY_VERSION=v1.6.0-rc1
curl -L --retry 5 --retry-connrefused --retry-all-errors https://raw.githubusercontent.com/foundry-rs/foundry/${FOUNDRY_VERSION}/foundryup/install > /tmp/foundry-install-script
# Set the foundry version in the install script
# Avoid using sed -i due to macos m1 incompatibility
sed "s/\/foundry-rs\/foundry\/master\/foundryup/\/foundry-rs\/foundry\/${FOUNDRY_VERSION}\/foundryup/g" /tmp/foundry-install-script
bash < /tmp/foundry-install-script

export PATH=$PATH:$HOME/.foundry/bin:$HOME/.foundry:$HOME/.cargo/bin

# foundryup verifies release attestations, and GitHub's download endpoints
# intermittently return 504 to CI runners, which aborts the whole install.
for attempt in $(seq 1 "$FOUNDRYUP_ATTEMPTS"); do
    if foundryup --install "${FOUNDRY_VERSION}"; then
        exit 0
    fi
    if [ "$attempt" -lt "$FOUNDRYUP_ATTEMPTS" ]; then
        echo "install_foundry: foundryup failed (attempt ${attempt}/${FOUNDRYUP_ATTEMPTS}), retrying in ${FOUNDRYUP_RETRY_DELAY}s"
        sleep "$FOUNDRYUP_RETRY_DELAY"
    fi
done

echo "install_foundry: foundryup could not install ${FOUNDRY_VERSION} after ${FOUNDRYUP_ATTEMPTS} attempts" >&2
exit 1
