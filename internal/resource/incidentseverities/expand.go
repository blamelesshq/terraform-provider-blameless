package incidentseverities

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expandSettings(config cty.Value) *model.IncidentSeveritySettings {
	settings := &model.IncidentSeveritySettings{
		Severities: expandIncidentSeverties(config.GetAttr("severities")),
	}
	return settings
}

func expandIncidentSeverties(severities cty.Value) []*model.IncidentSeverity {
	var results []*model.IncidentSeverity
	severities.ForEachElement(func(_, sev cty.Value) (stop bool) {
		results = []*model.IncidentSeverity{
			{
				Level: 0,
				Label: *value.String(sev.GetAttr("sev0_label")),
			},
			{
				Level: 1,
				Label: *value.String(sev.GetAttr("sev1_label")),
			},
			{
				Level: 2,
				Label: *value.String(sev.GetAttr("sev2_label")),
			},
			{
				Level: 3,
				Label: *value.String(sev.GetAttr("sev3_label")),
			},
		}
		return stop
	})

	return results
}
