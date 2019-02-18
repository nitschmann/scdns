package cloudflare

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientServices(t *testing.T) {
	credentials := NewCredentials("email", "authKey")
	client := NewClient(credentials)
	services := []string{
		"DnsRecords",
		"Zones",
	}

	cp := reflect.ValueOf(client)
	cv := reflect.Indirect(cp)

	for _, s := range services {
		assert.NotNil(t, cv.FieldByName(s))
	}
}
