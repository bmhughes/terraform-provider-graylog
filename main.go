package main

import (
	"flag"

	"github.com/bmhughes/terraform-provider-graylog/graylog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	opts := &plugin.ServeOpts{
		ProviderAddr: "app.terraform.io/bmhughes-net/graylog",
		ProviderFunc: graylog.Provider,
	}

	flag.BoolVar(&opts.Debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(opts)
}
