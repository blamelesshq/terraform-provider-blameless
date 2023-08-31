package service

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetIncidentRoleSettings() (*model.IncidentRoleSettings, error) {
	return getSettings[model.IncidentRoleSettings](s, sectionIncidentRole)
}

func (s *Svc) UpdateIncidentRoleSettings(settings *model.IncidentRoleSettings) error {
	return updateSettings[model.IncidentRoleSettings](s, sectionIncidentRole, settings)
}
