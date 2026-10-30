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

    function packDynamicStruct(DynamicStruct memory obj) public pure returns (bytes memory) {
        bytes memory Addresses_bytes;
        require(obj.Hash.length == Lengths.HASH_LENGTH, "packDynamicStruct: invalid Hash length");
        require(bytes(obj.Tag).length == 3, "packDynamicStruct: invalid Tag length");
        require(obj.Pair.length == Lengths.PAIR_LENGTH, "packDynamicStruct: invalid Pair length");
        bytes memory Pair_bytes;
        bytes memory ParsedHash_bytes = abi.encodePacked(obj.ParsedHash);
        require(ParsedHash_bytes.length == 32, "packDynamicStruct: invalid ParsedHash length");
        uint256 Addresses_0_count = obj.Addresses.length;
        Addresses_bytes = abi.encodePacked(uint32(Addresses_0_count));
        for (uint256 i_0 = 0; i_0 < Addresses_0_count;) {
            bytes memory Addresses_1_bytes;
            Addresses_1_bytes = abi.encodePacked(obj.Addresses[i_0]);
            Addresses_bytes = abi.encodePacked(Addresses_bytes, Addresses_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        uint256 Pair_0_count = obj.Pair.length;
        for (uint256 i_0 = 0; i_0 < Pair_0_count;) {
            bytes memory Pair_1_bytes;
            Pair_1_bytes = abi.encodePacked(obj.Pair[i_0]);
            Pair_bytes = abi.encodePacked(Pair_bytes, Pair_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        return abi.encodePacked(
            abi.encodePacked(uint8(obj.Bytes.length), obj.Bytes),
            abi.encodePacked(uint16(bytes(obj.String).length), obj.String),
            Addresses_bytes,
            abi.encodePacked(obj.Hash),
            abi.encodePacked(obj.Tag),
            Pair_bytes,
            ParsedHash_bytes
        );
    }
}
