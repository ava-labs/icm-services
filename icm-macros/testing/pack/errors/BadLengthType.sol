pragma solidity ^0.8.30;

// #[pack()]
struct HasBadLength {
    // #[pack(length = bool)]
    bytes data;
}