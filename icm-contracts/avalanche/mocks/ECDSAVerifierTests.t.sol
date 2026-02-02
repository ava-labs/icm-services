// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Reference: This is core logic from the Succinct Telepathy Library, which can be found at https://github.com/succinctlabs/telepathy-contracts/blob/main/test/libraries/SimpleSerialize.t.sol
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {TeleporterMessageReceipt} from "@teleporter/ITeleporterMessenger.sol";
import {ICMMessage, TeleporterMessageV2} from "../../common/ITeleporterMessengerV2.sol";
import {ECDSAVerifier} from "./ECDSAVerifier.sol";

contract ECDSAVerifierTest is Test {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    ECDSAVerifier public verifier;

    uint256 public signerPrivateKey;
    address public signerAddress;

    uint256 public attackerPrivateKey;
    address public attackerAddress;

    function setUp() public {
        (signerAddress, signerPrivateKey) = makeAddrAndKey("trustedSigner");
        (attackerAddress, attackerPrivateKey) = makeAddrAndKey("attacker");

        // Deploy the contract with the trusted signer
        verifier = new ECDSAVerifier(signerAddress);
    }

    function testVerifyMessageSuccess() public view {
        TeleporterMessageV2 memory tMsg = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: address(0xABC),
            originTeleporterAddress: address(0xDEF),
            destinationBlockchainID: bytes32(uint256(999)),
            destinationAddress: address(0x456),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"1122"
        });

        bytes32 sourceBlockchainID = bytes32(uint256(998));
        bytes memory signature = _sign(tMsg, sourceBlockchainID, signerPrivateKey);

        ICMMessage memory icmMsg = ICMMessage({
            message: tMsg,
            sourceNetworkID: 0,
            sourceBlockchainID: sourceBlockchainID,
            attestation: signature
        });

        assertTrue(verifier.verifyMessage(icmMsg));
    }

    function testVerifyMessageFailWrongSigner() public view {
        TeleporterMessageV2 memory tMsg = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: address(0xABC),
            originTeleporterAddress: address(0xDEF),
            destinationBlockchainID: bytes32(uint256(999)),
            destinationAddress: address(0x456),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"1122"
        });

        bytes32 sourceBlockchainID = bytes32(uint256(998));
        // Sign with the wrong private key
        bytes memory signature = _sign(tMsg, sourceBlockchainID, attackerPrivateKey);

        ICMMessage memory icmMsg = ICMMessage({
            message: tMsg,
            sourceNetworkID: 0,
            sourceBlockchainID: sourceBlockchainID,
            attestation: signature
        });

        assertFalse(verifier.verifyMessage(icmMsg));
    }

    function testVerifyMessageFailTamperedData() public view {
        TeleporterMessageV2 memory message = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: address(0xABC),
            originTeleporterAddress: address(0xDEF),
            destinationBlockchainID: bytes32(uint256(999)),
            destinationAddress: address(0x456),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"1122"
        });

        bytes32 sourceBlockchainID = bytes32(uint256(998));
        // Sign the real data
        bytes memory signature = _sign(message, sourceBlockchainID, signerPrivateKey);
        // Create tampered data
        TeleporterMessageV2 memory wrongMessage = TeleporterMessageV2({
            messageNonce: 1,
            originSenderAddress: address(0xABC),
            originTeleporterAddress: address(0xDEF),
            destinationBlockchainID: bytes32(uint256(999)),
            destinationAddress: address(0x456),
            requiredGasLimit: 0,
            allowedRelayerAddresses: new address[](0),
            receipts: new TeleporterMessageReceipt[](0),
            message: hex"DEADBEEF"
        });

        ICMMessage memory icmMsg = ICMMessage({
            message: wrongMessage,
            sourceNetworkID: 0,
            sourceBlockchainID: sourceBlockchainID,
            attestation: signature
        });

        assertFalse(verifier.verifyMessage(icmMsg));
    }

    /**
     * @dev Helper function to generate ECDSA signatures.
     */
    function _sign(
        TeleporterMessageV2 memory message,
        bytes32 chainID,
        uint256 privateKey
    ) internal pure returns (bytes memory) {
        // Reconstruct the digest
        bytes32 dataHash = keccak256(abi.encode(message, chainID));

        // Apply EIP-191 prefix
        bytes32 digest = dataHash.toEthSignedMessageHash();

        // Sign the message using Foundry
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privateKey, digest);

        return abi.encodePacked(r, s, v);
    }
}
