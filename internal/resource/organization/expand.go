package organization

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expandSettings(config cty.Value) *model.OrgSettings {
	settings := &model.OrgSettings{
		Name:          value.String(config.GetAttr("name")),
		Timezone:      value.String(config.GetAttr("timezone")),
		Description:   value.String(config.GetAttr("description")),
		IncidentRoles: expandIncidentRoles(config.GetAttr("incident_roles")),
		Severities:    expandIncidentSeverties(config.GetAttr("incident_severities")),
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

func expandIncidentSeverties(severities cty.Value) []*model.IncidentSeverity {
	var results []*model.IncidentSeverity
	severities.ForEachElement(func(_, sev cty.Value) (stop bool) {
		results = []*model.IncidentSeverity{
			{
				Level: 0,
				Label: value.String(sev.GetAttr("sev0_label")),
			},
			{
				Level: 1,
				Label: value.String(sev.GetAttr("sev1_label")),
			},
			{
				Level: 2,
				Label: value.String(sev.GetAttr("sev2_label")),
			},
			{
				Level: 3,
				Label: value.String(sev.GetAttr("sev3_label")),
			},
		}
		return stop
	})

	return results
}
