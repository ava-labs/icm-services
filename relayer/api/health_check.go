package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/alexliesenfeld/health"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"go.uber.org/atomic"
)

const HealthAPIPath = "/health"

func HandleHealthCheck(
	logger logging.Logger,
	relayerHealth map[ids.ID]*atomic.Bool,
	networkHealth func(context.Context) error,
) {
	http.Handle(HealthAPIPath, healthCheckHandler(logger, relayerHealth, networkHealth))
}

func healthCheckHandler(
	logger logging.Logger,
	relayerHealth map[ids.ID]*atomic.Bool,
	networkHealth func(context.Context) error,
) http.Handler {
	return health.NewHandler(health.NewChecker(
		health.WithCheck(health.Check{
			Name: "relayers-all",
			Check: func(context.Context) error {
				// Store the IDs as the cb58 encoding
				var unhealthyRelayers []string
				for id, health := range relayerHealth {
					if !health.Load() {
						unhealthyRelayers = append(unhealthyRelayers, id.String())
					}
				}

				if len(unhealthyRelayers) > 0 {
					return fmt.Errorf("relayers are unhealthy for blockchains %v", unhealthyRelayers)
				}
				return nil
			},
		}),
		health.WithCheck(health.Check{
			Name:  "network-all",
			Check: networkHealth,
		}),
	))
}
