package rest

import (
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL *url.URL
	Http    *http.Client
}

func NewClient(baseUrlStr string, httpClient *http.Client) (*Client, error) {
	client := &Client{Http: httpClient}
	err := client.SetBaseURL(baseUrlStr)
	if err != nil {
		return nil, err
	}

	if client.Http == nil {
		client.Http = &http.Client{
			Timeout: time.Duration(30 * time.Second),
		}
	}

	return client, nil
}

func (c Client) SetBaseURL(baseUrlStr string) error {
	baseUrl, err := url.Parse(baseUrlStr)
	if err != nil {
		return err
	} else {
		c.BaseURL = baseUrl
		return nil
	}
}
