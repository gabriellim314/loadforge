package main

import (
	"fmt"
	"flag"	
	"time"

	"github.com/gabriellim314/loadforge/app/load-engine/internal/metrics"
	"github.com/gabriellim314/loadforge/app/load-engine/internal/worker"
)

func main() {
	url := flag.String("url", "http://localhost:8000/health", "URL to test")
	concurrency := flag.Int("concurrency", 10, "Number of concurrent requests")
	totalRequests := flag.Int("total-requests", 100, "Total number of requests")
    timeout := flag.Duration("timeout", 5*time.Second, "Timeout for each request")
	
	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required")
		return
	}

	if *concurrency <= 0 {
		fmt.Println("Concurrency must be greater than 0")
		return
	}

	if *totalRequests <= 0 {
		fmt.Println("Total requests must be greater than 0")
		return
	}

	if *timeout <= 0 {
		fmt.Println("Timeout must be greater than 0")
		return
	}

    collector := metrics.New()

	err := worker.Run(worker.Config{
		URL: *url,
		Concurrency: *concurrency,
		TotalRequests: *totalRequests,
		Timeout: *timeout,
	}, collector)

	if err != nil {
		fmt.Printf("Error running worker: %v\n", err)
		return
	}

	fmt.Println(collector.Report())
}