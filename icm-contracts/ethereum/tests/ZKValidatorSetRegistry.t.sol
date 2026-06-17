// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {ZKValidatorSetRegistry} from "../ZKValidatorSetRegistry.sol";
import {ValidatorSets, ValidatorSetMerkleCommitment} from "../utils/ValidatorSets.sol";
import {ISP1Verifier} from "@sp1-contracts/ISP1Verifier.sol";
import {ICMMessage} from "../../common/ICM.sol";
import {
    TeleporterMessageV2,
    TeleporterMessageV2Parsing,
    TeleporterICMMessage
} from "../../common/TeleporterMessageV2.sol";

/**
 * @dev This mock verifier does not check the vkey, public values, or
 * proof bytes. The proof soundness is out of scope for these unit tests.
 */
contract MockSP1Verifier is ISP1Verifier {
    bool public shouldRevert;

    function setShouldRevert(
        bool v
    ) external {
        shouldRevert = v;
    }

    function verifyProof(bytes32, bytes calldata, bytes calldata) external view {
        require(!shouldRevert, "mock: invalid proof");
    }
}

contract ZKValidatorSetRegistryCommon is Test {
    uint32 public constant NETWORK_ID = 1;
    bytes32 public constant PCHAIN_BLOCKCHAIN_ID =
        0x3d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7;
    bytes32 public constant UNREGISTERED_BLOCKCHAIN_ID =
        0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef;

    // Opaque to the contract: it only compares these bytes against the committed public values.
    bytes32 public constant PCHAIN_GENESIS_ROOT = bytes32(uint256(0xC0FFEE));
    uint64 public constant PCHAIN_TOTAL_WEIGHT = 100;
    bytes32 public constant PROGRAM_VKEY = bytes32(uint256(0x1111));

    MockSP1Verifier internal _verifier;

    /**
     * @dev Sets up and deploys a mock verifier and ZKValidatorSetRegistry bootstrapped with a
     * P-chain commitment under PCHAIN_GENESIS_ROOT.
     */
    function _deployRegistry(
        bool allowPChainFallback
    ) internal returns (ZKValidatorSetRegistry) {
        _verifier = new MockSP1Verifier();
        return new ZKValidatorSetRegistry({
            avalancheNetworkID_: NETWORK_ID,
            pChainID_: PCHAIN_BLOCKCHAIN_ID,
            pChainGenesisRoot: PCHAIN_GENESIS_ROOT,
            pChainTotalWeight: PCHAIN_TOTAL_WEIGHT,
            pChainHeight: 1,
            pChainTimestamp: 1,
            allowPChainFallback_: allowPChainFallback,
            sp1Verifier_: _verifier,
            attestationProgramVKey_: PROGRAM_VKEY
        });
    }

    function _publicValues(
        bytes32 sourceBlockchainID,
        bytes32 root,
        bytes memory signedData,
        uint64 signedWeight
    ) internal pure returns (bytes memory) {
        return abi.encode(
            ZKValidatorSetRegistry.PublicValues({
                sourceBlockchainID: sourceBlockchainID,
                root: root,
                messageHash: sha256(signedData),
                signedWeight: signedWeight
            })
        );
    }

    function _attestation(
        bytes memory publicValues
    ) internal pure returns (bytes memory) {
        return abi.encode(publicValues, bytes("proof"));
    }

    function _validAttestation(
        bytes32 sourceBlockchainID,
        bytes32 root,
        bytes memory signedData,
        uint64 signedWeight
    ) internal pure returns (bytes memory) {
        return _attestation(_publicValues(sourceBlockchainID, root, signedData, signedWeight));
    }

    /**
     * @dev Builds an ICMMessage for register and update. Uses address(0) as the origin sender,
     * matching the warp preimage that verifyICMMessage reconstructs. pv.sourceBlockchainID is
     * set to the warp source (which equals the signingChainID in these tests), and pvRoot is
     * the stored root of the signing chain.
     */
    function _signedICMMessage(
        bytes32 warpSourceBlockchainID,
        bytes memory rawMessage,
        bytes32 pvRoot,
        uint64 signedWeight
    ) internal pure returns (ICMMessage memory) {
        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            NETWORK_ID, warpSourceBlockchainID, address(0), rawMessage
        );
        return ICMMessage({
            rawMessage: rawMessage,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: warpSourceBlockchainID,
            attestation: _validAttestation(warpSourceBlockchainID, pvRoot, signedData, signedWeight)
        });
    }

    function _commitment(
        bytes32 blockchainID,
        bytes32 root,
        uint64 totalWeight
    ) internal pure returns (bytes memory) {
        return ValidatorSets.serializeMerkleCommitment(
            ValidatorSetMerkleCommitment({
                avalancheBlockchainID: blockchainID,
                root: root,
                totalWeight: totalWeight,
                pChainHeight: 2,
                pChainTimestamp: 2
            })
        );
    }

    function _emptyMessage() internal pure returns (TeleporterICMMessage memory) {
        TeleporterMessageV2 memory inner;
        return TeleporterICMMessage({
            message: inner,
            sourceNetworkID: 0,
            sourceBlockchainID: bytes32(0),
            attestation: new bytes(0)
        });
    }
}

/**
 * @dev Tests for the ZKValidatorSetRegistry.verifyMessage pipeline. The registry is
 * bootstrapped with a P-chain commitment and the SP1 verifier is mocked, so each test isolates
 * either the binding checks (sourceBlockchainID, root, messageHash, quorum) or the proof verification.
 */
contract ZKValidatorSetRegistryVerifyMessageTest is ZKValidatorSetRegistryCommon {
    ZKValidatorSetRegistry private _registry;

    function setUp() public {
        _registry = _deployRegistry(true);
    }

    /// @dev verifyMessage reverts when the source blockchain ID has no commitment.
    function testVerifyMessageUnregisteredSourceReverts() public {
        TeleporterICMMessage memory message = _emptyMessage();
        message.sourceNetworkID = NETWORK_ID;
        message.sourceBlockchainID = UNREGISTERED_BLOCKCHAIN_ID;

        vm.expectRevert(bytes("No validator set registered to given ID"));
        _registry.verifyMessage(message);
    }

    /// @dev verifyMessage reverts on a network ID mismatch (source is registered).
    function testVerifyMessageNetworkMismatchReverts() public {
        TeleporterICMMessage memory message = _emptyMessage();
        message.sourceNetworkID = NETWORK_ID + 1;
        message.sourceBlockchainID = PCHAIN_BLOCKCHAIN_ID;

        vm.expectRevert(bytes("Network ID mismatch"));
        _registry.verifyMessage(message);
    }

    /// @dev A malformed attestation fails to abi.decode and reverts.
    function testVerifyMessageMalformedAttestationReverts() public {
        (TeleporterICMMessage memory message,) = _buildICMMessage(PCHAIN_BLOCKCHAIN_ID);
        message.attestation = hex"deadbeef";

        vm.expectRevert();
        _registry.verifyMessage(message);
    }

    /// @dev Returns false when the committed sourceBlockchainID doesn't match the message.
    function testVerifyMessageWrongSourceBlockchainID() public {
        (TeleporterICMMessage memory message, bytes memory signedData) =
            _buildICMMessage(PCHAIN_BLOCKCHAIN_ID);
        message.attestation =
            _validAttestation(UNREGISTERED_BLOCKCHAIN_ID, PCHAIN_GENESIS_ROOT, signedData, 100);

        vm.expectRevert(bytes("public value chain ID mismatch"));
        _registry.verifyMessage(message);
    }

    /// @dev Returns false when the committed root doesn't match the stored commitment.
    function testVerifyMessageWrongRootReverts() public {
        (TeleporterICMMessage memory message, bytes memory signedData) =
            _buildICMMessage(PCHAIN_BLOCKCHAIN_ID);
        message.attestation =
            _validAttestation(PCHAIN_BLOCKCHAIN_ID, bytes32(uint256(0xBAD)), signedData, 100);

        vm.expectRevert(bytes("public value root mismatch"));
        _registry.verifyMessage(message);
    }

    /// @dev Returns false when the committed messageHash doesn't match sha256(signedData).
    function testVerifyMessageWrongMessageHashReverts() public {
        (TeleporterICMMessage memory message,) = _buildICMMessage(PCHAIN_BLOCKCHAIN_ID);
        // Hash a different preimage so the binding mismatches.
        bytes memory pv = abi.encode(
            ZKValidatorSetRegistry.PublicValues({
                sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
                root: PCHAIN_GENESIS_ROOT,
                messageHash: sha256(bytes("wrong signed data")),
                signedWeight: 100
            })
        );
        message.attestation = _attestation(pv);

        vm.expectRevert(bytes("public value message hash mismatch"));
        _registry.verifyMessage(message);
    }

    /// @dev Returns false when signed weight is below quorum threshold
    function testVerifyMessageBelowQuorumReverts() public {
        (TeleporterICMMessage memory message, bytes memory signedData) = _buildICMMessage(PCHAIN_BLOCKCHAIN_ID);
        message.attestation =
            _validAttestation(PCHAIN_BLOCKCHAIN_ID, PCHAIN_GENESIS_ROOT, signedData, 66);

        vm.expectRevert(bytes("stake-weighted quorum threshold not met"));
        _registry.verifyMessage(message);
    }

    /// @dev Reverts with an invalid mocked proof
    function testVerifyMessageInvalidProofReverts() public {
        (TeleporterICMMessage memory message, bytes memory signedData) =
            _buildICMMessage(PCHAIN_BLOCKCHAIN_ID);
        message.attestation =
            _validAttestation(PCHAIN_BLOCKCHAIN_ID, PCHAIN_GENESIS_ROOT, signedData, 100);

        _verifier.setShouldRevert(true);
        vm.expectRevert(bytes("mock: invalid proof"));
        _registry.verifyMessage(message);
    }

    /// @dev Happy path: the public bindings are correct, weight is above threshold, and the proof is valid.
    function testVerifyMessageSuccess() public {
        (TeleporterICMMessage memory message, bytes memory signedData) =
            _buildICMMessage(PCHAIN_BLOCKCHAIN_ID);
        message.attestation =
            _validAttestation(PCHAIN_BLOCKCHAIN_ID, PCHAIN_GENESIS_ROOT, signedData, 100);

        _verifier.setShouldRevert(false);
        assertTrue(_registry.verifyMessage(message));
    }

    /// @dev Quorum boundary: exactly 67 of 100 meets the threshold
    function testVerifyMessageQuorumBoundarySuccess() public {
        (TeleporterICMMessage memory message, bytes memory signedData) =
            _buildICMMessage(PCHAIN_BLOCKCHAIN_ID);
        message.attestation =
            _validAttestation(PCHAIN_BLOCKCHAIN_ID, PCHAIN_GENESIS_ROOT, signedData, 67);

        _verifier.setShouldRevert(false);
        assertTrue(_registry.verifyMessage(message));
    }

    /// @dev Builds the verifyMessage warp preimage (address(this) sender) and a base message.
    function _buildICMMessage(
        bytes32 sourceBlockchainID
    ) internal view returns (TeleporterICMMessage memory, bytes memory) {
        TeleporterMessageV2 memory inner;
        bytes memory innerSerialized =
            TeleporterMessageV2Parsing.serializeTeleporterMessageV2(inner);
        bytes memory signedData = ValidatorSets.buildUnsignedWarpMessage(
            NETWORK_ID, sourceBlockchainID, address(_registry), innerSerialized
        );
        TeleporterICMMessage memory message = TeleporterICMMessage({
            message: inner,
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: sourceBlockchainID,
            attestation: new bytes(0)
        });
        return (message, signedData);
    }
}

/**
 * @dev Tests for register and update. Register tests prove new chains can be added under P-chain
 * attestations. Update tests prove an existing chain's commitment can be replaced under its own
 * or with fallback the P-chain's attestation. The SP1 verifier is mocked throughout.
 */
contract ZKValidatorSetRegistryRegisterUpdateTest is ZKValidatorSetRegistryCommon {
    bytes32 internal constant _NEW_CHAIN_ID =
        0xabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcdabcd;

    ZKValidatorSetRegistry private _registry;

    function setUp() public {
        _registry = _deployRegistry(true);
        _verifier.setShouldRevert(false);
    }

    /// @dev First registration of a new chain, attested by the P-chain set.
    function testRegisterNewChainSuccess() public {
        bytes memory raw = _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10);
        ICMMessage memory message =
            _signedICMMessage(PCHAIN_BLOCKCHAIN_ID, raw, PCHAIN_GENESIS_ROOT, 100);

        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);

        ValidatorSetMerkleCommitment memory stored =
            _registry.getValidatorSetCommitment(_NEW_CHAIN_ID);
        assertEq(stored.avalancheBlockchainID, _NEW_CHAIN_ID);
        assertEq(stored.root, bytes32(uint256(0xABABA)));
        assertEq(stored.totalWeight, 10);
    }

    /// @dev registerValidatorSet reverts when the message's network ID doesn't match.
    function testRegisterRevertsOnNetworkMismatch() public {
        ICMMessage memory message = ICMMessage({
            rawMessage: _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10),
            sourceNetworkID: NETWORK_ID + 1,
            sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            attestation: new bytes(0)
        });

        vm.expectRevert(bytes("Network ID mismatch"));
        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);
    }

    /// @dev Update of the P-chain's own commitment, attested by the P-chain set.
    function testUpdateExistingChainSuccess() public {
        bytes memory raw = _commitment(PCHAIN_BLOCKCHAIN_ID, bytes32(uint256(0xABABA)), 10);
        ICMMessage memory message =
            _signedICMMessage(PCHAIN_BLOCKCHAIN_ID, raw, PCHAIN_GENESIS_ROOT, 100);

        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);

        ValidatorSetMerkleCommitment memory stored =
            _registry.getValidatorSetCommitment(PCHAIN_BLOCKCHAIN_ID);
        assertEq(stored.root, bytes32(uint256(0xABABA)));
        assertEq(stored.totalWeight, 10);
    }

    /// @dev Reverts when the payload chain isn't registered and the signer isn't the P-chain.
    function testUpdateRevertsOnUnregisteredChain() public {
        ICMMessage memory message = ICMMessage({
            rawMessage: _commitment(UNREGISTERED_BLOCKCHAIN_ID, bytes32(uint256(0xABABA)), 10),
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: UNREGISTERED_BLOCKCHAIN_ID,
            attestation: new bytes(0)
        });

        vm.expectRevert(bytes("Initial registration must be signed by the P-Chain"));
        _registry.registerValidatorSet(message, UNREGISTERED_BLOCKCHAIN_ID);
    }

    /// @dev Reverts when a first-time registration is self-signed rather than P-chain signed.
    function testRegisterRevertsOnSelfSignedFirstTimeRegistration() public {
        ICMMessage memory message = ICMMessage({
            rawMessage: _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10),
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: _NEW_CHAIN_ID,
            attestation: new bytes(0)
        });

        vm.expectRevert(bytes("Initial registration must be signed by the P-Chain"));
        _registry.registerValidatorSet(message, _NEW_CHAIN_ID);
    }

    /// @dev Once registered, an update signed by neither the chain itself nor the P-chain reverts.
    function testRegisterRevertsOnInvalidSigningChain() public {
        bytes memory raw = _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10);
        _registry.registerValidatorSet(
            _signedICMMessage(PCHAIN_BLOCKCHAIN_ID, raw, PCHAIN_GENESIS_ROOT, 100),
            PCHAIN_BLOCKCHAIN_ID
        );

        ICMMessage memory update = ICMMessage({
            rawMessage: _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xCAFE)), 20),
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            attestation: new bytes(0)
        });

        vm.expectRevert(bytes("Invalid signing chain"));
        _registry.registerValidatorSet(update, UNREGISTERED_BLOCKCHAIN_ID);
    }

    /// @dev A binding mismatch (wrong root) makes _verifyZKAttestation return false.
    function testRegisterRevertsOnBindingMismatch() public {
        bytes memory raw = _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10);
        // pvRoot is wrong
        ICMMessage memory message =
            _signedICMMessage(PCHAIN_BLOCKCHAIN_ID, raw, bytes32(uint256(0xBAD)), 100);

        vm.expectRevert(bytes("public value root mismatch"));
        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);
    }

    /// @dev Below-quorum signed weight fails the weight binding
    function testRegisterRevertsBelowQuorum() public {
        bytes memory raw = _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10);
        ICMMessage memory message =
            _signedICMMessage(PCHAIN_BLOCKCHAIN_ID, raw, PCHAIN_GENESIS_ROOT, 66);

        vm.expectRevert(bytes("stake-weighted quorum threshold not met"));
        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);
    }

    /// @dev With bindings correct, an invalid proof reverts inside the verifier (distinct from
    /// the binding-mismatch revert string above).
    function testRegisterRevertsOnInvalidProof() public {
        bytes memory raw = _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10);
        ICMMessage memory message =
            _signedICMMessage(PCHAIN_BLOCKCHAIN_ID, raw, PCHAIN_GENESIS_ROOT, 100);

        _verifier.setShouldRevert(true);
        vm.expectRevert(bytes("mock: invalid proof"));
        _registry.registerValidatorSet(message, PCHAIN_BLOCKCHAIN_ID);
    }

    /// @dev Fallback path: a registered L1's commitment is updated via P-chain attestation.
    function testUpdateRegisteredChainViaPChainSignature() public {
        _registry.registerValidatorSet(
            _signedICMMessage(
                PCHAIN_BLOCKCHAIN_ID,
                _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10),
                PCHAIN_GENESIS_ROOT,
                100
            ),
            PCHAIN_BLOCKCHAIN_ID
        );

        _registry.registerValidatorSet(
            _signedICMMessage(
                PCHAIN_BLOCKCHAIN_ID,
                _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xCAFE)), 20),
                PCHAIN_GENESIS_ROOT,
                100
            ),
            PCHAIN_BLOCKCHAIN_ID
        );

        ValidatorSetMerkleCommitment memory stored =
            _registry.getValidatorSetCommitment(_NEW_CHAIN_ID);
        assertEq(stored.root, bytes32(uint256(0xCAFE)));
        assertEq(stored.totalWeight, 20);
    }

    /// @dev With fallback disabled, a P-chain attested update of a registered L1 reverts.
    function testUpdateRevertsPChainFallbackDisabled() public {
        ZKValidatorSetRegistry registry = _deployRegistry(false);
        _verifier.setShouldRevert(false);

        registry.registerValidatorSet(
            _signedICMMessage(
                PCHAIN_BLOCKCHAIN_ID,
                _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xABABA)), 10),
                PCHAIN_GENESIS_ROOT,
                100
            ),
            PCHAIN_BLOCKCHAIN_ID
        );

        ICMMessage memory update = ICMMessage({
            rawMessage: _commitment(_NEW_CHAIN_ID, bytes32(uint256(0xCAFE)), 20),
            sourceNetworkID: NETWORK_ID,
            sourceBlockchainID: PCHAIN_BLOCKCHAIN_ID,
            attestation: new bytes(0)
        });

        vm.expectRevert(bytes("Invalid signing chain"));
        registry.registerValidatorSet(update, PCHAIN_BLOCKCHAIN_ID);
    }

    /// @dev Happy path: P-chain registers an L1, then the L1 self-signs an update to its own
    /// commitment (attested against the L1's stored root).
    function testL1SelfSignedUpdateSuccess() public {
        bytes32 l1Root = bytes32(uint256(0x1111));
        _registry.registerValidatorSet(
            _signedICMMessage(
                PCHAIN_BLOCKCHAIN_ID,
                _commitment(_NEW_CHAIN_ID, l1Root, 100),
                PCHAIN_GENESIS_ROOT,
                100
            ),
            PCHAIN_BLOCKCHAIN_ID
        );

        // Self-signed update: signer is the L1, pvRoot is the L1's stored root.
        _registry.registerValidatorSet(
            _signedICMMessage(
                _NEW_CHAIN_ID, _commitment(_NEW_CHAIN_ID, bytes32(uint256(0x2222)), 50), l1Root, 100
            ),
            _NEW_CHAIN_ID
        );

        ValidatorSetMerkleCommitment memory stored =
            _registry.getValidatorSetCommitment(_NEW_CHAIN_ID);
        assertEq(stored.root, bytes32(uint256(0x2222)));
        assertEq(stored.totalWeight, 50);
    }
}
