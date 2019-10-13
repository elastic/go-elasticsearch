package estransport

import (
	"sync"
)

// Metrics represents the client metrics.
//
type Metrics struct {
	sync.RWMutex

	NumRequests int `json:"num_requests"`
	NumFailures int `json:"num_failures"`

	// TODO(karmi): Serialize connection URL as plain string
	Pool []*Connection `json:"pool"`
	Dead []*Connection `json:"dead"`
}

// EnableMetrics makes the metrics enabled.
//
func EnableMetrics() {
	metrics = &Metrics{}
}

// IsMetricsEnabled returns true when metrics are enabled.
//
func IsMetricsEnabled() bool {
	return metrics != nil
}

// MetricFunc returns object for the "expvar" package.
//
func MetricFunc() interface{} {
	metrics.RLock()
	metrics.RUnlock()
	return metrics
}
