package graylog

import (
	"github.com/bmhughes/terraform-provider-graylog/graylog/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform resource provider for graylog.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema:         provider.SchemaMap(),
		ResourcesMap:   resourceMap,
		DataSourcesMap: dataSourcesMap,
		ConfigureFunc:  provider.Configure,
	}
}
