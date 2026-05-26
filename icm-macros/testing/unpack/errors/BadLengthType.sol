pragma solidity ^0.8.30;

// #[unpack()]
struct HasBadLength {
    // #[unpack(length = int32)]
    bytes data;
}