package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/raelix/terraform-provider-kind/kind"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: kind.Provider})
}
