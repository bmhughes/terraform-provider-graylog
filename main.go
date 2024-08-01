package main

import (
	"flag"

	"github.com/bmhughes/terraform-provider-graylog/graylog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := &plugin.ServeOpts{
		Debug:        debug,
		ProviderAddr: "app.terraform.io/bmhughes-net/graylog",
		ProviderFunc: graylog.Provider,
	}

	plugin.Serve(opts)
}
