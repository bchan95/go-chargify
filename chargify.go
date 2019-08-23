package chargify

import (
	"context"
	"errors"
	"net/http"
	"os"
)

func NewClient(ctx context.Context, subdomain string) (*Client, error) {
	apiKey := os.Getenv("CHARGIFY_DEFAULT_CREDENTIALS")
	if apiKey == "" {
		return nil, errors.New("CHARGIFY_DEFAULT_CREDENTIALS env not found")
	}
	if subdomain == "" {
		return nil, errors.New("no subdomain specified")
	}
	url := constructUrl(subdomain)
	httpClient := http.DefaultClient

	rt := WithBasicAuth(httpClient.Transport, apiKey)
	httpClient.Transport = rt
	return &Client{
		url:   url,
		httpClient: httpClient,
	}, nil
}
