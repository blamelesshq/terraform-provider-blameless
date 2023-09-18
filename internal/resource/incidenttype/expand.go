package incidenttype

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expand(config cty.Value) *model.IncidentTypeSettings {
	settings := &model.IncidentTypeSettings{
		Name:   value.String(config.GetAttr("name")),
		Active: *value.Bool(config.GetAttr("active")),
	}
	return settings
}
