pragma solidity ^0.8.30;

// #[unpack()]
struct HasFunctionArray {
    function() external[] handlers;
}