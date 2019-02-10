package cloudflare

import (
	"fmt"
	"reflect"
)

type Credentials struct {
	authKey string `http-header:"X-Auth-Key"`
	email   string `http-header:"X-Auth-Email"`
}

func NewCredentials(email string, authKey string) *Credentials {
	c := &Credentials{
		authKey: authKey,
		email:   email,
	}

	return c
}

func (c Credentials) httpHeaders() (list [][2]string) {
	t := reflect.ValueOf(c)

	for i := 0; i < t.NumField(); i++ {
		value := t.Field(i)
		typeField := t.Type().Field(i)
		httpHeaderTag := typeField.Tag.Get("http-header")

		if httpHeaderTag != "" {
			var entry [2]string
			entry[0] = httpHeaderTag
			entry[1] = fmt.Sprint(value)
			list = append(list, entry)
		}
	}

	return
}
