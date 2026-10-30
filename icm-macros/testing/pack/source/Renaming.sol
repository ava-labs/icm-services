pragma solidity ^0.8.30;

struct FreeStanding {
    string text;
    bool flag;
}

library PackMethods {

    enum Enum {First, Second}

    // #[pack(name="serializeStruct", visibility="internal")]
    struct Struct {
        bytes inner;
        bool flag;
        // #[pack(ignore)]
        FreeStanding free;
        // #[pack(method="abi.encodePacked")]
        Enum choice;
    }
}