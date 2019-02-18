package output

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/nitschmann/scdns/pkg/cloudflare"
	"github.com/nitschmann/scdns/pkg/util/output"
)

func RenderCloudflareApiErrorsTable(httpResponse *http.Response, errors []cloudflare.ResponseError) {
	errMsg := "Cloudflare API request failed (HTTP status %v), check the errors below.\n"
	fmt.Printf(errMsg, httpResponse.Status)

	table := output.Table([]string{"Code", "Error-Message"})
	for _, e := range errors {
		line := []string{strconv.Itoa(e.Code), e.Message}
		table.Append(line)
	}
	table.Render()

	os.Exit(1)
}
