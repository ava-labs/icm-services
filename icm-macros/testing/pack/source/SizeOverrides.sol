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
    }
}