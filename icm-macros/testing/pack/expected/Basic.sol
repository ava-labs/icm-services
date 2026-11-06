pragma solidity ^0.8.30;

// #[pack()]
struct FreeStanding {
    string text;
    bool flag;
}

function packFreeStanding(FreeStanding memory obj) pure returns (bytes memory) {
    return abi.encodePacked(abi.encodePacked(bytes(obj.text).length, obj.text), abi.encodePacked(obj.flag));
}

library PackMethods {
    // #[pack()]
    enum Enum {
        First,
        Second
    }

    // #[pack()]
    struct Struct {
        FreeStanding free;
        string[] names;
    }

    function packStruct(Struct memory obj) public pure returns (bytes memory) {
        bytes memory names_bytes;
        uint256 names_0_count = obj.names.length;
        names_bytes = abi.encodePacked(names_0_count);
        for (uint256 i_0 = 0; i_0 < names_0_count;) {
            bytes memory names_1_bytes;
            names_1_bytes = abi.encodePacked(bytes(obj.names[i_0]).length, obj.names[i_0]);
            names_bytes = abi.encodePacked(names_bytes, names_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        return abi.encodePacked(packFreeStanding(obj.free), names_bytes);
    }

    function packEnum(Enum obj) public pure returns (bytes memory) {
        return abi.encodePacked(obj);
    }
}
