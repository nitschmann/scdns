package cmd

import (
	"log"

	"github.com/nitschmann/scdns/util/output"

	"github.com/spf13/cobra"
)

var dnsDetailsCmd = &cobra.Command{
	Use:   "details [DNS-RECORD ID]",
	Short: "Show details of a single DNS record",
	Long:  "Show details of a single DNS record in the given zone",
	Args:  cobra.ExactArgs(1),
	Run:   executeDnsDetailsCmd,
}

func init() {
	dnsCmd.AddCommand(dnsDetailsCmd)
}

func executeDnsDetailsCmd(cmd *cobra.Command, args []string) {
	client := cloudflareClient()
	zoneId := getZoneIdForDnsCmdOrFail(cmd)
	dnsRecordId := args[0]

	dnsRecord, httpResponse, err := client.DnsRecords.Details(zoneId, dnsRecordId)
	if err != nil {
		log.Fatalf("An unexpected error occured: %s", err)
	}

	if httpResponse.StatusCode == 200 {
		table := output.Table(dnsRecordResultTableHeaders())
		line := dnsRecordResultTableLine(dnsRecord.Result)
		table.Append(line)
		table.Render()
	} else {
		output.RenderCloudflareApiErrorsTable(httpResponse, dnsRecord.Errors)
	}
}
