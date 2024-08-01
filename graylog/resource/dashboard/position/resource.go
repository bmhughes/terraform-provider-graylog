package position

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
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			// Required
			"dashboard_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"positions": {
				Required:         true,
				Type:             schema.TypeString,
				DiffSuppressFunc: util.SchemaDiffSuppressJSONString,
				ValidateFunc:     util.ValidateIsMapJSON,
			},
		},
	}
}
