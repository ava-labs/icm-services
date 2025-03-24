// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package signer

import (
	"fmt"
	"math/big"

	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ethereum/go-ethereum/common"
)

type Signer interface {
	SignTx(tx *types.Transaction, evmChainID *big.Int) (*types.Transaction, error)
	Address() common.Address
}

func NewSigners(destinationBlockchain *config.DestinationBlockchain) ([]Signer, error) {
	txSigners, err := NewTxSigners(destinationBlockchain.AccountPrivateKeys)
	if err != nil {
		return nil, err
	}
	kmsSigners, err := NewKMSSigners(destinationBlockchain.KMSAWSRegions, destinationBlockchain.KMSKeyIDs)
	if err != nil {
		return nil, err
	}
	return append(txSigners, kmsSigners...), nil
}

func NewTxSigners(pks []string) ([]Signer, error) {
	var signers []Signer
	for _, pk := range pks {
		signer, err := NewTxSigner(pk)
		if err != nil {
			return signers, err
		}
		signers = append(signers, signer)
	}
	return signers, nil
}

func NewKMSSigners(awsRegions []string, keyIDs []string) ([]Signer, error) {
	if len(keyIDs) != len(awsRegions) {
		return nil, fmt.Errorf("length of key IDs %d not equal to length of awsRegions %d", len(keyIDs), len(awsRegions))
	}

	var signers []Signer
	for i := range keyIDs {
		signer, err := NewKMSSigner(awsRegions[i], keyIDs[i])
		if err != nil {
			return signers, err
		}
		signers = append(signers, signer)
	}
	return signers, nil
}
