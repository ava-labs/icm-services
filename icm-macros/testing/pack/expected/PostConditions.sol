pragma solidity ^0.8.30;

library PackMethods {
    // #[pack(assert = "|b| { b.length == 20 }")]
    struct Fixed {
        address addr;
    }

    // #[pack(
    //   assert = "|b| { b.length > 0 }",
    //   assert = "|b| { b.length % 2 == 0 }",
    // )]
    struct MultiAssert {
        bytes data;
    }

    // #[pack(name = "serializeNamed", assert = "|b| { b.length == 48 }")]
    struct Named {
        bytes data;
    }

    // #[pack(assert = "|b| { b.length == 1 }")]
    enum Status {
        Active,
        Inactive
    }

    function packFixed(Fixed memory obj) public pure returns (bytes memory) {
        bytes memory result = abi.encodePacked(abi.encodePacked(obj.addr));
        {
            require(result.length == 20);
        }
        return result;
    }

    function packMultiAssert(MultiAssert memory obj) public pure returns (bytes memory) {
        bytes memory result = abi.encodePacked(abi.encodePacked(obj.data.length, obj.data));
        {
            require(result.length > 0);
        }
        {
            require(result.length % 2 == 0);
        }
        return result;
    }

    function serializeNamed(Named memory obj) public pure returns (bytes memory) {
        bytes memory result = abi.encodePacked(abi.encodePacked(obj.data.length, obj.data));
        {
            require(result.length == 48);
        }
        return result;
    }

    function packStatus(Status obj) public pure returns (bytes memory) {
        bytes memory result = abi.encodePacked(obj);
        {
            require(result.length == 1);
        }
        return result;
    }
}
