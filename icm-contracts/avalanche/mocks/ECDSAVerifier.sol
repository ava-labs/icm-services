// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import {IMessageVerifier, ICMMessage} from "../../common/ITeleporterMessengerV2.sol";

contract ECDSAVerifier is IMessageVerifier {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    address public immutable trustedSigner;

    /**
     * @notice Sets the trusted signer address at deployment time.
     * @param signer The address corresponding to the off-chain private key.
     */
    constructor(
        address signer
    ) {
        require(signer != address(0), "Invalid signer address");
        trustedSigner = signer;
    }

    /**
     * @notice Verifies that the provided message was signed by TRUSTED_SIGNER.
     * @dev The 'attestation' field is expected to contain the 65-byte ECDSA signature.
     */
    function verifyMessage(
        ICMMessage calldata message
    ) external view override returns (bool) {
        bytes32 dataHash = keccak256(abi.encode(message.message, message.sourceBlockchainID));
        // Apply EIP-191 prefix. See https://github.com/OpenZeppelin/openzeppelin-contracts/blob/75973f63b5a84dd2fc998b5f329f1e254b0fdc77/contracts/utils/cryptography/ECDSA.sol#L50
        bytes32 digest = dataHash.toEthSignedMessageHash();
        address recoveredSigner = digest.recover(message.attestation);
        return recoveredSigner == trustedSigner;
    }
}
