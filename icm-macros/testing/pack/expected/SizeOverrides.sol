pragma solidity ^0.8.30;

library PackMethods {
    // #[pack()]
    struct DynamicStruct {
        // #[pack(length=uint8)]
        bytes Bytes;
        // #[pack(length=uint16)]
        string String;
        // #[pack(length=uint32)]
        address[] Addresses;
        // #[pack(length=drop)]
        bytes Hash;
        // #[pack(length=drop)]
        address[] Pair;
    }

    function packDynamicStruct(DynamicStruct memory obj) public pure returns (bytes memory) {
        bytes memory Addresses_bytes;
        bytes memory Pair_bytes;
        Addresses_bytes = abi.encodePacked(uint32(obj.Addresses.length));
        for (uint256 i_0 = 0; i_0 < obj.Addresses.length;) {
            bytes memory Addresses_1_bytes;
            Addresses_1_bytes = abi.encodePacked(obj.Addresses[i_0]);
            Addresses_bytes = abi.encodePacked(Addresses_bytes, Addresses_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        Pair_bytes = abi.encodePacked();
        for (uint256 i_0 = 0; i_0 < obj.Pair.length;) {
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
            Pair_bytes
        );
    }
}
