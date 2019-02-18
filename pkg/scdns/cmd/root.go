package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	authKey string
	cfgFile string
	email   string
)

func Execute() {
	cobra.OnInitialize(initConfig)

	err := newCmdRoot().Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scdns",
		Short: "Simple Cloudflare DNS",
		Long:  "Simple management for Cloudflare DNS entries via CLI",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(1)
			}
		},
	}

	cmd.PersistentFlags().StringP("auth-key", "", "", "Cloudflare API AuthKey (Note: ignores authkey from config file if given)")
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file for Cloudflare API (default: /etc/cloudflare/config.yaml)")
	cmd.PersistentFlags().StringP("email", "", "", "Cloudflare API Email (Note: ignores email from config file if given)")

	viper.BindPFlag("authKey", cmd.PersistentFlags().Lookup("auth-key"))
	viper.BindPFlag("email", cmd.PersistentFlags().Lookup("email"))

	cmd.AddCommand(newCmdPublicIp())
	cmd.AddCommand(newZonesCmd())

	return cmd
}

func initConfig() {
	authKey = viper.GetString("authKey")
	email = viper.GetString("email")

	// TODO: Update if configs change
	if authKey == "" || email == "" {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			viper.AddConfigPath(filepath.Join("/etc", "cloudflare"))
			viper.AddConfigPath(filepath.Join(home, ".cloudflare"))
			viper.AddConfigPath(".")
			viper.SetConfigName("config")
		}

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Can't read config:", err)
			os.Exit(1)
		}

		authKey = viper.GetString("authKey")
		email = viper.GetString("email")
	}
}
