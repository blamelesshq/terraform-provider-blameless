package service

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetIncidentTypeSeveritySettings(ctx context.Context) (*model.IncidentTypeSeverity, error) {
	return getSettings[model.IncidentTypeSeverity](ctx, s, sectionIncidentTypeSeverity)
}

func (s *Svc) UpdateIncidentTypeSeveritySettings(ctx context.Context, settings *model.IncidentTypeSeverity) error {
	return updateSettings[model.IncidentTypeSeverity](ctx, s, sectionIncidentTypeSeverity, settings)
}
