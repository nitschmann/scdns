package cloudflare

import "github.com/nitschmann/scdns/pkg/util/rest"

const (
	BASE_URL          = "https://api.cloudflare.com/client/v4/"
	HTTP_CONTENT_TYPE = "application/json"
)

type Client struct {
	*rest.Client

	credentials *Credentials

	DnsRecords DnsRecordsService
	Zones      ZonesService
}

func NewClient(credentials *Credentials) (*Client, error) {
	restClient, err := rest.NewClient(BASE_URL, nil)
	if err != nil {
		return nil, err
	}
	restClient.CustomHttpHeaders = clientCustomHttpHeaders(credentials)

	client := &Client{
		Client:      restClient,
		credentials: credentials,
	}

	client.DnsRecords = &DnsRecordsServiceOperator{client: client}
	client.Zones = &ZonesServiceOperator{client: client}

	return client, nil
}

func clientCustomHttpHeaders(credentials *Credentials) *rest.CustomHttpHeaders {
	list := make(map[string]string)
	list["Content-Type"] = HTTP_CONTENT_TYPE

	for _, header := range credentials.httpHeaders() {
		list[header[0]] = header[1]
	}

	return &rest.CustomHttpHeaders{List: list}
}
