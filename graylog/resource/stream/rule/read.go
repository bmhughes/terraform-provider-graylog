package rule

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

	sID := d.Get(keyStreamID).(string)
	rID := d.Get(keyRuleID).(string)
	data, resp, err := cl.StreamRule.Get(ctx, sID, rID)
	if err != nil {
		return util.HandleGetResourceError(
			d, resp, fmt.Errorf(
				"failed to get a stream rule (stream id: %s, rule id: %s): %w", sID, rID, err))
	}
	return setDataToResourceData(d, data)
}
