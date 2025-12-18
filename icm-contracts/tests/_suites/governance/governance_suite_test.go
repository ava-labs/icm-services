package governance_test

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	suites "github.com/ava-labs/icm-services/icm-contracts/tests/_suites"
	governanceFlows "github.com/ava-labs/icm-services/icm-contracts/tests/flows/governance"
	localnetwork "github.com/ava-labs/icm-services/icm-contracts/tests/network"
	"github.com/ava-labs/icm-services/icm-contracts/tests/utils"
	"github.com/ava-labs/icm-services/log"
	"github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	validatorSetSigLabel = "ValidatorSetSig"
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

func TestGovernance(t *testing.T) {
	RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Governance e2e test")
}

// Define the before and after suite functions.
var _ = ginkgo.BeforeSuite(func(ctx context.Context) {
	localNetworkInstance = suites.StartDefaultNetwork(ctx, "governance-test-local-network", e2eFlags)
	log.Info("Started local network")
})

var _ = ginkgo.AfterSuite(func() {
	localNetworkInstance.TearDownNetwork()
	localNetworkInstance = nil
})

var _ = ginkgo.Describe("[Governance integration tests]", func() {
	// Governance tests
	ginkgo.It("Deliver ValidatorSetSig signed message",
		ginkgo.Label(validatorSetSigLabel),
		func(ctx context.Context) {
			governanceFlows.ValidatorSetSig(ctx, localNetworkInstance)
		})
})
