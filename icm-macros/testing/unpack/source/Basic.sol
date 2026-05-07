pragma solidity ^0.8.30;

// #[unpack()]
struct Primitives {
    bool a;
    address b;
    uint256 c;
    int256 d;
    bytes32 e;
    bytes f;
    string g;
}

library TestUnpack {
    // #[unpack()]
    struct Inner {
        bool boolean;
        Primitives primitives;
        uint32 id;
    }

    // #[unpack()]
    struct OtherInner {
        bool boolean;
        // #[unpack(method="OtherTestUnpack.unpackOtherPrimitives")]
        OtherTestUnpack.OtherPrimitives primitives;
        uint32 id;
    }
}

library OtherTestUnpack {
    // #[unpack(calldata)]
    struct OtherPrimitives {
        bool a;
        address b;
        uint256 c;
        int256 d;
        bytes32 e;
        bytes f;
        string g;
        Primitives primitives;
    }
}