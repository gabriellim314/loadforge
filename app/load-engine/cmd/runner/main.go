package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	start := time.Now()

	resp, err := http.Get("http://localhost:8000/health")

	duration := time.Since(start)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status:", resp.Status)
	fmt.Println("Latency:", duration)
}