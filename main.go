package main

import (
	"github.com/bweir/terraform-provider-netlify/netlify"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: netlify.Provider})
}
