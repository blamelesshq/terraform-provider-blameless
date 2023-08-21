package incidentroles

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expandSettings(config cty.Value) *model.IncidentRoleSettings {
	settings := &model.IncidentRoleSettings{
		IncidentRoles: expandIncidentRoles(config.GetAttr("roles")),
	}
	return settings
}

func expandIncidentRoles(roles cty.Value) []string {
	incidentRoles := make([]string, roles.LengthInt())
	if roles.IsNull() {
		return nil
	}
	roles.ForEachElement(func(key, val cty.Value) (stop bool) {
		incidentRoles = append(incidentRoles, value.String(val))
		return stop
	})
	return incidentRoles
}
