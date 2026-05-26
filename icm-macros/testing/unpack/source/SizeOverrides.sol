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
    }

    // #[unpack(calldata)]
    struct DynamicStructCalldata {
        // #[unpack(length=uint8)]
        bytes Bytes;
        // #[unpack(length=uint16)]
        string String;
        // #[unpack(length=uint32)]
        address[] Addresses;
    }
}