package incidentroles

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expandSettings(config cty.Value) *model.IncidentRoleSettings {
	settings := &model.IncidentRoleSettings{
		Roles: value.StringArray(config.GetAttr("roles")),
	}
	return settings
}
