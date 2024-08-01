package extractor

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

	inputID := d.Get(keyInputID).(string)
	eID := d.Get(keyExtractorID).(string)
	if _, err := cl.Extractor.Delete(ctx, inputID, eID); err != nil {
		return fmt.Errorf("failed to delete a extractor (input id: %s, id: %s): %w", inputID, eID, err)
	}
	return nil
}
