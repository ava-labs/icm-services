pragma solidity ^0.8.30;

// #[pack()]
struct HasEmptyLength {
    // #[pack(length = )]
    bytes data;
}
