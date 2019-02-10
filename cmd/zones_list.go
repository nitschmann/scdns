package cmd

import (
	"log"

	"github.com/nitschmann/scdns/cloudflare"
	"github.com/nitschmann/scdns/util/output"

	"github.com/spf13/cobra"
)

var zonesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all zones",
	Long:  "List all zones within an account in a table",
	Run:   executeZonesListCmd,
}

var zonesListCmdOptionalLocalFlags = []string{
	"name",
	"status",
	"page",
	"per_page",
	"order",
	"direction",
	"match",
}

func init() {
	initLocalZonesListCmdFlags()
	zonesCmd.AddCommand(zonesListCmd)
}

func executeZonesListCmd(cmd *cobra.Command, args []string) {
	client := cloudflareClient()
	requestParams := parseZonesListCmdParamsForRequest(cmd)

	list, httpResponse, err := client.Zones.List(requestParams)
	if err != nil {
		log.Fatalf("An unexpected error occured: %s", err)
	}

	if httpResponse.StatusCode == 200 {
		table := output.Table([]string{"id", "name", "status"})
		for _, entry := range list.Result {
			line := []string{entry.Id, entry.Name, entry.Status}
			table.Append(line)
		}

		table.Render()
	} else {
		output.RenderCloudflareApiErrorsTable(httpResponse, list.Errors)
	}
}

func initLocalZonesListCmdFlags() {
	for _, param := range zonesListCmdOptionalLocalFlags {
		description := "Optional request paramater '" + param + "'"
		zonesListCmd.Flags().StringP(param, "", "", description)
	}
}

func parseZonesListCmdParamsForRequest(cmd *cobra.Command) *cloudflare.RequestParams {
	list := make(map[string]string)

	for _, param := range zonesListCmdOptionalLocalFlags {
		paramValue, err := cmd.Flags().GetString(param)

		if err == nil && paramValue != "" {
			list[param] = paramValue
		}
	}

	return &cloudflare.RequestParams{List: list}
}
