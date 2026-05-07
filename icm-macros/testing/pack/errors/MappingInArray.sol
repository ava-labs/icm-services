pragma solidity ^0.8.30;

// #[pack()]
struct HasMappingArray {
    mapping(uint256 => uint256)[] data;
}