package output

import (
	"strconv"

	"github.com/nitschmann/scdns/pkg/cloudflare"
)

func CloudflareDnsRecordTableHeaders() []string {
	return []string{
		"id",
		"type",
		"name",
		"content",
		"proxiable",
		"proxied",
		"ttl",
		"locked",
		"created on",
		"modified on",
	}
}

func CloudflareDnsRecordTableLine(dnsRecord cloudflare.DnsRecord) []string {
	return []string{
		dnsRecord.Id,
		dnsRecord.Type,
		dnsRecord.Name,
		dnsRecord.Content,
		strconv.FormatBool(dnsRecord.Proxiable),
		strconv.FormatBool(dnsRecord.Proxied),
		strconv.Itoa(dnsRecord.Ttl),
		strconv.FormatBool(dnsRecord.Locked),
		dnsRecord.CreatedOn,
		dnsRecord.ModifiedOn,
	}
}
