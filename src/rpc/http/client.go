package http

import (
	"github.com/go-kit/log"
	"net/http"
)

type RPC struct {
	client *http.Client
	logger log.Logger
}

// NewRPC New returns a RPC backed by HTTP
func NewRPC(client *http.Client, logger log.Logger) *RPC {
	return &RPC{
		client: client,
		logger: log.With(logger, "rpc", "http"),
	}
}

// InitHTTPClient Initialize HTTP client
func InitHTTPClient() *http.Client {
	httpClient := &http.Client{
		// Timeout:   time.Second * 5,
		Transport: &http.Transport{},
	}

	return httpClient
}
