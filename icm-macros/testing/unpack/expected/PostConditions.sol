pragma solidity ^0.8.30;

library PackMethods {
    // #[unpack(
    //    calldata,
    //    assert = "|t| { t.Bytes.length == bytes(t.String).length }",
    //  )]
    struct Struct {
        bytes Bytes;
        string String;
        // #[unpack(
        //    assert = "|inner| { inner.length > 0 }",
        //    assert = "|each v| { v > 0 }",
        //  )]
        Inner[] inner;
        // #[unpack(
        //    assert = "|m| { m.flag }",
        //    method = "parseMethod",
        //  )]
        Method method;
    }

    // #[unpack()]
    struct Inner {
        //#[unpack(assert = "|v| { v < 10 }")]
        uint256 value;
    }

    struct Method {
        bool flag;
    }

    function parseMethod(bytes calldata data) public pure returns (uint256, Method memory) {
        Method memory res;
        if (data[0]) {
            res.flag = true;
        } else {
            res.flag = false;
        }
        return (1, res);
    }

    function unpackStruct(bytes calldata data) public pure returns (uint256, Struct memory) {
        uint256 _initial_length = data.length;
        Struct memory result;

        bytes memory Bytes;
        {
            uint256 length = uint256(bytes32(data[0:32]));
            Bytes = bytes(data[32:length + 32]);
            data = data[32 + length:];
        }
        result.Bytes = Bytes;
        string memory String;
        {
            uint256 length = uint256(bytes32(data[0:32]));
            String = string(data[32:length + 32]);
            data = data[32 + length:];
        }
        result.String = String;
        Inner[] memory inner;
        {
            uint256 length = uint256(bytes32(data[0:32]));
            data = data[32:];
            inner = new Inner[](length);
            for (uint256 i = 0; i < length;) {
                Inner memory inner_1;
                {
                    uint256 read;
                    (read, inner_1) = unpackInner(data);
                    data = data[read:];
                }
                inner[i] = inner_1;
                require(inner_1 > 0);
                unchecked {
                    ++i;
                }
            }
        }
        result.inner = inner;
        require(inner.length > 0);
        Method memory method;
        {
            uint256 read;
            (read, method) = parseMethod(data);
            data = data[read:];
        }
        result.method = method;
        require(method.flag);
        require(result.Bytes.length == bytes(result.String).length);
        return (_initial_length - data.length, result);
    }

    function unpackInner(bytes memory data) public pure returns (uint256, Inner memory) {
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        Inner memory result;

        uint256 value;
        assembly {
            value := shr(0, mload(add(data, 32)))
        }
        assembly {
            let _data_new_len := sub(mload(data), 32)
            data := add(data, 32)
            mstore(data, _data_new_len)
        }
        result.value = value;
        require(value < 10);
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
    }
}
