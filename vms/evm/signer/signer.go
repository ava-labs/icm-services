// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package signer

import (
	"math/big"

	"github.com/ava-labs/icm-services/relayer/config"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
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
	kmsSigners, err := NewKMSSigners(destinationBlockchain.KMSKeys)
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

func NewKMSSigners(kmsKeys []config.KMSKey) ([]Signer, error) {
	var signers []Signer
	for i := range kmsKeys {
		signer, err := NewKMSSigner(kmsKeys[i].AWSRegion, kmsKeys[i].KeyID)
		if err != nil {
			return signers, err
		}
		signers = append(signers, signer)
	}
	return signers, nil
}
