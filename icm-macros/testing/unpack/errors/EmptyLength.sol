pragma solidity ^0.8.30;

// #[unpack()]
struct HasEmptyLength {
    // #[unpack(length = )]
    bytes data;
}
