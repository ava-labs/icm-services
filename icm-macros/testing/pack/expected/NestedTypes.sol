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
    struct Struct {
        FreeStanding[][] free;
        bool[][][] flags;
    }

    function packStruct(Struct memory obj) public pure returns (bytes memory) {
        /* solhint-disable var-name-mixedcase */
        bytes memory free_bytes;
        bytes memory flags_bytes;
        free_bytes = abi.encodePacked(obj.free.length);
        for (uint256 i_0 = 0; i_0 < obj.free.length;) {
            bytes memory free_1_bytes;

            free_1_bytes = abi.encodePacked(obj.free[i_0].length);
            for (uint256 i_1 = 0; i_1 < obj.free[i_0].length;) {
                bytes memory free_2_bytes;
                free_2_bytes = packFreeStanding(obj.free[i_0][i_1]);
                free_1_bytes = abi.encodePacked(free_1_bytes, free_2_bytes);
                unchecked {
                    ++i_1;
                }
            }
            free_bytes = abi.encodePacked(free_bytes, free_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        flags_bytes = abi.encodePacked(obj.flags.length);
        for (uint256 i_0 = 0; i_0 < obj.flags.length;) {
            bytes memory flags_1_bytes;

            flags_1_bytes = abi.encodePacked(obj.flags[i_0].length);
            for (uint256 i_1 = 0; i_1 < obj.flags[i_0].length;) {
                bytes memory flags_2_bytes;

                flags_2_bytes = abi.encodePacked(obj.flags[i_0][i_1].length);
                for (uint256 i_2 = 0; i_2 < obj.flags[i_0][i_1].length;) {
                    bytes memory flags_3_bytes;
                    flags_3_bytes = abi.encodePacked(obj.flags[i_0][i_1][i_2]);
                    flags_2_bytes = abi.encodePacked(flags_2_bytes, flags_3_bytes);
                    unchecked {
                        ++i_2;
                    }
                }
                flags_1_bytes = abi.encodePacked(flags_1_bytes, flags_2_bytes);
                unchecked {
                    ++i_1;
                }
            }
            flags_bytes = abi.encodePacked(flags_bytes, flags_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        return abi.encodePacked(free_bytes, flags_bytes);
        /* solhint-enable var-name-mixedcase */
    }
}
