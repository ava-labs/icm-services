// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {CallUtils} from "../CallUtils.sol";

/// @dev Exposes the internal library functions so they run in their own call frame with a
/// controllable gas allowance.
contract CallUtilsHarness {
    receive() external payable {}

    function callWithExactGas(
        uint256 gasAmount,
        address target,
        bytes memory data
    ) external returns (bool) {
        return CallUtils._callWithExactGas(gasAmount, target, data);
    }

    function callWithExactGasAndValue(
        uint256 gasAmount,
        uint256 value,
        address target,
        bytes memory data
    ) external returns (bool) {
        return CallUtils._callWithExactGasAndValue(gasAmount, value, target, data);
    }
}

/// @dev Records how much gas it was given on entry, so tests can check the gas that actually
/// reached the _target rather than the gas the caller believed it forwarded.
contract GasRecorder {
    uint256 public lastGasObserved;
    uint256 public callCount;
    uint256 public lastPayloadLength;
    bytes32 public lastPayloadHash;

    function record() external payable {
        lastGasObserved = gasleft();
        callCount++;
    }

    /// @dev Same as {record}, but takes an arbitrarily large payload so tests can check that the
    /// gas guarantee holds regardless of calldata size and that the payload arrives intact.
    /// gasleft() is read before anything proportional to the payload size is done.
    function recordPayload(
        bytes calldata payload
    ) external payable {
        lastGasObserved = gasleft();
        callCount++;
        lastPayloadLength = payload.length;
        lastPayloadHash = keccak256(payload);
    }
}

// solhint-disable avoid-low-level-calls
contract CallUtilsTest is Test {
    uint256 private constant _GAS_AMOUNT = 1_000_000;

    CallUtilsHarness private _harness;
    GasRecorder private _target;
    bytes private _recordCall;

    // Gas the _target consumes between being entered and reading gasleft() in {GasRecorder-record}.
    uint256 private _targetEntryOverhead;
    // The same for {GasRecorder-recordPayload}, whose ABI decoding costs slightly more.
    uint256 private _payloadEntryOverhead;

    function setUp() public {
        _harness = new CallUtilsHarness();
        _target = new GasRecorder();
        _recordCall = abi.encodeCall(GasRecorder.record, ());
        vm.deal(address(_harness), 100 ether);

        // Measure the _target's entry overhead with a direct call that forwards exactly this much gas.
        uint256 probeGas = 200_000;
        (bool ok,) = address(_target).call{gas: probeGas}(_recordCall);
        require(ok, "probe call failed");
        _targetEntryOverhead = probeGas - _target.lastGasObserved();
        assertLt(_targetEntryOverhead, 1_000, "unexpected entry overhead");

        // Decoding a bytes calldata parameter costs a fixed amount regardless of its length, so
        // probing with an empty payload measures the overhead for every payload size.
        (ok,) = address(_target).call{gas: probeGas}(_payloadCall(0));
        require(ok, "payload probe call failed");
        _payloadEntryOverhead = probeGas - _target.lastGasObserved();
        assertLt(_payloadEntryOverhead, 1_000, "unexpected payload entry overhead");
    }

    function testTargetReceivesGasAmountWithAmpleGas() public {
        bool success = _harness.callWithExactGas(_GAS_AMOUNT, address(_target), _recordCall);
        assertTrue(success);
        assertGe(_target.lastGasObserved(), _GAS_AMOUNT - _targetEntryOverhead);
    }

    /// @dev Regression test: before the EIP-150 adjustment, a caller supplying only slightly more
    /// than gasAmount passed the gas check, yet the CALL could forward only 63/64 of what remained,
    /// so the _target silently received less than gasAmount. That must now revert instead.
    function testRevertsWhenOnlyGasAmountPlusSmallMarginAvailable() public {
        bytes memory callData =
            abi.encodeCall(_harness.callWithExactGas, (_GAS_AMOUNT, address(_target), _recordCall));
        uint256 callsBefore = _target.callCount();
        (bool ok, bytes memory ret) = address(_harness).call{gas: _GAS_AMOUNT + 5_000}(callData);
        assertFalse(ok);
        assertEq(ret, abi.encodeWithSignature("Error(string)", "CallUtils: insufficient gas"));
        assertEq(_target.callCount(), callsBefore, "target must not have been called");
    }

    /// @dev For any gas allowance, the library must either revert or hand the _target at least
    /// gasAmount. It must never complete having given the _target less.
    function testFuzzTargetNeverReceivesLessThanGasAmount(
        uint256 outerGas
    ) public {
        outerGas = bound(outerGas, 0, 2 * _GAS_AMOUNT);
        bytes memory callData =
            abi.encodeCall(_harness.callWithExactGas, (_GAS_AMOUNT, address(_target), _recordCall));
        uint256 callsBefore = _target.callCount();
        (bool ok, bytes memory ret) = address(_harness).call{gas: outerGas}(callData);
        if (!ok) {
            return;
        }
        assertTrue(abi.decode(ret, (bool)), "call reported failure");
        assertEq(_target.callCount(), callsBefore + 1);
        assertGe(_target.lastGasObserved(), _GAS_AMOUNT - _targetEntryOverhead);
    }

    function testFuzzTargetNeverReceivesLessThanGasAmountWithValue(
        uint256 outerGas
    ) public {
        outerGas = bound(outerGas, 0, 2 * _GAS_AMOUNT);
        uint256 value = 1 ether;
        bytes memory callData = abi.encodeCall(
            _harness.callWithExactGasAndValue, (_GAS_AMOUNT, value, address(_target), _recordCall)
        );
        uint256 callsBefore = _target.callCount();
        (bool ok, bytes memory ret) = address(_harness).call{gas: outerGas}(callData);
        if (!ok) {
            return;
        }
        assertTrue(abi.decode(ret, (bool)), "call reported failure");
        assertEq(_target.callCount(), callsBefore + 1);
        assertGe(_target.lastGasObserved(), _GAS_AMOUNT - _targetEntryOverhead);
        assertEq(address(_target).balance, value);
    }

    /// @dev The minimum gas the library accepts is tight: only a little more than
    /// gasAmount + gasAmount / 63 plus the documented overhead is needed.
    function testSucceedsJustAboveMinimum() public {
        // Gas the _harness itself spends before reaching the library's check (dispatch and copying
        // calldata into memory), plus the library's own overhead constant, generously bounded.
        uint256 slack = 20_000;
        uint256 outerGas = _GAS_AMOUNT + _GAS_AMOUNT / 63 + slack;
        bytes memory callData =
            abi.encodeCall(_harness.callWithExactGas, (_GAS_AMOUNT, address(_target), _recordCall));
        (bool ok, bytes memory ret) = address(_harness).call{gas: outerGas}(callData);
        assertTrue(ok, "expected call to be accepted");
        assertTrue(abi.decode(ret, (bool)));
        assertGe(_target.lastGasObserved(), _GAS_AMOUNT - _targetEntryOverhead);
    }

    function testRevertsWhenGasAmountExceedsAvailable() public {
        bytes memory callData =
            abi.encodeCall(_harness.callWithExactGas, (_GAS_AMOUNT, address(_target), _recordCall));
        (bool ok, bytes memory ret) = address(_harness).call{gas: _GAS_AMOUNT / 2}(callData);
        assertFalse(ok);
        assertEq(ret, abi.encodeWithSignature("Error(string)", "CallUtils: insufficient gas"));
    }

    function testRevertsWhenGasAmountIsAbsurd() public {
        vm.expectRevert("CallUtils: insufficient gas");
        _harness.callWithExactGas(type(uint256).max, address(_target), _recordCall);
    }

    function testRevertsOnInsufficientValue() public {
        vm.expectRevert("CallUtils: insufficient value");
        _harness.callWithExactGasAndValue(_GAS_AMOUNT, 1_000 ether, address(_target), _recordCall);
    }

    function testReturnsFalseForTargetWithoutCode() public {
        address noCode = makeAddr("noCode");
        bool success = _harness.callWithExactGas(_GAS_AMOUNT, noCode, _recordCall);
        assertFalse(success);
    }

    function testReturnsFalseWhenTargetReverts() public {
        RevertingTarget reverting = new RevertingTarget();
        bool success = _harness.callWithExactGas(
            _GAS_AMOUNT, address(reverting), abi.encodeWithSignature("boom()")
        );
        assertFalse(success);
    }

    // -------------------------------------------------------------------------------------------
    // Large calldata
    //
    // CALL charges nothing per byte of input: the input already lives in this frame's memory, so
    // no memory expansion is paid at the CALL either. The only cost that grows with the payload is
    // paid by the _harness copying calldata into memory before the library ever reads gasleft(),
    // so _CALL_OVERHEAD_GAS does not need to scale with payload size. These tests pin that down
    // by finding the smallest gas allowance the library accepts for each payload size and checking
    // that the target still receives the full gasAmount at exactly that boundary.
    // -------------------------------------------------------------------------------------------

    function testTargetReceivesGasAmountWithLargeCalldata() public {
        uint256[3] memory sizes = [uint256(1 * 1024), 32 * 1024, 128 * 1024];
        for (uint256 i = 0; i < sizes.length; i++) {
            bytes memory payload = _payload(sizes[i]);
            uint256 callsBefore = _target.callCount();
            bool success = _harness.callWithExactGas(
                _GAS_AMOUNT, address(_target), abi.encodeCall(GasRecorder.recordPayload, (payload))
            );
            assertTrue(success);
            assertEq(_target.callCount(), callsBefore + 1);
            assertGe(_target.lastGasObserved(), _GAS_AMOUNT - _payloadEntryOverhead);
            assertEq(_target.lastPayloadLength(), sizes[i], "payload length");
            assertEq(_target.lastPayloadHash(), keccak256(payload), "payload contents");
        }
    }

    function testTargetReceivesGasAmountAtMinimumWithLargeCalldata() public {
        uint256[4] memory sizes = [uint256(0), 1 * 1024, 32 * 1024, 128 * 1024];
        for (uint256 i = 0; i < sizes.length; i++) {
            _assertGasAmountForwardedAtMinimum(sizes[i], 0);
        }
    }

    function testTargetReceivesGasAmountAtMinimumWithLargeCalldataAndValue() public {
        uint256[4] memory sizes = [uint256(0), 1 * 1024, 32 * 1024, 128 * 1024];
        for (uint256 i = 0; i < sizes.length; i++) {
            _assertGasAmountForwardedAtMinimum(sizes[i], 1 ether);
        }
    }

    function testFuzzTargetNeverReceivesLessThanGasAmountWithLargeCalldata(
        uint256 outerGas,
        uint256 payloadSize
    ) public {
        payloadSize = bound(payloadSize, 0, 64 * 1024);
        outerGas = bound(outerGas, 0, 3 * _GAS_AMOUNT);
        bytes memory callData = abi.encodeCall(
            _harness.callWithExactGas, (_GAS_AMOUNT, address(_target), _payloadCall(payloadSize))
        );
        uint256 callsBefore = _target.callCount();
        (bool ok, bytes memory ret) = address(_harness).call{gas: outerGas}(callData);
        if (!ok) {
            return;
        }
        assertTrue(abi.decode(ret, (bool)), "call reported failure");
        assertEq(_target.callCount(), callsBefore + 1);
        assertGe(_target.lastGasObserved(), _GAS_AMOUNT - _payloadEntryOverhead);
        assertEq(_target.lastPayloadLength(), payloadSize);
    }

    /// @dev Finds the smallest gas allowance for which the _harness call is accepted, then checks
    /// that at exactly that allowance the _target still received gasAmount and the whole payload.
    /// If CALL charged for its input, the library's check would pass at this boundary while the
    /// _target received less than gasAmount, and the gas assertion would fail.
    function _assertGasAmountForwardedAtMinimum(
        uint256 payloadSize,
        uint256 value
    ) private {
        bytes memory payload = _payload(payloadSize);
        bytes memory callData = abi.encodeCall(
            _harness.callWithExactGasAndValue,
            (
                _GAS_AMOUNT,
                value,
                address(_target),
                abi.encodeCall(GasRecorder.recordPayload, (payload))
            )
        );

        uint256 minGas = _minimumAcceptedGas(callData, 3 * _GAS_AMOUNT);
        assertFalse(_harnessAccepts(callData, minGas - 1), "minimum is not tight");

        uint256 callsBefore = _target.callCount();
        uint256 balanceBefore = address(_target).balance;
        (bool ok, bytes memory ret) = address(_harness).call{gas: minGas}(callData);
        assertTrue(ok, "expected call to be accepted");
        assertTrue(abi.decode(ret, (bool)), "call reported failure");
        assertEq(_target.callCount(), callsBefore + 1);
        assertGe(
            _target.lastGasObserved(),
            _GAS_AMOUNT - _payloadEntryOverhead,
            "target received less than gasAmount at the minimum allowance"
        );
        assertEq(_target.lastPayloadLength(), payloadSize, "payload length");
        assertEq(_target.lastPayloadHash(), keccak256(payload), "payload contents");
        assertEq(address(_target).balance, balanceBefore + value);
    }

    /// @dev Binary search for the smallest gas allowance at which the _harness call does not
    /// revert. Whether the call reverts is monotonic in the allowance, so the search is valid.
    function _minimumAcceptedGas(
        bytes memory callData,
        uint256 hi
    ) private returns (uint256) {
        require(_harnessAccepts(callData, hi), "upper bound too low");
        uint256 lo = 0;
        while (lo < hi) {
            uint256 mid = (lo + hi) / 2;
            if (_harnessAccepts(callData, mid)) {
                hi = mid;
            } else {
                lo = mid + 1;
            }
        }
        return hi;
    }

    function _harnessAccepts(
        bytes memory callData,
        uint256 outerGas
    ) private returns (bool) {
        (bool ok,) = address(_harness).call{gas: outerGas}(callData);
        return ok;
    }

    function _payloadCall(
        uint256 size
    ) private pure returns (bytes memory) {
        return abi.encodeCall(GasRecorder.recordPayload, (_payload(size)));
    }

    /// @dev A payload of the given size filled with a non-repeating pattern, so a forwarding bug
    /// that zeroed or truncated the data would change its hash.
    function _payload(
        uint256 size
    ) private pure returns (bytes memory payload) {
        payload = new bytes(size);
        // Each word is the hash of the word before it, starting from the length word. `new bytes`
        // allocates a whole number of words, so writing the final partial word is in bounds.
        // solhint-disable-next-line no-inline-assembly
        assembly {
            let ptr := add(payload, 0x20)
            let end := add(ptr, size)
            for {} lt(ptr, end) {} {
                mstore(ptr, keccak256(sub(ptr, 0x20), 0x20))
                ptr := add(ptr, 0x20)
            }
        }
    }
}

contract RevertingTarget {
    function boom() external pure {
        revert("boom");
    }
}
