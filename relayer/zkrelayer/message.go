package zkrelayer

import (
	"github.com/ava-labs/libevm/common"
)

// MessageState represents a message's position in the relayer pipeline.
type MessageState string

const (
	// MessageStateDetected: event is observed on the source chain and is awaiting
	// delivery to destination chain, e.g., if the Boundless proof for its anchor slot
	// has not yet been submitted to the destination chain.
	MessageStateDetected MessageState = "detected"
	// MessageStateSubmitted: the event (i.e., its delivery transaction) has been
	// submitted to the destination chain.
	MessageStateSubmitted MessageState = "submitted"
	// MessageStateDelivered: delivery confirmed on the destination chain.
	MessageStateDelivered MessageState = "delivered"
	// MessageStateFailed: gave up after maxAttempts delivery failures.
	MessageStateFailed MessageState = "failed"
)

// PendingMessage tracks one Ethereum-originated message through the pipeline.
// Records are held in memory only and rebuilt on startup by rescanning the
// source chain from the persisted slot cursor and the destination-chain for confirmed
// deliveries. The relayer does not persist the message state, attempts, or last error. 
type PendingMessage struct {
	// TxHash of the source-chain transaction that emitted the message event.
	// Used as the dedup key for the in-memory pending set.
	TxHash common.Hash
	// BlockNumber of the emitting transaction on the source chain.
	BlockNumber uint64
	// Slot is the beacon chain slot corresponding to BlockNumber.
	Slot uint64
	// LogIndex of the message event within its transaction receipt.
	LogIndex uint
	// State is the message's position in the pipeline.
	State MessageState
	// Attempts counts delivery failures.
	Attempts uint32
	// DeliveryTx is the destination-chain transaction hash once Submitted.
	DeliveryTx common.Hash
	// LastError records the most recent failure, for observability.
	LastError string
}
