package validatorupdater

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp/message"
)

// ValidatorSetMerkleCommitment is the warp payload for Merkle validator set
// registration. Wire layout (linear codec type ID 6) matches
// platformvm/warp/message on avalanchego branches that register
// … ValidatorSetMerkleCommitment in seventh position.
//
// serialized holds the full codec output after
// [initializeValidatorSetMerkleCommitment]; it is not part of the struct
// encoding.
type ValidatorSetMerkleCommitment struct {
	serialized []byte `serialize:"false"`

	AvalancheBlockchainID ids.ID   `serialize:"true" json:"blockchainID"`
	RootHash              [32]byte `serialize:"true" json:"rootHash"`
	TotalWeight           uint64   `serialize:"true" json:"totalWeight"`
	PChainHeight          uint64   `serialize:"true" json:"pChainHeight"`
	PChainTimestamp       uint64   `serialize:"true" json:"pChainTimestamp"`
}

func (v *ValidatorSetMerkleCommitment) initialize(b []byte) {
	v.serialized = b
}

type validatorSetMerkleCommitmentPayload interface {
	Bytes() []byte
	initialize([]byte)
}

const merkleCommitmentCodecVersion = 0

// merkleCommitmentCodec registers the same first seven warp message types as
// current avalanchego development branches, so ValidatorSetMerkleCommitment
// keeps type ID 6.
var merkleCommitmentCodec codec.Manager

func init() {
	merkleCommitmentCodec = codec.NewManager(math.MaxInt)
	lc := linearcodec.NewDefault()
	err := errors.Join(
		lc.RegisterType(&message.SubnetToL1Conversion{}),
		lc.RegisterType(&message.RegisterL1Validator{}),
		lc.RegisterType(&message.L1ValidatorRegistration{}),
		lc.RegisterType(&message.L1ValidatorWeight{}),
		lc.RegisterType(&ValidatorSetMetadata{}),
		lc.RegisterType(&ValidatorSetDiff{}),
		lc.RegisterType(&ValidatorSetMerkleCommitment{}),
		merkleCommitmentCodec.RegisterCodec(merkleCommitmentCodecVersion, lc),
	)
	if err != nil {
		panic(err)
	}
}

type ValidatorSetMerkleAttestation struct {
	Signers            []*Validator
	Proof              [][32]byte
	ProofFlags         []bool
	AggregateSignature [192]byte
}

// Node Keeps track if a node is on a path from the
// root of the merkle tree to one of the leaves
// being included in a multi-proof.
type Node struct {
	Hash   [32]byte
	OnPath bool
}

func (v *ValidatorSetMerkleCommitment) Bytes() []byte {
	return v.serialized
}

// NewValidatorSetMerkleCommitment builds and serializes a
// ValidatorSetMerkleCommitment payload.
func NewValidatorSetMerkleCommitment(
	avalancheBlockchainID ids.ID,
	validators []*Validator,
	pChainHeight uint64,
	pChainTimestamp uint64,
) (*ValidatorSetMerkleCommitment, error) {
	rootHash := BuildMerkleRoot(validators)
	totalWeight := sumWeights(validators)
	msg := &ValidatorSetMerkleCommitment{
		AvalancheBlockchainID: avalancheBlockchainID,
		RootHash:              rootHash,
		TotalWeight:           totalWeight,
		PChainHeight:          pChainHeight,
		PChainTimestamp:       pChainTimestamp,
	}
	return msg, initializeValidatorSetMerkleCommitment(msg)
}

func initializeValidatorSetMerkleCommitment(v *ValidatorSetMerkleCommitment) error {
	var p validatorSetMerkleCommitmentPayload = v
	b, err := merkleCommitmentCodec.Marshal(merkleCommitmentCodecVersion, &p)
	if err != nil {
		return fmt.Errorf("couldn't marshal ValidatorSetMerkleCommitment payload: %w", err)
	}
	v.initialize(b)
	return nil
}

// NewValidatorSetMerkleAttestation constructs a ValidatorSetMerkleAttestation from a full
// ordered validator set and a BitSetSignature. It uses the signature's signer bitset to
// identify which validators signed, then walks the Merkle tree bottom-up to produce a
// multi-proof (proof nodes and proof flags) that allows a verifier to reconstruct the root
// from only the signing leaves.
func NewValidatorSetMerkleAttestation(
	validators []*Validator,
	bitSetSig *avalancheWarp.BitSetSignature,
) (*ValidatorSetMerkleAttestation, error) {
	sig, err := bls.SignatureFromBytes(bitSetSig.Signature[:])
	if err != nil {
		return &ValidatorSetMerkleAttestation{}, fmt.Errorf("failed to decompress BLS signature: %w", err)
	}

	// Pad to the next power of two so the tree is balanced and the resulting
	// multi-proof is compatible with OpenZeppelin's MerkleProof.multiProofVerify.
	n := nextPow2(len(validators))
	signers := make([]*Validator, 0)
	layer := make([]Node, n)
	// AvalancheGo encodes BitSetSignature.Signers as a big-endian big.Int
	// (bytes[0] is most significant). Use set.BitsFromBytes to correctly map
	// bit index i to validator i regardless of how many bytes are present.
	signerBits := set.BitsFromBytes(bitSetSig.Signers)
	for i, v := range validators {
		hash := validatorHash(validators[i])
		if signerBits.Contains(i) {
			signers = append(signers, v)
			layer[i] = Node{Hash: hash, OnPath: true}
		} else {
			layer[i] = Node{Hash: hash, OnPath: false}
		}
	}
	// Null padding — never on the signing path.
	for i := len(validators); i < n; i++ {
		layer[i] = Node{Hash: nullLeafHash, OnPath: false}
	}

	// Build the multi-proof bottom-up over the (always-even) padded layer.
	proof := make([][32]byte, 0)
	proofFlags := make([]bool, 0)
	for len(layer) > 1 {
		nextLen := len(layer) / 2 // always even
		nextLayer := make([]Node, nextLen)
		for i := 0; i < nextLen; i++ {
			left, right := layer[2*i], layer[2*i+1]
			nextHash := sha256Pair(left.Hash, right.Hash)
			onPath := left.OnPath || right.OnPath
			nextLayer[i] = Node{Hash: nextHash, OnPath: onPath}
			if onPath {
				// Flag indicates whether both children contributed to the path.
				proofFlags = append(proofFlags, left.OnPath && right.OnPath)
				if !left.OnPath {
					proof = append(proof, left.Hash)
				} else if !right.OnPath {
					proof = append(proof, right.Hash)
				}
			}
		}
		layer = nextLayer
	}

	aggregateSig := sig.Serialize()
	return &ValidatorSetMerkleAttestation{
		Signers:            signers,
		Proof:              proof,
		ProofFlags:         proofFlags,
		AggregateSignature: [192]byte(aggregateSig),
	}, nil
}

func (v *ValidatorSetMerkleAttestation) Bytes() []byte {
	numFlags := len(v.ProofFlags)
	flagBytesLen := (numFlags + 7) / 8

	buf := make([]byte, 0, 2+4+len(v.Signers)*(96+8)+4+len(v.Proof)*32+4+flagBytesLen+192)

	// codec (2 zero bytes)
	buf = append(buf, 0, 0)

	// number of signers, uncompressed public key (96 bytes), weight (8 bytes) each
	buf = binary.BigEndian.AppendUint32(buf, uint32(len(v.Signers)))
	for _, s := range v.Signers {
		buf = append(buf, s.UncompressedPublicKeyBytes[:]...)
		buf = binary.BigEndian.AppendUint64(buf, s.Weight)
	}

	// length of proof, hashes (32 bytes) each
	buf = binary.BigEndian.AppendUint32(buf, uint32(len(v.Proof)))
	for _, p := range v.Proof {
		buf = append(buf, p[:]...)
	}

	// number of proof flags, then flags packed as a bitset (LSB-first within each byte)
	buf = binary.BigEndian.AppendUint32(buf, uint32(numFlags))
	packedFlags := make([]byte, flagBytesLen)
	for i, flag := range v.ProofFlags {
		if flag {
			packedFlags[i>>3] |= 1 << uint(i&7)
		}
	}
	buf = append(buf, packedFlags...)

	// aggregate BLS signature (192 bytes)
	buf = append(buf, v.AggregateSignature[:]...)

	return buf
}

// nullLeafHash is the leaf hash used to pad the validator set to the next power
// of two. It is sha256 of 136 zero bytes (the null padded key + zero weight),
// which cannot be produced by any real validator.
var nullLeafHash = sha256.Sum256(make([]byte, 136))

// nextPow2 returns the smallest power of two >= n (minimum 1).
func nextPow2(n int) int {
	p := 1
	for p < n {
		p <<= 1
	}
	return p
}

// BuildMerkleRoot returns the root hash of the Merkle tree of the given validators.
// Validators must be sorted by public key. The leaf set is padded to the next
// power of two with nullLeafHash so that the resulting tree is always balanced
// and compatible with OpenZeppelin's MerkleProof.multiProofVerify.
func BuildMerkleRoot(validators []*Validator) [32]byte {
	if len(validators) == 0 {
		return [32]byte{}
	}
	n := nextPow2(len(validators))
	layer := make([][32]byte, n)
	for i, v := range validators {
		layer[i] = validatorHash(v)
	}
	for i := len(validators); i < n; i++ {
		layer[i] = nullLeafHash
	}
	for len(layer) > 1 {
		nextLen := len(layer) / 2 // always even — no promotion needed
		nextLayer := make([][32]byte, nextLen)
		for i := 0; i < nextLen; i++ {
			nextLayer[i] = sha256Pair(layer[2*i], layer[2*i+1])
		}
		layer = nextLayer
	}
	return layer[0]
}

// validatorHash computes the Merkle leaf hash for a validator using the 128-byte
// BLST-padded public key format: [16 zeros][48 bytes X][16 zeros][48 bytes Y].
// This matches Solidity's sha256(abi.encodePacked(paddedKey, weight)) in
// ValidatorSets.verifyMerkleAttestation.
func validatorHash(validator *Validator) [32]byte {
	var buf [136]byte // 128 bytes padded key + 8 bytes weight
	// X coordinate (bytes 0-47 of uncompressed key) → bytes 16-63 of padded key
	copy(buf[16:64], validator.UncompressedPublicKeyBytes[:48])
	// Y coordinate (bytes 48-95 of uncompressed key) → bytes 80-127 of padded key
	copy(buf[80:128], validator.UncompressedPublicKeyBytes[48:96])
	binary.BigEndian.PutUint64(buf[128:], validator.Weight)
	return sha256.Sum256(buf[:])
}

// sha256Pair hashes two 32-byte values together, sorting them lexicographically
// before hashing so that sha256Pair(a, b) == sha256Pair(b, a).
func sha256Pair(a, b [32]byte) [32]byte {
	if bytes.Compare(a[:], b[:]) > 0 {
		a, b = b, a
	}
	return sha256.Sum256(append(a[:], b[:]...))
}
