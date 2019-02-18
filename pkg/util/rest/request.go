package rest

import (
	"io"
	"net/http"
	"net/url"
)

type RequestUrlParams struct {
	List map[string]string
}

type Request struct {
	Client        ClientService
	Body          io.Reader
	HttpRequest   *http.Request
	Method        string
	Path          string
	UrlParams     *RequestUrlParams
	BeforeExecRun func(r *Request)
}

func (r *Request) addCustomHttpHeadersToRequest() {
	customHttpHeaders := r.Client.GetCustomHttpHeaders()

	if customHttpHeaders != nil && len(customHttpHeaders.List) > 0 {
		for key, value := range customHttpHeaders.List {
			r.HttpRequest.Header.Add(key, value)
		}
	}
}

func (r *Request) Exec() (*http.Response, error) {
	if r.HttpRequest == nil {
		req, err := http.NewRequest(r.Method, r.FullURL(), r.Body)
		if err != nil {
			return nil, err
		}

		r.HttpRequest = req
	}

	r.addCustomHttpHeadersToRequest()

	if r.BeforeExecRun != nil {
		r.BeforeExecRun(r)
	}

	httpClient := r.Client.GetHttp()
	httpResponse, err := httpClient.Do(r.HttpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

func (r Request) FullURL() string {
	u, _ := url.Parse(r.Client.GetBaseURL().String() + r.Path)
	q, _ := url.ParseQuery(u.RawQuery)

	if r.UrlParams != nil && len(r.UrlParams.List) > 0 {
		for key, value := range r.UrlParams.List {
			q.Add(key, value)
		}
		u.RawQuery = q.Encode()
	}

	return u.String()
}
