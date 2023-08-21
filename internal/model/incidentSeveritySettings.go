package model

type IncidentSeveritySettings struct {
	Severities []*IncidentSeverity `json:"severities"`
}

type IncidentSeverity struct {
	Level int    `json:"level"`
	Label string `json:"label"`
}
