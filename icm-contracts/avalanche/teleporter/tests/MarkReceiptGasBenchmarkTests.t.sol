// (c) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: LicenseRef-Ecosystem

pragma solidity 0.8.30;

import {
    TeleporterMessengerTest,
    TeleporterMessageReceipt,
    TeleporterMessage,
    WarpMessage
} from "./TeleporterMessengerTest.t.sol";

/**
 * @dev Measures the marginal on-chain gas cost of marking one additional receipt during
 * receiveCrossChainMessage. The relayer in icm-services budgets for this with
 * MarkMessageReceiptGasCost in utils/gas-utils/gas_utils.go; if the constant there is below what
 * this benchmark reports, an attacker can craft a receipt-heavy message (via the permissionless
 * sendSpecifiedReceipts) that the relayer under-provisions and then pays for as an out-of-gas
 * revert.
 *
 * Worst case per receipt, which is what the relayer must budget for:
 *  - 3 cold SLOADs of the SentMessageInfo struct
 *  - 3 SSTOREs clearing it (refunds are credited only if the transaction succeeds and never
 *    reduce the upfront gas requirement, so they cannot be counted here)
 *  - a cold zero-to-nonzero SSTORE crediting _relayerRewardAmounts, which requires a distinct
 *    relayer reward address per receipt and a fee-bearing sent message
 *  - the ReceiptReceived LOG4
 *
 * The sends that create this state happen in setUp() rather than in the test body. Foundry runs
 * setUp() as its own call, so the slots they touch are cold again when the measured receive runs
 * - matching a real delivery, which is always a separate transaction. Measuring the sends and the
 * receive in one test body instead leaves the SentMessageInfo slots warm and understates the cost
 * by the cold-access surcharge on every receipt.
 */
contract MarkReceiptGasBenchmarkTest is TeleporterMessengerTest {
    // Sample sizes. Marginal cost comes from differences between them, so the fixed cost of the
    // receive path cancels out. The cost is expected to be linear in the receipt count, since
    // every receipt does the same fixed amount of work. Three sizes rather than two so the result
    // also confirms that expectation: if the cost were superlinear, no per-receipt constant in
    // gas_utils.go could cover it and the relayer would need a hard cap instead.
    uint256 private constant _FEW_RECEIPTS = 5;
    uint256 private constant _MID_RECEIPTS = 105;
    uint256 private constant _MANY_RECEIPTS = 405;

    // MarkMessageReceiptGasCost in utils/gas-utils/gas_utils.go. The formula there budgets this
    // plus the 64-byte receipt's share of the per-byte decode term, so measuring against this
    // constant alone leaves that decode budget as extra headroom.
    uint256 private constant _RELAYER_PER_RECEIPT_BUDGET = 45_000;

    TeleporterMessageReceipt[] private _fewReceipts;
    TeleporterMessageReceipt[] private _midReceipts;
    TeleporterMessageReceipt[] private _manyReceipts;

    // Seed for deriving relayer reward addresses. Incremented on every use so that no two
    // receipts, within or across batches, can share an address.
    uint256 private _nextRewardAddressSeed;

    function setUp() public virtual override {
        TeleporterMessengerTest.setUp();
        _stageReceipts(_fewReceipts, _FEW_RECEIPTS);
        _stageReceipts(_midReceipts, _MID_RECEIPTS);
        _stageReceipts(_manyReceipts, _MANY_RECEIPTS);
    }

    function testMarginalGasPerReceipt() public {
        // Deliver the largest batch first and the smallest last. The batches touch disjoint
        // per-receipt slots, but the fixed receive-path slots are shared: measuring larger
        // batches while those are still cold biases the differences upward, so the reported
        // marginal costs err on the conservative side.
        uint256 gasMany = _measureReceive(_manyReceipts, 3);
        uint256 gasMid = _measureReceive(_midReceipts, 2);
        uint256 gasFew = _measureReceive(_fewReceipts, 1);

        assertGt(gasMany, gasMid);
        assertGt(gasMid, gasFew);
        uint256 marginalSmall = (gasMid - gasFew) / (_MID_RECEIPTS - _FEW_RECEIPTS);
        uint256 marginalLarge = (gasMany - gasMid) / (_MANY_RECEIPTS - _MID_RECEIPTS);

        emit log_named_uint("gas: receive  405 receipts", gasMany);
        emit log_named_uint("gas: receive  105 receipts", gasMid);
        emit log_named_uint("gas: receive    5 receipts", gasFew);
        emit log_named_uint("marginal gas/receipt 5-105", marginalSmall);
        emit log_named_uint("marginal gas/receipt 105-405", marginalLarge);

        // Sanity floor: a receipt that does real work cannot be cheaper than its cold SLOADs
        // plus the SSTOREs clearing the sent-message state. Catches a silently short-circuited
        // benchmark - a buggy vm.cool made an earlier version of this report an impossibly low
        // number while still passing.
        assertGt(marginalSmall, 15_000, "receipts appear not to be doing real work");

        // The per-receipt cost must not grow with the receipt count, otherwise a single
        // per-receipt constant cannot bound it at any size.
        assertApproxEqRel(
            marginalLarge,
            marginalSmall,
            0.05e18,
            "per-receipt cost is not linear in the receipt count"
        );

        // Guard rail: the relayer's per-receipt budget must cover the measured worst case at
        // every size. If a contract change pushes the real cost past it, this fails and
        // MarkMessageReceiptGasCost must be re-derived from the numbers logged above.
        assertLe(
            marginalSmall,
            _RELAYER_PER_RECEIPT_BUDGET,
            "marginal per-receipt gas exceeds the relayer's MarkMessageReceiptGasCost budget"
        );
        assertLe(
            marginalLarge,
            _RELAYER_PER_RECEIPT_BUDGET,
            "marginal per-receipt gas exceeds the relayer's MarkMessageReceiptGasCost budget"
        );
    }

    /**
     * @dev Sends `count` fee-bearing messages and records a receipt for each into `receipts`.
     * Every receipt names a distinct, never-before-used relayer reward address, derived from the
     * monotonically increasing _nextRewardAddressSeed, so that marking it pays a cold
     * zero-to-nonzero SSTORE on _relayerRewardAmounts.
     */
    function _stageReceipts(TeleporterMessageReceipt[] storage receipts, uint256 count) private {
        for (uint256 i; i < count; ++i) {
            uint256 nonce = _getNextMessageNonce();
            _sendTestMessageWithFee(DEFAULT_SOURCE_BLOCKCHAIN_ID, 1 ether);
            receipts.push(
                TeleporterMessageReceipt({
                    receivedMessageNonce: nonce,
                    relayerRewardAddress: address(
                        uint160(uint256(keccak256(abi.encode(_nextRewardAddressSeed++))))
                    )
                })
            );
        }
    }

    /**
     * @dev Delivers one message carrying `receipts` and returns the gas the call consumed,
     * asserting that the receipts were actually marked rather than short-circuited.
     */
    function _measureReceive(
        TeleporterMessageReceipt[] storage receipts,
        uint256 deliveryNonce
    ) private returns (uint256) {
        TeleporterMessage memory messageToReceive =
            _createMockTeleporterMessage(deliveryNonce, new bytes(0));
        messageToReceive.receipts = receipts;
        WarpMessage memory warpMessage =
            _createDefaultWarpMessage(DEFAULT_SOURCE_BLOCKCHAIN_ID, abi.encode(messageToReceive));
        _setUpSuccessGetVerifiedWarpMessageMock(0, warpMessage);

        uint256 gasBefore = gasleft();
        teleporterMessenger.receiveCrossChainMessage(0, DEFAULT_RELAYER_REWARD_ADDRESS);
        uint256 gasUsed = gasBefore - gasleft();

        // Every receipt must have taken the full path: reward credited and sent-message state
        // cleared. Without this the benchmark would happily measure an early-returning loop.
        for (uint256 i; i < receipts.length; ++i) {
            assertEq(
                teleporterMessenger.checkRelayerRewardAmount(
                    receipts[i].relayerRewardAddress, address(_mockFeeAsset)
                ),
                1 ether,
                "receipt did not credit a reward: benchmark is measuring a no-op"
            );
        }

        return gasUsed;
    }
}
