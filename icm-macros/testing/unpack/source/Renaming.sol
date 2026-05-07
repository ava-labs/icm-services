pragma solidity ^0.8.30;

struct FreeStanding {
    string text;
    bool flag;
}

library PackMethods {
    // #[unpack(solhint-disable)]
    enum Choice {
        First,
        Second
    }

    enum Enum {
        First,
        Second
    }

    // #[unpack(name="deserializeStruct", visibility="internal")]
    struct Struct {
        bytes inner;
        bool flag;
        // #[unpack(default)]
        FreeStanding free;
        Choice choice;
        // #[unpack(method="decodeEnum")]
        Enum other_choice;
    }

    function decodeEnum(bytes memory data) public pure returns (uint256, Enum) {
        return (1, Enum(uint8(data[0])));
    }

}
