package lookupdataadapter

import (
	"github.com/bmhughes/terraform-provider-graylog/graylog/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Resource() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: destroy,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_error_ttl_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"custom_error_ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"custom_error_ttl_unit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"content_pack": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},
		},
	}
}
