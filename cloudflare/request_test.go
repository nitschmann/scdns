package cloudflare

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestAttributeAssignments(t *testing.T) {
	credentials := NewCredentials("email", "authKey")
	client := NewClient(credentials)
	method := "GET"
	path := "zones"

	testReq := &Request{
		Client: client,
		Method: method,
		Path:   path,
	}

	assert.Equal(t, testReq.Client, client)
	assert.Equal(t, testReq.Method, method)
	assert.Equal(t, testReq.Path, path)
	assert.Nil(t, testReq.Params)
}

func TestRequestAddHttpRequestClientHeaders(t *testing.T) {
	email := "email"
	authKey := "authKey"
	credentials := NewCredentials(email, authKey)
	client := NewClient(credentials)
	path := "dns"
	method := "GET"

	testReq := &Request{
		Client: client,
		Method: "GET",
		Path:   path,
		Params: &RequestParams{},
	}
	httpReq, _ := http.NewRequest(method, testReq.fullURL(), testReq.Body)
	testReq.HttpRequest = httpReq
	testReq.addHttpRequestClientHeaders()
	headers := httpReq.Header

	// TODO: Some more specific tests here
	assert.Len(t, headers, 3)
	assert.Contains(t, headers, "Content-Type")
	assert.Contains(t, headers, "X-Auth-Key")
	assert.Contains(t, headers, "X-Auth-Email")
}

func TestRequestFullURL(t *testing.T) {
	credentials := NewCredentials("email", "authKey")
	client := NewClient(credentials)
	path := "zones"

	testReq := &Request{
		Client: client,
		Method: "GET",
		Path:   path,
		Params: &RequestParams{},
	}

	expectedURL := client.BaseURL.String() + path
	assert.Equal(t, expectedURL, testReq.fullURL())

	params := &RequestParams{
		List: map[string]string{"name": "test-domain.com"},
	}
	testReq.Params = params

	expectedURLWithParams := expectedURL + "?name=test-domain.com"
	assert.Equal(t, expectedURLWithParams, testReq.fullURL())
}
