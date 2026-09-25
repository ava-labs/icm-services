// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem
pragma solidity 0.8.30;

library CallUtils {
    /**
     * @dev Upper bound on the gas this frame spends between measuring gasleft() and the target
     * beginning execution, which must be available on top of the gas reserved for the target:
     *  - the CALL opcode's base cost for an address already warmed by the code-size check: 100
     *  - the CALL opcode's value-transfer surcharge, charged whenever value is non-zero: 9,000
     *  - the comparison, require, and memory reads leading up to the CALL: well under 100
     * Rounded up to leave a margin.
     */
    uint256 private constant _CALL_OVERHEAD_GAS = 10_000;

    /**
     * @dev calls target address with exactly gasAmount gas and data as calldata
     * or reverts if at least gasAmount gas is not available.
     */
    function _callWithExactGas(
        uint256 gasAmount,
        address target,
        bytes memory data
    ) internal returns (bool) {
        return _callWithExactGasAndValue(gasAmount, 0, target, data);
    }

    /**
     * @dev calls target address with exactly gasAmount gas and data as calldata
     * or reverts if at least gasAmount gas is not available.
     */
    function _callWithExactGasAndValue(
        uint256 gasAmount,
        uint256 value,
        address target,
        bytes memory data
    ) internal returns (bool) {
        // This early check also bounds gasAmount by the gas actually available, so the arithmetic in
        // the EIP-150 check below cannot overflow.
        require(gasleft() >= gasAmount, "CallUtils: insufficient gas");
        require(address(this).balance >= value, "CallUtils: insufficient value");

        // If there is no code at the target, automatically consider the call to have failed since it
        // doesn't have any effect on state.
        if (target.code.length == 0) {
            return false;
        }

        // EIP-150 caps the gas a CALL can forward at 63/64 of the gas remaining after the CALL's own
        // base cost is paid. Checking gasleft() >= gasAmount alone is therefore not enough: a caller
        // can supply just enough gas to pass that check while the target receives up to 1/64 less
        // than gasAmount, and a failure caused that way is indistinguishable from a genuine one.
        // Require enough gas for the target's full gasAmount, the 1/64 retained by this frame, and
        // the overhead spent before the target starts executing. gasAmount / 63 rounds down; the
        // margin in _CALL_OVERHEAD_GAS covers the rounding.
        require(
            gasleft() >= gasAmount + gasAmount / 63 + _CALL_OVERHEAD_GAS,
            "CallUtils: insufficient gas"
        );

        // Call the target address of the message with the provided data and amount of gas.
        //
        // Assembly is used for the low-level call to avoid unnecessary expansion of the return data in memory.
        // This prevents possible "return bomb" vectors where the external contract could force the caller
        // to use an arbitrary amount of gas. See Solidity issue here: https://github.com/ethereum/solidity/issues/12306
        bool success;
        // solhint-disable-next-line no-inline-assembly
        assembly {
            success :=
                call(
                    gasAmount, // gas provided to the call
                    target, // call target
                    value, // value transferred
                    add(data, 0x20), // input data - 0x20 needs to be added to an array because the first 32-byte slot contains the array length (0x20 in hex is 32 in decimal).
                    mload(data), // input data size - mload returns mem[p..(p+32)], which is the first 32-byte slot of the array. In this case, the array length.
                    0, // output
                    0 // output size
                )
        }
        return success;
    }
}
