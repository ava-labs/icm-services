package validator_manager_test

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	suites "github.com/ava-labs/icm-services/icm-contracts/tests/_suites"
	validatorManagerFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/validator-manager"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	validatorManagerLabel = "ValidatorManager"
)

var (
	localNetworkInstance *localnetwork.LocalNetwork
	teleporterInfo       utils.TeleporterTestInfo
	e2eFlags             *e2e.FlagVars
)

func TestMain(m *testing.M) {
	e2eFlags = e2e.RegisterFlags()
	flag.Parse()
	os.Exit(m.Run())
}

func TestValidatorManager(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Validator Manager e2e test")
}

// Define the before and after suite functions.
var _ = ginkgo.BeforeEach(func(ctx context.Context) {
	localNetworkInstance = suites.StartDefaultNetwork(ctx, "validator-manager-test-local-network", e2eFlags)
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
