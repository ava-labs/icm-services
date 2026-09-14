pragma solidity ^0.8.30;

// #[pack()]
struct FreeStanding {
    string text;
    bool flag;
}

library PackMethods {

    // #[pack()]
    enum Enum { First, Second }

    // #[pack()]
    struct Struct {
        FreeStanding free;
        string[] names;
    }
}