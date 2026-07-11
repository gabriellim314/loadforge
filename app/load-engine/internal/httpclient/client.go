package httpclient

import (
	"net/http"
	"time"
)

// HTTPClient is a struct that represents the HTTP client
type HTTPClient struct {
	URL string
}

// Result contains the outcome of a single HTTP request
type Result struct {
	StatusCode int
	Latency    time.Duration
	Error      error
}

func New(url string) *HTTPClient {
	return &HTTPClient{URL: url}
} 

func (client *HTTPClient) SendRequest() Result {

	start := time.Now()
	resp, err := http.Get(client.URL)
	duration := time.Since(start)

    if err != nil {
		return Result{Error: err}
	}

	defer resp.Body.Close()

	return Result{
		StatusCode: resp.StatusCode,
		Latency:    duration,
		Error:      nil,
	}
}