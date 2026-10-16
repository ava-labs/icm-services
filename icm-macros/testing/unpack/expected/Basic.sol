pragma solidity ^0.8.30;

// #[unpack()]
struct Primitives {
    bool a;
    address b;
    uint256 c;
    int256 d;
    bytes32 e;
    bytes f;
    string g;
}

function unpackPrimitives(bytes memory data) pure returns (uint256, Primitives memory) {
    uint256 _initial_length;
    assembly { _initial_length := mload(data) }
    Primitives memory result;

    bool a;
    assembly {
        a := shr(248, mload(add(data, 32)))
    }
    assembly {
        let _data_new_len := sub(mload(data), 1)
        data := add(data, 1)
        mstore(data, _data_new_len)
    }
    result.a = a;
    address b;
    assembly {
        b := shr(96, mload(add(data, 32)))
    }
    assembly {
        let _data_new_len := sub(mload(data), 20)
        data := add(data, 20)
        mstore(data, _data_new_len)
    }
    result.b = b;
    uint256 c;
    assembly {
        c := shr(0, mload(add(data, 32)))
    }
    assembly {
        let _data_new_len := sub(mload(data), 32)
        data := add(data, 32)
        mstore(data, _data_new_len)
    }
    result.c = c;
    int256 d;
    assembly {
        d := sar(0, mload(add(data, 32)))
    }
    assembly {
        let _data_new_len := sub(mload(data), 32)
        data := add(data, 32)
        mstore(data, _data_new_len)
    }
    result.d = d;
    bytes32 e;
    assembly {
        e := and(mload(add(data, 32)), shl(0, not(0)))
    }
    assembly {
        let _data_new_len := sub(mload(data), 32)
        data := add(data, 32)
        mstore(data, _data_new_len)
    }
    result.e = e;
    bytes memory f;
    {
        uint256 data_length;
        assembly { data_length := mload(add(data, 32)) }
        f = new bytes(data_length);
        assembly {
            mcopy(add(f, 32), add(data, 64), data_length)
            let _data_orig_len := mload(data)
            data := add(data, add(32, data_length))
            mstore(data, sub(_data_orig_len, add(32, data_length)))
        }
    }
    result.f = f;
    string memory g;
    {
        uint256 data_length;
        assembly { data_length := mload(add(data, 32)) }
        g = string(new bytes(data_length));
        assembly {
            mcopy(add(g, 32), add(data, 64), data_length)
            let _data_orig_len := mload(data)
            data := add(data, add(32, data_length))
            mstore(data, sub(_data_orig_len, add(32, data_length)))
        }
    }
    result.g = g;
    uint256 _final_length;
    assembly { _final_length := mload(data) }
    return (_initial_length - _final_length, result);
}

library TestUnpack {
    // #[unpack()]
    struct Inner {
        bool boolean;
        Primitives primitives;
        uint32 id;
    }

    // #[unpack()]
    struct OtherInner {
        bool boolean;
        // #[unpack(method="OtherTestUnpack.unpackOtherPrimitives")]
        OtherTestUnpack.OtherPrimitives primitives;
        uint32 id;
    }

    function unpackInner(bytes memory data) public pure returns (uint256, Inner memory) {
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        Inner memory result;

        bool boolean;
        assembly {
            boolean := shr(248, mload(add(data, 32)))
        }
        assembly {
            let _data_new_len := sub(mload(data), 1)
            data := add(data, 1)
            mstore(data, _data_new_len)
        }
        result.boolean = boolean;
        Primitives memory primitives;
        {
            uint256 _len_before;
            assembly { _len_before := mload(data) }
            uint256 read;
            (read, primitives) = unpackPrimitives(data);
            assembly {
                data := add(data, read)
                mstore(data, sub(_len_before, read))
            }
        }
        result.primitives = primitives;
        uint32 id;
        assembly {
            id := shr(224, mload(add(data, 32)))
        }
        assembly {
            let _data_new_len := sub(mload(data), 4)
            data := add(data, 4)
            mstore(data, _data_new_len)
        }
        result.id = id;
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
    }

    function unpackOtherInner(bytes memory data) public pure returns (uint256, OtherInner memory) {
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        OtherInner memory result;

        bool boolean;
        assembly {
            boolean := shr(248, mload(add(data, 32)))
        }
        assembly {
            let _data_new_len := sub(mload(data), 1)
            data := add(data, 1)
            mstore(data, _data_new_len)
        }
        result.boolean = boolean;
        OtherTestUnpack.OtherPrimitives memory primitives;
        {
            uint256 _len_before;
            assembly { _len_before := mload(data) }
            uint256 read;
            (read, primitives) = OtherTestUnpack.unpackOtherPrimitives(data);
            assembly {
                data := add(data, read)
                mstore(data, sub(_len_before, read))
            }
        }
        result.primitives = primitives;
        uint32 id;
        assembly {
            id := shr(224, mload(add(data, 32)))
        }
        assembly {
            let _data_new_len := sub(mload(data), 4)
            data := add(data, 4)
            mstore(data, _data_new_len)
        }
        result.id = id;
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
    }
}

library OtherTestUnpack {
    // #[unpack(calldata)]
    struct OtherPrimitives {
        bool a;
        address b;
        uint256 c;
        int256 d;
        bytes32 e;
        bytes f;
        string g;
        Primitives primitives;
    }

    function unpackOtherPrimitives(bytes calldata data) public pure returns (uint256, OtherPrimitives memory) {
        uint256 _initial_length = data.length;
        OtherPrimitives memory result;

        bool a = bytes1(data[0:1]) != 0x00;
        data = data[1:];
        result.a = a;
        address b = address(bytes20(data[0:20]));
        data = data[20:];
        result.b = b;
        uint256 c = uint256(bytes32(data[0:32]));
        data = data[32:];
        result.c = c;
        int256 d = int256(uint256(bytes32(data[0:32])));
        data = data[32:];
        result.d = d;
        bytes32 e = bytes32(data[0:32]);
        data = data[32:];
        result.e = e;
        bytes memory f;
        {
            uint256 length = uint256(bytes32(data[0:32]));
            f = bytes(data[32:length + 32]);
            data = data[32 + length:];
        }
        result.f = f;
        string memory g;
        {
            uint256 length = uint256(bytes32(data[0:32]));
            g = string(data[32:length + 32]);
            data = data[32 + length:];
        }
        result.g = g;
        Primitives memory primitives;
        {
            uint256 read;
            (read, primitives) = unpackPrimitives(data);
            data = data[read:];
        }
        result.primitives = primitives;

        return (_initial_length - data.length, result);
    }
}
