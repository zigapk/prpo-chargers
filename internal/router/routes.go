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
			Name:           "chargers",
			Path:           "/chargers/{id:\\d+}/reservations",
			AuthorizedOnly: true,
			GET:            http.HandlerFunc(handle.GetReservationsHandle),
		},
		{
			Name:           "reservations",
			Path:           "/reservations",
			AuthorizedOnly: true,
			POST:           http.HandlerFunc(handle.CreateReservationRequest),
		},
		{
			Name:           "reservations",
			Path:           "/reservations/{id:\\d+}",
			AuthorizedOnly: true,
			PUT:            http.HandlerFunc(handle.UpdateReservationRequest),
			DELETE:         http.HandlerFunc(handle.DeleteReservationRequest),
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
