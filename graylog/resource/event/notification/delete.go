package notification

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
	if _, err := cl.EventNotification.Delete(ctx, d.Id()); err != nil {
		return fmt.Errorf("failed to delete a event notification %s: %w", d.Id(), err)
	}
	return nil
}
