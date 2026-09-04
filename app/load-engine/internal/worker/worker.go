package worker

import (
	"context"
	"sync"
	"time"

	"github.com/gabriellim314/loadforge/app/load-engine/internal/httpclient"
	"github.com/gabriellim314/loadforge/app/load-engine/internal/metrics"
)

type Config struct {
	URL string
	Concurrency int
	TotalRequests int
	Timeout time.Duration
}

func Run(config Config, collector *metrics.Metrics) error {
    collector.Start()

	httpClient := httpclient.New(config.URL)

	wg := sync.WaitGroup{}
	semaphore := make(chan struct{}, config.Concurrency)

	for i := 0; i < config.TotalRequests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()

			ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
			defer cancel()

			result := httpClient.SendRequest(ctx)
			collector.Add(result)
		}()
	}

	wg.Wait()
	collector.Stop()
	return nil
}