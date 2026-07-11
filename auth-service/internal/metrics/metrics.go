package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_http_request_total",
			Help: "Total number of requests processed by the API",
		},
		[]string{"path", "status"},
	)

	HttpRequestErrorTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_http_request_error_total",
			Help: "Total number of error requests",
		},
		[]string{"path", "status"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests",
			Buckets: []float64{
				0.005,
				0.01,
				0.025,
				0.05,
				0.1,
				0.25,
				0.5,
				1,
				2.5,
				5,
				10,
			},
		},
		[]string{"method", "route", "status"},
	)
)

func init() {
	prometheus.MustRegister(
		HttpRequestTotal,
		HttpRequestErrorTotal,
		HTTPRequestDuration,
	)
}
