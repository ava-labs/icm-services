pragma solidity ^0.8.30;

library Alpha {
    struct Thing {
        bool flag;
    }
}

library Beta {
    // #[unpack()]
    struct Holder {
        Alpha.Thing thing;
    }
}
