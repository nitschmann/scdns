package output

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
