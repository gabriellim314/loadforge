package httpclient

import (
	"context"
	"net/http"
	"time"
)

// HTTPClient is a struct that represents the HTTP client
type HTTPClient struct {
	URL string
	Client *http.Client
}

// Result contains the outcome of a single HTTP request
type Result struct {
	StatusCode int
	Latency    time.Duration
	Error      error
}

func New(url string) *HTTPClient {
	return &HTTPClient{URL: url, Client: &http.Client{}}
} 

func (client *HTTPClient) SendRequest(ctx context.Context) Result {

	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, client.URL, nil)
	if err != nil {
		return Result{Error: err}
	}

	resp, err := client.Client.Do(req)
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