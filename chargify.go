package chargify

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func NewClient(subdomain string) (Client, error) {
	apiKey, err := apiKey()
	if err != nil {
		return nil, err
	}
	if subdomain == "" {
		return nil, errors.New("no subdomain specified")
	}
	url := constructUrl(subdomain)
	httpClient := http.DefaultClient

	rt := WithBasicAuth(httpClient.Transport, apiKey)
	httpClient.Transport = rt
	return &client{url, httpClient}, nil
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
	return string(b), nil
}