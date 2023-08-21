package service

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetOrgSettings() (*model.OrgSettings, error) {
	return getSettings[model.OrgSettings](s, sectionOrg)
}

func (s *Svc) UpdateOrgSettings(settings *model.OrgSettings) error {
	return updateSettings[model.OrgSettings](s, sectionOrg, settings)
}
