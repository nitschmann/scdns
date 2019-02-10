package cloudflare

import (
	"net/http"
	"net/url"
	"time"
)

const (
	defaultBaseURL         = "https://api.cloudflare.com/client/v4/"
	defaultHttpContentType = "application/json"
)

type Client struct {
	client      *http.Client
	contentType string
	credentials *Credentials

	BaseURL *url.URL

	Zones ZonesService
}

// TODO: Timeout maybe dynamically configurable?
func NewClient(credentials *Credentials) *Client {
	httpClient := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:      httpClient,
		contentType: defaultHttpContentType,
		credentials: credentials,
		BaseURL:     baseURL,
	}
	c.Zones = &ZonesServiceOperator{client: c}

	return c
}
