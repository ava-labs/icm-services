pragma solidity ^0.8.30;

// #[unpack()]
struct HasEachOnNonContainer {
    // #[unpack(assert = "|each x| { x > 0 }")]
    uint256 value;
}