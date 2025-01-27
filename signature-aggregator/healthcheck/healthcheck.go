package healthcheck

import (
	"context"
	"net/http"

	"github.com/alexliesenfeld/health"
)

func HandleHealthCheckRequest(checkFunc func(context.Context) error) {
	healthChecker := health.NewChecker(
		health.WithCheck(health.Check{
			Name:  "signature-aggregator-health",
			Check: checkFunc,
		}),
	)

	http.Handle("/health", health.NewHandler(healthChecker))
}
