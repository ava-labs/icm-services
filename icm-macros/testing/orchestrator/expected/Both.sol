pragma solidity ^0.8.30;

library Messages {
    // #[pack()]
    // #[unpack()]
    enum Kind {
        First,
        Second
    }

    // #[pack(name="encodeHeader", visibility="internal")]
    // #[unpack(name="decodeHeader", visibility="internal", calldata)]
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

    function decodeHeader(bytes calldata data) internal pure returns (uint256, Header memory) {
        /* solhint-disable var-name-mixedcase */
        uint256 _initial_length = data.length;
        Header memory result;

        uint64 height = uint64(bytes8(data[0:8]));
        data = data[8:];
        result.height = height;
        Kind kind;
        {
            uint256 read;
            (read, kind) = unpackKind(data);
            data = data[read:];
        }
        result.kind = kind;
        bytes4 tag;
        {
            uint256 read;
            (read, tag) = decodeTag(data);
            data = data[read:];
        }
        result.tag = tag;

        return (_initial_length - data.length, result);
        /* solhint-enable var-name-mixedcase */
    }

    function unpackKind(bytes memory data) public pure returns (uint256, Kind) {
        return (1, Kind(uint8(data[0])));
    }
}