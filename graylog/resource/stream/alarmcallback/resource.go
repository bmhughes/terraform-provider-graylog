package alarmcallback

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
			StateContext: util.GenStateFunc("stream_id", "alarm_callback_id"),
		},

		Schema: map[string]*schema.Schema{
			// Required
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"stream_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"configuration": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsMapJSON,
			},
			"alarmcallback_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
