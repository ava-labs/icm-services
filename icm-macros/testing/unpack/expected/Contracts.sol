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

    function unpackFreeStanding(bytes memory data) public pure returns (uint256, FreeStanding memory) {
        /* solhint-disable */
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
        /* solhint-enable */
    }

    function unpackStruct(bytes memory data) public pure returns (uint256, Struct memory) {
        /* solhint-disable */
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        Struct memory result;

        FreeStanding memory free;
        {
            uint256 _len_before;
            assembly { _len_before := mload(data) }
            uint256 read;
            (read, free) = unpackFreeStanding(data);
            assembly {
                data := add(data, read)
                mstore(data, sub(_len_before, read))
            }
        }
        result.free = free;
        string[] memory names;
        {
            uint256 length;
            assembly { length := mload(add(data, 32)) }
            assembly {
                let _data_new_len := sub(mload(data), 32)
                data := add(data, 32)
                mstore(data, _data_new_len)
            }
            names = new string[](length);
            for (uint256 i = 0; i < length;) {
                string memory names_1;
                {
                    uint256 data_length;
                    assembly { data_length := mload(add(data, 32)) }
                    names_1 = string(new bytes(data_length));
                    assembly {
                        mcopy(add(names_1, 32), add(data, 64), data_length)
                        let _data_orig_len := mload(data)
                        data := add(data, add(32, data_length))
                        mstore(data, sub(_data_orig_len, add(32, data_length)))
                    }
                }
                names[i] = names_1;
                unchecked {
                    ++i;
                }
            }
        }
        result.names = names;
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
        /* solhint-enable */
    }
}

library SecondContract {
    // #[unpack()]
    struct Struct {
        // #[unpack(method="FirstContract.packFreeStanding")]
        FreeStanding free;
        string[] names;
    }

    function serializeOther(bytes memory data) public pure returns (uint256, FirstContract.Other memory) {
        /* solhint-disable */
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        FirstContract.Other memory result;

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
        /* solhint-enable */
    }

    function unpackStruct(bytes memory data) public pure returns (uint256, Struct memory) {
        /* solhint-disable */
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        Struct memory result;

        FreeStanding memory free;
        {
            uint256 _len_before;
            assembly { _len_before := mload(data) }
            uint256 read;
            (read, free) = FirstContract.packFreeStanding(data);
            assembly {
                data := add(data, read)
                mstore(data, sub(_len_before, read))
            }
        }
        result.free = free;
        string[] memory names;
        {
            uint256 length;
            assembly { length := mload(add(data, 32)) }
            assembly {
                let _data_new_len := sub(mload(data), 32)
                data := add(data, 32)
                mstore(data, _data_new_len)
            }
            names = new string[](length);
            for (uint256 i = 0; i < length;) {
                string memory names_1;
                {
                    uint256 data_length;
                    assembly { data_length := mload(add(data, 32)) }
                    names_1 = string(new bytes(data_length));
                    assembly {
                        mcopy(add(names_1, 32), add(data, 64), data_length)
                        let _data_orig_len := mload(data)
                        data := add(data, add(32, data_length))
                        mstore(data, sub(_data_orig_len, add(32, data_length)))
                    }
                }
                names[i] = names_1;
                unchecked {
                    ++i;
                }
            }
        }
        result.names = names;
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
        /* solhint-enable */
    }
}
