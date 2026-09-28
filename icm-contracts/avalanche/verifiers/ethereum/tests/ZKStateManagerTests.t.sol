// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity ^0.8.30;

import {Test} from "forge-std/Test.sol";
import {IRiscZeroVerifier, Receipt as RiscZeroReceipt} from "@risc0/IRiscZeroVerifier.sol";
import {ZKStateManager, ConsensusData, Journal} from "../ZKStateManager.sol";
import {Consensus, Execution} from "../StateManagerLibrary.sol";

/// @dev Accepts every seal, so the tests exercise the contract's own checks rather than proof
/// validity.
// solhint-disable no-empty-blocks
contract AcceptAllVerifier is IRiscZeroVerifier {
    function verify(bytes calldata, bytes32, bytes32) external pure {}

    function verifyIntegrity(
        RiscZeroReceipt calldata
    ) external pure {}
}
// solhint-enable no-empty-blocks

contract ZKStateManagerTest is Test {
    uint64 private constant _SLOT_PER_EPOCH = 32;
    uint64 private constant _SECONDS_PER_SLOT = 12;
    uint64 private constant _EPOCH_SECONDS = _SLOT_PER_EPOCH * _SECONDS_PER_SLOT;
    // Ethereum mainnet beacon genesis.
    uint256 private constant _GENESIS_TIME = 1_606_824_023;
    // One day, which is 225 epochs.
    uint24 private constant _PERMISSIBLE_TIMESPAN = 86_400;
    uint64 private constant _MAX_EPOCH_SPAN = _PERMISSIBLE_TIMESPAN / _EPOCH_SECONDS;
    uint64 private constant _START_EPOCH = 1_000;

    ZKStateManager private _manager;
    address private _admin = makeAddr("admin");

    function setUp() public {
        // Execution-layer verification is not exercised here, so the beacon config can be empty.
        Execution.BeaconConfig memory config;
        _manager = new ZKStateManager({
            newSourceChainId: 1,
            startingState: _state(_START_EPOCH),
            beaconConfig: config,
            genesisTime_: _GENESIS_TIME,
            permissibleTimespan_: _PERMISSIBLE_TIMESPAN,
            verifier_: address(new AcceptAllVerifier()),
            imageID_: bytes32(uint256(1)),
            admin: _admin,
            superAdmin: _admin
        });
        // Finality lags the chain head by about two epochs, so start the clock there.
        vm.warp(_manager.epochTimestamp(_START_EPOCH + 2));
    }

    function testConstructorRejectsZeroGenesisTime() public {
        Execution.BeaconConfig memory config;
        address verifier = address(new AcceptAllVerifier());
        vm.expectRevert("Invalid genesis time");
        new ZKStateManager({
            newSourceChainId: 1,
            startingState: _state(_START_EPOCH),
            beaconConfig: config,
            genesisTime_: 0,
            permissibleTimespan_: _PERMISSIBLE_TIMESPAN,
            verifier_: verifier,
            imageID_: bytes32(uint256(1)),
            admin: _admin,
            superAdmin: _admin
        });
    }

    function testTransitionSucceedsWhenRecent() public {
        uint64 postEpoch = _START_EPOCH + 1;
        vm.warp(_manager.epochTimestamp(postEpoch + 2));

        _manager.transition(_consensusData(_START_EPOCH, postEpoch));

        (bytes32 root, bool valid) = _manager.getBeaconBlockRoot(postEpoch * _SLOT_PER_EPOCH);
        assertTrue(valid);
        assertEq(root, _checkpoint(postEpoch).root);
    }

    function testAcceptsPostStateExactlyAtPermissibleTimespan() public {
        uint64 postEpoch = _START_EPOCH + 1;
        vm.warp(_manager.epochTimestamp(postEpoch) + _PERMISSIBLE_TIMESPAN);

        _manager.transition(_consensusData(_START_EPOCH, postEpoch));

        (, bool valid) = _manager.getBeaconBlockRoot(postEpoch * _SLOT_PER_EPOCH);
        assertTrue(valid);
    }

    function testRevertsWhenPostStateOlderThanPermissibleTimespan() public {
        uint64 postEpoch = _START_EPOCH + 1;
        vm.warp(_manager.epochTimestamp(postEpoch) + _PERMISSIBLE_TIMESPAN + 1);

        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(_consensusData(_START_EPOCH, postEpoch));
    }

    function testRevertsWhenPostStateIsInTheFuture() public {
        uint64 postEpoch = _START_EPOCH + 1;
        vm.warp(_manager.epochTimestamp(postEpoch) - 1);

        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(_consensusData(_START_EPOCH, postEpoch));
    }

    function testRevertsWhenFinalityDoesNotAdvance() public {
        // The pre-state matches the stored state, so only the finality check can reject this.
        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(_consensusData(_START_EPOCH, _START_EPOCH));
    }

    function testAcceptsTransitionSpanningExactlyPermissibleTimespan() public {
        uint64 postEpoch = _START_EPOCH + _MAX_EPOCH_SPAN;
        vm.warp(_manager.epochTimestamp(postEpoch + 2));

        _manager.transition(_consensusData(_START_EPOCH, postEpoch));

        (, bool valid) = _manager.getBeaconBlockRoot(postEpoch * _SLOT_PER_EPOCH);
        assertTrue(valid);
    }

    function testRevertsWhenTransitionSpansMoreThanPermissibleTimespan() public {
        uint64 postEpoch = _START_EPOCH + _MAX_EPOCH_SPAN + 1;
        // The post-state itself is recent; only the span is too long.
        vm.warp(_manager.epochTimestamp(postEpoch + 2));

        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(_consensusData(_START_EPOCH, postEpoch));
    }

    function testRevertsOnPreStateMismatch() public {
        uint64 postEpoch = _START_EPOCH + 2;
        vm.warp(_manager.epochTimestamp(postEpoch + 2));

        vm.expectRevert(ZKStateManager.InvalidPreState.selector);
        _manager.transition(_consensusData(_START_EPOCH + 1, postEpoch));
    }

    /// @dev The scenario from the staleness finding. A proof that was valid when first applied
    /// becomes applicable again after an admin resets the contract to that proof's pre-state.
    /// While the proof's post-state is still recent, re-applying it is harmless and allowed. Once
    /// the post-state is older than permissibleTimespan, replaying it would move the contract onto
    /// an outdated view of the chain, so it must be rejected.
    function testReplayedProofRejectedOnceStaleAfterManualRecovery() public {
        uint64 postEpoch = _START_EPOCH + 1;
        ConsensusData memory proof = _consensusData(_START_EPOCH, postEpoch);
        vm.warp(_manager.epochTimestamp(postEpoch + 2));
        _manager.transition(proof);

        // Admin recovery back to the proof's pre-state, then an immediate replay is still fresh.
        _manualTransition(postEpoch, _START_EPOCH);
        _manager.transition(proof);

        // Recover again, let the post-state go stale, and the replay is now rejected.
        _manualTransition(postEpoch, _START_EPOCH);
        vm.warp(_manager.epochTimestamp(postEpoch) + _PERMISSIBLE_TIMESPAN + 1);
        vm.expectRevert(ZKStateManager.PermissibleTimespanLapsed.selector);
        _manager.transition(proof);
    }

    function testManualTransitionBypassesStalenessCheck() public {
        // Far in the future relative to the starting state, and to an old post-state.
        vm.warp(_manager.epochTimestamp(_START_EPOCH) + 365 days);
        _manualTransition(_START_EPOCH, _START_EPOCH - 100);

        (, bool valid) = _manager.getBeaconBlockRoot((_START_EPOCH - 100) * _SLOT_PER_EPOCH);
        assertTrue(valid);
    }

    function testEpochTimestamp() public view {
        assertEq(_manager.epochTimestamp(0), _GENESIS_TIME);
        assertEq(_manager.epochTimestamp(10), _GENESIS_TIME + 10 * _EPOCH_SECONDS);
        assertEq(_manager.genesisTime(), _GENESIS_TIME);
    }

    // ---------------------------------------------------------------------------------------
    // Helpers
    // ---------------------------------------------------------------------------------------

    function _manualTransition(uint64 preEpoch, uint64 postEpoch) private {
        Journal memory journal = Journal({
            preState: _state(preEpoch),
            postState: _state(postEpoch),
            finalizedSlot: postEpoch * _SLOT_PER_EPOCH
        });
        vm.prank(_admin);
        _manager.manualTransition(abi.encode(journal), postEpoch * _SLOT_PER_EPOCH);
    }

    function _consensusData(
        uint64 preEpoch,
        uint64 postEpoch
    ) private pure returns (ConsensusData memory) {
        Journal memory journal = Journal({
            preState: _state(preEpoch),
            postState: _state(postEpoch),
            finalizedSlot: postEpoch * _SLOT_PER_EPOCH
        });
        return ConsensusData({
            journalData: abi.encode(journal),
            seal: "",
            finalizedSlot: postEpoch * _SLOT_PER_EPOCH
        });
    }

    /// @dev A consensus state finalized at `finalizedEpoch`, with justification one epoch later.
    function _state(
        uint64 finalizedEpoch
    ) private pure returns (Consensus.State memory) {
        return Consensus.State({
            currentJustifiedCheckpoint: _checkpoint(finalizedEpoch + 1),
            finalizedCheckpoint: _checkpoint(finalizedEpoch)
        });
    }

    function _checkpoint(
        uint64 epoch
    ) private pure returns (Consensus.Checkpoint memory) {
        return Consensus.Checkpoint({epoch: epoch, root: keccak256(abi.encode("root", epoch))});
    }
}
