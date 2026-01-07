// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {BLST} from "../utils/BLST.sol";

contract BLSTUtilsTest is Test {
    function testCreateAndVerifySignature(bytes calldata message) public view {
        uint256 secretKey = 273952;
        bytes memory pk = BLST.getPublicKeyFromSecret(secretKey);
        bytes memory sig = BLST.createSignature(secretKey, message);
        bool valid = BLST.verifySignature(pk, sig, message);
        assertTrue(valid);
        // check that everything works the same when making an aggregate signature with a single secret key
        uint256[] memory secretKeys = new uint256[](1);
        secretKeys[0] = secretKey;
        bytes memory sigAggregate = BLST.createAggregateSignature(secretKeys, message);
        assertEq(sig, sigAggregate);
        bool validAggregate = BLST.verifySignature(pk, sigAggregate, message);
        assertTrue(validAggregate);
    }

    function testCreateAndVerifyAggregateSignature(bytes calldata message) public view {
        uint256[] memory secretKeys = new uint256[](2);
        secretKeys[0] = 273952;
        secretKeys[1] = 93293485;
        bytes[] memory pks = new bytes[](2);
        pks[0] = BLST.getPublicKeyFromSecret(secretKeys[0]);
        pks[1] = BLST.getPublicKeyFromSecret(secretKeys[1]);
        bytes memory sig = BLST.createAggregateSignature(secretKeys, message);
        bool valid = BLST.verifyAggregateSignature(pks, sig, message);
        assertTrue(valid);
    }

    function testNoAggregatePublicKeysSuccess() public view {
        // Test data from bls_formatter.go output
        bytes memory pk1 =
            hex"00000000000000000000000000000000105fdda3f4e22ea318d3831af1de2ef487f4b0506dfe61dca9cbd5c042484b3ecad5565855a557c237081456979433ea0000000000000000000000000000000013711c822adbbcb4f5811d48576f14c5ab78a9e5313636bf5a3722ad60ecdb96629f7e7d76873291960f6d08ab312869";

        bytes[] memory publicKeys = new bytes[](1);
        publicKeys[0] = pk1;
        bytes memory aggregated = BLST.aggregatePublicKeys(publicKeys);
        assertEq(aggregated, pk1);
    }

    function testAggregatePublicKeysSuccess() public view {
        // Test data from bls_formatter.go output
        bytes memory pk1 =
            hex"00000000000000000000000000000000105fdda3f4e22ea318d3831af1de2ef487f4b0506dfe61dca9cbd5c042484b3ecad5565855a557c237081456979433ea0000000000000000000000000000000013711c822adbbcb4f5811d48576f14c5ab78a9e5313636bf5a3722ad60ecdb96629f7e7d76873291960f6d08ab312869";
        bytes memory pk2 =
            hex"000000000000000000000000000000000f872fdc2f552676f3297330d13d888f9921038cce6cc087f1087d315363c79847e43f85d836a796cf6ddf101d13e2d9000000000000000000000000000000000b1737099061d418482275164132e3cea936ebb202661cac3df9c2dfba028fd164a947b0f62b38cae6dba6784ba475d8";
        bytes memory expectedAggregated =
            hex"00000000000000000000000000000000086401cb5b1276627747ca377b2de206fd0b6427bac06fdbd21e4cf5c71b57676367e5cc198c54c8ab0e4d0e5fa05663000000000000000000000000000000001436d02c536eb6e244867b806bc9453b367e923edcf1ada35d743122c5d2503d4114d8c776b64e67d22f3c5c6e4583e7";

        bytes[] memory publicKeys = new bytes[](2);
        publicKeys[0] = pk1;
        publicKeys[1] = pk2;

        bytes memory aggregated = BLST.aggregatePublicKeys(publicKeys);
        assertEq(aggregated, expectedAggregated);
    }

    function testVerifyBLSAggregateSignatureSuccess() public view {
        bytes memory message = "Hello Ethereum, it's Avalanche.";
        bytes[] memory pks = new bytes[](2);
        pks[0] =
            hex"000000000000000000000000000000000d89fe5ee3754167f889ff991c81e1717b665690955ad302a50e285dbe68aac93d03177fba54622bf5c4926ce90d0ac60000000000000000000000000000000015874901c0288afcdf442b402199d6a2ed38c5ac28d055637e7bafb44dc4a46a56db1c74caf0e339f6ed665dbc0c0c60";
        pks[1] =
            hex"00000000000000000000000000000000001f2946023547012bf42ca45962136ebca74aeea4291a25e9bf46bf37c5fdd37ecad6d5a389bcecf3f3d1f4f73c44db000000000000000000000000000000000385c2753673a29e199f0a0dc02772b5e2dec7a3aa70e012b313195906a27ac95d0a4963d96340df6f4d13cf40f7c78f";
        bytes memory sig =
            hex"0ace4f38b5e22e8cfd27951b8327d5260def6e6370e334c3f74fb9b5ccd209d0d80d1a34bb804f6ac40dd991141d87b5147a0a08008b085fd27341456bb43654fed88ebca4407d0b5ae66a67df8141edc58fced0468b5e34cec33dc2d58f023d1770616fe829ef9af2350c7c9d7b72ee1c05edefc7cd831ac13114e8ae7e1466c7ef2990a8e6387f9def145be2c9be8b0cbc8f7e6ebcb1b099a1785ea2001d70887faaaf86c937a292049134d02c7bc9aeaed7a3354673985a88a6e3aff66ccd";
        bool valid = BLST.verifyAggregateSignature(pks, sig, message);
        assertTrue(valid);
    }

    function testFormatUncompressedBLSPublicKey() public pure {
        bytes memory unformatted =
            hex"074a926af04042ab5876a9ba4297a806a34732d7df7a33b2d34d8e77313c34c49f365eba61e923acfbdfa160c3e51d2710b93b3ff2351cc019178b142ecb2eed17b0c99c2b6e9c30bedae8cfc26f8f7cdf75b9f0dbb2ea31e11117a04dbe962f";
        bytes memory padded =
            hex"00000000000000000000000000000000074a926af04042ab5876a9ba4297a806a34732d7df7a33b2d34d8e77313c34c49f365eba61e923acfbdfa160c3e51d270000000000000000000000000000000010b93b3ff2351cc019178b142ecb2eed17b0c99c2b6e9c30bedae8cfc26f8f7cdf75b9f0dbb2ea31e11117a04dbe962f";

        assertEq(padded, BLST.padUncompressedBLSPublicKey(unformatted));
    }
}
