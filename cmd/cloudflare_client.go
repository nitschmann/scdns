package cmd

import "github.com/nitschmann/scdns/cloudflare"

// TODO: Maybe this needes to be moved somewhere else?
func cloudflareClient() *cloudflare.Client {
	credentials := cloudflare.NewCredentials(email, authKey)
	client := cloudflare.NewClient(credentials)

	return client
}
