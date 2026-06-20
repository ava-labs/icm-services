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
}