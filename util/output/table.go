package output

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/nitschmann/scdns/cloudflare"

	"github.com/olekukonko/tablewriter"
)

func Table(header []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	return table
}

func RenderCloudflareApiErrorsTable(httpResponse *http.Response, errors []cloudflare.ResponseError) {
	errMsg := "Cloudflare API request failed (HTTP status %v), check the errors below.\n"
	fmt.Printf(errMsg, httpResponse.Status)

	table := Table([]string{"Code", "Error-Message"})
	for _, e := range errors {
		line := []string{strconv.Itoa(e.Code), e.Message}
		table.Append(line)
	}
	table.Render()

	os.Exit(1)
}
