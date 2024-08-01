package connection

import (
	"context"
	"fmt"

	"github.com/bmhughes/terraform-provider-graylog/graylog/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func destroy(d *schema.ResourceData, m interface{}) error {
	ctx := context.Background()
	cl, err := client.New(m)
	if err != nil {
		return err
	}

	if _, err := cl.PipelineConnection.ConnectPipelinesToStream(ctx, map[string]interface{}{
		"stream_id":    d.Id(),
		"pipeline_ids": []string{},
	}); err != nil {
		return fmt.Errorf(
			"failed to delete connections between a stream and pipelines (stream id: %s): %w", d.Id(), err)
	}
	return nil
}
