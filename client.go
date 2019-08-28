package chargify

import (
	"bytes"
	"fmt"
	"net/http"
)

type Client interface {
	Get(string) (*http.Response, error)
	Post([]byte) (*http.Response, error)
	Put([]byte, string) (*http.Response, error)
	Delete([]byte, string) (*http.Response, error)
}

type client struct {
	url        string
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

func (c *client) Get(uri string) (*http.Response, error) {
	return c.httpClient.Get(fmt.Sprintf("%s/%s", c.url, uri))
}

func (c *client) Post(body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func (c *client) Put(body []byte, uri string) (*http.Response, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", c.url, uri), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}

func (c *client) Delete(body []byte, uri string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", c.url, uri), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	return c.httpClient.Do(req)
}
