package internal

import (
	"github.com/spf13/viper"
)

type CloudflareConfig struct {
	AuthKey string
	Email   string
}

func LoadCloudflareConfig(cfgFile string) (*CloudflareConfig, error) {
	handler := viper.New()

	if cfgFile != "" {
		handler.SetConfigFile(cfgFile)
	} else {
		handler.SetConfigName("config")
		handler.AddConfigPath("/etc/cloudflare")
		handler.AddConfigPath("$HOME/.cloudflare")
		handler.AddConfigPath(".")
	}

	err := handler.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config *CloudflareConfig = &CloudflareConfig{}
	handler.Unmarshal(config)

	return config, nil
}
