package service

import (
	"context"
	"fmt"

	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetIncidentTypeSettings(ctx context.Context, id string) (*model.IncidentTypeSettings, error) {
	return getSettings[model.IncidentTypeSettings](ctx, s, fmt.Sprintf("%s/%s", sectionIncidentType, id))
}

func (s *Svc) CreateIncidentTypeSettings(ctx context.Context, settings *model.IncidentTypeSettings) (string, error) {
	resp, err := createSettings[model.IncidentTypeSettings, model.IncidentTypeSettings](ctx, s, sectionIncidentType, settings)
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

func (s *Svc) UpdateIncidentTypeSettings(ctx context.Context, id string, settings *model.IncidentTypeSettings) error {
	return updateSettings[model.IncidentTypeSettings](ctx, s, fmt.Sprintf("%s/%s", sectionIncidentType, id), settings)
}

func (s *Svc) DeleteIncidentTypeSettings(ctx context.Context, id string) error {
	return deleteSettings[string](ctx, s, sectionIncidentType, id)
}
