// Copyright (C) 2026, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package zk

import (
	"bytes"
	"fmt"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

// findMessageLog finds our message's log in the receipt and returns its
// position in the receipt's log list. A transaction can emit many logs, so
// we match by emitter, event type, and payload bytes. Note that this is the
// position within the receipt, not log.Index.
func findMessageLog(
	receipt *types.Receipt,
	protocolAddress common.Address,
	payload []byte,
) (uint, error) {
	for i, log := range receipt.Logs {
		if log.Address == protocolAddress &&
			len(log.Topics) > 0 &&
			log.Topics[0] == messageSentTopic &&
			bytes.Equal(log.Data, payload) {
			return uint(i), nil
		}
	}
	return 0, fmt.Errorf("no matching TeleporterV2MessageSent log in receipt")
}
