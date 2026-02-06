// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

/**
 * @title ByteComparator
 * @notice Library for comparing byte arrays lexicographically
 */
library ByteComparator {
    /**
     * @notice Compares two byte arrays (lexicographically).
     * @return result -1 if a < b, 0 if a == b, 1 if a > b
     */
    function _compareBytes(bytes memory a, bytes memory b) internal pure returns (int256) {
        if (a.length == b.length) {
             bool samePointer;
             assembly { samePointer := eq(a, b) }
             if (samePointer) return 0;
        }

        uint256 minLength = a.length < b.length ? a.length : b.length;
        uint256 aPtr;
        uint256 bPtr;

        assembly {
            aPtr := add(a, 32) 
            bPtr := add(b, 32)
        }

        for (uint256 i = 0; i < minLength; i += 32) {
            bytes32 wordA;
            bytes32 wordB;

            assembly {
                wordA := mload(add(aPtr, i))
                wordB := mload(add(bPtr, i))
            }

            if (wordA != wordB) {
                return wordA < wordB ? int256(-1) : int256(1);
            }
        }

        if (a.length < b.length) return -1;
        if (a.length > b.length) return 1;
        
        return 0;
    }
}
