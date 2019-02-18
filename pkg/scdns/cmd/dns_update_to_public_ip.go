package cmd

import (
	"fmt"
	"os"

	"github.com/nitschmann/scdns/pkg/cloudflare"
	"github.com/nitschmann/scdns/pkg/ipify"
	scdnsOutput "github.com/nitschmann/scdns/pkg/scdns/output"
	"github.com/nitschmann/scdns/pkg/util/output"

	"github.com/spf13/cobra"
)

func newDnsUpdateToPublicIpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-to-public-ip [ID]",
		Short: "Update DNS record content to public IP",
		Long:  "Update content of a existing DNS record with the public IPv4 of this machine",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			credentials := cloudflare.NewCredentials(email, authKey)
			client, err := cloudflare.NewClient(credentials)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			id := args[0]
			zoneId, _ := cmd.Flags().GetString("zone")
			if zoneId == "" {
				fmt.Println("Missing required flag zone")
				os.Exit(1)
			}

			dnsRecordResult, httpResponse, err := client.DnsRecords.Details(zoneId, id)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if httpResponse.StatusCode == 200 {
				ip, err := ipify.GetIp()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				modifiedDnsRecord := dnsRecordResult.Result.Modifiable()
				modifiedDnsRecord.Content = ip

				result, httpResponse, err := client.DnsRecords.Update(zoneId, id, modifiedDnsRecord)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				if httpResponse.StatusCode == 200 {
					tableHeaders := scdnsOutput.CloudflareDnsRecordTableHeaders()
					table := output.Table(tableHeaders)
					line := scdnsOutput.CloudflareDnsRecordTableLine(result.Result)
					table.Append(line)

					table.Render()
				} else {
					scdnsOutput.RenderCloudflareApiErrorsTable(httpResponse, result.Errors)
				}
			} else {
				scdnsOutput.RenderCloudflareApiErrorsTable(httpResponse, dnsRecordResult.Errors)
			}
		},
	}

	return cmd
}
