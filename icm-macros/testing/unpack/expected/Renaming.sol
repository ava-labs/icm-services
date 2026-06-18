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

    function deserializeStruct(bytes memory data) internal pure returns (uint256, Struct memory) {
        /* solhint-disable no-inline-assembly */
        /* solhint-disable var-name-mixedcase */
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        Struct memory result;

        bytes memory inner;
        {
            uint256 data_length;
            assembly { data_length := mload(add(data, 32)) }
            inner = new bytes(data_length);
            assembly {
                mcopy(add(inner, 32), add(data, 64), data_length)
                let _data_orig_len := mload(data)
                data := add(data, add(32, data_length))
                mstore(data, sub(_data_orig_len, add(32, data_length)))
            }
        }
        result.inner = inner;
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
        Choice choice;
        {
            uint256 _len_before;
            assembly { _len_before := mload(data) }
            uint256 read;
            (read, choice) = unpackChoice(data);
            assembly {
                data := add(data, read)
                mstore(data, sub(_len_before, read))
            }
        }
        result.choice = choice;
        Enum other_choice;
        {
            uint256 _len_before;
            assembly { _len_before := mload(data) }
            uint256 read;
            (read, other_choice) = decodeEnum(data);
            assembly {
                data := add(data, read)
                mstore(data, sub(_len_before, read))
            }
        }
        result.other_choice = other_choice;
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
        /* solhint-enable no-inline-assembly */
        /* solhint-enable var-name-mixedcase */
    }

    function unpackChoice(bytes memory data) public pure returns (uint256, Choice) {
        return (1, Choice(uint8(data[0])));
    }
}
