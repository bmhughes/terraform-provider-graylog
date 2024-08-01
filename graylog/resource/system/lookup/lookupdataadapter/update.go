package lookupdataadapter

import (
	"context"
	"fmt"

	"github.com/bmhughes/terraform-provider-graylog/graylog/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func update(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}
	if _, _, err := cl.LookupDataAdapter.Update(ctx, d.Id(), data); err != nil {
		return fmt.Errorf("failed to update a lookup data adapter %s: %w", d.Id(), err)
	}
	return nil
}
