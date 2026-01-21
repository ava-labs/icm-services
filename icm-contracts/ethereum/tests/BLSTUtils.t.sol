// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

import {Test} from "@forge-std/Test.sol";
import {BLST} from "../utils/BLST.sol";

contract BLSTUtilsTest is Test {
    /*
     * @dev Test that signatures created by this library pass it's own verification
     * checks, i.e. self-consistency. Creating a signature vs. an aggregate with
     * a single key should produce the same results.
     */
    function testCreateAndVerifySignature(bytes calldata message, uint256 secretKey) public view {
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

    /*
     * @dev Test self-consistency with regard to an aggregate signature from two separate
     * keys.
     */
    function testCreateAndVerifyAggregateSignature(
        bytes calldata message,
        uint256 secretKey1,
        uint256 secretKey2
    ) public view {
        uint256[] memory secretKeys = new uint256[](2);
        secretKeys[0] = secretKey1;
        secretKeys[1] = secretKey2;
        bytes[] memory pks = new bytes[](2);
        pks[0] = BLST.getPublicKeyFromSecret(secretKeys[0]);
        pks[1] = BLST.getPublicKeyFromSecret(secretKeys[1]);
        bytes memory sig = BLST.createAggregateSignature(secretKeys, message);
        bool valid = BLST.verifyAggregateSignature(pks, sig, message);
        assertTrue(valid);
    }

    /*
     * @dev Values obtained via https://iancoleman.io/eip2333/
     * This was the first derived address (index m/0) from the following seed:
     * 4c4f7f21e38afd4c586cbd1e5854450b25149ed8d9d71ca4372cb810e58a827c197cb337e0afbfedec7a0c849e405fea4e54316daf01a5b7e03a6b0a523e2fe3
     */
    function testSecretToPublicAgainstIanColeman() public view {
        uint256 secretKey = 0x316cb723e4bbdbf536d82384efe04b15484fd44afb5e579e04718c7e7eb83e0c;
        bytes memory expectedPk =
            hex"97d5726528eef5a2da8aa09bee99b04fbb3f3b7893a2988e42bfeb5af1163525c9d3832bed9e5237885339ff48d6c9fa";
        bytes memory unCompressedPK = BLST.getPublicKeyFromSecret(secretKey);
        bytes memory pk = BLST.compressPublicKey(unCompressedPK);
        assertEq(expectedPk, pk);
    }

    /*
     * @dev Using the mini-app at the bottom of https://paulmillr.com/noble/
     */
    function testSecretToPublicAgainstNoble() public view {
        uint256 secretKey = 0xf0c5bf519a6ede6be1ab684f6ecc1b129b0fc2ed95bd294bb2967077ae38a378;
        bytes memory expectedPk =
            hex"855e5129c94bb05d0bcdf0ba1e56750f9fac3da8d272baec0ce3f1fec6f22a91b84b33032a99dee48844feefc37739dc";
        bytes memory unCompressedPK = BLST.getPublicKeyFromSecret(secretKey);
        bytes memory pk = BLST.compressPublicKey(unCompressedPK);
        assertEq(expectedPk, pk);
    }

    /*
     * @dev Check signature create/verification against a test vector generated using the BLS utils in AvalancheGo.
     * https://github.com/ava-labs/avalanchego/blob/ec05f574dd61f8fbfbd53064d81f8a3a46b866ec/utils/crypto/bls/signer/localsigner/bls_test.go#L42
     */
    function testSigAgainstAvalancheGo() public view {
        uint256 secretKey = 0x42ad1a9752f8c12a295e97f8bc3d689cef5e6f2cfab204ab3ed8ebb36ee335d3;
        bytes memory expectedPk = BLST.padUncompressedBLSPublicKey(
            hex"11b840dd28418722ee77fcebc6d172065a9a6407b925486c8be7926c49f3886ad50831b74499cf6cecc629d95799f6931353dfdd0a91084d9371f541b3caba467d0a7db2e3d12011b86ffea28a814e594d6ac38c22c397afe0d5b281ca88937e"
        );

        bytes memory pk = BLST.getPublicKeyFromSecret(secretKey);
        assertEq(expectedPk, pk);

        bytes memory message = "TestVerifyValidSignature local signer";
        bytes memory expectedSig =
            hex"10c58df90f242a5ccd301e6c4d9dfa26ac53ec59674ea23e593675aa2b96e89110a23b16b4bf353e887e733382de16ce166e98b7aa50dbe6e59e756c2c7cd2d7bb9e9e3ef69747785753f04e09af0b826365288e58860bbc618ba9636e1d78a211c351ca2514d57c8b48962791c33d0c402e5377cfeb510660c29185eb2859872948208c04fc4897266ae4569aa50f0d16b25ab59f3b1f1ab1d494b74b9639097d9c4f5429e0db3d6fd29cd6c5b0a5b1122a7175749d567668a44e96f372dba3";

        bytes memory sig = BLST.createSignature(secretKey, message);
        assertEq(expectedSig, sig);
    }

    /*
     * @dev Check aggregate signature create/verification against a test vector generated using the BLS utils in
      * AvalancheGo.
     * https://github.com/ava-labs/avalanchego/blob/ec05f574dd61f8fbfbd53064d81f8a3a46b866ec/utils/crypto/bls/signer/localsigner/bls_test.go#L42
     */
    function testAggregateSigAgainstAvalancheGo() public view {
        uint256[] memory secretKeys = new uint256[](3);
        secretKeys[0] = 0x561fee6432efc4b109c69a2c4e1bd9cc474ce78ed7572fc4a220783030cdbcac;
        secretKeys[1] = 0x18a52533cf7debd4b90b3d35d7c2a3783b524f69a03c4d3fa64cccf635338808;
        secretKeys[2] = 0x4981010927d0c2233d810328805d451b8119f2c6ff50fc36f56ed5a56aa8690c;
        bytes memory expectedPk1 = BLST.padUncompressedBLSPublicKey(
            hex"11ae567e08314b059b33771ff6fe3dc25703b9f01e2794b0695da79285ac069e7d7471737dddde2e6fa36bc2557bc01111bcd774c0251f8c4a146287983692c1af1e869075ae8ac60107714c3c6fe91b0aa730f6f2d070026cca642515691802"
        );
        bytes memory expectedPk2 = BLST.padUncompressedBLSPublicKey(
            hex"1789b8b9a7c411eff217e3699073cf7184daa825aed5faeed58d37ebf7b4c77b95e477a9d0901a62ee09260c86081a87097cb335c95ac0d10df57a0114cfcb34142a1ce162466b5a628c074882d1837420124d7a317813f912dff063f5890e9d"
        );
        bytes memory expectedPk3 = BLST.padUncompressedBLSPublicKey(
            hex"01bcc815aee18ec84ae166176bc7ccccf42cbb44c4984f3b0aa4f8eaba18dc6bc40aaac8cc1ed4c80251a2598d4be63b0214fedd1f4c73548782e05a138863355a1a676c8d8e68a014c1120522924bc1a8cdc57a194cf918b6e6e2bad2e1da63"
        );

        bytes memory pk = BLST.getPublicKeyFromSecret(secretKeys[0]);
        assertEq(expectedPk1, pk);
        pk = BLST.getPublicKeyFromSecret(secretKeys[1]);
        assertEq(expectedPk2, pk);
        pk = BLST.getPublicKeyFromSecret(secretKeys[2]);
        assertEq(expectedPk3, pk);

        bytes memory expectedSig =
            hex"13548fe35a1be72e79719ce4631989efdaa29a88873d949201e510a6117b8c7b91dea1d5e6070f9996a31cef00afb0f404254a8979c8c7617710dfc324c989460be9c874c2584dbf114e04bcb1a9dd5894c788990948d1909f40cc692b9e160406be23ffd1d1ece752ba9c6f9e7ee7dce91c312eb3580ca4c6a72b5a4c9911ca2287a06ee6e4953a2f2cc106019c06fc19c09662244295626f93b8dcd20c09c1fd520d7245b659080e1defe1cf0c33bbb054ba462b244d3c0f23b1d398d5db7e";
        bytes memory message = "TestValidAggregation local signer";
        bytes memory aggSig = BLST.createAggregateSignature(secretKeys, message);
        assertEq(expectedSig, aggSig);
    }

    /*
     * @dev A small check that aggregating a single public key is a no-op
     */
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

    /*
     * @dev Test that padding out a public key behaves correctly and that the round trip
     * is a no-op
     */
    function testPadUncompressedBLSPublicKey() public pure {
        bytes memory unpadded =
            hex"074a926af04042ab5876a9ba4297a806a34732d7df7a33b2d34d8e77313c34c49f365eba61e923acfbdfa160c3e51d2710b93b3ff2351cc019178b142ecb2eed17b0c99c2b6e9c30bedae8cfc26f8f7cdf75b9f0dbb2ea31e11117a04dbe962f";
        bytes memory padded =
            hex"00000000000000000000000000000000074a926af04042ab5876a9ba4297a806a34732d7df7a33b2d34d8e77313c34c49f365eba61e923acfbdfa160c3e51d270000000000000000000000000000000010b93b3ff2351cc019178b142ecb2eed17b0c99c2b6e9c30bedae8cfc26f8f7cdf75b9f0dbb2ea31e11117a04dbe962f";

        assertEq(padded, BLST.padUncompressedBLSPublicKey(unpadded));
        assertEq(unpadded, BLST.unPadUncompressedBlsPublicKey(padded));
    }

    /*
     * @dev Test that padding out a signature behaves correctly and that the round trip
     * is a no-op
     */
    function testPadUncompressedBLSTSignature() public pure {
        bytes memory unpadded =
            hex"0ace4f38b5e22e8cfd27951b8327d5260def6e6370e334c3f74fb9b5ccd209d0d80d1a34bb804f6ac40dd991141d87b5147a0a08008b085fd27341456bb43654fed88ebca4407d0b5ae66a67df8141edc58fced0468b5e34cec33dc2d58f023d1770616fe829ef9af2350c7c9d7b72ee1c05edefc7cd831ac13114e8ae7e1466c7ef2990a8e6387f9def145be2c9be8b0cbc8f7e6ebcb1b099a1785ea2001d70887faaaf86c937a292049134d02c7bc9aeaed7a3354673985a88a6e3aff66ccd";
        bytes memory padded =
            hex"00000000000000000000000000000000147a0a08008b085fd27341456bb43654fed88ebca4407d0b5ae66a67df8141edc58fced0468b5e34cec33dc2d58f023d000000000000000000000000000000000ace4f38b5e22e8cfd27951b8327d5260def6e6370e334c3f74fb9b5ccd209d0d80d1a34bb804f6ac40dd991141d87b5000000000000000000000000000000000cbc8f7e6ebcb1b099a1785ea2001d70887faaaf86c937a292049134d02c7bc9aeaed7a3354673985a88a6e3aff66ccd000000000000000000000000000000001770616fe829ef9af2350c7c9d7b72ee1c05edefc7cd831ac13114e8ae7e1466c7ef2990a8e6387f9def145be2c9be8b";
        assertEq(padded, BLST.padUncompressedBLSTSignature(unpadded));
        assertEq(unpadded, BLST.unPadUncompressedBLSTSignature(padded));
    }

    /*
     * @dev Test the lexicographic ordering of 96 byte uncompressed public keys
     */
    function testComparePublicKeys() public pure {
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(2, 1, 3), _createPublicKeyFromWords(1, 21, 4)
            ),
            1
        );
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(2, 1, 3), _createPublicKeyFromWords(2, 0, 5)
            ),
            1
        );
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(2, 1, 3), _createPublicKeyFromWords(2, 1, 1)
            ),
            1
        );
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(1, 21, 4), _createPublicKeyFromWords(2, 1, 3)
            ),
            -1
        );
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(2, 0, 5), _createPublicKeyFromWords(2, 1, 3)
            ),
            -1
        );
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(2, 1, 1), _createPublicKeyFromWords(2, 1, 3)
            ),
            -1
        );
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(0, 0, 1), _createPublicKeyFromWords(0, 0, 1)
            ),
            0
        );
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(0, 1, 1), _createPublicKeyFromWords(0, 1, 1)
            ),
            0
        );
        assertEq(
            BLST.comparePublicKeys(
                _createPublicKeyFromWords(1, 1, 1), _createPublicKeyFromWords(1, 1, 1)
            ),
            0
        );
    }

    /*
     * @dev Create 96 bytes from three 32 bytes words
     */
    function _createPublicKeyFromWords(
        uint256 x,
        uint256 y,
        uint256 z
    ) private pure returns (bytes memory) {
        bytes memory pk = new bytes(96);
        /* solhint-disable-next-line no-inline-assembly */
        assembly {
            mstore(add(pk, 0x20), x)
            mstore(add(pk, 0x40), y)
            mstore(add(pk, 0x60), z)
        }
        return pk;
    }
}
