package chargify

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func NewClient(subdomain string) (Client, error) {
	apiKey, err := apiKey()
	if err != nil {
		return nil, err
	}
	selfServiceKey, err := selfServiceKey()
	if err != nil {
		return nil, err
	}
	if subdomain == "" {
		return nil, errors.New("no subdomain specified")
	}
	url := constructUrl(subdomain)
	httpClient := http.DefaultClient
	keepAliveTimeout := 600 * time.Second
	timeout := 2 * time.Second
	defaultTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAliveTimeout,
		}).DialContext,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	rt := WithBasicAuth(defaultTransport, apiKey)
	httpClient.Transport = rt
	return &client{url, selfServiceKey, httpClient}, nil
}

func apiKey() (string, error) {
	apiKeyFile := os.Getenv("CHARGIFY_DEFAULT_CREDENTIALS")
	if apiKeyFile == "" {
		return "", errors.New("CHARGIFY_DEFAULT_CREDENTIALS env not found")
	}
	f, err := os.Open(apiKeyFile)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func selfServiceKey() (string, error) {
	apiKeyFile := os.Getenv("CHARGIFY_SITE_SHARED_KEY")
	if apiKeyFile == "" {
		return "", nil
	}
	f, err := os.Open(apiKeyFile)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}
