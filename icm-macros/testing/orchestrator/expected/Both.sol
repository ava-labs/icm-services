pragma solidity ^0.8.30;

library Messages {
    // #[pack()]
    // #[unpack()]
    enum Kind {
        First,
        Second
    }

    // #[pack(name="encodeHeader", visibility="internal")]
    // #[unpack(name="decodeHeader", visibility="internal")]
    struct Header {
        uint64 height;
        Kind kind;
        // #[pack(ignore)]
        // #[unpack(default)]
        bytes cache;
        // #[pack(method="encodeTag")]
        // #[unpack(method="decodeTag")]
        bytes4 tag;
    }

    function encodeTag(bytes4 tag) internal pure returns (bytes memory) {
        return abi.encodePacked(tag);
    }

    function decodeTag(bytes memory data) internal pure returns (uint256, bytes4) {
        return (4, bytes4(data));
    }

    function encodeHeader(Header memory obj) internal pure returns (bytes memory) {
        return abi.encodePacked(abi.encodePacked(obj.height), packKind(obj.kind), encodeTag(obj.tag));
    }

    function packKind(Kind obj) public pure returns (bytes memory) {
        return abi.encodePacked(obj);
    }

    function decodeHeader(bytes memory data) internal pure returns (uint256, Header memory) {
        /* solhint-disable no-inline-assembly */
        /* solhint-disable var-name-mixedcase */
        uint256 _initial_length;
        assembly { _initial_length := mload(data) }
        Header memory result;

        uint64 height;
        assembly {
            height := shr(192, mload(add(data, 32)))
        }
        assembly {
            let _data_new_len := sub(mload(data), 8)
            data := add(data, 8)
            mstore(data, _data_new_len)
        }
        result.height = height;
        Kind kind;
        {
            uint256 _len_before;
            assembly { _len_before := mload(data) }
            uint256 read;
            (read, kind) = unpackKind(data);
            assembly {
                data := add(data, read)
                mstore(data, sub(_len_before, read))
            }
        }
        result.kind = kind;
        bytes4 tag;
        {
            uint256 _len_before;
            assembly { _len_before := mload(data) }
            uint256 read;
            (read, tag) = decodeTag(data);
            assembly {
                data := add(data, read)
                mstore(data, sub(_len_before, read))
            }
        }
        result.tag = tag;
        uint256 _final_length;
        assembly { _final_length := mload(data) }
        return (_initial_length - _final_length, result);
        /* solhint-enable no-inline-assembly */
        /* solhint-enable var-name-mixedcase */
    }

    function unpackKind(bytes memory data) public pure returns (uint256, Kind) {
        return (1, Kind(uint8(data[0])));
    }
}