package output

import (
	"context"
	"fmt"

	"github.com/bmhughes/terraform-provider-graylog/graylog/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func create(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}
	data, err := getDataFromResourceData(d)
	if err != nil {
		return err
	}

	output, _, err := cl.Output.Create(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to create a output: %w", err)
	}
	d.SetId(output[keyID].(string))
	return nil
}
