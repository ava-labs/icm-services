pragma solidity ^0.8.30;

struct Inner {
    uint256 value;
}

// #[unpack()]
struct HasEachOnCustom {
    // #[unpack(assert = "|each x| { x.value > 0 }")]
    Inner inner;
}