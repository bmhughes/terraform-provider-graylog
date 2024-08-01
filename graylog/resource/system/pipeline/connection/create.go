package connection

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

	sID := data[keyStreamID].(string)
	if _, err := cl.PipelineConnection.ConnectPipelinesToStream(ctx, data); err != nil {
		return fmt.Errorf("failed to connect pipelines to a stream (stream id: %s): %w", sID, err)
	}
	d.SetId(sID)
	return nil
}
