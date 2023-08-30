package service

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetIncidentSeveritySettings(ctx context.Context) (*model.IncidentSeveritySettings, error) {
	return getSettings[model.IncidentSeveritySettings](ctx, s, sectionIncidentSeverity)
}

func (s *Svc) UpdateIncidentSeveritySettings(ctx context.Context, settings *model.IncidentSeveritySettings) error {
	return updateSettings[model.IncidentSeveritySettings](ctx, s, sectionIncidentSeverity, settings)
}
