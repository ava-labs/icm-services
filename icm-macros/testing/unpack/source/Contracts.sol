pragma solidity ^0.8.30;

// #[unpack(contract="FirstContract")]
struct FreeStanding {
    string text;
    bool flag;
}

library FirstContract {

    // #[unpack()]
    struct Struct {
        FreeStanding free;
        string[] names;
    }

    // #[unpack(contract="SecondContract", name="serializeOther")]
    struct Other {
        bool flag;
    }
}

library SecondContract {

    // #[unpack()]
    struct Struct {
        // #[unpack(method="FirstContract.packFreeStanding")]
        FreeStanding free;
        string[] names;
    }
}