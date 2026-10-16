pragma solidity ^0.8.30;

// #[pack()]
struct HasBadLength {
    // #[pack(length = int32)]
    bytes data;
}