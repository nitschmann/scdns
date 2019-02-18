package cmd

import "github.com/spf13/cobra"

func newDnsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dns",
		Aliases: []string{"d"},
		Short:   "DNS records for a zone",
		Long:    "DNS Records:\nRepresents DNS Records for a specific Zone",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
			}
		},
	}

	cmd.PersistentFlags().StringP("zone", "z", "", "Zone ID")

	cmd.AddCommand(newDnsListCmd())

	return cmd
}
