package lookuptable

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"content_pack": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_adapter_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cache_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_single_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_single_value_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "NULL",
				ValidateFunc: validation.StringInSlice([]string{"STRING", "NUMBER", "OBJECT", "BOOLEAN", "NULL"}, false),
			},
			"default_multi_value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_multi_value_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "NULL",
				ValidateFunc: validation.StringInSlice([]string{"STRING", "NUMBER", "OBJECT", "BOOLEAN", "NULL"}, false),
			},
		},
	}
}
