// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity ^0.8.30;

/**
 * THIS IS LIBRARY IS UN-AUDITED CODE.
 * DO NOT USE THIS CODE IN PRODUCTION.
 */

// #[pack(contract="ICM", name="serializeICMMessage")]
// #[unpack(contract="ICM", calldata, name="parseICMMessage")]
struct ICMMessage {
    // The serialized bytes of raw message. The data and serializations formats
    // for this data will be app / contract specific
    bytes rawMessage;
    // used to distinguish between mainnet and testnets
    uint32 sourceNetworkID;
    // The blockchain on which the message originated
    bytes32 sourceBlockchainID;
    // Arbitrary bytes that is used by receiving contracts to
    // authenticate this message.
    bytes attestation;
}

/**
 * @title ICM
 * @notice Utility library for Interchain Messaging (ICM) messages. Mainly (de)serialization.
 */
library ICM {}
