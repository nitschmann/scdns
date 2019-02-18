package cloudflare

import (
	"net/http"
	"path"

	"github.com/nitschmann/scdns/pkg/util/rest"
)

type ZonesService interface {
	Details(id string) (*ZoneResult, *http.Response, error)
	List(params *rest.RequestUrlParams) (*ZoneResultList, *http.Response, error)
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
	req := o.client.NewRequest()
	req.Method = "GET"
	req.Path = path.Join("zones", id)

	var result *ZoneResult = &ZoneResult{}
	httpResponse, err := ExecRequestAndUnmarshalJson(req, &result)

	return result, httpResponse, err
}

func (o ZonesServiceOperator) List(params *rest.RequestUrlParams) (*ZoneResultList, *http.Response, error) {
	req := o.client.NewRequest()
	req.Method = "GET"
	req.Path = "zones"
	req.UrlParams = params

	var result *ZoneResultList = &ZoneResultList{}
	httpResponse, err := ExecRequestAndUnmarshalJson(req, &result)

	return result, httpResponse, err
}
