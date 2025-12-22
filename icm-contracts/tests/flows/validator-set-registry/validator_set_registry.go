// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validator_set_registry

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
	simplevalidatorsetregistry "github.com/ava-labs/icm-services/abi-bindings/go/SimpleValidatorSetRegistry"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/vms/evm"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/crypto"
	"github.com/ava-labs/libevm/log"
	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	. "github.com/onsi/gomega"
)

// ValidatorSetRegistryTest tests the deployment and validator set registration/update
// on the SimpleValidatorSetRegistry contract on an external Ethereum network.
//
// This test:
// 1. Deploys a SimpleValidatorSetRegistry contract on the Ethereum network
// 2. Registers a validator set (mock data, no signature verification)
// 3. Updates the validator set via direct contract call
// 4. Updates the validator set via ExternalEVMDestinationClient.SendTx
// 5. Verifies all updates were applied correctly
func ValidatorSetRegistryTest(
	avalancheNetwork *network.LocalNetwork,
	ethereumNetwork *network.LocalEthereumNetwork,
) {
	ctx := context.Background()
	_, ethereumFundedKey := ethereumNetwork.GetFundedAccountInfo()

	// Get L1 info from Avalanche network
	l1Infos := avalancheNetwork.GetL1Infos()
	Expect(len(l1Infos)).Should(BeNumerically(">", 0), "No L1s found in Avalanche network")
	l1Info := l1Infos[0]

	// Create transaction options
	opts, err := bind.NewKeyedTransactorWithChainID(ethereumFundedKey, ethereumNetwork.ChainID)
	Expect(err).Should(BeNil())

	// Deploy SimpleValidatorSetRegistry contract on the Ethereum network
	log.Info("Deploying SimpleValidatorSetRegistry to Ethereum network",
		"avalancheNetworkID", avalancheNetwork.GetNetworkID(),
		"l1BlockchainID", l1Info.BlockchainID.String(),
	)

	registryAddress, tx, registry, err := simplevalidatorsetregistry.DeploySimplevalidatorsetregistry(
		opts,
		ethereumNetwork.Client,
		avalancheNetwork.GetNetworkID(),
		l1Info.BlockchainID,
	)
	Expect(err).Should(BeNil())

	// Wait for deployment transaction
	receipt := utils.WaitForTransactionSuccessWithClient(ctx, ethereumNetwork.Client, tx.Hash())
	log.Info("Deployed SimpleValidatorSetRegistry contract",
		"address", registryAddress.Hex(),
		"txHash", receipt.TxHash.Hex(),
	)

	// Verify the contract was deployed correctly
	Expect(registryAddress).ShouldNot(BeZero())
	Expect(registry).ShouldNot(BeNil())

	// Verify contract state
	callOpts := &bind.CallOpts{Context: ctx}

	networkID, err := registry.AvalancheNetworkID(callOpts)
	Expect(err).Should(BeNil())
	Expect(networkID).Should(Equal(avalancheNetwork.GetNetworkID()))

	blockchainID, err := registry.AvalancheBlockChainID(callOpts)
	Expect(err).Should(BeNil())
	Expect(ids.ID(blockchainID)).Should(Equal(l1Info.BlockchainID))

	nextValidatorSetID, err := registry.NextValidatorSetID(callOpts)
	Expect(err).Should(BeNil())
	Expect(nextValidatorSetID).Should(Equal(uint32(0)))

	log.Info("Contract state verified",
		"networkID", networkID,
		"blockchainID", ids.ID(blockchainID).String(),
		"nextValidatorSetID", nextValidatorSetID,
	)

	// Step 1: Register initial validator set
	log.Info("Registering initial validator set")

	initialValidators := []simplevalidatorsetregistry.SimpleValidatorSetRegistryValidator{
		{
			BlsPublicKey: make([]byte, 48), // Mock 48-byte BLS public key
			Weight:       1000,
		},
		{
			BlsPublicKey: make([]byte, 48),
			Weight:       2000,
		},
	}
	// Set different keys to satisfy sorting requirement
	initialValidators[0].BlsPublicKey[0] = 0x01
	initialValidators[1].BlsPublicKey[0] = 0x02

	initialPChainHeight := uint64(100)

	tx, err = registry.RegisterValidatorSet(opts, initialValidators, initialPChainHeight)
	Expect(err).Should(BeNil())

	receipt = utils.WaitForTransactionSuccessWithClient(ctx, ethereumNetwork.Client, tx.Hash())
	log.Info("Registered initial validator set",
		"txHash", receipt.TxHash.Hex(),
		"gasUsed", receipt.GasUsed,
	)

	// Verify the validator set was registered
	nextValidatorSetID, err = registry.NextValidatorSetID(callOpts)
	Expect(err).Should(BeNil())
	Expect(nextValidatorSetID).Should(Equal(uint32(1)))

	// Get the registered validator set
	validatorSet, err := registry.GetValidatorSet(callOpts, big.NewInt(0))
	Expect(err).Should(BeNil())
	Expect(validatorSet.TotalWeight).Should(Equal(uint64(3000)))
	Expect(validatorSet.PChainHeight).Should(Equal(initialPChainHeight))
	Expect(len(validatorSet.Validators)).Should(Equal(2))

	log.Info("Verified registered validator set",
		"validatorSetID", 0,
		"totalWeight", validatorSet.TotalWeight,
		"pChainHeight", validatorSet.PChainHeight,
		"validatorCount", len(validatorSet.Validators),
	)

	// Step 2: Update the validator set via direct contract call
	log.Info("Updating validator set via direct contract call")

	updatedValidators := []simplevalidatorsetregistry.SimpleValidatorSetRegistryValidator{
		{
			BlsPublicKey: make([]byte, 48),
			Weight:       1500,
		},
		{
			BlsPublicKey: make([]byte, 48),
			Weight:       2500,
		},
		{
			BlsPublicKey: make([]byte, 48),
			Weight:       1000,
		},
	}
	// Set different keys
	updatedValidators[0].BlsPublicKey[0] = 0x01
	updatedValidators[1].BlsPublicKey[0] = 0x02
	updatedValidators[2].BlsPublicKey[0] = 0x03

	updatedPChainHeight := uint64(200) // Must be greater than initial

	tx, err = registry.UpdateValidatorSet(opts, big.NewInt(0), updatedValidators, updatedPChainHeight)
	Expect(err).Should(BeNil())

	receipt = utils.WaitForTransactionSuccessWithClient(ctx, ethereumNetwork.Client, tx.Hash())
	log.Info("Updated validator set via direct call",
		"txHash", receipt.TxHash.Hex(),
		"gasUsed", receipt.GasUsed,
	)

	// Verify the validator set was updated
	validatorSet, err = registry.GetValidatorSet(callOpts, big.NewInt(0))
	Expect(err).Should(BeNil())
	Expect(validatorSet.TotalWeight).Should(Equal(uint64(5000))) // 1500 + 2500 + 1000
	Expect(validatorSet.PChainHeight).Should(Equal(updatedPChainHeight))
	Expect(len(validatorSet.Validators)).Should(Equal(3))

	log.Info("Verified updated validator set",
		"validatorSetID", 0,
		"totalWeight", validatorSet.TotalWeight,
		"pChainHeight", validatorSet.PChainHeight,
		"validatorCount", len(validatorSet.Validators),
	)

	// Get current validator set (should be the same as validator set 0)
	currentValidatorSet, err := registry.GetCurrentValidatorSet(callOpts)
	Expect(err).Should(BeNil())
	Expect(currentValidatorSet.TotalWeight).Should(Equal(validatorSet.TotalWeight))
	Expect(currentValidatorSet.PChainHeight).Should(Equal(validatorSet.PChainHeight))

	log.Info("Direct contract call test completed successfully!",
		"registryAddress", registryAddress.Hex(),
		"finalValidatorCount", len(currentValidatorSet.Validators),
		"finalTotalWeight", currentValidatorSet.TotalWeight,
		"finalPChainHeight", currentValidatorSet.PChainHeight,
	)

	// =========================================================================
	// Step 3: Test ExternalEVMDestinationClient.SendTx
	// This simulates how the ValidatorSetUpdater would call the contract
	// =========================================================================
	log.Info("Testing ExternalEVMDestinationClient.SendTx for validator set update")

	// Get the private key hex (strip "0x" if present)
	privateKeyBytes := crypto.FromECDSA(ethereumFundedKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	// Create ExternalEVMDestinationClient
	logger := logging.NoLog{}

	externalClient, err := evm.NewExternalEVMDestinationClient(
		logger,
		ethereumNetwork.ChainID.String(), // chainID
		ethereumNetwork.BaseURL,          // rpcEndpointURL
		registryAddress,                  // registryAddress
		[]string{privateKeyHex},          // privateKeyHexes
		8000000,                          // blockGasLimit
		nil,                              // maxBaseFee (use default)
		big.NewInt(1000000000),           // suggestedPriorityFeeBuffer (1 gwei)
		big.NewInt(100000000000),         // maxPriorityFeePerGas (100 gwei)
		60,                               // txInclusionTimeoutSeconds
	)
	Expect(err).Should(BeNil())

	log.Info("Created ExternalEVMDestinationClient",
		"chainID", ethereumNetwork.ChainID.String(),
		"senderAddresses", externalClient.SenderAddresses(),
	)

	// Prepare new validator set for update via SendTx
	newValidators := []simplevalidatorsetregistry.SimpleValidatorSetRegistryValidator{
		{
			BlsPublicKey: make([]byte, 48),
			Weight:       2000,
		},
		{
			BlsPublicKey: make([]byte, 48),
			Weight:       3000,
		},
	}
	newValidators[0].BlsPublicKey[0] = 0x01
	newValidators[1].BlsPublicKey[0] = 0x02

	newPChainHeight := uint64(300) // Must be greater than previous (200)

	// ABI-encode the updateValidatorSet call
	parsedABI, err := abi.JSON(strings.NewReader(simplevalidatorsetregistry.SimplevalidatorsetregistryABI))
	Expect(err).Should(BeNil())

	callData, err := parsedABI.Pack(
		"updateValidatorSet",
		big.NewInt(0),   // validatorSetID
		newValidators,   // validators
		newPChainHeight, // pChainHeight
	)
	Expect(err).Should(BeNil())

	log.Info("Encoded updateValidatorSet callData",
		"callDataLength", len(callData),
		"newPChainHeight", newPChainHeight,
		"newTotalWeight", 5000,
	)

	// Send transaction via ExternalEVMDestinationClient.SendTx
	txReceipt, err := externalClient.SendTx(
		nil,                       // signedMessage (not needed for SimpleValidatorSetRegistry)
		set.Set[common.Address]{}, // deliverers (empty = any sender)
		registryAddress.Hex(),     // toAddress
		500000,                    // gasLimit
		callData,                  // pre-encoded callData
	)
	Expect(err).Should(BeNil())
	Expect(txReceipt).ShouldNot(BeNil())
	Expect(txReceipt.Status).Should(Equal(uint64(1)), "Transaction should succeed")

	log.Info("ExternalEVMDestinationClient.SendTx succeeded",
		"txHash", txReceipt.TxHash.Hex(),
		"gasUsed", txReceipt.GasUsed,
		"status", txReceipt.Status,
	)

	// Verify the validator set was updated
	validatorSet, err = registry.GetValidatorSet(callOpts, big.NewInt(0))
	Expect(err).Should(BeNil())
	Expect(validatorSet.TotalWeight).Should(Equal(uint64(5000))) // 2000 + 3000
	Expect(validatorSet.PChainHeight).Should(Equal(newPChainHeight))
	Expect(len(validatorSet.Validators)).Should(Equal(2))

	log.Info("Verified validator set updated via ExternalEVMDestinationClient.SendTx",
		"validatorSetID", 0,
		"totalWeight", validatorSet.TotalWeight,
		"pChainHeight", validatorSet.PChainHeight,
		"validatorCount", len(validatorSet.Validators),
	)

	log.Info("All tests completed successfully!",
		"registryAddress", registryAddress.Hex(),
	)
}

