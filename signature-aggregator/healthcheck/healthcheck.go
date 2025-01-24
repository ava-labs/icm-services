package healthcheck

import (
	"context"
	"fmt"
	"math/big"
	"net/http"

	"github.com/alexliesenfeld/health"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/coreth/precompile/contracts/warp"
	"github.com/ava-labs/icm-services/peers"
	"github.com/ava-labs/icm-services/signature-aggregator/aggregator"
	"github.com/ava-labs/icm-services/utils"
)

func HandleHealthCheckRequest(network peers.AppRequestNetwork) {
	healthChecker := health.NewChecker(
		health.WithCheck(health.Check{
			Name: "signature-aggregator-health",
			Check: func(context.Context) error {
				connectedValidators, err := network.ConnectToCanonicalValidators(constants.PrimaryNetworkID)
				if err != nil {
					return fmt.Errorf("Failed to connect to primary network validators: %w", err)
				}
				if !utils.CheckStakeWeightExceedsThreshold(
					big.NewInt(0).SetUint64(connectedValidators.ConnectedWeight),
					connectedValidators.TotalValidatorWeight,
					warp.WarpDefaultQuorumNumerator,
				) {
					return aggregator.ErrNotEnoughConnectedStake
				}
				return nil
			},
		}),
	)

	http.Handle("/health", health.NewHandler(healthChecker))
}
