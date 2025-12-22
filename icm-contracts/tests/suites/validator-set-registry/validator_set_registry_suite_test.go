// Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validator_set_registry_test

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	validatorSetRegistry "github.com/ava-labs/icm-services/icm-contracts/tests/flows/validator-set-registry"
	"github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	warpGenesisTemplateFile    = "./icm-contracts/tests/utils/warp-genesis-template.json"
	validatorSetRegistryLabel  = "validator-set-registry"
)

var (
	LocalAvalancheNetworkInstance *network.LocalNetwork
	LocalEthereumNetworkInstance  *network.LocalEthereumNetwork
	e2eFlags                      *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestValidatorSetRegistry(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Validator Set Registry e2e test")
}

var _ = ginkgo.BeforeSuite(func() {
	// Create the local Avalanche network instance
	// We don't need Teleporter for this test, just the P-chain validator set
	ctx, cancel := context.WithTimeout(context.Background(), 240*time.Second)
	defer cancel()

	LocalAvalancheNetworkInstance = network.NewLocalNetwork(
		ctx,
		"validator-set-registry-test-local-network",
		warpGenesisTemplateFile,
		[]network.L1Spec{
			{
				Name:       "L1",
				EVMChainID: 12345,
				// Empty Teleporter values - not needed for validator set registry test
				TeleporterContractAddress:  common.Address{},
				TeleporterDeployedBytecode: "",
				TeleporterDeployerAddress:  common.Address{},
				NodeCount:                  2,
			},
		},
		2,
		2,
		e2eFlags,
	)
	log.Info("Started local Avalanche network", "networkID", LocalAvalancheNetworkInstance.GetNetworkID())

	// Connect to local Ethereum network (must be started separately)
	LocalEthereumNetworkInstance = network.NewLocalEthereumNetwork(context.Background())
	log.Info("Connected to local Ethereum network", "chainID", LocalEthereumNetworkInstance.ChainID)

	log.Info("Set up ginkgo before suite")
})

var _ = ginkgo.AfterSuite(func() {
	if LocalAvalancheNetworkInstance != nil {
		LocalAvalancheNetworkInstance.TearDownNetwork()
		LocalAvalancheNetworkInstance = nil
	}
	if LocalEthereumNetworkInstance != nil {
		LocalEthereumNetworkInstance.TearDownNetwork()
		LocalEthereumNetworkInstance = nil
	}
})

var _ = ginkgo.Describe("[Validator Set Registry integration tests]", func() {
	ginkgo.It("Deploy and test validator set registry update",
		ginkgo.Label(validatorSetRegistryLabel),
		func() {
			validatorSetRegistry.ValidatorSetRegistryTest(
				LocalAvalancheNetworkInstance,
				LocalEthereumNetworkInstance,
			)
		})
})

