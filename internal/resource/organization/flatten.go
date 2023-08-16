package organization

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
)

func flattenIncidentSeverities(severities []*model.IncidentSeverity) []interface{} {
	var result []interface{}
	for _, severity := range severities {
		result = append(result, map[string]interface{}{
			"level": severity.Level,
			"label": severity.Label,
		})
	}
	return result
}
