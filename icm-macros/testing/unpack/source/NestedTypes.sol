pragma solidity ^0.8.30;

// #[unpack()]
struct FreeStanding {
    string text;
    bool flag;
}

// #[unpack(calldata)]
struct FreeStandingCalldata {
    string text;
    bool flag;
}

library UnpackMethods {

    // #[unpack()]
    struct Struct {
        FreeStanding[][] free;
        bool[][][] flags;
    }

    // #[unpack(calldata)]
    struct StructCalldata {
        FreeStandingCalldata[][] free;
        bool[][][] flags;
    }
}
