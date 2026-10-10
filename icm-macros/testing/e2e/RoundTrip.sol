pragma solidity ^0.8.30;

// #[pack()]
// #[unpack()]
struct Primitives {
    bool a;
    address b;
    uint32 c;
    int64 d;
    bytes16 e;
    bytes f;
    string g;
}

// #[pack()]
// #[unpack()]
struct Sizes {
    // #[pack(length=uint8)]
    // #[unpack(length=uint8)]
    bytes Bytes;
    // #[pack(length=uint16)]
    // #[unpack(length=uint16)]
    string String;
    // #[pack(length=uint32)]
    // #[unpack(length=uint32)]
    address[] Addresses;
}

// #[pack(contract="RoundTrip")]
// #[unpack(contract="RoundTrip")]
enum Choice {
    First,
    Second
}

library RoundTrip {
    // #[pack()]
    // #[unpack(calldata, name="deserialize")]
    struct OtherPrimitives {
        // #[pack(ignore)]
        // #[unpack(default)]
        bool a;
        address b;
        uint32 c;
        int64 d;
        bytes16 e;
        bytes f;
        string g;
        Primitives[] primitives;
        Choice choice;
    }

    // #[pack()]
    // #[unpack(calldata)]
    struct Sizes {
        // #[pack(length=uint8)]
        // #[unpack(length=uint8)]
        bytes Bytes;
        // #[pack(length=uint16)]
        // #[unpack(length=uint16)]
        string String;
        // #[pack(length=uint32)]
        // #[unpack(length=uint32)]
        address[] Addresses;
    }


}