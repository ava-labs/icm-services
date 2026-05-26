pragma solidity ^0.8.30;

library PackMethods {
    // #[unpack()]
    struct DynamicStruct {
        // #[unpack(length=uint8)]
        bytes Bytes;
        // #[unpack(length=uint16)]
        string String;
        // #[unpack(length=uint32)]
        address[] Addresses;
        // #[unpack(length=32)]
        bytes Hash;
    }

    // #[unpack(calldata)]
    struct DynamicStructCalldata {
        // #[unpack(length=uint8)]
        bytes Bytes;
        // #[unpack(length=uint16)]
        string String;
        // #[unpack(length=uint32)]
        address[] Addresses;
        // #[unpack(length=32, method="parseHash")]
        bytes Hash;
    }

    function parseHash(bytes calldata hashBytes) internal pure returns (bytes memory) {
        if (hashBytes.length == 32) {
            return hashBytes;
        } else {
            return new bytes(32);
        }
    }

    function unpackDynamicStruct(bytes memory data) public pure returns (uint256, DynamicStruct memory) {
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        DynamicStruct memory result;

        bytes memory Bytes;
        {
            uint256 data_length;
            assembly { data_length := shr(248, mload(add(data, 32))) }
            Bytes = new bytes(data_length);
            assembly {
                mcopy(add(Bytes, 32), add(data, 33), data_length)
                let _data_orig_len := mload(data)
                data := add(data, add(1, data_length))
                mstore(data, sub(_data_orig_len, add(1, data_length)))
            }
        }
        result.Bytes = Bytes;
        string memory String;
        {
            uint256 data_length;
            assembly { data_length := shr(240, mload(add(data, 32))) }
            String = string(new bytes(data_length));
            assembly {
                mcopy(add(String, 32), add(data, 34), data_length)
                let _data_orig_len := mload(data)
                data := add(data, add(2, data_length))
                mstore(data, sub(_data_orig_len, add(2, data_length)))
            }
        }
        result.String = String;
        address[] memory Addresses;
        {
            uint256 length;
            assembly { length := shr(224, mload(add(data, 32))) }
            assembly {
                let _data_new_len := sub(mload(data), 4)
                data := add(data, 4)
                mstore(data, _data_new_len)
            }
            Addresses = new address[](length);
            for (uint256 i = 0; i < length;) {
                address Addresses_1;
                assembly {
                    Addresses_1 := shr(96, mload(add(data, 32)))
                }
                assembly {
                    let _data_new_len := sub(mload(data), 20)
                    data := add(data, 20)
                    mstore(data, _data_new_len)
                }
                Addresses[i] = Addresses_1;
                unchecked {
                    ++i;
                }
            }
        }
        result.Addresses = Addresses;
        bytes memory Hash;
        {
            Hash = new bytes(32);
            assembly {
                mcopy(add(Hash, 32), add(data, 32), 32)
                let _data_orig_len := mload(data)
                data := add(data, 32)
                mstore(data, sub(_data_orig_len, 32))
            }
        }
        result.Hash = Hash;
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
    }

    function unpackDynamicStructCalldata(bytes calldata data)
        public
        pure
        returns (uint256, DynamicStructCalldata memory)
    {
        uint256 _initial_length = data.length;
        DynamicStructCalldata memory result;

        bytes memory Bytes;
        {
            uint256 length = uint256(uint8(bytes1(data[0:1])));
            Bytes = bytes(data[1:length + 1]);
            data = data[1 + length:];
        }
        result.Bytes = Bytes;
        string memory String;
        {
            uint256 length = uint256(uint16(bytes2(data[0:2])));
            String = string(data[2:length + 2]);
            data = data[2 + length:];
        }
        result.String = String;
        address[] memory Addresses;
        {
            uint256 length = uint256(uint32(bytes4(data[0:4])));
            data = data[4:];
            Addresses = new address[](length);
            for (uint256 i = 0; i < length;) {
                address Addresses_1 = address(bytes20(data[0:20]));
                data = data[20:];
                Addresses[i] = Addresses_1;
                unchecked {
                    ++i;
                }
            }
        }
        result.Addresses = Addresses;
        bytes memory Hash;
        {
            Hash = parseHash(data[0:32]);
            data = data[32:];
        }
        result.Hash = Hash;

        return (_initial_length - data.length, result);
    }
}
