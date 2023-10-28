package server

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	totalRequestsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_request_count",
			Help: "Num. of requests handled by handlers",
		},
	)
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duration of HTTP requests.",
		},
		[]string{"method"},
	)
	badRequestsCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "bad_request_count",
			Help: "Num. of bad requests handled by handlers",
		},
	)
)
