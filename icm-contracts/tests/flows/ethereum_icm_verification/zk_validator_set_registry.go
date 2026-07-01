package ethereum_icm_verification

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"os"
	"path/filepath"

	zkregistry "github.com/ava-labs/icm-services/abi-bindings/go/ZKValidatorSetRegistry"
	sp1groth16verifier "github.com/ava-labs/icm-services/abi-bindings/go/external/sp1-contracts/v6.1.0/SP1VerifierGroth16"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/libevm/accounts/abi"
	"github.com/ava-labs/libevm/accounts/abi/bind"
	"github.com/ava-labs/libevm/common"
	. "github.com/onsi/gomega"
)

// warpEnvelopeHeaderLen is the fixed header buildUnsignedWarpMessage prepends (address(0) path):
// codec(2)|networkID(4)|sourceChainID(32)|payloadLen(4)|acCodec(2)|typeID(4)|srcAddrLen(4)=0|innerLen(4)
// = 56 bytes. The inner payload the contract re-wraps is signedData[56:].
const warpEnvelopeHeaderLen = 56

// zkFixture is the fixture format read from tests/testdata/zk-groth16-fixture.json
type zkFixture struct {
	Vkey         string `json:"vkey"`
	PublicValues string `json:"publicValues"`
	Proof        string `json:"proof"`
	SignedData   string `json:"signedData"`
}

// decodedPublicValues are the four values the guest commits, recovered from the fixture so the
// test can seed a commitment the proof's on-chain checks pass against.
type decodedPublicValues struct {
	sourceBlockchainID [32]byte
	root               [32]byte
	messageHash        [32]byte
	signedWeight       uint64
}

/**
* ZK analog of the MerkleValidatorSetRegistry roundtrip. Instead of building a Merkle
* multi-inclusion proof and verifying an aggregate BLS signature on-chain, a single committed
* SP1 Groth16 proof (generated offline by zk-proofs-gen and checked in under tests/testdata)
* is verified on-chain by ZKValidatorSetRegistry via SP1VerifierGroth16.
*
* The committed proof is over a FIXED message and a synthetic validator set (known keys), so this
* test exercises the ZK VERIFICATION PATH only — not live signature aggregation, which is left as
* future work (it would require generating a fresh proof per message via the SP1 network prover).
*
* Steps:
* 1. Load the committed fixture and decode its public values.
* 2. Deploy a real SP1VerifierGroth16 and a ZKValidatorSetRegistry whose stored P-Chain commitment
*    (root, totalWeight, chain ID, networkID) matches the proof's public values and signed message.
* 3. Submit the fixture proof (over the matching inner payload) through verifyICMMessage, a view that
*    reverts unless the SP1 proof verifies and all four public-value bindings hold.
 */
func ZKValidatorSetRegistry(
	ctx context.Context,
	ethereumNetwork *localnetwork.LocalEthereumNetwork,
) {
	_, ethFundedKey := ethereumNetwork.GetFundedAccountInfo()
	ethereumOpts, err := bind.NewKeyedTransactorWithChainID(ethFundedKey, ethereumNetwork.ChainID)
	Expect(err).Should(BeNil())

	// load and decode the test fixture
	fixture := loadZKFixture()
	publicValues := decodeZKPublicValues(fixture.PublicValues)
	var vkey [32]byte
	copy(vkey[:], common.FromHex(fixture.Vkey))

	signedData := common.FromHex(fixture.SignedData)
	Expect(len(signedData)).Should(BeNumerically(">", warpEnvelopeHeaderLen))

	networkID := binary.BigEndian.Uint32(signedData[2:6])

	// obtain the raw message (inner payload) since this is what the contract rebuilds the warp envelope from.
	rawMessage := signedData[warpEnvelopeHeaderLen:]

	// sanity check the fixture's signedData hashes to the committed messageHash
	Expect(sha256.Sum256(signedData)).Should(Equal(publicValues.messageHash))

	// deploy the Groth16 verifier (no constructor args) and wait for it to land
	sp1VerifierAddr, tx, _, err := sp1groth16verifier.DeploySP1VerifierGroth16(ethereumOpts, ethereumNetwork.EthClient)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, ethereumNetwork.EthClient, tx.Hash())

	// deploy the ZK registry, seeding its P-Chain commitment so the proof's public values bind:
	// root/sourceBlockchainID match the stored commitment, and totalWeight == signedWeight clears quorum.
	pChainTotalWeight := publicValues.signedWeight
	_, tx, zkRegistry, err := zkregistry.DeployZKValidatorSetRegistry(
		ethereumOpts,
		ethereumNetwork.EthClient,
		networkID,
		publicValues.sourceBlockchainID,
		publicValues.root,
		pChainTotalWeight,
		uint64(1),
		uint64(1),
		false,
		sp1VerifierAddr,
		vkey,
	)
	Expect(err).Should(BeNil())
	utils.WaitForTransactionSuccess(ctx, ethereumNetwork.EthClient, tx.Hash())

	// build the ICM message and submit it to verifyICMMessage, the address(0) path that matches
	// how the proof's signed data was built. It reverts unless the proof verifies and all four
	// public-value bindings (sourceBlockchainID, root, messageHash, signedWeight) hold.
	attestation := packZKAttestation(fixture)
	icmMessage := zkregistry.ICMMessage{
		RawMessage:         rawMessage,
		SourceNetworkID:    networkID,
		SourceBlockchainID: publicValues.sourceBlockchainID,
		Attestation:        attestation,
	}
	err = zkRegistry.VerifyICMMessage(&bind.CallOpts{}, icmMessage, publicValues.sourceBlockchainID)
	Expect(err).Should(BeNil())
}

// loadZKFixture reads the committed Groth16 fixture from tests/testdata/zk-groth16-fixture.json
func loadZKFixture() zkFixture {
	path := filepath.Join("tests", "testdata", "zk-groth16-fixture.json")
	data, err := os.ReadFile(path)
	Expect(err).Should(BeNil())
	var fixture zkFixture
	Expect(json.Unmarshal(data, &fixture)).Should(BeNil())
	return fixture
}

// decodeZKPublicValues abi-decodes the committed PublicValues struct
func decodeZKPublicValues(publicValuesHex string) decodedPublicValues {
	b32, err := abi.NewType("bytes32", "", nil)
	Expect(err).Should(BeNil())
	u64, err := abi.NewType("uint64", "", nil)
	Expect(err).Should(BeNil())

	args := abi.Arguments{{Type: b32}, {Type: b32}, {Type: b32}, {Type: u64}}
	vals, err := args.Unpack(common.FromHex(publicValuesHex))
	Expect(err).Should(BeNil())

	return decodedPublicValues{
		sourceBlockchainID: vals[0].([32]byte),
		root:               vals[1].([32]byte),
		messageHash:        vals[2].([32]byte),
		signedWeight:       vals[3].(uint64),
	}
}

// packZKAttestation produces the attestation blob that _verifyZKAttestation decodes into (publicValues, proofBytes)
func packZKAttestation(fixture zkFixture) []byte {
	bytesT, err := abi.NewType("bytes", "", nil)
	Expect(err).Should(BeNil())

	args := abi.Arguments{{Type: bytesT}, {Type: bytesT}}
	packed, err := args.Pack(common.FromHex(fixture.PublicValues), common.FromHex(fixture.Proof))
	Expect(err).Should(BeNil())
	return packed
}
