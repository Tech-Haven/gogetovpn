package configs

import (
	"net/http"
	"time"
)

type HttpClient struct {
	HttpClient *http.Client
}

type Config struct {
	HttpClient HttpClient
}

func New() (*Config, error) {
	httpClient := newHttpClient()

	return &Config{
		HttpClient: *httpClient,
	}, nil
}

func newHttpClient() *HttpClient {
	// Define HTTP Client transport options
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100

	// Create HTTP client
	client := &http.Client{
		Timeout:   time.Second * 60,
		Transport: t,
	}

	httpClient := &HttpClient{
		HttpClient: client,
	}

	return httpClient
}
