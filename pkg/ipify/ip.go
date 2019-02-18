package ipify

import (
	"io/ioutil"

	"github.com/nitschmann/scdns/pkg/util/rest"
)

const BASE_URL = "https://api.ipify.org"

func GetIp() (string, error) {
	client, err := rest.NewClient(BASE_URL, nil)
	if err != nil {
		return "", err
	}

	req := client.NewRequest()
	req.Method = "GET"

	httpResponse, err := req.Exec()
	if err != nil {
		return "", err
	}

	ip, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
