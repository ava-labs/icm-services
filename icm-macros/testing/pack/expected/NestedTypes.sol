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
        bytes memory free_bytes;
        bytes memory flags_bytes;
        uint256 free_0_count = obj.free.length;
        free_bytes = abi.encodePacked(free_0_count);
        for (uint256 i_0 = 0; i_0 < free_0_count;) {
            bytes memory free_1_bytes;

            uint256 free_1_count = obj.free[i_0].length;
            free_1_bytes = abi.encodePacked(free_1_count);
            for (uint256 i_1 = 0; i_1 < free_1_count;) {
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
        uint256 flags_0_count = obj.flags.length;
        flags_bytes = abi.encodePacked(flags_0_count);
        for (uint256 i_0 = 0; i_0 < flags_0_count;) {
            bytes memory flags_1_bytes;

            uint256 flags_1_count = obj.flags[i_0].length;
            flags_1_bytes = abi.encodePacked(flags_1_count);
            for (uint256 i_1 = 0; i_1 < flags_1_count;) {
                bytes memory flags_2_bytes;

                uint256 flags_2_count = obj.flags[i_0][i_1].length;
                flags_2_bytes = abi.encodePacked(flags_2_count);
                for (uint256 i_2 = 0; i_2 < flags_2_count;) {
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
    }
}
