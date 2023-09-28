package incidentseverities

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
)

func flattenIncidentSeverities(severities *model.IncidentSeveritySettings) []interface{} {
	result := []interface{}{}
	for _, severity := range severities.Severities {
		result = append(result, map[string]interface{}{
			"level": severity.Level,
			"label": severity.Label,
		})
	}
	return result
}
