package main

import (
	"github.com/lukekalbfleisch/terraform-provider-kind/kind"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kind.Provider})
}
