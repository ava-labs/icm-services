// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {WarpMessage, IWarpMessenger} from "@subnet-evm/IWarpMessenger.sol";
import {
    IAdapter,
    IMessageSender,
    IMessageVerifier,
    TeleporterMessageV2,
    TeleporterICMMessage
} from "@common/ITeleporterMessengerV2.sol";

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
 * @notice IAdapter implementation that delivers validator-attested oracle messages via TeleporterV2.
 *
 * Oracle messages originate from external chains (e.g. Solana). Validators on this L1 attest
 * to them by signing a warp message whose sourceChainID equals this chain's own blockchain ID.
 * The signed message rides through TeleporterMessengerV2.receiveCrossChainMessage using
 * sourceBlockchainID == thisChainID (self-origin).
 *
 * ## Attestation encoding
 *
 * TeleporterICMMessage.attestation = abi.encode(uint32 warpIndex, OracleMessage oracleMsg)
 *
 * ## Message payload encoding (TeleporterMessageV2.message)
 *
 * abi.encode(sourceType, sourceAddress, sourceBlockHeight, nonce, payload)
 * Destination contracts receive these fields via receiveTeleporterMessage.
 *
 * @custom:security-contact https://github.com/ava-labs/icm-contracts/blob/main/SECURITY.md
 */
contract OracleAdapter is IAdapter {
    IWarpMessenger public constant WARP_MESSENGER =
        IWarpMessenger(0x0200000000000000000000000000000000000005);

    address public owner;

    // keccak256(abi.encode(sourceType, sourceAddress)) => allowed
    mapping(bytes32 => bool) private _allowedSources;

    // keccak256(abi.encode(sourceType, sourceAddress, nonce)) => delivered
    // Kept alongside Teleporter's own replay protection because oracle nonce is unique
    // per (sourceType, sourceAddress) pair, not globally across all sources.
    mapping(bytes32 => bool) private _processedMessages;

    // -------------------------------------------------------------------------
    // Events
    // -------------------------------------------------------------------------

    /**
     * @notice Emitted when an oracle message passes verification and is marked for delivery.
     * @param messageID     Replay-protection key: keccak256(sourceType, sourceAddress, nonce).
     * @param sourceType    External chain type (e.g. "solana").
     * @param sourceAddress Source program/contract address.
     * @param destContract  Destination contract that will receive the payload.
     */
    event OracleMessageVerified(
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
    // Modifiers
    // -------------------------------------------------------------------------

    modifier onlyOwner() {
        if (msg.sender != owner) revert Unauthorized();
        _;
    }

    // -------------------------------------------------------------------------
    // Constructor
    // -------------------------------------------------------------------------

    constructor(
        address initialOwner
    ) {
        if (initialOwner == address(0)) revert ZeroAddress();
        owner = initialOwner;
        emit OwnershipTransferred(address(0), initialOwner);
    }

    // -------------------------------------------------------------------------
    // IAdapter
    // -------------------------------------------------------------------------

    /**
     * @inheritdoc IMessageSender
     */
    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external override {
        WARP_MESSENGER.sendWarpMessage(abi.encode(message));
    }

    /**
     * @notice Verify a validator-attested oracle message.
     *
     * @dev The calling transaction MUST include the signed warp oracle message in its
     *      access list. The warp precompile verifies the BLS aggregate during block
     *      execution before this function runs.
     *
     *      message.attestation must be abi.encode(uint32 warpIndex, OracleMessage oracleMsg).
     *      message.sourceBlockchainID must equal this chain's blockchain ID.
     *
     * @inheritdoc IMessageVerifier
     */
    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external override returns (bool) {
        (uint32 warpIndex, OracleMessage memory oracleMsg) =
            abi.decode(message.attestation, (uint32, OracleMessage));

        // 1. Read the precompile-verified warp message. The BLS aggregate was already
        //    checked against this L1's validator set during block execution.
        (WarpMessage memory warp, bool valid) = WARP_MESSENGER.getVerifiedWarpMessage(warpIndex);
        if (!valid) revert InvalidWarpMessage();

        // 2. Oracle validators sign with this L1's own blockchain ID as source.
        bytes32 thisChainID = WARP_MESSENGER.getBlockchainID();
        if (warp.sourceChainID != thisChainID) {
            revert WrongSourceChain(warp.sourceChainID, thisChainID);
        }

        // 3. Bind the BLS-verified payload to the oracle fields provided in the attestation.
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

        // 5. Replay protection keyed on (sourceType, sourceAddress, nonce). Independent of
        //    Teleporter's messageNonce-based replay protection because oracle nonce is unique
        //    per (sourceType, sourceAddress) pair, not globally.
        bytes32 messageID =
            keccak256(abi.encode(oracleMsg.sourceType, oracleMsg.sourceAddress, oracleMsg.nonce));
        if (_processedMessages[messageID]) revert AlreadyProcessed(messageID);
        _processedMessages[messageID] = true;

        emit OracleMessageVerified(
            messageID, oracleMsg.sourceType, oracleMsg.sourceAddress, oracleMsg.destContract
        );

        return true;
    }

    // -------------------------------------------------------------------------
    // Admin
    // -------------------------------------------------------------------------

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
