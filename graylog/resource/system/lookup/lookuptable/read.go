package lookuptable

import (
	"context"
	"fmt"

	"github.com/bmhughes/terraform-provider-graylog/graylog/client"
	"github.com/bmhughes/terraform-provider-graylog/graylog/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func read(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, resp, err := cl.LookupTable.Get(ctx, d.Id())
	if err != nil {
		return util.HandleGetResourceError(d, resp, fmt.Errorf("failed to get a lookup table %s: %w", d.Id(), err))
	}

	// data = data["lookup_tables"].(map[string]interface{})
	// data = data[0]
	return setDataToResourceData(d, data)
}
