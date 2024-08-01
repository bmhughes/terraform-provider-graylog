package graylog

import (
	"github.com/bmhughes/terraform-provider-graylog/graylog/datasource/dashboard"
	"github.com/bmhughes/terraform-provider-graylog/graylog/datasource/sidecar"
	"github.com/bmhughes/terraform-provider-graylog/graylog/datasource/stream"
	"github.com/bmhughes/terraform-provider-graylog/graylog/datasource/system/indices/indexset"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var dataSourcesMap = map[string]*schema.Resource{
	"graylog_dashboard": dashboard.DataSource(),
	"graylog_index_set": indexset.DataSource(),
	"graylog_sidecar":   sidecar.DataSource(),
	"graylog_stream":    stream.DataSource(),
}
