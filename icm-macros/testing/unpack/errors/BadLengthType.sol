pragma solidity ^0.8.30;

// #[unpack()]
struct HasBadLength {
    // #[unpack(length = bool)]
    bytes data;
}