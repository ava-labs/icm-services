package zkrelayer

import (
	"github.com/ava-labs/libevm/common"
)

// MessageState represents a message's position in the relayer pipeline.
// There are 4 states: detected, submitted, delivered, and failed.
type MessageState string

const (
	// MessageStateDetected: event is observed on the source chain and is awaiting
	// delivery to destination chain, e.g., if the Boundless proof for its anchor slot
	// has not yet been submitted to the destination chain.
	MessageStateDetected MessageState = "detected"
	// MessageStateSubmitted: the event (i.e. its delivery transaction) has been
	// submitted to the destination chain.
	MessageStateSubmitted MessageState = "submitted"
	// MessageStateDelivered: delivery confirmed on the destination chain.
	MessageStateDelivered MessageState = "delivered"
	// MessageStateFailed: gave up after maxAttempts delivery failures.
	MessageStateFailed MessageState = "failed"
)

// PendingMessage tracks one Ethereum-originated message through the pipeline.
// Fields are exported so the record survives JSON persistence in RelayerState.
type PendingMessage struct {
	// TxHash of the source-chain transaction that emitted the message event.
	// Used as the dedup key in RelayerState.PendingMessages.
	TxHash common.Hash `json:"txHash"`
	// BlockNumber of the emitting transaction on the source chain.
	BlockNumber uint64 `json:"blockNumber"`
	// Slot is the beacon chain slot corresponding to BlockNumber.
	Slot uint64 `json:"slot"`
	// LogIndex of the message event within its transaction receipt.
	LogIndex uint `json:"logIndex"`
	// State is the message's position in the pipeline.
	State MessageState `json:"state"`
	// Attempts counts delivery failures
	Attempts uint32 `json:"attempts"`
	// DeliveryTx is the destination-chain transaction hash once Submitted.
	DeliveryTx common.Hash `json:"deliveryTx,omitempty"`
	// LastError records the most recent failure, for observability.
	LastError string `json:"lastError,omitempty"`
}
