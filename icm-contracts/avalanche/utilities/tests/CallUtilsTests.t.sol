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

    function record() external payable {
        lastGasObserved = gasleft();
        callCount++;
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
}

contract RevertingTarget {
    function boom() external pure {
        revert("boom");
    }
}
