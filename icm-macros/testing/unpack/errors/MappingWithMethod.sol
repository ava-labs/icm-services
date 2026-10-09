pragma solidity ^0.8.30;

// #[unpack()]
struct HasMappingWithMethod {
    // #[unpack(method="readData")]
    mapping(uint256 => uint256) data;
}
