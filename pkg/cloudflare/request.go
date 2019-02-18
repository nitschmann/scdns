package cloudflare

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nitschmann/scdns/pkg/util/rest"
)

func ExecRequestAndUnmarshalJson(r *rest.Request, v interface{}) (*http.Response, error) {
	httpResponse, err := r.Exec()
	if err != nil {
		return httpResponse, err
	}

	values, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return httpResponse, err
	}

	json.Unmarshal(values, &v)

	return httpResponse, nil
}
