package cmd

import (
	"log"

	"github.com/nitschmann/scdns/cloudflare"
	"github.com/nitschmann/scdns/util/output"

	"github.com/spf13/cobra"
)

var dnsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all DNS records for a zone",
	Long:  "List all DNS records for a zone in a table",
	Run:   executeDnsListCmd,
}

var dnsListCmdOptionalLocalFlags = []string{
	"type",
	"name",
	"content",
	"page",
	"per_page",
	"order",
	"direction",
	"match",
}

func init() {
	initLocalDnsListCmdFlags()
	dnsCmd.AddCommand(dnsListCmd)
}

func executeDnsListCmd(cmd *cobra.Command, args []string) {
	client := cloudflareClient()
	zoneId := getZoneIdForDnsCmdOrFail(cmd)
	requestParams := parseDnsListCmdParamsForRequest(cmd)

	list, httpResponse, err := client.DnsRecords.List(zoneId, requestParams)
	if err != nil {
		log.Fatalf("An unexpected error occured: %s", err)
	}

	if httpResponse.StatusCode == 200 {
		table := output.Table(dnsRecordResultTableHeaders())
		for _, entry := range list.Result {
			line := dnsRecordResultTableLine(entry)
			table.Append(line)
		}

		table.Render()
	} else {
		output.RenderCloudflareApiErrorsTable(httpResponse, list.Errors)
	}
}

func initLocalDnsListCmdFlags() {
	for _, param := range dnsListCmdOptionalLocalFlags {
		description := "Optional request paramater '" + param + "'"
		dnsListCmd.Flags().StringP(param, "", "", description)
	}
}

func parseDnsListCmdParamsForRequest(cmd *cobra.Command) *cloudflare.RequestParams {
	list := make(map[string]string)

	for _, param := range dnsListCmdOptionalLocalFlags {
		paramValue, err := cmd.Flags().GetString(param)

		if err == nil && paramValue != "" {
			list[param] = paramValue
		}
	}

	return &cloudflare.RequestParams{List: list}
}
