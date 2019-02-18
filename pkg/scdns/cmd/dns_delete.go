package cmd

import (
	"fmt"
	"os"

	"github.com/nitschmann/scdns/pkg/cloudflare"
	scdnsOutput "github.com/nitschmann/scdns/pkg/scdns/output"
	// "github.com/nitschmann/scdns/pkg/util/cli"
	// "github.com/nitschmann/scdns/pkg/util/output"

	"github.com/spf13/cobra"
)

func newDnsDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [ID]",
		Aliases: []string{"d"},
		Short:   "Delete a DNS record",
		Long:    "Delete a DNS record in the given zone",
		Args:    cobra.ExactArgs(1),
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

			result, httpResponse, err := client.DnsRecords.Delete(zoneId, id)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if httpResponse.StatusCode == 200 {
				fmt.Printf("Successfully deleted DNS record '%s' for zone '%s'\n", id, zoneId)
			} else {
				scdnsOutput.RenderCloudflareApiErrorsTable(httpResponse, result.Errors)
			}
		},
	}

	return cmd
}
