package cmd

import (
	"log"

	"github.com/nitschmann/scdns/util/output"

	"github.com/spf13/cobra"
)

var dnsDeleteCmd = &cobra.Command{
	Use:   "delete [DNS-RECORD ID]",
	Short: "Delete a DNS record",
	Long:  "Delete a DNS record in the given zone",
	Args:  cobra.ExactArgs(1),
	Run:   executeDnsDeleteCmd,
}

func init() {
	dnsCmd.AddCommand(dnsDeleteCmd)
}

func executeDnsDeleteCmd(cmd *cobra.Command, args []string) {
	client := cloudflareClient()
	zoneId := getZoneIdForDnsCmdOrFail(cmd)
	dnsRecordId := args[0]

	dnsRecord, httpResponse, err := client.DnsRecords.Delete(zoneId, dnsRecordId)
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
