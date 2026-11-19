pragma solidity ^0.8.30;

// #[unpack()]
struct HasAssertOnDefault {
    // #[unpack(default, assert = "|x| { x > 0 }")]
    uint256 value;
}