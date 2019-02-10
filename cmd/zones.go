package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var zonesCmd = &cobra.Command{
	Use:     "zones",
	Aliases: []string{"z"},
	Short:   "Overview for Zones within an account",
	Long:    "Zones: \n- A Zone is a domain name along with its subdomains and other identities",
	Run:     executeZonesCmd,
}

func init() {
	rootCmd.AddCommand(zonesCmd)
}

func executeZonesCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}
}
