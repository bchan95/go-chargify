package chargify

import (
	"fmt"
	"net/http"
)

type Client struct {
	url   string
	httpClient *http.Client
}

type withBasicAuth struct {
	apiKey string
	rt     http.RoundTripper
}

func constructUrl(subdomain string) string {
	return fmt.Sprintf("https://%s.chargify.com", subdomain)
}

func WithBasicAuth(rt http.RoundTripper, apiKey string) *withBasicAuth {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &withBasicAuth{
		apiKey: apiKey,
		rt:     rt,
	}
}
func (ba *withBasicAuth) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(ba.apiKey, "X")
	return ba.rt.RoundTrip(req)
}
