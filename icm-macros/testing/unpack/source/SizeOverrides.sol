pragma solidity ^0.8.30;

library Lengths {
    uint256 internal constant HASH_LENGTH = 32;
    uint256 internal constant PAIR_LENGTH = 2;
}

library PackMethods {
    // #[unpack()]
    struct DynamicStruct {
        // #[unpack(length=uint8)]
        bytes Bytes;
        // #[unpack(length=uint16)]
        string String;
        // #[unpack(length=uint32)]
        address[] Addresses;
        // #[unpack(length=Lengths.HASH_LENGTH)]
        bytes Hash;
        // #[unpack(length=Lengths.PAIR_LENGTH)]
        address[] Pair;
        // #[unpack(length=Lengths.HASH_LENGTH, method="parseHashMemory")]
        bytes ParsedHash;
    }

    // #[unpack(calldata)]
    struct DynamicStructCalldata {
        // #[unpack(length=uint8)]
        bytes Bytes;
        // #[unpack(length=uint16)]
        string String;
        // #[unpack(length=uint32)]
        address[] Addresses;
        // #[unpack(length=32)]
        bytes Hash;
        // #[unpack(length=2)]
        address[] Pair;
        // #[unpack(length=32, method="parseHash")]
        bytes ParsedHash;
    }

    function parseHash(bytes calldata hashBytes) internal pure returns (bytes memory) {
        if (hashBytes.length == 32) {
            return hashBytes;
        } else {
            return new bytes(32);
        }
    }

    function parseHashMemory(bytes memory hashBytes) internal pure returns (bytes memory) {
        if (hashBytes.length == 32) {
            return hashBytes;
        } else {
            return new bytes(32);
        }
    }
}
