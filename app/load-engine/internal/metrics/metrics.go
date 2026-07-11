package metrics

import (
	"fmt"
	"time"

	"github.com/gabriellim314/loadforge/app/load-engine/internal/httpclient"
)

type Metrics struct {
    TotalRequests int
    Successes     int
    Failures      int

    TotalLatency time.Duration
}

func New() *Metrics {
	return &Metrics{}
}

func (m *Metrics) Add(result httpclient.Result) {
	m.TotalRequests++

	if result.Error != nil {
		m.Failures++
		return
	}

	m.Successes++
	m.TotalLatency += result.Latency
}

func (m *Metrics) AverageLatency() time.Duration {
	if m.TotalRequests == 0 {
		return 0
	}

	return m.TotalLatency / time.Duration(m.TotalRequests)
}

func (m *Metrics) SuccessRate() float64 {
	if m.TotalRequests == 0 {
		return 0
	}

	return float64(m.Successes) / float64(m.TotalRequests)
}

func (m *Metrics) Report() string {
	return fmt.Sprintf(
		"Total Requests: %d\nSuccesses: %d\nFailures: %d\nAverage Latency: %s\nSuccess Rate: %.2f%%",
		m.TotalRequests,
		m.Successes,
		m.Failures,
		m.AverageLatency(),
		m.SuccessRate()*100,
	)
}