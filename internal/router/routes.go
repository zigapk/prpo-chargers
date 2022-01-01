package router

import (
	"net/http"

	"github.com/zigapk/prpo-chargers/internal/handle"
)

func apiRoutes() []Route {
	return []Route{
		{
			Name:           "chargers",
			Path:           "/chargers",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.GetChargersHandle),
		},
		{
			Name: "liveness_probe",
			Path: "/health/liveness",
			GET:  http.HandlerFunc(handle.LivenessHandle),
		},
		{
			Name: "readiness_probe",
			Path: "/health/readiness",
			GET:  http.HandlerFunc(handle.ReadinessHandle),
		},

		{
			Name: "metrics",
			Path: "/metrics",
			GET:  http.HandlerFunc(handle.MetricsHandle),
		},
	}
}
