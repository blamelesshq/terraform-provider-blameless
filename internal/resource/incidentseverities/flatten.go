package incidentseverities

import (
	"fmt"

	"github.com/blamelesshq/terraform-provider/internal/model"
)

func flattenIncidentSeverities(severities *model.IncidentSeveritySettings) []interface{} {
	result := make(map[string]string)
	for _, severity := range severities.Severities {
		result[fmt.Sprintf("sev%d_label", severity.Level)] = severity.Label
	}
	return []interface{}{result}
}
