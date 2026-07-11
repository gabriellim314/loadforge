package main

import (
	"fmt"
	
	"github.com/gabriellim314/loadforge/app/load-engine/internal/httpclient"
)

func main() {

	client := httpclient.New("http://localhost:8000/health")
	result := client.SendRequest()

	if result.Error != nil {
		fmt.Println("Request failed:", result.Error)
		return
	}

	fmt.Println("Status:", result.StatusCode)
	fmt.Println("Latency:", result.Latency)
}