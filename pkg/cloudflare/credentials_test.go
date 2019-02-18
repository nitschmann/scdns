package cloudflare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCredentials(t *testing.T) {
	authKey := "auth-key"
	email := "tester@testing.com"
	credentials := NewCredentials(email, authKey)

	assert.NotNil(t, credentials)
	assert.Equal(t, authKey, credentials.authKey)
	assert.Equal(t, email, credentials.email)
}

func TestHttpHeaders(t *testing.T) {
	authKey := "auth-keys"
	email := "tester@testing.com"
	expected := make([][2]string, 2)
	expected[0] = [2]string{"X-Auth-Key", authKey}
	expected[1] = [2]string{"X-Auth-Email", email}
	headers := NewCredentials(email, authKey).httpHeaders()

	assert.Len(t, headers, 2)
	assert.ElementsMatch(t, expected, headers)
}
