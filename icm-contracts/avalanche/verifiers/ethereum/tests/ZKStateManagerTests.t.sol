// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {ISP1Verifier} from "@sp1-contracts/ISP1Verifier.sol";
import {ZKStateManager, ConsensusData, ProofOutputs, StorageSlot} from "../ZKStateManager.sol";
import {Consensus, Execution} from "../StateManagerLibrary.sol";

/// @dev Accepts every proof unless told to reject, so the tests exercise the contract's own
/// checks rather than proof validity.
contract MockSP1Verifier is ISP1Verifier {
    bool public shouldRevert;

    function setShouldRevert(
        bool v
    ) external {
        shouldRevert = v;
    }

    function verifyProof(
        bytes32,
        bytes calldata,
        bytes calldata
    ) external view {
        require(!shouldRevert, "mock: invalid proof");
    }
}

contract ZKStateManagerTest is Test {
    uint64 private constant _SLOT_PER_EPOCH = 32;
    uint64 private constant _SLOTS_PER_PERIOD = 8192;
    uint64 private constant _SECONDS_PER_SLOT = 12;
    // Ethereum mainnet beacon genesis.
    uint256 private constant _GENESIS_TIME = 1_606_824_023;
    // One day, which is 7200 slots.
    uint24 private constant _PERMISSIBLE_TIMESPAN = 86_400;
    uint64 private constant _MAX_SLOT_SPAN = _PERMISSIBLE_TIMESPAN / _SECONDS_PER_SLOT;
    // Four checkpoints before the start of period 100, so tests can cross into the next period.
    uint64 private constant _START_SLOT = 100 * _SLOTS_PER_PERIOD - 4 * _SLOT_PER_EPOCH;
    bytes32 private constant _VKEY = bytes32(uint256(1));

    ZKStateManager private _manager;
    MockSP1Verifier private _verifier;
    address private _admin = makeAddr("admin");

    function setUp() public {
        // Finality lags the chain head by about two epochs, so start the clock there. The
        // constructor rejects a starting head that lies in the future, so warp first.
        vm.warp(_slotTimestamp(_START_SLOT + 2 * _SLOT_PER_EPOCH));
        // Deployed once here so that `vm.expectRevert` in the constructor tests targets the
        // manager's creation rather than the verifier's.
        _verifier = new MockSP1Verifier();
        _manager = _deploy(_START_SLOT, _GENESIS_TIME);
    }

    function testConstructorRejectsZeroGenesisTime() public {
        vm.expectRevert("Invalid genesis time");
        _deploy(_START_SLOT, 0);
    }

    function testConstructorRejectsStartingStateInTheFuture() public {
        vm.warp(_slotTimestamp(_START_SLOT) - 1);
        vm.expectRevert("Starting state in the future");
        _deploy(_START_SLOT, _GENESIS_TIME);
    }

    function testConstructorAcceptsStartingStateAtCurrentTime() public {
        vm.warp(_slotTimestamp(_START_SLOT));
        _deploy(_START_SLOT, _GENESIS_TIME);
    }

    /// @dev A genesis time that is too far ahead makes the starting head read as future.
    function testConstructorRejectsGenesisTimeTooFarAhead() public {
        uint256 age = block.timestamp - _slotTimestamp(_START_SLOT);
        vm.expectRevert("Starting state in the future");
        _deploy(_START_SLOT, _GENESIS_TIME + age + 1);
    }

    function testTransitionSucceedsWhenRecent() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));

        _manager.transition(_consensusData(_START_SLOT, newHead));

        (bytes32 root, bool valid) = _manager.getBeaconBlockRoot(newHead);
        assertTrue(valid);
        assertEq(root, _header(newHead));
        assertEq(_manager.lastConfirmedSlot(), newHead);
        _assertState(newHead);
    }

    function testTransitionVerifiesPublicValuesAgainstVkey() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));
        ConsensusData memory data = _consensusData(_START_SLOT, newHead);

        vm.expectCall(
            address(_verifier),
            abi.encodeCall(ISP1Verifier.verifyProof, (_VKEY, data.publicValues, data.proof))
        );
        _manager.transition(data);
    }

    function testRevertsWhenProofInvalid() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));
        _verifier.setShouldRevert(true);

        vm.expectRevert("mock: invalid proof");
        _manager.transition(_consensusData(_START_SLOT, newHead));
    }

    function testStoresNextSyncCommittee() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));
        uint256 nextPeriod = _period(newHead) + 1;
        assertEq(_manager.syncCommittees(nextPeriod), bytes32(0));

        _manager.transition(_consensusData(_START_SLOT, newHead));

        assertEq(_manager.syncCommittees(nextPeriod), _committee(nextPeriod));
    }

    /// @dev Once the head crosses into a new period, the state tracks that period's committee.
    function testTransitionIntoNextPeriod() public {
        uint64 newHead = _START_SLOT + 4 * _SLOT_PER_EPOCH;
        assertEq(_period(newHead), _period(_START_SLOT) + 1);
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));

        _manager.transition(_consensusData(_START_SLOT, newHead));

        _assertState(newHead);
        assertEq(_manager.syncCommittees(_period(newHead) + 1), _committee(_period(newHead) + 1));
    }

    function testRevertsOnNextSyncCommitteeMismatch() public {
        uint64 firstHead = _START_SLOT + _SLOT_PER_EPOCH;
        uint64 secondHead = firstHead + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(secondHead + 2 * _SLOT_PER_EPOCH));
        _manager.transition(_consensusData(_START_SLOT, firstHead));

        ProofOutputs memory outputs = _outputs(firstHead, secondHead);
        bytes32 wrongNext = keccak256("wrong");
        outputs.nextSyncCommitteeHash = wrongNext;
        vm.expectRevert(
            abi.encodeWithSelector(
                ZKStateManager.NextSyncCommitteeMismatch.selector,
                wrongNext,
                _committee(_period(secondHead) + 1)
            )
        );
        _manager.transition(_toConsensusData(outputs));
    }

    function testRevertsOnNonCheckpointSlot() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH + 1;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));

        vm.expectRevert(abi.encodeWithSelector(ZKStateManager.NonCheckpointSlot.selector, newHead));
        _manager.transition(_consensusData(_START_SLOT, newHead));
    }

    function testAcceptsHeadExactlyAtPermissibleTimespan() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead) + _PERMISSIBLE_TIMESPAN);

        _manager.transition(_consensusData(_START_SLOT, newHead));

        (, bool valid) = _manager.getBeaconBlockRoot(newHead);
        assertTrue(valid);
    }

    function testRevertsWhenHeadOlderThanPermissibleTimespan() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead) + _PERMISSIBLE_TIMESPAN + 1);

        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(_consensusData(_START_SLOT, newHead));
    }

    function testRevertsWhenHeadIsInTheFuture() public {
        uint64 newHead = _START_SLOT + 4 * _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead) - 1);

        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(_consensusData(_START_SLOT, newHead));
    }

    function testRevertsWhenHeadDoesNotAdvance() public {
        // The pre-state matches the stored state, so only the advance check can reject this.
        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(_consensusData(_START_SLOT, _START_SLOT));
    }

    function testAcceptsTransitionSpanningExactlyPermissibleTimespan() public {
        uint64 newHead = _START_SLOT + _MAX_SLOT_SPAN;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));

        _manager.transition(_consensusData(_START_SLOT, newHead));

        (, bool valid) = _manager.getBeaconBlockRoot(newHead);
        assertTrue(valid);
    }

    function testRevertsWhenTransitionSpansMoreThanPermissibleTimespan() public {
        uint64 newHead = _START_SLOT + _MAX_SLOT_SPAN + _SLOT_PER_EPOCH;
        // The new head itself is recent; only the span is too long.
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));

        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(_consensusData(_START_SLOT, newHead));
    }

    function testRevertsOnPrevHeadMismatch() public {
        uint64 newHead = _START_SLOT + 2 * _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));

        vm.expectRevert(ZKStateManager.InvalidPreState.selector);
        _manager.transition(_consensusData(_START_SLOT + _SLOT_PER_EPOCH, newHead));
    }

    function testRevertsOnPrevHeaderMismatch() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));
        ProofOutputs memory outputs = _outputs(_START_SLOT, newHead);
        outputs.prevHeader = keccak256("wrong");

        vm.expectRevert(ZKStateManager.InvalidPreState.selector);
        _manager.transition(_toConsensusData(outputs));
    }

    function testRevertsOnPrevSyncCommitteeMismatch() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));
        ProofOutputs memory outputs = _outputs(_START_SLOT, newHead);
        outputs.prevSyncCommitteeHash = keccak256("wrong");

        vm.expectRevert(ZKStateManager.InvalidPreState.selector);
        _manager.transition(_toConsensusData(outputs));
    }

    /// @dev A proof that was valid when first applied becomes applicable again after an admin
    /// resets the contract to that proof's pre-state. While the proof's new head is still recent,
    /// re-applying it is harmless and allowed. Once the new head is older than
    /// permissibleTimespan, replaying it would move the contract onto an outdated view of the
    /// chain, so it must be rejected.
    function testReplayedProofRejectedOnceStaleAfterManualRecovery() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        ConsensusData memory proof = _consensusData(_START_SLOT, newHead);
        vm.warp(_manager.slotTimestamp(newHead + 2 * _SLOT_PER_EPOCH));
        _manager.transition(proof);

        // Admin recovery back to the proof's pre-state, then an immediate replay is still fresh.
        _manualTransition(_START_SLOT);
        _manager.transition(proof);

        // Recover again, let the new head go stale, and the replay is now rejected.
        _manualTransition(_START_SLOT);
        vm.warp(_manager.slotTimestamp(newHead) + _PERMISSIBLE_TIMESPAN + 1);
        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(proof);
    }

    function testManualTransitionBypassesStalenessCheck() public {
        // Far in the future relative to the starting state, and to an old head.
        vm.warp(_manager.slotTimestamp(_START_SLOT) + 365 days);
        uint64 oldHead = _START_SLOT - 100 * _SLOT_PER_EPOCH;
        _manualTransition(oldHead);

        (, bool valid) = _manager.getBeaconBlockRoot(oldHead);
        assertTrue(valid);
        _assertState(oldHead);
    }

    /// @dev Admin recovery must be able to replace a stored sync committee, for example after
    /// no update landed within a period.
    function testManualTransitionOverwritesSyncCommittees() public {
        uint64 newHead = _START_SLOT + _SLOT_PER_EPOCH;
        uint256 period = _period(newHead);
        ProofOutputs memory outputs = _outputs(_START_SLOT, newHead);
        outputs.syncCommitteeHash = keccak256("replaced current");
        outputs.nextSyncCommitteeHash = keccak256("replaced next");

        vm.prank(_admin);
        _manager.manualTransition(abi.encode(outputs));

        assertEq(_manager.syncCommittees(period), outputs.syncCommitteeHash);
        assertEq(_manager.syncCommittees(period + 1), outputs.nextSyncCommitteeHash);
    }

    function testManualTransitionRequiresAdmin() public {
        vm.expectRevert();
        _manager.manualTransition(abi.encode(_outputs(_START_SLOT, _START_SLOT + _SLOT_PER_EPOCH)));
    }

    function testConstructorStoresStartingState() public view {
        _assertState(_START_SLOT);
        assertEq(_manager.syncCommittees(_period(_START_SLOT)), _committee(_period(_START_SLOT)));
        assertEq(_manager.lastConfirmedSlot(), _START_SLOT);

        // The trusted starting header anchors messages before the first transition.
        (bytes32 root, bool valid) = _manager.getBeaconBlockRoot(_START_SLOT);
        assertTrue(valid);
        assertEq(root, _header(_START_SLOT));
    }

    function testSlotTimestamp() public view {
        assertEq(_manager.slotTimestamp(0), _GENESIS_TIME);
        assertEq(_manager.slotTimestamp(10), _GENESIS_TIME + 10 * _SECONDS_PER_SLOT);
        assertEq(_manager.genesisTime(), _GENESIS_TIME);
    }

    // ---------------------------------------------------------------------------------------
    // Helpers
    // ---------------------------------------------------------------------------------------

    function _deploy(
        uint64 startSlot,
        uint256 genesisTime
    ) private returns (ZKStateManager) {
        // Execution-layer verification is not exercised here, so the beacon config can be empty.
        Execution.BeaconConfig memory config;
        return new ZKStateManager({
            newSourceChainId: 1,
            startingState: _state(startSlot),
            beaconConfig: config,
            genesisTime_: genesisTime,
            permissibleTimespan_: _PERMISSIBLE_TIMESPAN,
            verifier_: address(_verifier),
            lightClientVkey_: _VKEY,
            admin: _admin,
            superAdmin: _admin
        });
    }

    function _manualTransition(
        uint64 newHead
    ) private {
        vm.prank(_admin);
        _manager.manualTransition(abi.encode(_outputs(newHead, newHead)));
    }

    function _assertState(
        uint64 head
    ) private view {
        Consensus.State memory state = _manager.getCurrentState();
        assertEq(state.head, head);
        assertEq(state.header, _header(head));
        assertEq(state.syncCommitteeHash, _committee(_period(head)));
    }

    /// @dev Mirrors {ZKStateManager.slotTimestamp} for use before the manager is deployed.
    function _slotTimestamp(
        uint64 slot
    ) private pure returns (uint256) {
        return _GENESIS_TIME + uint256(slot) * _SECONDS_PER_SLOT;
    }

    function _consensusData(
        uint64 prevHead,
        uint64 newHead
    ) private pure returns (ConsensusData memory) {
        return _toConsensusData(_outputs(prevHead, newHead));
    }

    function _toConsensusData(
        ProofOutputs memory outputs
    ) private pure returns (ConsensusData memory) {
        return ConsensusData({publicValues: abi.encode(outputs), proof: ""});
    }

    /// @dev Public values for an update from `prevHead` to `newHead`, consistent with {_state}.
    function _outputs(
        uint64 prevHead,
        uint64 newHead
    ) private pure returns (ProofOutputs memory) {
        return ProofOutputs({
            prevHeader: _header(prevHead),
            prevHead: prevHead,
            prevSyncCommitteeHash: _committee(_period(prevHead)),
            newHead: newHead,
            newHeader: _header(newHead),
            executionStateRoot: bytes32(0),
            executionBlockNumber: 0,
            syncCommitteeHash: _committee(_period(newHead)),
            nextSyncCommitteeHash: _committee(_period(newHead) + 1),
            storageSlots: new StorageSlot[](0)
        });
    }

    function _state(
        uint64 head
    ) private pure returns (Consensus.State memory) {
        return Consensus.State({
            head: head, header: _header(head), syncCommitteeHash: _committee(_period(head))
        });
    }

    function _period(
        uint64 slot
    ) private pure returns (uint256) {
        return slot / _SLOTS_PER_PERIOD;
    }

    function _header(
        uint64 slot
    ) private pure returns (bytes32) {
        return keccak256(abi.encode("header", slot));
    }

    function _committee(
        uint256 period
    ) private pure returns (bytes32) {
        return keccak256(abi.encode("committee", period));
    }
}
