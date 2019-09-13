package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/isometry/terraform-provider-keychain/keychain"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: keychain.Provider,
	})
}
