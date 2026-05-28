pragma solidity ^0.8.30;

struct FreeStanding {
    string text;
    bool flag;
}

library PackMethods {
    enum Enum {
        First,
        Second
    }

    // #[pack(name="serializeStruct", visibility="internal")]
    struct Struct {
        bytes inner;
        bool flag;
        // #[pack(ignore)]
        FreeStanding free;
        // #[pack(method="abi.encodePacked")]
        Enum choice;
    }

    // #[pack()]
    struct DynamicStruct {
        // #[pack(length=uint8)]
        bytes Bytes;
        // #[pack(length=uint16)]
        string String;
        // #[pack(length=uint32)]
        address[] Addresses;
    }

    function serializeStruct(Struct memory obj) internal pure returns (bytes memory) {
        return abi.encodePacked(
            abi.encodePacked(obj.inner.length, obj.inner), abi.encodePacked(obj.flag), abi.encodePacked(obj.choice)
        );
    }

    function packDynamicStruct(DynamicStruct memory obj) public pure returns (bytes memory) {
        bytes memory Addresses_bytes;
        Addresses_bytes = abi.encodePacked(uint32(obj.Addresses.length));
        for (uint256 i_0 = 0; i_0 < obj.Addresses.length;) {
            bytes memory Addresses_1_bytes;
            Addresses_1_bytes = abi.encodePacked(obj.Addresses[i_0]);
            Addresses_bytes = abi.encodePacked(Addresses_bytes, Addresses_1_bytes);
            unchecked {
                ++i_0;
            }
        }
        return abi.encodePacked(
            abi.encodePacked(uint8(obj.Bytes.length), obj.Bytes),
            abi.encodePacked(uint16(bytes(obj.String).length), obj.String),
            Addresses_bytes
        );
    }
}
