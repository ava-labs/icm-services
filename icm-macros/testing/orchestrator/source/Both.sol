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
}