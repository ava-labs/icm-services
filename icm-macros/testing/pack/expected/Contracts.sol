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

    function packFreeStanding(FreeStanding memory obj) public pure returns (bytes memory) {
        /* solhint-disable */

        return abi.encodePacked(abi.encodePacked(bytes(obj.text).length, obj.text), abi.encodePacked(obj.flag));
        /* solhint-enable */
    }

    function packStruct(Struct memory obj) public pure returns (bytes memory) {
        /* solhint-disable */
        bytes memory names_bytes;
        names_bytes = abi.encodePacked(obj.names.length);
        for (uint256 i_0 = 0; i_0 < obj.names.length;) {
            bytes memory names_1_bytes;
            names_1_bytes = abi.encodePacked(bytes(obj.names[i_0]).length, obj.names[i_0]);
            names_bytes = abi.encodePacked(names_bytes, names_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        return abi.encodePacked(packFreeStanding(obj.free), names_bytes);
        /* solhint-enable */
    }
}

library SecondContract {
    // #[pack()]
    struct Struct {
        // #[pack(method="FirstContract.packFreeStanding")]
        FreeStanding free;
        string[] names;
    }

    function serializeOther(FirstContract.Other memory obj) public pure returns (bytes memory) {
        /* solhint-disable */

        return abi.encodePacked(abi.encodePacked(obj.flag));
        /* solhint-enable */
    }

    function packStruct(Struct memory obj) public pure returns (bytes memory) {
        /* solhint-disable */
        bytes memory names_bytes;
        names_bytes = abi.encodePacked(obj.names.length);
        for (uint256 i_0 = 0; i_0 < obj.names.length;) {
            bytes memory names_1_bytes;
            names_1_bytes = abi.encodePacked(bytes(obj.names[i_0]).length, obj.names[i_0]);
            names_bytes = abi.encodePacked(names_bytes, names_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        return abi.encodePacked(FirstContract.packFreeStanding(obj.free), names_bytes);
        /* solhint-enable */
    }
}
