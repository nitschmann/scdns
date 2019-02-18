package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/nitschmann/scdns/pkg/cloudflare"
	scdnsOutput "github.com/nitschmann/scdns/pkg/scdns/output"
	"github.com/nitschmann/scdns/pkg/util/cli"
	"github.com/nitschmann/scdns/pkg/util/output"

	"github.com/spf13/cobra"
)

func newDnsCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Short:   "Create a new DNS record",
		Long:    "Create a new DNS record in the given zone",
		Run: func(cmd *cobra.Command, args []string) {
			credentials := cloudflare.NewCredentials(email, authKey)
			client, err := cloudflare.NewClient(credentials)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			zoneId, _ := cmd.Flags().GetString("zone")
			if zoneId == "" {
				fmt.Println("Missing required flag zone")
				os.Exit(1)
			}

			newDnsRecord := &cloudflare.ModifiedDnsRecord{}
			cli.SetInterfaceFieldsFromFlags(cmd.Flags(), newDnsRecord)

			result, httpResponse, err := client.DnsRecords.Create(zoneId, newDnsRecord)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if httpResponse.StatusCode == 200 {
				tableHeaders := scdnsOutput.CloudflareDnsRecordTableHeaders()
				table := output.Table(tableHeaders)
				entry := result.Result
				line := []string{
					entry.Id,
					entry.Type,
					entry.Name,
					entry.Content,
					strconv.FormatBool(entry.Proxiable),
					strconv.FormatBool(entry.Proxied),
					strconv.Itoa(entry.Ttl),
					strconv.FormatBool(entry.Locked),
					entry.CreatedOn,
					entry.ModifiedOn,
				}
				table.Append(line)

				table.Render()
			} else {
				scdnsOutput.RenderCloudflareApiErrorsTable(httpResponse, result.Errors)
			}
		},
	}

	cli.AssignFlagsFromInterfaceFields(cmd.Flags(), &cloudflare.ModifiedDnsRecord{})

	return cmd
}
