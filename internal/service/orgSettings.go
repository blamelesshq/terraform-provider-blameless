package service

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetOrgSettings(ctx context.Context) (*model.OrgSettings, error) {
	return getSettings[model.OrgSettings](ctx, s, sectionOrg)
}

func (s *Svc) UpdateOrgSettings(ctx context.Context, settings *model.OrgSettings) error {
	return updateSettings[model.OrgSettings](ctx, s, sectionOrg, settings)
}
