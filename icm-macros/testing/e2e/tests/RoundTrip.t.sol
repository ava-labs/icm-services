pragma solidity ^0.8.30;

import {Test} from  "forge-std/Test.sol";
import {Primitives, Sizes, packSizes, unpackSizes, Choice, RoundTrip} from "../RoundTrip.sol";

contract TestRoundTrip is Test {

    function testRoundTrip(
        bool a,
        address b,
        uint32 c,
        int64 d,
        bytes16 e,
        bytes memory f,
        string memory g,
        uint8 numPrimitives
    ) public pure {
        vm.assume(numPrimitives <= 10);

        Primitives memory primitive = Primitives({
            a: a,
            b: b,
            c: c,
            d: d,
            e: e,
            f: f,
            g: g
        });

        Primitives[] memory primitives = new Primitives[](numPrimitives);
        for (uint256 i = 0; i < numPrimitives; i++) {
            primitives[i] = primitive;
        }

        RoundTrip.OtherPrimitives memory original = RoundTrip.OtherPrimitives({
            a: true,
            b: b,
            c: c,
            d: d,
            e: e,
            f: f,
            g: g,
            primitives: primitives,
            choice: Choice.Second
        });

        bytes memory data = RoundTrip.packOtherPrimitives(original);
        (, RoundTrip.OtherPrimitives memory deserialized) = RoundTrip.deserialize(data);

        assertEq(deserialized.a, false);
        assertEq(deserialized.b, original.b);
        assertEq(deserialized.c, original.c);
        assertEq(deserialized.d, original.d);
        assertEq(bytes32(deserialized.e), bytes32(original.e));
        assertEq(deserialized.f, original.f);
        assertEq(deserialized.g, original.g);
        assertEq(uint8(deserialized.choice), uint8(original.choice));

        assertEq(deserialized.primitives.length, original.primitives.length);
        for (uint256 i = 0; i < original.primitives.length; i++) {
            assertEq(deserialized.primitives[i].a, original.primitives[i].a);
            assertEq(deserialized.primitives[i].b, original.primitives[i].b);
            assertEq(deserialized.primitives[i].c, original.primitives[i].c);
            assertEq(deserialized.primitives[i].d, original.primitives[i].d);
            assertEq(bytes32(deserialized.primitives[i].e), bytes32(original.primitives[i].e));
            assertEq(deserialized.primitives[i].f, original.primitives[i].f);
            assertEq(deserialized.primitives[i].g, original.primitives[i].g);
        }
    }

    function testRoundTripSizes(
        bytes memory b,
        string memory s,
        address addr,
        uint8 numAddresses,
        bytes32 hash
    ) public pure {
        vm.assume(b.length <= 255);
        vm.assume(bytes(s).length <= 65535);
        vm.assume(numAddresses <= 10);

        address[] memory addrs = new address[](numAddresses);
        for (uint256 i = 0; i < numAddresses; i++) {
            addrs[i] = addr;
        }

        address[] memory pair = new address[](2);
        pair[0] = addr;
        pair[1] = address(uint160(uint256(hash)));

        Sizes memory original = Sizes({
            Bytes: b,
            String: s,
            Addresses: addrs,
            Pair: pair,
            ReverseHash: abi.encodePacked(hash)
        });

        bytes memory data = packSizes(original);
        (, Sizes memory deserialized) = unpackSizes(data);

        assertEq(deserialized.Bytes, original.Bytes);
        assertEq(deserialized.String, original.String);
        assertEq(deserialized.Addresses.length, original.Addresses.length);
        for (uint256 i = 0; i < original.Addresses.length; i++) {
            assertEq(deserialized.Addresses[i], original.Addresses[i]);
        }
        assertEq(deserialized.Pair.length, 2);
        assertEq(deserialized.Pair[0], original.Pair[0]);
        assertEq(deserialized.Pair[1], original.Pair[1]);
        assertEq(deserialized.ReverseHash, original.ReverseHash);
    }

    function testRoundTripSizesCalldata(
        bytes memory b,
        string memory s,
        address addr,
        uint8 numAddresses,
        bytes32 hash
    ) public pure {
        vm.assume(b.length <= 255);
        vm.assume(bytes(s).length <= 65535);
        vm.assume(numAddresses <= 10);

        address[] memory addrs = new address[](numAddresses);
        for (uint256 i = 0; i < numAddresses; i++) {
            addrs[i] = addr;
        }

        address[] memory pair = new address[](2);
        pair[0] = addr;
        pair[1] = address(uint160(uint256(hash)));

        RoundTrip.Sizes memory original = RoundTrip.Sizes({
            Bytes: b,
            String: s,
            Addresses: addrs,
            Pair: pair,
            ReverseHash: abi.encodePacked(hash)
        });

        bytes memory data = RoundTrip.packSizes(original);
        (, RoundTrip.Sizes memory deserialized) = RoundTrip.unpackSizes(data);

        assertEq(deserialized.Bytes, original.Bytes);
        assertEq(deserialized.String, original.String);
        assertEq(deserialized.Addresses.length, original.Addresses.length);
        for (uint256 i = 0; i < original.Addresses.length; i++) {
            assertEq(deserialized.Addresses[i], original.Addresses[i]);
        }
        assertEq(deserialized.Pair.length, 2);
        assertEq(deserialized.Pair[0], original.Pair[0]);
        assertEq(deserialized.Pair[1], original.Pair[1]);
        assertEq(deserialized.ReverseHash, original.ReverseHash);
    }

    // External helpers so vm.expectRevert() can intercept the generated functions.
    function callPackSizes(Sizes memory sizes) external pure returns (bytes memory) {
        return packSizes(sizes);
    }

    function callPackSizesLibrary(RoundTrip.Sizes memory sizes) external pure returns (bytes memory) {
        return RoundTrip.packSizes(sizes);
    }

    function callUnpackSizes(bytes memory data) external pure returns (uint256, Sizes memory) {
        return unpackSizes(data);
    }

    function callUnpackSizesCalldata(bytes calldata data) external pure returns (uint256, RoundTrip.Sizes memory) {
        return RoundTrip.unpackSizes(data);
    }

    function _validSizes() internal pure returns (Sizes memory) {
        address[] memory addrs = new address[](1);
        addrs[0] = address(0x1234);
        address[] memory pair = new address[](2);
        pair[0] = address(0x1);
        pair[1] = address(0x2);
        return Sizes({
            Bytes: hex"0102",
            String: "abc",
            Addresses: addrs,
            Pair: pair,
            ReverseHash: abi.encodePacked(keccak256("hash"))
        });
    }

    function _validLibrarySizes() internal pure returns (RoundTrip.Sizes memory) {
        Sizes memory sizes = _validSizes();
        return RoundTrip.Sizes({
            Bytes: sizes.Bytes,
            String: sizes.String,
            Addresses: sizes.Addresses,
            Pair: sizes.Pair,
            ReverseHash: sizes.ReverseHash
        });
    }

    // A fixed-count array that does not have the agreed count must not be packed, since the
    // stream carries no count for the receiver to check.
    function testPackSizesRejectsWrongPairCount() public {
        Sizes memory original = _validSizes();
        original.Pair = new address[](3);
        vm.expectRevert(bytes("packSizes: invalid Pair length"));
        this.callPackSizes(original);
    }

    // With method + constant the check applies to the method's output.
    function testPackSizesRejectsShortHash() public {
        RoundTrip.Sizes memory original = _validLibrarySizes();
        original.ReverseHash = new bytes(20);
        vm.expectRevert(bytes("packSizes: invalid ReverseHash length"));
        this.callPackSizesLibrary(original);
    }

    // A memory buffer that ends inside a fixed-length field reverts at that field.
    function testUnpackSizesRejectsShortData() public {
        bytes memory data = packSizes(_validSizes());
        // Drop the final byte, which belongs to the 32-byte ReverseHash.
        assembly {
            mstore(data, sub(mload(data), 1))
        }
        vm.expectRevert(bytes("unpackSizes: data too short for ReverseHash"));
        this.callUnpackSizes(data);
    }

    function testUnpackSizesCalldataRejectsShortData() public {
        bytes memory data = RoundTrip.packSizes(_validLibrarySizes());
        assembly {
            mstore(data, sub(mload(data), 1))
        }
        vm.expectRevert();
        this.callUnpackSizesCalldata(data);
    }
}
