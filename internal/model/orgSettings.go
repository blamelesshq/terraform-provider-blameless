package model

type OrgSettings struct {
	Name          string              `json:"name"`
	Timezone      string              `json:"timezone"`
	Description   string              `json:"description"`
	IncidentRoles []string            `json:"incidentRoles"`
	Severities    []*IncidentSeverity `json:"severities"`
}

type IncidentSeverity struct {
	Level int    `json:"level"`
	Label string `json:"label"`
}
