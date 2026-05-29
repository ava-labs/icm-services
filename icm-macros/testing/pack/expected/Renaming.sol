pragma solidity ^0.8.30;

struct FreeStanding {
    string text;
    bool flag;
}

library PackMethods {
    enum Enum {
        First,
        Second
    }

    // #[pack(name="serializeStruct", visibility="internal")]
    struct Struct {
        bytes inner;
        bool flag;
        // #[pack(ignore)]
        FreeStanding free;
        // #[pack(method="abi.encodePacked")]
        Enum choice;
    }

    function serializeStruct(Struct memory obj) internal pure returns (bytes memory) {
        /* solhint-disable */

        return abi.encodePacked(
            abi.encodePacked(obj.inner.length, obj.inner), abi.encodePacked(obj.flag), abi.encodePacked(obj.choice)
        );
        /* solhint-enable */
    }
}
