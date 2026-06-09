package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "authorpedro_requests_total",
		Help: "Total number of agent requests",
	})
	RequestsInFlight = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "authorpedro_requests_in_flight",
		Help: "Number of requests in flight",
	})
	RequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "authorpedro_request_duration_seconds",
		Help:    "Request duration in seconds",
		Buckets: prometheus.DefBuckets,
	})
	TTFT = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "authorpedro_ttft_seconds",
		Help:    "Time to first token in seconds",
		Buckets: prometheus.DefBuckets,
	})
	ErrorsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Name: "authorpedro_errors_total",
		Help: "Total number of errors",
	})
	TokensGenerated = promauto.NewCounter(prometheus.CounterOpts{
		Name: "authorpedro_tokens_generated_total",
		Help: "Total tokens generated",
	})
	Iterations = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "authorpedro_iterations",
		Help:    "Number of iterations per request",
		Buckets: prometheus.DefBuckets,
	})
)

func RecordRequest(ttft, duration float64, iterations int, tokens float64, isError bool) {
	RequestsTotal.Inc()
	if isError {
		ErrorsTotal.Inc()
	}
	TTFT.Observe(ttft)
	RequestDuration.Observe(duration)
	Iterations.Observe(float64(iterations))
	TokensGenerated.Add(tokens)
}

func IncInFlight(delta float64) {
	RequestsInFlight.Add(delta)
}
