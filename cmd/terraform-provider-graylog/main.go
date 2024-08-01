package main

import (
	"github.com/bmhughes/terraform-provider-graylog/graylog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: graylog.Provider,
	})
}
