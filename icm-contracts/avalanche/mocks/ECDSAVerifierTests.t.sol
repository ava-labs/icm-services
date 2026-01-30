// (c) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

// Reference: This is core logic from the Succinct Telepathy Library, which can be found at https://github.com/succinctlabs/telepathy-contracts/blob/main/test/libraries/SimpleSerialize.t.sol
pragma solidity ^0.8.30;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./ECDSAVerifier.sol";
import {Test} from "@forge-std/Test.sol";

contract ECDSAVerifierTest is Test {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;
    
    ECDSAVerifier verifier;
    
    uint256 signerPrivateKey;
    address signerAddress;
    
    uint256 attackerPrivateKey;
    address attackerAddress;

    function setUp() public {
        (signerAddress, signerPrivateKey) = makeAddrAndKey("trustedSigner");
        (attackerAddress, attackerPrivateKey) = makeAddrAndKey("attacker");
        
        // Deploy the contract with the trusted signer
        verifier = new ECDSAVerifier(signerAddress);
    }

    /**
     * @dev Helper function to generate ECDSA signatures. 
     */
    function _sign(
        TeleporterMessage memory message, 
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

    function test_VerifyMessage_Success() public {
        TeleporterMessage memory message = TeleporterMessage(1, address(0xABC), hex"1122");
        bytes32 chainID = bytes32(uint256(999));
        bytes memory signature = _sign(message, chainID, signerPrivateKey);

        ICMMessage memory icmMessage = ICMMessage({
            unsignedMessage: message,
            sourceBlockchainID: chainID,
            attestation: signature
        });

        assertTrue(verifier.verifyMessage(icmMessage));
    }

    function test_VerifyMessage_Fail_WrongSigner() public {
        TeleporterMessage memory message = TeleporterMessage(1, address(0xABC), hex"1122");
        bytes32 chainID = bytes32(uint256(999));
        // Sign with the wrong private key
        bytes memory signature = _sign(message, chainID, attackerPrivateKey); 

        ICMMessage memory icmMessage = ICMMessage({
            unsignedMessage: message,
            sourceBlockchainID: chainID,
            attestation: signature
        });

        assertFalse(verifier.verifyMessage(icmMessage));
    }

    function test_VerifyMessage_Fail_TamperedData() public {
        TeleporterMessage memory message = TeleporterMessage(1, address(0xABC), hex"1122");
        bytes32 chainID = bytes32(uint256(999));
        // Sign the real data
        bytes memory signature = _sign(message, chainID, signerPrivateKey);
        // Create tampered data
        TeleporterMessage memory wrongMessage = TeleporterMessage(1, address(0xABC), hex"DEADBEEF");

        ICMMessage memory icmMessage = ICMMessage({
            unsignedMessage: wrongMessage, 
            sourceBlockchainID: chainID,
            attestation: signature
        });

        assertFalse(verifier.verifyMessage(icmMessage));
    }
}
