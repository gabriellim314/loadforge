package main

import (
	"fmt"
	
	"github.com/gabriellim314/loadforge/app/load-engine/internal/metrics"
	"github.com/gabriellim314/loadforge/app/load-engine/internal/worker"
)

func main() {

    collector := metrics.New()

	err := worker.Run(worker.Config{
		URL: "http://localhost:8000/health",
		Concurrency: 10,
		TotalRequests: 100,
	}, collector)

	if err != nil {
		fmt.Printf("Error running worker: %v\n", err)
		return
	}

	fmt.Println(collector.Report())
}