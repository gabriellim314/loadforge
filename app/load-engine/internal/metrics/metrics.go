package metrics

import (
	"fmt"
	"sync"
	"time"

	"github.com/gabriellim314/loadforge/app/load-engine/internal/httpclient"
)

type Metrics struct {
    mu sync.Mutex

	StartedAt time.Time
	EndedAt time.Time

    TotalRequests int
    Successes     int
    Failures      int
    StatusCodeStats map[int]int
    TotalLatency time.Duration
}

func (m *Metrics) Start() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.StartedAt = time.Now()
}

func (m *Metrics) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.EndedAt = time.Now()
}

func (m *Metrics) Duration() time.Duration {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.EndedAt.Sub(m.StartedAt)
}

func (m *Metrics) RPS() float64 {
	d := m.Duration()
	if d == 0 {
		return 0
	}
	return float64(m.TotalRequests) / d.Seconds()
}

func New() *Metrics {
	return &Metrics{StatusCodeStats: make(map[int]int)}
}

func (m *Metrics) Add(result httpclient.Result) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.TotalRequests++

	if result.Error != nil {
		m.Failures++
		return
	}

	m.StatusCodeStats[result.StatusCode]++
	m.TotalLatency += result.Latency

	if result.StatusCode >= 200 && result.StatusCode < 300 {
		m.Successes++
	} else {
		m.Failures++
	}
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
		"Total Requests: %d\nSuccesses: %d\nFailures: %d\nAverage Latency: %s\nSuccess Rate: %.2f%%\nStatus Code Stats: %v\nRPS: %.2f\nDuration: %s",
		m.TotalRequests,
		m.Successes,
		m.Failures,
		m.AverageLatency(),
		m.SuccessRate()*100,
		m.StatusCodeStats,
		m.RPS(),
		m.Duration(),
	)
}