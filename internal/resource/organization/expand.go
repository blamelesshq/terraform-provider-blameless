package organization

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expandSettings(config cty.Value) *model.OrgSettings {
	settings := &model.OrgSettings{
		Name:     *value.String(config.GetAttr("name")),
		Timezone: *value.String(config.GetAttr("timezone")),
	}
	return settings
}
