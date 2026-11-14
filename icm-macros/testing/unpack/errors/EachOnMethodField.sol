pragma solidity ^0.8.30;

// #[unpack()]
struct HasEachOnMethodField {
    // #[unpack(method="readValues", assert = "|each x| { x > 0 }")]
    uint256[] values;
}
