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
	Client        *Client
	Body          io.Reader
	HttpRequest   *http.Request
	Method        string
	Path          string
	UrlParams     *RequestUrlParams
	BeforeExecRun func(r *Request)
}

func (r *Request) Exec() (*http.Response, error) {
	req, err := http.NewRequest(r.Method, r.FullUrl(), r.Body)
	if err != nil {
		return nil, err
	}

	r.HttpRequest = req
	// r.addHttpRequestClientHeaders()

	if r.BeforeExecRun != nil {
		r.BeforeExecRun(r)
	}

	httpResponse, err := r.Client.Http.Do(r.HttpRequest)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

func (r Request) FullUrl() string {
	u, _ := url.Parse(r.Client.BaseURL.String() + r.Path)
	q, _ := url.ParseQuery(u.RawQuery)

	if len(r.UrlParams.List) > 0 {
		for key, value := range r.UrlParams.List {
			q.Add(key, value)
		}
		u.RawQuery = q.Encode()
	}

	return u.String()
}
