package incidentseverities

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expandSettings(config cty.Value) *model.IncidentSeveritySettings {
	settings := &model.IncidentSeveritySettings{
		Severities: expandIncidentSeverties(config.GetAttr("severity")),
	}
	return settings
}

func expandIncidentSeverties(severities cty.Value) []*model.IncidentSeverity {
	var results []*model.IncidentSeverity
	severities.ForEachElement(func(_, sev cty.Value) (stop bool) {
		results = append(results, &model.IncidentSeverity{
			Level: *value.Int(sev.GetAttr("level")),
			Label: *value.String(sev.GetAttr("label")),
		})
		return stop
	})

	return results
}
