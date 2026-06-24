// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {WarpMessage, IWarpMessenger} from "@subnet-evm/IWarpMessenger.sol";
import {IOracleMessageReceiver} from "./IOracleMessageReceiver.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

/**
 * @dev On-chain oracle message struct. Mirrors the Go oracle.OracleMessage encoding.
 *
 * The warp payload signed by validators is abi.encode of the individual fields
 * (sourceType, sourceAddress, destContract, sourceBlockHeight, nonce, payload),
 * which is identical to abi.encode of this struct when passed as a tuple.
 */
struct OracleMessage {
    // Identifies the external chain or data source (e.g. "solana", "bitcoin").
    string sourceType;
    // Program or contract address on the source chain.
    string sourceAddress;
    // Destination contract on this L1 that receives the decoded payload.
    address destContract;
    // Block or slot number on the source chain at which the event occurred.
    uint64 sourceBlockHeight;
    // Monotonically increasing per (sourceType, sourceAddress). Enforces replay protection.
    uint64 nonce;
    // Application-level data from the source chain event.
    bytes payload;
}

/**
 * @notice Standalone adapter that delivers validator-attested oracle messages to destination
 *         contracts on this L1.
 *
 * ## How it works
 *
 * 1. A relayer observes an event on an external chain (e.g. Solana).
 * 2. The relayer encodes the event as an OracleMessage and constructs a warp UnsignedMessage
 *    whose payload is abi.encode(sourceType, sourceAddress, destContract, sourceBlockHeight, nonce, payload).
 *    The warp message's SourceChainID is set to this L1's own chain ID.
 * 3. The relayer fans out ACP-118 signature requests to this L1's validators via handler ID 4.
 *    Each validator's sidecar independently verifies the event against the source chain before signing.
 * 4. Once a quorum of validators (≥ 2/3 stake) has signed, the relayer aggregates the BLS
 *    signatures and includes the signed warp message as a predicate in the access list of the
 *    delivery transaction.
 * 5. During block execution, the warp precompile verifies the BLS aggregate against this L1's
 *    validator set and stores the verified WarpMessage in predicate storage.
 * 6. The relayer calls receiveOracleMessage(warpIndex, oracleMsg). This contract reads the
 *    precompile result, hashes the payload to bind the BLS verification to the oracle fields,
 *    enforces the source allowlist and replay protection, then calls the receiver.
 *
 * ## Security model
 *
 * - BLS security: same as Avalanche consensus — requires ≥ 2/3 stake to be adversarial.
 * - Source allowlist: only configured (sourceType, sourceAddress) pairs are accepted, preventing
 *   a compromised validator from attesting to arbitrary sources.
 * - Replay protection: keyed on keccak256(sourceType, sourceAddress, nonce). Once delivered,
 *   a message ID can never be reused.
 * - No originSenderAddress check: oracle warp messages are constructed off-chain by the relayer
 *   (there is no sendWarpMessage call on the source), so originSenderAddress is address(0).
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract OracleAdapter {
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    address public owner;

    // keccak256(abi.encode(sourceType, sourceAddress)) => allowed
    mapping(bytes32 => bool) private _allowedSources;

    // keccak256(abi.encode(sourceType, sourceAddress, nonce)) => delivered
    mapping(bytes32 => bool) private _processedMessages;

    // -------------------------------------------------------------------------
    // Events
    // -------------------------------------------------------------------------

    /**
     * @notice Emitted when an oracle message is successfully delivered.
     * @param messageID    Replay-protection key: keccak256(sourceType, sourceAddress, nonce).
     * @param sourceType   External chain type (e.g. "solana").
     * @param sourceAddress Source program/contract address.
     * @param destContract Destination contract that received the payload.
     */
    event OracleMessageReceived(
        bytes32 indexed messageID,
        string sourceType,
        string sourceAddress,
        address indexed destContract
    );

    /**
     * @notice Emitted when an allowed source is added or removed.
     */
    event AllowedSourceUpdated(string sourceType, string sourceAddress, bool allowed);

    /**
     * @notice Emitted when ownership is transferred.
     */
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    // -------------------------------------------------------------------------
    // Errors
    // -------------------------------------------------------------------------

    error InvalidWarpMessage();
    error WrongSourceChain(bytes32 got, bytes32 want);
    error PayloadMismatch();
    error SourceNotAllowed(string sourceType, string sourceAddress);
    error AlreadyProcessed(bytes32 messageID);
    error Unauthorized();
    error ZeroAddress();

    // -------------------------------------------------------------------------
    // Constructor
    // -------------------------------------------------------------------------

    constructor(
        address _owner
    ) {
        if (_owner == address(0)) revert ZeroAddress();
        owner = _owner;
        emit OwnershipTransferred(address(0), _owner);
    }

    // -------------------------------------------------------------------------
    // Admin
    // -------------------------------------------------------------------------

    modifier onlyOwner() {
        if (msg.sender != owner) revert Unauthorized();
        _;
    }

    /**
     * @notice Add or remove a (sourceType, sourceAddress) pair from the allowlist.
     * @dev Only allowed sources can be delivered through this adapter. This is the
     *      on-chain counterpart of the Go-side AllowedSources config on each validator.
     */
    function setAllowedSource(
        string calldata sourceType,
        string calldata sourceAddress,
        bool allowed
    ) external onlyOwner {
        bytes32 key = keccak256(abi.encode(sourceType, sourceAddress));
        _allowedSources[key] = allowed;
        emit AllowedSourceUpdated(sourceType, sourceAddress, allowed);
    }

    /**
     * @notice Transfer ownership to a new address.
     */
    function transferOwnership(
        address newOwner
    ) external onlyOwner {
        if (newOwner == address(0)) revert ZeroAddress();
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }

    // -------------------------------------------------------------------------
    // Core
    // -------------------------------------------------------------------------

    /**
     * @notice Deliver a validator-attested oracle message to its destination contract.
     *
     * @dev The calling transaction MUST include the signed warp oracle message in its
     *      access list at index `warpIndex`. The warp precompile verifies the BLS aggregate
     *      during block execution before this function runs.
     *
     *      The warp payload must equal abi.encode(
     *          oracleMsg.sourceType,
     *          oracleMsg.sourceAddress,
     *          oracleMsg.destContract,
     *          oracleMsg.sourceBlockHeight,
     *          oracleMsg.nonce,
     *          oracleMsg.payload
     *      ). The relayer constructs the warp message with this exact encoding.
     *
     * @param warpIndex Index of the verified warp message in predicate storage.
     * @param oracleMsg The oracle message, provided as calldata by the relayer. Its hash
     *                  is checked against the warp payload to bind BLS verification to content.
     */
    function receiveOracleMessage(uint32 warpIndex, OracleMessage calldata oracleMsg) external {
        // 1. Read the precompile-verified warp message. The BLS aggregate was already
        //    checked against this L1's validator set during block execution.
        (WarpMessage memory warp, bool valid) = WARP_MESSENGER.getVerifiedWarpMessage(warpIndex);
        if (!valid) revert InvalidWarpMessage();

        // 2. The warp SourceChainID must be this chain. Oracle validators sign with the
        //    L1's own warp signer (SourceChainID = this chain's blockchain ID), so a message
        //    from a different chain cannot be accepted here.
        bytes32 thisChainID = WARP_MESSENGER.getBlockchainID();
        if (warp.sourceChainID != thisChainID) {
            revert WrongSourceChain(warp.sourceChainID, thisChainID);
        }

        // 3. Bind the BLS-verified payload to the oracle message fields provided by the relayer.
        //    The warp payload is abi.encode of the individual fields (NOT abi.encode of the struct,
        //    which would add an extra indirection for dynamic types).
        bytes32 warpPayloadHash = keccak256(warp.payload);
        bytes32 msgHash = keccak256(
            abi.encode(
                oracleMsg.sourceType,
                oracleMsg.sourceAddress,
                oracleMsg.destContract,
                oracleMsg.sourceBlockHeight,
                oracleMsg.nonce,
                oracleMsg.payload
            )
        );
        if (warpPayloadHash != msgHash) revert PayloadMismatch();

        // 4. Source allowlist check. Validators also enforce this per-node, but the on-chain
        //    check ensures a rogue validator cannot deliver to an unconfigured source.
        bytes32 sourceKey = keccak256(abi.encode(oracleMsg.sourceType, oracleMsg.sourceAddress));
        if (!_allowedSources[sourceKey]) {
            revert SourceNotAllowed(oracleMsg.sourceType, oracleMsg.sourceAddress);
        }

        // 5. Replay protection. MessageID commits to (sourceType, sourceAddress, nonce).
        //    The nonce must be unique per source; the contract does not enforce monotonicity,
        //    only uniqueness. Callers should treat nonces as opaque replay-protection tokens.
        bytes32 messageID =
            keccak256(abi.encode(oracleMsg.sourceType, oracleMsg.sourceAddress, oracleMsg.nonce));
        if (_processedMessages[messageID]) revert AlreadyProcessed(messageID);
        _processedMessages[messageID] = true;

        // 6. Deliver to destination. The destContract is part of the signed warp payload
        //    (verified in step 3), so the relayer cannot redirect to an arbitrary contract.
        IOracleMessageReceiver(oracleMsg.destContract).receiveOracleMessage(
            warp.sourceChainID,
            oracleMsg.sourceType,
            oracleMsg.sourceAddress,
            oracleMsg.nonce,
            oracleMsg.payload
        );

        emit OracleMessageReceived(
            messageID, oracleMsg.sourceType, oracleMsg.sourceAddress, oracleMsg.destContract
        );
    }

    // -------------------------------------------------------------------------
    // Views
    // -------------------------------------------------------------------------

    /**
     * @notice Returns true if the (sourceType, sourceAddress) pair is on the allowlist.
     */
    function isAllowed(
        string calldata sourceType,
        string calldata sourceAddress
    ) external view returns (bool) {
        return _allowedSources[keccak256(abi.encode(sourceType, sourceAddress))];
    }

    /**
     * @notice Returns true if the message with the given ID has already been delivered.
     * @dev messageID = keccak256(abi.encode(sourceType, sourceAddress, nonce))
     */
    function isProcessed(
        bytes32 messageID
    ) external view returns (bool) {
        return _processedMessages[messageID];
    }
}
