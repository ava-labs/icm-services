// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./ECDSAVerifier.sol";

contract ECDSAVerifier is IMessageVerifier {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    address public immutable TRUSTED_SIGNER;

    /**
     * @notice Sets the trusted signer address at deployment time.
     * @param _trustedSigner The address corresponding to the off-chain private key.
     */
    constructor(address _trustedSigner) {
        require(_trustedSigner != address(0), "Invalid signer address");
        TRUSTED_SIGNER = _trustedSigner;
    }

    /**
     * @notice Verifies that the provided message was signed by TRUSTED_SIGNER.
     * @dev The 'attestation' field is expected to contain the 65-byte ECDSA signature.
     */
    function verifyMessage(
        ICMMessage calldata message
    ) external view override returns (bool) {
        
        // Reconstruct the digest 
        bytes32 dataHash = keccak256(abi.encode(
            message.unsignedMessage, 
            message.sourceBlockchainID
        ));

        // Apply EIP-191 prefix. See https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/cryptography/ECDSA.sol#L50
        bytes32 digest = dataHash.toEthSignedMessageHash();

        // Recover the signer from the attestation
        address recoveredSigner = digest.recover(message.attestation);

        // Check if the recovered signer matches the trusted signer
        return recoveredSigner == TRUSTED_SIGNER;
    }
}
