package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nitschmann/scdns/pkg/cloudflare"
	scdnsOutput "github.com/nitschmann/scdns/pkg/scdns/output"
	"github.com/nitschmann/scdns/pkg/util/cli"
	"github.com/nitschmann/scdns/pkg/util/output"
	"github.com/nitschmann/scdns/pkg/util/rest"

	"github.com/spf13/cobra"
)

func newZonesListCmd() *cobra.Command {
	cmdLocalFlags := []string{
		"name",
		"status",
		"page",
		"per_page",
		"direction",
		"match",
	}

	cmd := &cobra.Command{
		Use:     "list [ID]",
		Aliases: []string{"l"},
		Short:   "List zone(s)",
		Long:    "List all or single (when ID given) zone(s) within an account in a table",
		Args:    cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			credentials := cloudflare.NewCredentials(email, authKey)
			client, err := cloudflare.NewClient(credentials)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			var result *cloudflare.ZoneResult
			var resultList *cloudflare.ZoneResultList
			var httpResponse *http.Response

			if len(args) == 0 {
				params := &rest.RequestUrlParams{
					List: cli.ParseStringFlagList(cmd.Flags(), cmdLocalFlags),
				}

				resultList, httpResponse, err = client.Zones.List(params)
			} else {
				result, httpResponse, err = client.Zones.Details(args[0])
			}

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if result != nil {
				resultList = &cloudflare.ZoneResultList{
					Response: result.Response,
					Result:   []cloudflare.Zone{result.Result},
				}
			}

			if httpResponse.StatusCode == 200 {
				table := output.Table([]string{"id", "name", "status"})
				for _, entry := range resultList.Result {
					line := []string{entry.Id, entry.Name, entry.Status}
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
