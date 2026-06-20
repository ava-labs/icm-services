pragma solidity ^0.8.30;

import {Test} from  "forge-std/Test.sol";
import {Primitives, Choice, RoundTrip} from "../RoundTrip.sol";

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
}