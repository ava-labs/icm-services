// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

/**
 * @dev Interface that contracts must implement to receive oracle messages from OracleAdapter.
 *
 * The caller is always the OracleAdapter contract. Implementations should verify
 * msg.sender equals the known OracleAdapter address to prevent spoofing.
 */
interface IOracleMessageReceiver {
    /**
     * @notice Called by OracleAdapter after a validator-attested oracle message passes
     *         warp verification, source allowlist, and replay-protection checks.
     *
     * @param sourceChainID     Avalanche blockchain ID of the chain whose validators attested this message.
     * @param sourceType        External chain type (e.g. "solana").
     * @param sourceAddress     Program or contract address on the source chain.
     * @param nonce             Replay-protection nonce, unique per (sourceType, sourceAddress).
     * @param payload           Application-level data extracted from the source chain event.
     */
    function receiveOracleMessage(
        bytes32 sourceChainID,
        string calldata sourceType,
        string calldata sourceAddress,
        uint64 nonce,
        bytes calldata payload
    ) external;
}
