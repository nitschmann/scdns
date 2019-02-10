package cmd

import (
	"log"

	"github.com/nitschmann/scdns/util/output"

	"github.com/spf13/cobra"
)

var zonesDetailsCmd = &cobra.Command{
	Use:   "details [ZONE ID]",
	Short: "List detail info for specific zone",
	Long:  "List detail info for specific zone with given ID",
	Args:  cobra.ExactArgs(1),
	Run:   executeZonesDetailsCmd,
}

func init() {
	zonesCmd.AddCommand(zonesDetailsCmd)
}

func executeZonesDetailsCmd(cmd *cobra.Command, args []string) {
	client := cloudflareClient()
	id := args[0]

	zone, httpResponse, err := client.Zones.Details(id)
	if err != nil {
		log.Fatalf("An unexpected error occured: %s", err)
	}

	if httpResponse.StatusCode == 200 {
		table := output.Table([]string{"id", "name", "status"})
		line := []string{zone.Result.Id, zone.Result.Name, zone.Result.Status}
		table.Append(line)
		table.Render()
	} else {
		output.RenderCloudflareApiErrorsTable(httpResponse, zone.Errors)
	}
}
