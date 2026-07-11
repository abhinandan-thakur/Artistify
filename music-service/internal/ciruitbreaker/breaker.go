package circuitbreaker

import (
	"time"

	"github.com/sony/gobreaker"
)

func New(name string) *gobreaker.CircuitBreaker {
	settings := gobreaker.Settings{
		Name: name,

		// Open circuit after repeated failures
		MaxRequests: 3,

		// Reset timeout before trying again
		Timeout: 10 * time.Second,

		// Sliding window
		Interval: 30 * time.Second,

		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Don't trip too early
			if counts.Requests < 5 {
				return false
			}

			failureRatio := float64(counts.TotalFailures) /
				float64(counts.Requests)

			return failureRatio >= 0.6
		},

		OnStateChange: func(name string,
			from gobreaker.State,
			to gobreaker.State,
		) {
			println(
				"Circuit breaker:",
				name,
				from.String(),
				"->",
				to.String(),
			)
		},
	}

	return gobreaker.NewCircuitBreaker(settings)
}