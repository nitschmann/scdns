package cmd

import (
	"log"

	"github.com/nitschmann/scdns/cloudflare"
	"github.com/nitschmann/scdns/util/cli"
	"github.com/nitschmann/scdns/util/output"

	"github.com/spf13/cobra"
)

var dnsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new DNS record",
	Long:  "Create a new DNS record in the given zone",
	Run:   executeDnsCreateCmd,
}

func init() {
	initLocalDnsCreateCmdFlags()
	dnsCmd.AddCommand(dnsCreateCmd)
}

func executeDnsCreateCmd(cmd *cobra.Command, args []string) {
	client := cloudflareClient()
	zoneId := getZoneIdForDnsCmdOrFail(cmd)
	newDnsRecord := &cloudflare.ModifiedDnsRecord{}
	cli.SetInterfaceFieldsFromFlags(cmd.Flags(), newDnsRecord)

	dnsRecord, httpResponse, err := client.DnsRecords.Create(zoneId, newDnsRecord)
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

func initLocalDnsCreateCmdFlags() {
	cli.AssignFlagsFromInterfaceFields(dnsCreateCmd.Flags(), &cloudflare.ModifiedDnsRecord{})
}
