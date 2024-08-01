package role

import (
	"github.com/bmhughes/terraform-provider-graylog/graylog/convert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	keyName     = "name"
	keyReadOnly = "read_only"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	delete(data, keyReadOnly)
	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}
	d.SetId(data[keyName].(string))
	return nil
}
