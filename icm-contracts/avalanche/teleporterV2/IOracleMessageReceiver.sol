// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

/**
 * @dev Interface that destination contracts must implement to receive oracle-attested messages.
 *
 * The OracleAdapter calls this after verifying the BLS aggregate, enforcing the source allowlist,
 * and enforcing replay protection. The receiver can trust that:
 *   - The BLS aggregate over (sourceType, sourceAddress, destContract, sourceBlockHeight, nonce, payload)
 *     was produced by a quorum of this L1's validators.
 *   - The (sourceType, sourceAddress) pair is on the adapter's allowlist.
 *   - This (sourceType, sourceAddress, nonce) triple has never been delivered before.
 *
 * SECURITY: Implementations MUST check that msg.sender is a trusted OracleAdapter contract.
 * Anyone can deploy a contract that calls receiveOracleMessage directly.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
interface IOracleMessageReceiver {
    /**
     * @notice Called by OracleAdapter when a validator-attested oracle message is delivered.
     *
     * @param sourceBlockchainID The blockchain ID of the L1 whose validators attested the message.
     *                           For oracle messages this is always the receiving L1's own chain ID.
     * @param sourceType         Identifies the external chain or data source (e.g. "solana").
     * @param sourceAddress      The program or contract address on the source that emitted the event.
     * @param nonce              Monotonically increasing per (sourceType, sourceAddress); used for
     *                           replay protection. Uniqueness is enforced by OracleAdapter.
     * @param payload            Application-level data from the source chain event.
     */
    function receiveOracleMessage(
        bytes32 sourceBlockchainID,
        string calldata sourceType,
        string calldata sourceAddress,
        uint64 nonce,
        bytes calldata payload
    ) external;
}
