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

    function parseMethod(bytes calldata data) pure public returns (uint256, Method memory) {
        Method memory res;
        if (data[0]) {
            res.flag = true;
        } else {
            res.flag = false;
        }
        return (1, res);
    }
}