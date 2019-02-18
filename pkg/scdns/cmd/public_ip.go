package cmd

import (
	"fmt"
	"os"

	"github.com/nitschmann/scdns/pkg/ipify"

	"github.com/spf13/cobra"
)

func newCmdPublicIp() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "public-ip",
		Aliases: []string{"ip", "public"},
		Short:   "Get public IPv4",
		Long:    "Get the current public IPv4 of the network for this machine",
		Run: func(cmd *cobra.Command, args []string) {
			ip, err := ipify.GetIp()
			if err == nil {
				fmt.Println(ip)
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
