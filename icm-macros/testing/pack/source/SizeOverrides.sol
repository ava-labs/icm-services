pragma solidity ^0.8.30;

library Lengths {
    uint256 internal constant HASH_LENGTH = 32;
    uint256 internal constant PAIR_LENGTH = 2;
}

library PackMethods {
    // #[pack()]
    struct DynamicStruct {
        // #[pack(length=uint8)]
        bytes Bytes;
        // #[pack(length=uint16)]
        string String;
        // #[pack(length=uint32)]
        address[] Addresses;
        // #[pack(length=Lengths.HASH_LENGTH)]
        bytes Hash;
        // #[pack(length=3)]
        string Tag;
        // #[pack(length=Lengths.PAIR_LENGTH)]
        address[] Pair;
        // #[pack(length=32, method="abi.encodePacked")]
        bytes ParsedHash;
    }
}
