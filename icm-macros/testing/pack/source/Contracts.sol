pragma solidity ^0.8.30;

// #[pack(contract="FirstContract")]
struct FreeStanding {
    string text;
    bool flag;
}

library FirstContract {

    // #[pack()]
    struct Struct {
        FreeStanding free;
        string[] names;
    }

    // #[pack(contract="SecondContract", name="serializeOther")]
    struct Other {
        bool flag;
    }
}

library SecondContract {

    // #[pack()]
    struct Struct {
        // #[pack(method="FirstContract.packFreeStanding")]
        FreeStanding free;
        string[] names;
    }
}