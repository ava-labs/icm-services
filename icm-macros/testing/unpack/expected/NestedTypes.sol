pragma solidity ^0.8.30;

// #[unpack()]
struct FreeStanding {
    string text;
    bool flag;
}

function unpackFreeStanding(bytes memory data) pure returns (uint256, FreeStanding memory) {
    /* solhint-disable no-inline-assembly */
    /* solhint-disable var-name-mixedcase */
    uint256 _initial_length;
    assembly { _initial_length := mload(data) }
    FreeStanding memory result;

    string memory text;
    {
        uint256 data_length;
        assembly { data_length := mload(add(data, 32)) }
        text = string(new bytes(data_length));
        assembly {
            mcopy(add(text, 32), add(data, 64), data_length)
            let _data_orig_len := mload(data)
            data := add(data, add(32, data_length))
            mstore(data, sub(_data_orig_len, add(32, data_length)))
        }
    }
    result.text = text;
    bool flag;
    assembly {
        flag := shr(248, mload(add(data, 32)))
    }
    assembly {
        let _data_new_len := sub(mload(data), 1)
        data := add(data, 1)
        mstore(data, _data_new_len)
    }
    result.flag = flag;
    uint256 _final_length;
    assembly { _final_length := mload(data) }
    return (_initial_length - _final_length, result);
    /* solhint-enable no-inline-assembly */
    /* solhint-enable var-name-mixedcase */
}

// #[unpack(calldata)]
struct FreeStandingCalldata {
    string text;
    bool flag;
}

function unpackFreeStandingCalldata(bytes calldata data) pure returns (uint256, FreeStandingCalldata memory) {
    /* solhint-disable var-name-mixedcase */
    uint256 _initial_length = data.length;
    FreeStandingCalldata memory result;

    string memory text;
    {
        uint256 length = uint256(bytes32(data[0:32]));
        text = string(data[32:length + 32]);
        data = data[32 + length:];
    }
    result.text = text;
    bool flag = bytes1(data[0:1]) != 0x00;
    data = data[1:];
    result.flag = flag;

    return (_initial_length - data.length, result);
    /* solhint-enable var-name-mixedcase */
}

library UnpackMethods {
    // #[unpack()]
    struct Struct {
        FreeStanding[][] free;
        bool[][][] flags;
    }

    // #[unpack(calldata)]
    struct StructCalldata {
        FreeStandingCalldata[][] free;
        bool[][][] flags;
    }

    function unpackStruct(bytes memory data) public pure returns (uint256, Struct memory) {
        /* solhint-disable no-inline-assembly */
        /* solhint-disable var-name-mixedcase */
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        Struct memory result;

        FreeStanding[][] memory free;
        {
            uint256 length;
            assembly { length := mload(add(data, 32)) }
            assembly {
                let _data_new_len := sub(mload(data), 32)
                data := add(data, 32)
                mstore(data, _data_new_len)
            }
            free = new FreeStanding[][](length);
            for (uint256 i = 0; i < length;) {
                FreeStanding[] memory free_1;
                {
                    uint256 length;
                    assembly { length := mload(add(data, 32)) }
                    assembly {
                        let _data_new_len := sub(mload(data), 32)
                        data := add(data, 32)
                        mstore(data, _data_new_len)
                    }
                    free_1 = new FreeStanding[](length);
                    for (uint256 i = 0; i < length;) {
                        FreeStanding memory free_1_2;
                        {
                            uint256 _len_before;
                            assembly { _len_before := mload(data) }
                            uint256 read;
                            (read, free_1_2) = unpackFreeStanding(data);
                            assembly {
                                data := add(data, read)
                                mstore(data, sub(_len_before, read))
                            }
                        }
                        free_1[i] = free_1_2;
                        unchecked {
                            ++i;
                        }
                    }
                }
                free[i] = free_1;
                unchecked {
                    ++i;
                }
            }
        }
        result.free = free;
        bool[][][] memory flags;
        {
            uint256 length;
            assembly { length := mload(add(data, 32)) }
            assembly {
                let _data_new_len := sub(mload(data), 32)
                data := add(data, 32)
                mstore(data, _data_new_len)
            }
            flags = new bool[][][](length);
            for (uint256 i = 0; i < length;) {
                bool[][] memory flags_1;
                {
                    uint256 length;
                    assembly { length := mload(add(data, 32)) }
                    assembly {
                        let _data_new_len := sub(mload(data), 32)
                        data := add(data, 32)
                        mstore(data, _data_new_len)
                    }
                    flags_1 = new bool[][](length);
                    for (uint256 i = 0; i < length;) {
                        bool[] memory flags_1_2;
                        {
                            uint256 length;
                            assembly { length := mload(add(data, 32)) }
                            assembly {
                                let _data_new_len := sub(mload(data), 32)
                                data := add(data, 32)
                                mstore(data, _data_new_len)
                            }
                            flags_1_2 = new bool[](length);
                            for (uint256 i = 0; i < length;) {
                                bool flags_1_2_3;
                                assembly {
                                    flags_1_2_3 := shr(248, mload(add(data, 32)))
                                }
                                assembly {
                                    let _data_new_len := sub(mload(data), 1)
                                    data := add(data, 1)
                                    mstore(data, _data_new_len)
                                }
                                flags_1_2[i] = flags_1_2_3;
                                unchecked {
                                    ++i;
                                }
                            }
                        }
                        flags_1[i] = flags_1_2;
                        unchecked {
                            ++i;
                        }
                    }
                }
                flags[i] = flags_1;
                unchecked {
                    ++i;
                }
            }
        }
        result.flags = flags;
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
        /* solhint-enable no-inline-assembly */
        /* solhint-enable var-name-mixedcase */
    }

    function unpackStructCalldata(bytes calldata data) public pure returns (uint256, StructCalldata memory) {
        /* solhint-disable var-name-mixedcase */
        uint256 _initial_length = data.length;
        StructCalldata memory result;

        FreeStandingCalldata[][] memory free;
        {
            uint256 length = uint256(bytes32(data[0:32]));
            data = data[32:];
            free = new FreeStandingCalldata[][](length);
            for (uint256 i = 0; i < length;) {
                FreeStandingCalldata[] memory free_1;
                {
                    uint256 length = uint256(bytes32(data[0:32]));
                    data = data[32:];
                    free_1 = new FreeStandingCalldata[](length);
                    for (uint256 i = 0; i < length;) {
                        FreeStandingCalldata memory free_1_2;
                        {
                            uint256 read;
                            (read, free_1_2) = unpackFreeStandingCalldata(data);
                            data = data[read:];
                        }
                        free_1[i] = free_1_2;
                        unchecked {
                            ++i;
                        }
                    }
                }
                free[i] = free_1;
                unchecked {
                    ++i;
                }
            }
        }
        result.free = free;
        bool[][][] memory flags;
        {
            uint256 length = uint256(bytes32(data[0:32]));
            data = data[32:];
            flags = new bool[][][](length);
            for (uint256 i = 0; i < length;) {
                bool[][] memory flags_1;
                {
                    uint256 length = uint256(bytes32(data[0:32]));
                    data = data[32:];
                    flags_1 = new bool[][](length);
                    for (uint256 i = 0; i < length;) {
                        bool[] memory flags_1_2;
                        {
                            uint256 length = uint256(bytes32(data[0:32]));
                            data = data[32:];
                            flags_1_2 = new bool[](length);
                            for (uint256 i = 0; i < length;) {
                                bool flags_1_2_3 = bytes1(data[0:1]) != 0x00;
                                data = data[1:];
                                flags_1_2[i] = flags_1_2_3;
                                unchecked {
                                    ++i;
                                }
                            }
                        }
                        flags_1[i] = flags_1_2;
                        unchecked {
                            ++i;
                        }
                    }
                }
                flags[i] = flags_1;
                unchecked {
                    ++i;
                }
            }
        }
        result.flags = flags;

        return (_initial_length - data.length, result);
        /* solhint-enable var-name-mixedcase */
    }
}
