package cmd

import (
	"log"
	"os"

	"github.com/nitschmann/scdns/internal"

	"github.com/spf13/cobra"
)

var (
	authKey          string
	cfgFile          string
	cloudflareConfig *internal.CloudflareConfig
	email            string
)

var rootCmd = &cobra.Command{
	Use:   "scdns",
	Short: "Simple Cloudflare DNS",
	Long:  "Simple management for Cloudflare DNS entries via CLI",
	Run:   executeRootCmd,
}

func init() {
	initGlobalCommandFlags()
	initCloudflareConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func executeRootCmd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(1)
	}
}

func initCloudflareConfig() {
	// TODO: Change this assignment if config when config is extended
	if authKey != "" && email != "" {
		cloudflareConfig = &internal.CloudflareConfig{
			AuthKey: authKey,
			Email:   email,
		}
	} else {
		config, err := internal.LoadCloudflareConfig(cfgFile)
		if err != nil {
			log.Fatalf("An error occured while loading Cloudflare config: %s\n", err)
		}

		cloudflareConfig = config
		authKey = cloudflareConfig.AuthKey
		email = cloudflareConfig.Email
	}
}

func initGlobalCommandFlags() {
	rootCmd.PersistentFlags().StringVarP(&authKey, "auth-key", "", "", "Cloudflare API AuthKey (Note: ignores authkey from config file if given)")
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file for Cloudflare API (default: /etc/cloudflare/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&email, "email", "", "", "Cloudflare API Email (Note: ignores email from config file if given)")
}
