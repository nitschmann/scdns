package cloudflare

import (
	"bytes"
	"encoding/json"
	"net/http"
	"path"

	"github.com/nitschmann/scdns/pkg/util/rest"
)

type DnsRecordsService interface {
	Create(zoneId string, dnsRecord *ModifiedDnsRecord) (*DnsRecordResult, *http.Response, error)
	Delete(zoneId string, id string) (*DnsRecordResult, *http.Response, error)
	Details(zoneId string, id string) (*DnsRecordResult, *http.Response, error)
	List(zoneId string, params *rest.RequestUrlParams) (*DnsRecordResultList, *http.Response, error)
	Update(zoneId string, id string, dnsRecord *ModifiedDnsRecord) (*DnsRecordResult, *http.Response, error)
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

// Type for creation and updates
type ModifiedDnsRecord struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Ttl     int    `json:"ttl" default:"120"`
	Proxied bool   `json:"proxied" default:"false"`
}

func (d DnsRecord) Modifiable() *ModifiedDnsRecord {
	return &ModifiedDnsRecord{
		Type:    d.Type,
		Name:    d.Name,
		Content: d.Content,
		Ttl:     d.Ttl,
		Proxied: d.Proxied,
	}
}

func (o DnsRecordsServiceOperator) Create(zoneId string, dnsRecord *ModifiedDnsRecord) (*DnsRecordResult, *http.Response, error) {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(dnsRecord)

	req := o.client.NewRequest()
	req.Method = "POST"
	req.Body = buffer
	req.Path = path.Join("zones", zoneId, "dns_records")

	var result *DnsRecordResult = &DnsRecordResult{}
	httpResponse, err := ExecRequestAndUnmarshalJson(req, &result)

	return result, httpResponse, err
}

func (o DnsRecordsServiceOperator) Delete(zoneId string, id string) (*DnsRecordResult, *http.Response, error) {
	req := o.client.NewRequest()
	req.Method = "DELETE"
	req.Path = path.Join("zones", zoneId, "dns_records", id)

	var result *DnsRecordResult = &DnsRecordResult{}
	httpResponse, err := ExecRequestAndUnmarshalJson(req, &result)

	return result, httpResponse, err
}

func (o DnsRecordsServiceOperator) Details(zoneId string, id string) (*DnsRecordResult, *http.Response, error) {
	req := o.client.NewRequest()
	req.Method = "GET"
	req.Path = path.Join("zones", zoneId, "dns_records", id)

	var result *DnsRecordResult = &DnsRecordResult{}
	httpResponse, err := ExecRequestAndUnmarshalJson(req, &result)

	return result, httpResponse, err
}

func (o DnsRecordsServiceOperator) List(zoneId string, params *rest.RequestUrlParams) (*DnsRecordResultList, *http.Response, error) {
	req := o.client.NewRequest()
	req.Method = "GET"
	req.Path = path.Join("zones", zoneId, "dns_records")
	req.UrlParams = params

	var result *DnsRecordResultList = &DnsRecordResultList{}
	httpResponse, err := ExecRequestAndUnmarshalJson(req, &result)

	return result, httpResponse, err
}

func (o DnsRecordsServiceOperator) Update(zoneId string, id string, dnsRecord *ModifiedDnsRecord) (*DnsRecordResult, *http.Response, error) {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(dnsRecord)

	req := o.client.NewRequest()
	req.Method = "PUT"
	req.Body = buffer
	req.Path = path.Join("zones", zoneId, "dns_records", id)

	var result *DnsRecordResult = &DnsRecordResult{}
	httpResponse, err := ExecRequestAndUnmarshalJson(req, &result)

	return result, httpResponse, err
}
