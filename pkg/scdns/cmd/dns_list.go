package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/nitschmann/scdns/pkg/cloudflare"
	scdnsOutput "github.com/nitschmann/scdns/pkg/scdns/output"
	"github.com/nitschmann/scdns/pkg/util/cli"
	"github.com/nitschmann/scdns/pkg/util/output"
	"github.com/nitschmann/scdns/pkg/util/rest"

	"github.com/spf13/cobra"
)

func newDnsListCmd() *cobra.Command {
	cmdLocalFlags := []string{
		"type",
		"name",
		"content",
		"page",
		"per_page",
		"order",
		"direction",
		"match",
	}

	cmd := &cobra.Command{
		Use:     "list [ID]",
		Aliases: []string{"l"},
		Short:   "List all DNS record(s) for a zone",
		Long:    "List all or single (when ID given) DNS record(s) for a zone in a table",
		Args:    cobra.RangeArgs(0, 1),
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

			var result *cloudflare.DnsRecordResult
			var resultList *cloudflare.DnsRecordResultList
			var httpResponse *http.Response

			if len(args) == 0 {
				params := &rest.RequestUrlParams{
					List: cli.ParseStringFlagList(cmd.Flags(), cmdLocalFlags),
				}

				resultList, httpResponse, err = client.DnsRecords.List(zoneId, params)
			} else {
				result, httpResponse, err = client.DnsRecords.Details(zoneId, args[0])
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if result != nil {
				resultList = &cloudflare.DnsRecordResultList{
					Response: result.Response,
					Result:   []cloudflare.DnsRecord{result.Result},
				}
			}

			if httpResponse.StatusCode == 200 {
				tableHeaders := scdnsOutput.CloudflareDnsRecordTableHeaders()
				table := output.Table(tableHeaders)

				for _, entry := range resultList.Result {
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
				}

				table.Render()
			} else {
				scdnsOutput.RenderCloudflareApiErrorsTable(httpResponse, resultList.Errors)
			}
		},
	}

	cli.AddStringFlags(cmd.Flags(), cmdLocalFlags)

	return cmd
}
