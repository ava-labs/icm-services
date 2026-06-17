// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {IMerkleValidatorSetRegistry} from "./interfaces/IMerkleValidatorSetRegistry.sol";
import {ICMMessage} from "../common/ICM.sol";
import {ValidatorSetMerkleCommitment, ValidatorSets} from "./utils/ValidatorSets.sol";
import {
    TeleporterMessageV2Parsing,
    TeleporterICMMessage,
    TeleporterMessageV2
} from "../common/TeleporterMessageV2.sol";
import {IAdapter} from "../common/ITeleporterMessengerV2.sol";
import {IWarpMessenger} from "../avalanche/subnet-evm/IWarpMessenger.sol";
import {ISP1Verifier} from "@sp1-contracts/ISP1Verifier.sol";

/**
 * THIS IS AN EXAMPLE CONTRACT THAT USES UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

// A contract for managing Avalanche validator sets which can be used to verify ICM messages
// via a zero-knowledge (ZK) proof of a quorum of signatures.
//
// Like MerkleValidatorSetRegistry, validator sets are anchored on-chain by a Merkle root
// commitment over the set, rather than by storing the full validator set. Unlike that
// contract, the Merkle multi-inclusion proof and the aggregate BLS signature checks are not
// performed on-chain. Instead, a single SP1 ZK proof attests that:
//   (a) a set of validators is committed under the stored Merkle root, and
//   (b) their aggregate BLS signature over the signed message is valid,
// committing the total signing weight as a public value. The contract then checks those
// public values against its own state (the committed root matches storage, the message
// hash matches, the signing weight meets quorum) and verifies the ZK proof, trading on-chain
// proof-verification and per-message calldata for off-chain proving.
// As a result, the signing validators, the multi-inclusion proof, and the aggregate
// signature never appear in calldata; only the proof and its public values do.
contract ZKValidatorSetRegistry is IMerkleValidatorSetRegistry, IAdapter {
    // The public values committed by the SP1 program
    struct PublicValues {
        bytes32 sourceBlockchainID;
        bytes32 root;
        bytes32 messageHash;
        uint64 signedWeight;
    }

    address private constant _WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    uint32 public immutable avalancheNetworkID;
    // The Avalanche blockchain ID of the P-chain
    bytes32 public immutable pChainID;
    // Whether the P-Chain may sign off on updates to already registered L1s
    bool public immutable allowPChainFallback;
    // The SP1 verifier gateway for this chain, which routes a proof to the correct SNARK verifier based on its version
    ISP1Verifier public immutable sp1Verifier;
    // The verification key of the attestation guest program
    // TODO: Kept immutable here for simplicity; in production make this owner-upgradeable so the program or the
    // SP1 version can be updated without redeploying the registry
    bytes32 public immutable attestationProgramVKey;
    // Mapping of Avalanche blockchain IDs to their validator set commitments
    mapping(bytes32 => ValidatorSetMerkleCommitment) internal _valSetCommitments;

    // Constructs a new registry instance with the initial validator set commitment registered on the P-chain,
    // including the SP1 verifier gateway and the attestation program's verification key
    constructor(
        uint32 avalancheNetworkID_,
        bytes32 pChainID_,
        bytes32 pChainGenesisRoot,
        uint64 pChainTotalWeight,
        uint64 pChainHeight,
        uint64 pChainTimestamp,
        bool allowPChainFallback_,
        ISP1Verifier sp1Verifier_,
        bytes32 attestationProgramVKey_
    ) {
        avalancheNetworkID = avalancheNetworkID_;
        pChainID = pChainID_;
        allowPChainFallback = allowPChainFallback_;
        sp1Verifier = sp1Verifier_;
        attestationProgramVKey = attestationProgramVKey_;
        _valSetCommitments[pChainID_] = ValidatorSetMerkleCommitment({
            avalancheBlockchainID: pChainID_,
            root: pChainGenesisRoot,
            totalWeight: pChainTotalWeight,
            pChainHeight: pChainHeight,
            pChainTimestamp: pChainTimestamp
        });
    }

    /**
     * @dev sendMessage has no msg.sender == originTeleporterAddress check. This check must live in
     * the contract that the messenger (Teleporter) calls directly, currently in Adapter.sol,
     * in order to prevent unauthorized calls to sendMessage. To clarify, this contract cannot be used as the
     * direct messaging entry-point.
     */
    function sendMessage(
        TeleporterMessageV2 calldata message
    ) external {
        IWarpMessenger(_WARP_PRECOMPILE_ADDRESS).sendWarpMessage(
            TeleporterMessageV2Parsing.serializeTeleporterMessageV2(message)
        );
    }

    /**
     * @notice Registers or updates the Merkle commitment for a validator set keyed by Avalanche
     * blockchain ID.
     */
    function registerValidatorSet(ICMMessage calldata message, bytes32 signingChainID) external {
        require(pChainInitialized(), "No P-chain validator set registered.");
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        ValidatorSetMerkleCommitment memory newCommitment =
            ValidatorSets.parseMerkleCommitment(message.rawMessage);
        bytes32 payloadBlockchainID = newCommitment.avalancheBlockchainID;

        // Initial registration must always be signed by the P-Chain.
        // For subsequent updates, the signing authority must be either the target chain itself if already registered
        // or the P-Chain if allowPChainFalback is enabled
        bool selfSigned = (signingChainID == payloadBlockchainID);
        bool pChainSigned = (signingChainID == pChainID);

        if (!isRegistered(payloadBlockchainID)) {
            require(pChainSigned, "Initial registration must be signed by the P-Chain");
        } else {
            require(selfSigned || (pChainSigned && allowPChainFallback), "Invalid signing chain");
        }

        verifyICMMessage(message, signingChainID);

        _valSetCommitments[payloadBlockchainID] = newCommitment;
        emit ValidatorSetRegistered(payloadBlockchainID);
    }

    /**
     * @notice Verifies that a TeleporterICMMessage was signed by a quorum of the validator set
     * committed to under the Merkle root registered for `message.sourceBlockchainID`.
     *
     * @dev Rather than verifying a Merkle multi-inclusion proof and an aggregate BLS signature
     * on-chain, this checks a single SP1 proof carried in `message.attestation` whose public
     * values bind it to the stored commitment and to this message.
     */
    function verifyMessage(
        TeleporterICMMessage calldata message
    ) external view returns (bool) {
        require(pChainInitialized(), "No P-chain validator set registered.");
        require(isRegistered(message.sourceBlockchainID), "No validator set registered to given ID");
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            message.sourceNetworkID,
            message.sourceBlockchainID,
            address(this),
            TeleporterMessageV2Parsing.serializeTeleporterMessageV2(message.message)
        );
        return _verifyZKAttestation(message.attestation, signedData, message.sourceBlockchainID);
    }

    /**
     * @notice Returns the current validator set commitment registered for the given Avalanche blockchain ID.
     */
    function getValidatorSetCommitment(
        bytes32 avalancheBlockchainID
    ) external view returns (ValidatorSetMerkleCommitment memory) {
        return _valSetCommitments[avalancheBlockchainID];
    }

    /**
     * @notice Verify the message. Does the following checks:
     *   1. P-chain root of trust has been initialized.
     *   2. The given chain has a registered validator set commitment.
     *   3. The message is intended for this network.
     *   4. The attestation is verified against `avalancheBlockchainID`'s stored commitment.
     * Reverts if any check fails.
     *
     * @dev Used to verify validator set update messages and called by `registerValidatorSet`.
     * These messages differ from Teleporter application messages in that they are emitted
     * directly by the P-chain rather than from a contract invoking `sendWarpMessage`,
     * so the warp preimage's `originSenderAddress` is `address(0)` instead of `address(this)`.
     */
    function verifyICMMessage(
        ICMMessage calldata message,
        bytes32 avalancheBlockchainID
    ) public view {
        require(pChainInitialized(), "P-chain not initialized");
        require(isRegistered(avalancheBlockchainID), "No validator set registered to given ID");
        require(message.sourceNetworkID == avalancheNetworkID, "Network ID mismatch");

        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            message.sourceNetworkID, message.sourceBlockchainID, address(0), message.rawMessage
        );
        _verifyZKAttestation(message.attestation, signedData, avalancheBlockchainID);
    }

    /**
     * @dev Check if a P-chain validator set has been completely registered.
     * Until it has, no functions except updating this validator set are
     * permitted.
     */
    function pChainInitialized() public view returns (bool) {
        return _valSetCommitments[pChainID].totalWeight != 0;
    }

    /**
     * @notice Check if a validator set is registered.
     */
    function isRegistered(
        bytes32 avalancheBlockchainID
    ) public view returns (bool) {
        return _valSetCommitments[avalancheBlockchainID].totalWeight != 0;
    }

    /**
     * @notice Verify the ZK attestation against a message. Performs the following checks:
     *   1. The proof's committed source blockchain ID and validator set Merkle root match
     *      the stored on-chain commitment.
     *   2. The proof's committed message hash matches the message being verified.
     *   3. The committed signing weight meets the configured quorum of the stored total weight.
     *   4. The SP1 ZK proof itself is valid against the attestation program's verification key.
     * Reverts if any of the above checks fail, otherwise returns true.
     *
     * @dev `attestation` is `abi.encode(bytes publicValues, bytes proofBytes)`. The public values
     * are produced by the guest program, which performs the Merkle multi-inclusion proof and the
     * aggregate BLS12-381 signature check in-circuit; this function checks those public
     * values against the on-chain state. The final `verifyProof` call reverts on an invalid ZK proof.
     */
    function _verifyZKAttestation(
        bytes calldata attestation,
        bytes memory signedData,
        bytes32 blockchainID
    ) internal view returns (bool) {
        (bytes memory publicValues, bytes memory proofBytes) =
            abi.decode(attestation, (bytes, bytes));

        PublicValues memory pv = abi.decode(publicValues, (PublicValues));
        ValidatorSetMerkleCommitment storage commitment = _valSetCommitments[blockchainID];

        // Check the ZK proof's public values against the on-chain commitment
        require(pv.sourceBlockchainID == blockchainID, "public value chain ID mismatch");
        require(pv.root == commitment.root, "public value root mismatch");
        require(pv.messageHash == sha256(signedData), "public value message hash mismatch");
        require(
            ValidatorSets.verifyWeight(pv.signedWeight, commitment.totalWeight),
            "stake-weighted quorum threshold not met"
        );

        ISP1Verifier(sp1Verifier).verifyProof(attestationProgramVKey, publicValues, proofBytes);
        return true;
    }
}
