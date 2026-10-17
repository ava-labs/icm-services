pragma solidity ^0.8.30;

library PackMethods {
    // #[unpack()]
    struct DynamicStruct {
        // #[unpack(length=uint8)]
        bytes Bytes;
        // #[unpack(length=uint16)]
        string String;
        // #[unpack(length=uint32)]
        address[] Addresses;
        // #[unpack(length=32)]
        bytes Hash;
    }

    // #[unpack(calldata)]
    struct DynamicStructCalldata {
        // #[unpack(length=uint8)]
        bytes Bytes;
        // #[unpack(length=uint16)]
        string String;
        // #[unpack(length=uint32)]
        address[] Addresses;
        // #[unpack(length=32, method="parseHash")]
        bytes Hash;
    }

    function parseHash(bytes calldata hashBytes) internal pure returns (bytes memory) {
        if (hashBytes.length == 32) {
            return hashBytes;
        } else {
            return new bytes(32);
        }
    }
}