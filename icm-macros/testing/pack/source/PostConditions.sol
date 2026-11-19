pragma solidity ^0.8.30;

library PackMethods {

    // #[pack(assert = "|b| { b.length == 20 }")]
    struct Fixed {
        address addr;
    }

    // #[pack(
    //   assert = "|b| { b.length > 0 }",
    //   assert = "|b| { b.length % 2 == 0 }",
    // )]
    struct MultiAssert {
        bytes data;
    }

    // #[pack(name = "serializeNamed", assert = "|b| { b.length == 48 }")]
    struct Named {
        bytes data;
    }

    // #[pack(assert = "|b| { b.length == 1 }")]
    enum Status { Active, Inactive }
}