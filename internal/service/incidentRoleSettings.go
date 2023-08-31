package service

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetIncidentRoleSettings(ctx context.Context) (*model.IncidentRoleSettings, error) {
	return getSettings[model.IncidentRoleSettings](ctx, s, sectionIncidentRole)
}

func (s *Svc) UpdateIncidentRoleSettings(ctx context.Context, settings *model.IncidentRoleSettings) error {
	return updateSettings[model.IncidentRoleSettings](ctx, s, sectionIncidentRole, settings)
}
