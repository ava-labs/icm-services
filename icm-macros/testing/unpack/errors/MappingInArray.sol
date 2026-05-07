pragma solidity ^0.8.30;

// #[unpack()]
struct HasMappingArray {
    mapping(uint256 => uint256)[] data;
}