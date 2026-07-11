package main

import (
	"fmt"
	
	"github.com/gabriellim314/loadforge/app/load-engine/internal/httpclient"
	"github.com/gabriellim314/loadforge/app/load-engine/internal/metrics"
)

func main() {

	client := httpclient.New("http://localhost:8000/health")
    collector := metrics.New()

	for i :=0; i < 100; i++ {
		result := client.SendRequest()
		collector.Add(result)
	}

	fmt.Println(collector.Report())
}