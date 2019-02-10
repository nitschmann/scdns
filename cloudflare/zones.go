package cloudflare

import (
	"net/http"
	"path"
)

type ZonesService interface {
	Details(id string) (*ZoneResult, *http.Response, error)
	List(params *RequestParams) (*ZoneResultList, *http.Response, error)
}

type ZonesServiceOperator struct {
	client *Client
}

var _ ZonesService = &ZonesServiceOperator{}

type Zone struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ZoneResult struct {
	Response
	Result Zone `json:"result"`
}

type ZoneResultList struct {
	Response
	Result []Zone `json:"result"`
}

func (o ZonesServiceOperator) Details(id string) (*ZoneResult, *http.Response, error) {
	req := &Request{
		Client: o.client,
		Method: "GET",
		Path:   path.Join("zones", id),
		Params: &RequestParams{},
	}

	var result *ZoneResult = &ZoneResult{}
	httpResponse, err := req.ExecAndUnmarshalJson(&result)

	return result, httpResponse, err
}

func (o ZonesServiceOperator) List(params *RequestParams) (*ZoneResultList, *http.Response, error) {
	if params == nil {
		params = &RequestParams{}
	}

	req := &Request{
		Client: o.client,
		Method: "GET",
		Path:   "zones",
		Params: params,
	}

	var result *ZoneResultList = &ZoneResultList{}
	httpResponse, err := req.ExecAndUnmarshalJson(&result)

	return result, httpResponse, err
}
