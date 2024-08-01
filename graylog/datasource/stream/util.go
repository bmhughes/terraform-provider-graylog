package stream

import (
	"github.com/bmhughes/terraform-provider-graylog/graylog/convert"
	"github.com/bmhughes/terraform-provider-graylog/graylog/resource/stream"
	"github.com/bmhughes/terraform-provider-graylog/graylog/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	id, ok := util.RenameKey(data, "id", "stream_id")

	if err := convert.SetResourceData(d, stream.Resource(), data); err != nil {
		return err
	}

	if ok {
		d.SetId(id.(string))
	}
	return nil
}
