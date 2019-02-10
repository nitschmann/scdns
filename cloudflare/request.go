package cloudflare

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type RequestParams struct {
	List map[string]string
}

type Request struct {
	Client      *Client
	Body        io.Reader
	HttpRequest *http.Request
	Method      string
	Path        string
	Params      *RequestParams
}

func (r Request) addHttpRequestClientHeaders() {
	for _, header := range r.Client.credentials.httpHeaders() {
		key := header[0]
		value := header[1]
		r.HttpRequest.Header.Add(key, value)
	}

	r.HttpRequest.Header.Add("Content-Type", r.Client.contentType)
}

func (r Request) Exec() (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.fullURL(), r.Body)
	if err != nil {
		log.Fatalf("Couldn't initialize HTTP request due an error %s\n", err)
	}

	r.HttpRequest = req
	r.addHttpRequestClientHeaders()

	httpResponse, err := r.Client.client.Do(r.HttpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

func (r Request) ExecAndUnmarshalJson(v interface{}) (*http.Response, error) {
	httpResponse, err := r.Exec()
	if err != nil {
		return httpResponse, err
	}

	values, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return httpResponse, err
	}

	json.Unmarshal(values, &v)

	return httpResponse, nil
}

func (r Request) fullURL() string {
	u, _ := url.Parse(r.Client.BaseURL.String() + r.Path)
	q, _ := url.ParseQuery(u.RawQuery)

	if len(r.Params.List) > 0 {
		for key, value := range r.Params.List {
			q.Add(key, value)
		}
		u.RawQuery = q.Encode()
	}

	return u.String()
}
