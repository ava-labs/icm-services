pragma solidity ^0.8.30;

// #[pack()]
struct HasMalformedAnnotation {
    // #[pack(length = (1 + 2]
    bytes data;
}
