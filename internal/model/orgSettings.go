package model

type OrgSettings struct {
	Name          string
	Timezone      string
	Description   string
	IncidentRoles []string
	Severities    []*IncidentSeverity
}

type IncidentSeverity struct {
	Level int
	Label string
}
