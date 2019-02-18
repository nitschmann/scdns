package cmd

import "github.com/spf13/cobra"

func newZonesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "zones",
		Aliases: []string{"z", "zone"},
		Short:   "Overview for Zones within an account",
		Long:    "Zones:\nA Zone is a domain name along with its subdomains and other identities",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				newZonesListCmd().Run(cmd, args)
			}
		},
	}

	cmd.AddCommand(newZonesListCmd())

	return cmd
}
