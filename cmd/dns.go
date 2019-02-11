package cmd

import (
	"log"
	"os"
	"strconv"

	"github.com/nitschmann/scdns/cloudflare"

	"github.com/spf13/cobra"
)

var dnsCmd = &cobra.Command{
	Use:     "dns",
	Aliases: []string{"d"},
	Short:   "DNS records for a zone",
	Long:    "DNS Records:\nRepresents DNS Records for a specific Zone",
	Run:     executeDnsCmd,
}

func init() {
	initDnsCmdPersistentFlags()
	rootCmd.AddCommand(dnsCmd)
}

func dnsRecordResultTableHeaders() []string {
	return []string{
		"id",
		"type",
		"name",
		"content",
		"proxiable",
		"proxied",
		"ttl",
		"locked",
		"created on",
		"modified on",
	}
}

func dnsRecordResultTableLine(record cloudflare.DnsRecord) []string {
	return []string{
		record.Id,
		record.Type,
		record.Name,
		record.Content,
		strconv.FormatBool(record.Proxiable),
		strconv.FormatBool(record.Proxied),
		strconv.Itoa(record.Ttl),
		strconv.FormatBool(record.Locked),
		record.CreatedOn,
		record.ModifiedOn,
	}
}

func executeDnsCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}
}

func getZoneIdForDnsCmdOrFail(cmd *cobra.Command) string {
	zoneId, err := cmd.Flags().GetString("zone")
	if err != nil {
		log.Fatalf("An unexpected error occured: %s\n", err)
	}
	if zoneId == "" {
		log.Fatalf("Flag 'zone' is required")
	}

	return zoneId
}

func initDnsCmdPersistentFlags() {
	dnsCmd.PersistentFlags().StringP("zone", "z", "", "Zone ID")
}
