package service

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetIncidentSeveritySettings() (*model.IncidentSeveritySettings, error) {
	return getSettings[model.IncidentSeveritySettings](s, sectionIncidentSeverity)
}

func (s *Svc) UpdateIncidentSeveritySettings(settings *model.IncidentSeveritySettings) error {
	return updateSettings[model.IncidentSeveritySettings](s, sectionIncidentSeverity, settings)
}
