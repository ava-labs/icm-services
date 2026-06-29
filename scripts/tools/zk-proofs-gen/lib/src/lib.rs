//! (c) 2026, Ava Labs, Inc. All rights reserved.
//! See the file LICENSE for licensing terms.
//! SPDX-License-Identifier: LicenseRef-Ecosystem
//!
//! Validator-set Merkle attestation verification shared between the SP1 guest program and the host
//! prover. Parses a `ValidatorSetMerkleAttestation`, rebuilds the Merkle root, verifies the
//! aggregate BLS12-381 signature, and checks that the committed signing weight is the true sum of
//! the signers' weights. `PublicValues` is the ABI-encoded output the contract decodes.
//!
//! The attestation (signers, multiproof, flags, aggregate signature) and the signed message are
//! private inputs: they go into the proof but are never revealed on-chain. The proof attests that,
//! given those private inputs, the signer leaves hash up to a Merkle `root`, their aggregate
//! signature is valid over a message whose hash is `messageHash`, and their combined `signedWeight`
//! is the true sum of the signers' weights. Those values are committed publicly. The contract then
//! checks them against its own state — `root` and `sourceBlockchainID` against the stored
//! commitment, `messageHash` against the message it reconstructs — and applies the stake-weighted
//! quorum threshold to `signedWeight` on-chain. So a single proof plus a few small public values
//! stands in for verifying the multiproof and aggregate signature in calldata, while the signers
//! and signature themselves stay private.
//!
//! THIS IS AN EXAMPLE OF UNAUDITED CODE. DO NOT USE THIS IN PRODUCTION.

use bitvec::vec::BitVec;
use bls12_381::{
    hash_to_curve::{ExpandMsgXmd, HashToCurve},
    multi_miller_loop, G1Affine, G1Projective, G2Affine, G2Prepared, G2Projective, Gt,
};
use serde::Deserialize;
use sha2::{Digest, Sha256};

use alloy_sol_types::sol;

pub const BLS_PUBLIC_KEY_SIZE: usize = 96;
pub const BLS_SIGNATURE_SIZE: usize = 192;

pub type Hash = [u8; 32];

sol! {
    // ABI-encoded public values the guest commits and ZKValidatorSetRegistry decodes.
    // Field order/types must match the contract's PublicValues struct exactly.
    struct PublicValues {
        bytes32 sourceBlockchainID;
        bytes32 root;
        bytes32 messageHash;
        uint64  signedWeight;
    }
}

#[derive(Clone, Debug, PartialEq, Eq)]
pub struct Validator {
    pub public_key: G1Affine,
    pub weight: u64,
}

pub struct ValidatorSetMerkleAttestation {
    pub signers: Vec<Validator>,
    pub proof: Vec<Hash>,
    pub flags: BitVec<u8>,
    pub aggregate_sig: G2Affine,
}

impl<'de> Deserialize<'de> for ValidatorSetMerkleAttestation {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        let bytes = Vec::<u8>::deserialize(deserializer)?;
        parse_attestation(&bytes).ok_or_else(|| serde::de::Error::custom("invalid attestation"))
    }
}

pub fn verify(
    attestation: &ValidatorSetMerkleAttestation,
    message: &[u8],
    root: &Hash,
    message_hash: &Hash,
    signed_weight: u64,
) -> Option<()> {
    // Bind the proof to a specific message before any expensive crypto.
    let actual_hash: [u8; 32] = Sha256::digest(message).into();
    if &actual_hash != message_hash {
        return None;
    }

    // Aggregate signer public keys
    let agg_pk = G1Affine::from(
        attestation
            .signers
            .iter()
            .fold(G1Projective::identity(), |acc, v| {
                acc + G1Projective::from(v.public_key)
            }),
    );

    let flags_len = attestation.flags.len();
    let leaves_len = attestation.signers.len();
    let mut leaves = attestation.signers.iter().map(hash_validator);
    let mut hashes = vec![Hash::default(); flags_len];
    let mut proof_hashes = attestation.proof.iter();

    if leaves_len + attestation.proof.len() != flags_len + 1 {
        return None;
    }

    // walk up the tree, hashing the nodes as we go
    let mut hash_pos = 0usize;
    for i in 0..flags_len {
        let a = leaves.next().or_else(|| {
            let h = hashes.get(hash_pos).cloned();
            hash_pos += 1;
            h
        })?;
        let b = if attestation.flags.as_bitslice()[i] {
            leaves.next().or_else(|| {
                let h = hashes.get(hash_pos).cloned();
                hash_pos += 1;
                h
            })
        } else {
            proof_hashes.next().cloned()
        }?;
        hashes[i] = hash_node(&a, &b);
    }

    // Extract the reconstructed Merkle root from the multiproof walk.
    let computed_root = if flags_len > 0 {
        // Real tree: all proof hashes must have been consumed
        // and the root is the last computed hash.
        if proof_hashes.next().is_some() {
            return None;
        }
        hashes.last().cloned()?
    } else if leaves_len > 0 {
        // Single validator: no tree to walk, the root is its leaf hash.
        hash_validator(attestation.signers.first()?)
    } else {
        // No leaves: the root is the lone proof hash.
        attestation.proof.first().cloned()?
    };

    // Check the obtained root hash against the expected public value from the contract
    if &computed_root != root {
        return None;
    }

    // Verify the aggregate signature over the message
    let sig_valid = verify_signature(message, &agg_pk, attestation.aggregate_sig);
    if !sig_valid {
        return None;
    }

    // Assert the committed signing weight is the true sum of the signers' weights; the
    // contract applies the stake-weighted quorum threshold on-chain against its stored total.
    let signing_weight: u64 = attestation.signers.iter().map(|v| v.weight).sum();
    if signing_weight != signed_weight {
        None
    } else {
        Some(())
    }
}

pub fn verify_signature(message: &[u8], key: &G1Affine, signature: G2Affine) -> bool {
    const DST: &[u8] = b"BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_";
    let msg_hash = G2Affine::from(
        <G2Projective as HashToCurve<ExpandMsgXmd<Sha256>>>::hash_to_curve([message], DST),
    );

    // e(agg_pk, H(msg)) == e(G1_gen, sig)
    // ⟺ e(-G1_gen, sig) · e(agg_pk, H(msg)) == identity
    let neg_gen = G1Affine::from(-G1Projective::generator());
    multi_miller_loop(&[
        (&neg_gen, &G2Prepared::from(signature)),
        (key, &G2Prepared::from(msg_hash)),
    ])
    .final_exponentiation()
        == Gt::identity()
}

pub fn parse_attestation(mut attestation: &[u8]) -> Option<ValidatorSetMerkleAttestation> {
    const VALIDATOR_SIZE: usize = 8 + BLS_PUBLIC_KEY_SIZE;

    attestation
        .get(0..2)
        .map(|codec| codec[0] == 0x01 && codec[1] == 0x00)?;
    let num_signer_slice = attestation[2..6].try_into().ok()?;
    let num_signers = u32::from_be_bytes(num_signer_slice);
    attestation = &attestation[6..];

    // check ahead that we have enough bytes to parse the signers
    if num_signers * VALIDATOR_SIZE as u32 > attestation.len() as u32 {
        return None;
    }
    let mut signers = Vec::with_capacity(num_signers as usize);
    for _ in 0..num_signers {
        signers.push(Validator {
            public_key: G1Affine::from_uncompressed(
                attestation[..BLS_PUBLIC_KEY_SIZE].as_array().unwrap(),
            )
            .into_option()?,
            weight: u64::from_be_bytes(
                attestation[BLS_PUBLIC_KEY_SIZE..VALIDATOR_SIZE]
                    .try_into()
                    .ok()?,
            ),
        });
        attestation = &attestation[VALIDATOR_SIZE..];
    }

    let num_hashes_slice = attestation[..4].try_into().ok()?;
    let num_hashes = u32::from_be_bytes(num_hashes_slice);
    attestation = &attestation[4..];

    // check ahead that we have enough bytes to parse the hashes
    if num_hashes * 32 > attestation.len() as u32 {
        return None;
    }
    let mut hashes = Vec::<Hash>::with_capacity(num_hashes as usize);
    for _ in 0..num_hashes {
        hashes.push(attestation[..32].try_into().ok()?);
        attestation = &attestation[32..];
    }

    let num_flags_slice = attestation[..4].try_into().ok()?;
    let num_flags = u32::from_be_bytes(num_flags_slice);
    attestation = &attestation[4..];

    // check ahead that we have enough bytes to parse the flags
    let num_flag_bytes = num_flags.div_ceil(8) as usize;
    if num_flag_bytes > attestation.len() {
        return None;
    }

    let mut flags: BitVec<u8> = BitVec::from_slice(&attestation[..num_flag_bytes]);
    flags.truncate(num_flags as usize);
    attestation = &attestation[num_flag_bytes..];

    if attestation.len() != BLS_SIGNATURE_SIZE {
        None
    } else {
        Some(ValidatorSetMerkleAttestation {
            signers,
            proof: hashes,
            flags,
            aggregate_sig: {
                let agg_sig = attestation.try_into().ok()?;
                // Parse aggregate G2 signature
                G2Affine::from_uncompressed(&agg_sig).into_option()?
            },
        })
    }
}

pub fn hash_node(a: &Hash, b: &Hash) -> Hash {
    let (left, right) = if a <= b { (a, b) } else { (b, a) };
    let mut hasher = Sha256::new();
    hasher.update([0u8; 32]); // uint256(0) prefix for internal nodes
    hasher.update(left);
    hasher.update(right);
    hasher.finalize().into()
}

pub fn hash_validator(v: &Validator) -> Hash {
    let pk = v.public_key.to_uncompressed();
    let mut hasher = Sha256::new();
    hasher.update([0u8; 31]); // uint256(1) big-endian, high bytes
    hasher.update([1u8]); // uint256(1) low byte
    hasher.update([0u8; 16]); // x coordinate padding (48 → 64 bytes)
    hasher.update(&pk[..48]); // x coordinate
    hasher.update([0u8; 16]); // y coordinate padding (48 → 64 bytes)
    hasher.update(&pk[48..]); // y coordinate
    hasher.update(v.weight.to_be_bytes());
    hasher.finalize().into()
}

/// Fixtures derived from the icm-services test vectors. Enable with the `test-utils` feature.
#[cfg(any(test, feature = "test-utils"))]
pub mod test_fixtures {
    use super::*;
    use bls12_381::{G1Projective, Scalar};

    pub const ATTESTATION_HEX: &str = "0000000000030572cbea904d67468808c8eb50a9450c9721db309128012543902d0ac358a62ae28f75bb8f1c7c42c39a8c5529bf0f4e166a9d8cabc673a322fda673779d8e3822ba3ecb8670e461f73bb9021d5fd76a4c56d9d4cd16bd1bba86881979749d2800000000000000010c9b60d5afcbd5663a8a44b7c5a02f19e9a77ab0a35bd65809bb5c67ec582c897feb04decc694b13e08587f3ff9b5b60143be6d078c2b79a7d4f1d1b21486a030ec93f56aa54e1de880db5a66dd833a652a95bee27c824084006cb5644cbd43f000000000000000310e7791fb972fe014159aa33a98622da3cdc98ff707965e536d8636b5fcc5ac7a91a8c46e59a00dca575af0f18fb13dc16ba437edcc6551e30c10512367494bfb6b01cc6681e8a4c3cd2501832ab5c4abc40b4578b85cbaffbf0bcd70d67c6e20000000000000004000000017e5866de1dc819a4346a48c956a12e000bf26f3a6f61e0549e01d720f703d0550000000306055273f06e2a177237e5e949abe52c0194ed2a84a8bdb1b04986e721f1206c025f599458c7be4b7f9a133cb291f52900102f103b08360bb6142102804cb972f388a21142167d041b2d81dccfe29ef3ad7a1c1b7a5b725938167ebd089ee30e300fdadf227f40a0e9f6b168555871d5a992f0fbf8d9c65bde4f140f0552f38d009334987b6fcd56a051a63718716ff4220eb925350b4cb88d9d3abb10237f3da95c6d619ec19fa48bd8da6ec4469d8f71395e4681c9540c97a0afc5af8e457413";

    pub const SIGNED_DATA_HEX: &str = "0000000000013d0ad12b8ee8928edf248ca91ca55600fb383f07c32bff1d6dec472b25cf59a7000000c6000000000001000000145615deb798bb3e4dfa0139dfa1b3d433cc23b72f000000a40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000";

    /// Total weight of all 4 validators in the test set (weights 1, 2, 3, 4).
    pub const TOTAL_WEIGHT: u64 = 10;

    /// Weight of the 3 signers in the test attestation (v0=1, v2=3, v3=4).
    pub const SIGNING_WEIGHT: u64 = 8;

    /// The source blockchain ID committed in the signed warp message (bytes 6..38 of
    /// SIGNED_DATA_HEX, i.e. the warp message's sourceChainID).
    pub const SOURCE_BLOCKCHAIN_ID: Hash = [
        0x3d, 0x0a, 0xd1, 0x2b, 0x8e, 0xe8, 0x92, 0x8e, 0xdf, 0x24, 0x8c, 0xa9, 0x1c, 0xa5, 0x56,
        0x00, 0xfb, 0x38, 0x3f, 0x07, 0xc3, 0x2b, 0xff, 0x1d, 0x6d, 0xec, 0x47, 0x2b, 0x25, 0xcf,
        0x59, 0xa7,
    ];

    /// Computes the expected Merkle root for the 4-validator test set.
    ///
    /// Validator weights: v0=1, v1=2, v2=3, v3=4 (derived from scalars 2, 3, 4, 5).
    /// Tree: hash_node(hash_node(v0, v1), hash_node(v2, v3))
    pub fn expected_root() -> Hash {
        let validators: Vec<Validator> = [2u64, 3, 4, 5]
            .iter()
            .enumerate()
            .map(|(ix, i)| Validator {
                public_key: G1Affine::from(G1Projective::generator() * Scalar::from(*i)),
                weight: ix as u64 + 1,
            })
            .collect();

        hash_node(
            &hash_node(
                &hash_validator(&validators[0]),
                &hash_validator(&validators[1]),
            ),
            &hash_node(
                &hash_validator(&validators[2]),
                &hash_validator(&validators[3]),
            ),
        )
    }
}

#[cfg(test)]
mod test_utils {
    use super::*;

    pub fn scalar_from_be_hex(s: &str) -> bls12_381::Scalar {
        let bytes = hex::decode(s).expect("invalid hex");
        let mut arr = [0u8; 32];
        arr.copy_from_slice(&bytes);
        arr.reverse(); // Scalar::from_bytes expects little-endian
        bls12_381::Scalar::from_bytes(&arr).expect("scalar not in field")
    }

    pub fn get_public_key_from_secret(sk: bls12_381::Scalar) -> G1Affine {
        G1Affine::from(G1Projective::generator() * sk)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::test_utils::get_public_key_from_secret;

    const PK1_HEX: &str = "11ae567e08314b059b33771ff6fe3dc25703b9f01e2794b0695da79285ac069e7d7471737dddde2e6fa36bc2557bc01111bcd774c0251f8c4a146287983692c1af1e869075ae8ac60107714c3c6fe91b0aa730f6f2d070026cca642515691802";
    const PK2_HEX: &str = "1789b8b9a7c411eff217e3699073cf7184daa825aed5faeed58d37ebf7b4c77b95e477a9d0901a62ee09260c86081a87097cb335c95ac0d10df57a0114cfcb34142a1ce162466b5a628c074882d1837420124d7a317813f912dff063f5890e9d";
    const PK3_HEX: &str = "01bcc815aee18ec84ae166176bc7ccccf42cbb44c4984f3b0aa4f8eaba18dc6bc40aaac8cc1ed4c80251a2598d4be63b0214fedd1f4c73548782e05a138863355a1a676c8d8e68a014c1120522924bc1a8cdc57a194cf918b6e6e2bad2e1da63";
    const SIG_BYTES: &str = "13548fe35a1be72e79719ce4631989efdaa29a88873d949201e510a6117b8c7b91dea1d5e6070f9996a31cef00afb0f404254a8979c8c7617710dfc324c989460be9c874c2584dbf114e04bcb1a9dd5894c788990948d1909f40cc692b9e160406be23ffd1d1ece752ba9c6f9e7ee7dce91c312eb3580ca4c6a72b5a4c9911ca2287a06ee6e4953a2f2cc106019c06fc19c09662244295626f93b8dcd20c09c1fd520d7245b659080e1defe1cf0c33bbb054ba462b244d3c0f23b1d398d5db7e";

    const AGGREGATE_SIG: &str = "055273f06e2a177237e5e949abe52c0194ed2a84a8bdb1b04986e721f1206c025f599458c7be4b7f9a133cb291f52900102f103b08360bb6142102804cb972f388a21142167d041b2d81dccfe29ef3ad7a1c1b7a5b725938167ebd089ee30e300fdadf227f40a0e9f6b168555871d5a992f0fbf8d9c65bde4f140f0552f38d009334987b6fcd56a051a63718716ff4220eb925350b4cb88d9d3abb10237f3da95c6d619ec19fa48bd8da6ec4469d8f71395e4681c9540c97a0afc5af8e457413";
    const ATTESTATION: &str = "0000000000030572cbea904d67468808c8eb50a9450c9721db309128012543902d0ac358a62ae28f75bb8f1c7c42c39a8c5529bf0f4e166a9d8cabc673a322fda673779d8e3822ba3ecb8670e461f73bb9021d5fd76a4c56d9d4cd16bd1bba86881979749d2800000000000000010c9b60d5afcbd5663a8a44b7c5a02f19e9a77ab0a35bd65809bb5c67ec582c897feb04decc694b13e08587f3ff9b5b60143be6d078c2b79a7d4f1d1b21486a030ec93f56aa54e1de880db5a66dd833a652a95bee27c824084006cb5644cbd43f000000000000000310e7791fb972fe014159aa33a98622da3cdc98ff707965e536d8636b5fcc5ac7a91a8c46e59a00dca575af0f18fb13dc16ba437edcc6551e30c10512367494bfb6b01cc6681e8a4c3cd2501832ab5c4abc40b4578b85cbaffbf0bcd70d67c6e20000000000000004000000017e5866de1dc819a4346a48c956a12e000bf26f3a6f61e0549e01d720f703d0550000000306055273f06e2a177237e5e949abe52c0194ed2a84a8bdb1b04986e721f1206c025f599458c7be4b7f9a133cb291f52900102f103b08360bb6142102804cb972f388a21142167d041b2d81dccfe29ef3ad7a1c1b7a5b725938167ebd089ee30e300fdadf227f40a0e9f6b168555871d5a992f0fbf8d9c65bde4f140f0552f38d009334987b6fcd56a051a63718716ff4220eb925350b4cb88d9d3abb10237f3da95c6d619ec19fa48bd8da6ec4469d8f71395e4681c9540c97a0afc5af8e457413";

    /// Test that the utility function for getting public keys from secret keys matches the
    /// logic in BLST.t.sol in icm-services.
    #[test]
    fn test_get_pk_from_sk() {
        let secret_keys = [
            test_utils::scalar_from_be_hex(
                "561fee6432efc4b109c69a2c4e1bd9cc474ce78ed7572fc4a220783030cdbcac",
            ),
            test_utils::scalar_from_be_hex(
                "18a52533cf7debd4b90b3d35d7c2a3783b524f69a03c4d3fa64cccf635338808",
            ),
            test_utils::scalar_from_be_hex(
                "4981010927d0c2233d810328805d451b8119f2c6ff50fc36f56ed5a56aa8690c",
            ),
        ]
        .into_iter()
        .map(get_public_key_from_secret)
        .collect::<Vec<_>>();

        let public_keys = [PK1_HEX, PK2_HEX, PK3_HEX]
            .into_iter()
            .map(|pk| {
                G1Affine::from_uncompressed(
                    &hex::decode(pk)
                        .expect("Test failed")
                        .try_into()
                        .expect("Test failed"),
                )
                .into_option()
                .expect("Test failed")
            })
            .collect::<Vec<_>>();

        for i in 0..3 {
            assert_eq!(secret_keys[i], public_keys[i]);
        }
    }

    /// Mirrors the `testAggregateSigAgainstAvalancheGo` test in BLST.t.sol in icm-services. Ensures
    /// that the logic here behaves identically.
    #[test]
    fn test_aggregate_signature() {
        let pk1 = G1Affine::from_uncompressed(
            &hex::decode(PK1_HEX)
                .expect("Test failed")
                .try_into()
                .expect("Test failed"),
        )
        .into_option()
        .expect("Test failed");
        let pk2 = G1Affine::from_uncompressed(
            &hex::decode(PK2_HEX)
                .expect("Test failed")
                .try_into()
                .expect("Test failed"),
        )
        .into_option()
        .expect("Test failed");
        let pk3 = G1Affine::from_uncompressed(
            &hex::decode(PK3_HEX)
                .expect("Test failed")
                .try_into()
                .expect("Test failed"),
        )
        .into_option()
        .expect("Test failed");
        let signature = G2Affine::from_uncompressed(
            &hex::decode(SIG_BYTES)
                .expect("Test failed")
                .try_into()
                .expect("Test failed"),
        )
        .into_option()
        .expect("Test failed");

        let agg_pk = G1Affine::from(
            [pk1, pk2, pk3]
                .iter()
                .fold(G1Projective::identity(), |acc, v| {
                    acc + G1Projective::from(v)
                }),
        );

        let message = b"TestValidAggregation local signer";
        let sig_valid = verify_signature(message, &agg_pk, signature);
        assert!(sig_valid);

        let message = b"sudo rm -rf /";
        let sig_valid = verify_signature(message, &agg_pk, signature);
        assert!(!sig_valid);
    }

    /// Use an attestation generated by the icm-services and check that we parse it correctly.
    #[test]
    fn test_parse_merkle_attestation() {
        let validators = [2u64, 3, 4, 5]
            .iter()
            .enumerate()
            .map(|(ix, i)| Validator {
                public_key: get_public_key_from_secret(bls12_381::Scalar::from(*i)),
                weight: ix as u64 + 1,
            })
            .collect::<Vec<_>>();
        let signers = vec![
            validators[0].clone(),
            validators[2].clone(),
            validators[3].clone(),
        ];
        let mut flags = BitVec::<u8>::new();
        flags.push(false);
        flags.push(true);
        flags.push(true);
        let proof = vec![hash_validator(&validators[1])];

        let attestation =
            parse_attestation(hex::decode(ATTESTATION).expect("Test failed").as_slice())
                .expect("Test failed");

        assert_eq!(attestation.signers, signers);
        assert_eq!(attestation.flags, flags);
        assert_eq!(attestation.proof, proof);
        assert_eq!(
            attestation.aggregate_sig,
            G2Affine::from_uncompressed(
                &hex::decode(AGGREGATE_SIG)
                    .expect("Test failed")
                    .try_into()
                    .expect("Test failed")
            )
            .into_option()
            .expect("Test failed")
        )
    }

    /// Test that we can verify a merkle attestation generated by the icm-services.
    #[test]
    fn test_verify_merkle_attestation() {
        let attestation = parse_attestation(
            hex::decode(test_fixtures::ATTESTATION_HEX)
                .expect("Test failed")
                .as_slice(),
        )
        .expect("Test failed");
        let signed_data = hex::decode(test_fixtures::SIGNED_DATA_HEX).expect("Test failed");
        let expected = test_fixtures::expected_root();
        let signed_data_hash: [u8; 32] = Sha256::digest(&signed_data).into();

        verify(
            &attestation,
            &signed_data,
            &expected,
            &signed_data_hash,
            test_fixtures::SIGNING_WEIGHT,
        )
        .expect("Test failed");
    }
}
