package widget

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

		SchemaVersion:  schemaVersion,
		StateUpgraders: stateUpgraders,

		Importer: &schema.ResourceImporter{
			StateContext: util.GenStateFunc("dashboard_id", "widget_id"),
		},

		Schema: map[string]*schema.Schema{
			// Required
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dashboard_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},

			"config": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsJSON,
			},

			// Optional
			"cache_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"creator_user_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"widget_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
