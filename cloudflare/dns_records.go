package cloudflare

import (
	"net/http"
	"path"
)

type DnsRecordsService interface {
	Details(zoneId string, id string) (*DnsRecordResult, *http.Response, error)
	List(zoneId string, params *RequestParams) (*DnsRecordResultList, *http.Response, error)
}

type DnsRecordsServiceOperator struct {
	client *Client
}

var _ DnsRecordsService = &DnsRecordsServiceOperator{}

type DnsRecord struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Content    string `json:"content"`
	Proxiable  bool   `json:"proxiable"`
	Proxied    bool   `json:"proxied"`
	Ttl        int    `json:"ttl"`
	Locked     bool   `json:"locked"`
	ZoneId     string `json:"zone_id"`
	ZoneName   string `json:"zone_name"`
	CreatedOn  string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
}

type DnsRecordResult struct {
	Response
	Result DnsRecord `json:"result"`
}

type DnsRecordResultList struct {
	Response
	Result []DnsRecord `json:"result"`
}

func (o DnsRecordsServiceOperator) Details(zoneId string, id string) (*DnsRecordResult, *http.Response, error) {
	req := &Request{
		Client: o.client,
		Method: "GET",
		Path:   path.Join("zones", zoneId, "dns_records", id),
		Params: &RequestParams{},
	}

	var result *DnsRecordResult = &DnsRecordResult{}
	httpResponse, err := req.ExecAndUnmarshalJson(&result)

	return result, httpResponse, err
}

func (o DnsRecordsServiceOperator) List(zoneId string, params *RequestParams) (*DnsRecordResultList, *http.Response, error) {
	if params == nil {
		params = &RequestParams{}
	}

	req := &Request{
		Client: o.client,
		Method: "GET",
		Path:   path.Join("zones", zoneId, "dns_records"),
		Params: params,
	}

	var result *DnsRecordResultList = &DnsRecordResultList{}
	httpResponse, err := req.ExecAndUnmarshalJson(&result)

	return result, httpResponse, err
}
