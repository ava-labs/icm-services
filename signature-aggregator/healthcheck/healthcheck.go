package healthcheck

import (
	"context"
	"net/http"

	"github.com/alexliesenfeld/health"
)

const HealthAPIPath = "/health"

func HandleHealthCheckRequest(mux *http.ServeMux, checkFunc func(context.Context) error) {
	healthChecker := health.NewChecker(
		health.WithCheck(health.Check{
			Name:  "signature-aggregator-health",
			Check: checkFunc,
		}),
	)

	mux.Handle(HealthAPIPath, health.NewHandler(healthChecker))
}
