package validatorupdater

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

type ValidatorSetMerkleCommitment struct {
	AvalancheBlockchainID ids.ID
	RootHash              [32]byte
	TotalWeight           uint64
	PChainHeight          uint64
	PChainTimestamp       uint64
}

type ValidatorSetMerkleAttestation struct {
	Signers            []*Validator
	Proof              [][32]byte
	ProofFlags         []bool
	AggregateSignature [192]byte
}

// Node Keeps track if a node is on a path from the
// / root of the merkle tree to one of the leaves
// / being included in a multi-proof.
type Node struct {
	Hash   [32]byte
	OnPath bool
}

func NewValidatorSetMerkleCommitment(
	avalancheBlockchainID ids.ID,
	validators []*Validator,
	pChainHeight uint64,
	pChainTimestamp uint64,
) *ValidatorSetMerkleCommitment {
	rootHash := BuildMerkleRoot(validators)
	totalWeight := sumWeights(validators)
	return &ValidatorSetMerkleCommitment{
		AvalancheBlockchainID: avalancheBlockchainID,
		RootHash:              rootHash,
		TotalWeight:           totalWeight,
		PChainHeight:          pChainHeight,
		PChainTimestamp:       pChainTimestamp,
	}
}

func (v *ValidatorSetMerkleCommitment) Bytes() []byte {
	var buf [88]byte // 32 + 32 + 8 + 8 + 8
	copy(buf[0:32], v.AvalancheBlockchainID[:])
	copy(buf[32:64], v.RootHash[:])
	binary.BigEndian.PutUint64(buf[64:72], v.TotalWeight)
	binary.BigEndian.PutUint64(buf[72:80], v.PChainHeight)
	binary.BigEndian.PutUint64(buf[80:88], v.PChainTimestamp)
	return buf[:]
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

	signers := make([]*Validator, 0)
	layer := make([]Node, len(validators))
	for i, v := range validators {
		byteIdx := i / 8
		bitIdx := uint(i % 8)
		hash := validatorHash(validators[i])
		if byteIdx < len(bitSetSig.Signers) && bitSetSig.Signers[byteIdx]&(1<<bitIdx) != 0 {
			signers = append(signers, v)
			layer[i] = Node{Hash: hash, OnPath: true}
		} else {
			layer[i] = Node{Hash: hash, OnPath: false}
		}
	}

	proof := make([][32]byte, 0)
	proofFlags := make([]bool, 0)
	for len(layer) > 1 {
		nextLen := (len(layer) + 1) / 2
		nextLayer := make([]Node, nextLen)
		for i := 0; i < nextLen; i++ {
			if 2*i+1 < len(layer) {
				nextHash := sha256Pair(layer[2*i].Hash, layer[2*i+1].Hash)
				// A node is on the path if either of its children are on the path.
				nextLayer[i] = Node{Hash: nextHash, OnPath: layer[2*i].OnPath || layer[2*i+1].OnPath}
				if nextLayer[i].OnPath {
					// proof flags are only needed for nodes on the path to distinguish if both children are on the path
					// or only one.
					proofFlags = append(proofFlags, layer[2*i].OnPath && layer[2*i+1].OnPath)
					if !layer[2*i].OnPath {
						proof = append(proof, layer[2*i].Hash)
					} else if !layer[2*i+1].OnPath {
						proof = append(proof, layer[2*i+1].Hash)
					}
				}
			} else {
				nextLayer[i] = layer[2*i]
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

// BuildMerkleRoot returns the root hash of the merkle tree of the given validators along with
// the total weight. This function assumes that the validators are sorted by their public keys.
func BuildMerkleRoot(validators []*Validator) [32]byte {
	layer := make([][32]byte, len(validators))
	for i, v := range validators {
		layer[i] = validatorHash(v)
	}
	for len(layer) > 1 {
		nextLen := (len(layer) + 1) / 2
		nextLayer := make([][32]byte, nextLen)
		for i := 0; i < nextLen; i++ {
			if 2*i+1 < len(layer) {
				nextLayer[i] = sha256Pair(layer[2*i], layer[2*i+1])
			} else {
				nextLayer[i] = layer[2*i]
			}
		}
		layer = nextLayer
	}
	return layer[0]
}

func validatorHash(validator *Validator) [32]byte {
	var buf [104]byte // 96 bytes key + 8 bytes weight
	copy(buf[:96], validator.UncompressedPublicKeyBytes[:])
	binary.BigEndian.PutUint64(buf[96:], validator.Weight)
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
