package rule

import (
	"github.com/bmhughes/terraform-provider-graylog/graylog/convert"
	"github.com/bmhughes/terraform-provider-graylog/graylog/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	keyStreamID     = "stream_id"
	keyRuleID       = "rule_id"
	keyID           = "id"
	keyStreamRuleID = "streamrule_id"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	util.RenameKey(data, keyID, keyRuleID)
	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(d.Get(keyStreamID).(string) + "/" + d.Get(keyRuleID).(string))
	return nil
}
