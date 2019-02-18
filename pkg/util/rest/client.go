package rest

import (
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL           *url.URL
	CustomHttpHeaders *CustomHttpHeaders
	Http              *http.Client
}

type ClientService interface {
	GetBaseURL() *url.URL
	GetCustomHttpHeaders() *CustomHttpHeaders
	GetHttp() *http.Client
}

type CustomHttpHeaders struct {
	List map[string]string
}

func (c Client) GetBaseURL() *url.URL {
	return c.BaseURL
}

func (c Client) GetCustomHttpHeaders() *CustomHttpHeaders {
	return c.CustomHttpHeaders
}

func (c Client) GetHttp() *http.Client {
	return c.Http
}

func NewClient(baseUrlStr string, httpClient *http.Client) (*Client, error) {
	client := &Client{Http: httpClient}

	baseUrl, err := url.Parse(baseUrlStr)
	if err != nil {
		return nil, err
	}
	client.BaseURL = baseUrl

	if client.Http == nil {
		client.Http = &http.Client{
			Timeout: time.Duration(30 * time.Second),
		}
	}

	return client, nil
}

func (c Client) NewRequest() *Request {
	return &Request{
		Client:    c,
		UrlParams: &RequestUrlParams{},
	}
}
