package validator_manager_test

import (
	"context"
	"flag"
	"os"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	validatorManagerFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/validator-manager"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	warpGenesisTemplateFile = "./tests/utils/warp-genesis-template.json"
	validatorManagerLabel   = "ValidatorManager"
)

var (
	localNetworkInstance *localnetwork.LocalAvalancheNetwork
	e2eFlags             *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestValidatorManager(t *testing.T) {
	if os.Getenv("RUN_E2E") == "" {
		t.Skip("Environment variable RUN_E2E not set; skipping E2E tests")
	}

	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Validator Manager e2e test")
}

// Define the before and after suite functions.
var _ = ginkgo.BeforeEach(func(ctx context.Context) {
	// Create the local network instance
	ctx, cancel := context.WithTimeout(ctx, 240*time.Second)
	defer cancel()
	localNetworkInstance = localnetwork.NewLocalAvalancheNetwork(
		ctx,
		"validator-manager-test-local-network",
		warpGenesisTemplateFile,
		[]localnetwork.L1Spec{
			{
				Name:                         "A",
				EVMChainID:                   12345,
				NodeCount:                    2,
				RequirePrimaryNetworkSigners: true,
			},
			{
				Name:                         "B",
				EVMChainID:                   54321,
				NodeCount:                    2,
				RequirePrimaryNetworkSigners: true,
			},
		},
		2,
		2,
		e2eFlags,
	)
	log.Info("Started local network")
})

var _ = ginkgo.AfterEach(func() {
	localNetworkInstance.TearDownNetwork()
	localNetworkInstance = nil
})

var _ = ginkgo.Describe("[Validator manager integration tests]", func() {
	// Validator Manager tests
	ginkgo.It("Native token staking manager",
		ginkgo.Label(validatorManagerLabel),
		func(ctx context.Context) {
			validatorManagerFlows.NativeTokenStakingManager(ctx, localNetworkInstance)
		})
	ginkgo.It("ERC20 token staking manager",
		ginkgo.Label(validatorManagerLabel),
		func(ctx context.Context) {
			validatorManagerFlows.ERC20TokenStakingManager(ctx, localNetworkInstance)
		})
	ginkgo.It("PoA migration to PoS",
		ginkgo.Label(validatorManagerLabel),
		func(ctx context.Context) {
			validatorManagerFlows.PoAMigrationToPoS(ctx, localNetworkInstance)
		})
	ginkgo.It("Delegate disable validator",
		ginkgo.Label(validatorManagerLabel),
		func(ctx context.Context) {
			validatorManagerFlows.RemoveDelegatorInactiveValidator(ctx, localNetworkInstance)
		})
})
