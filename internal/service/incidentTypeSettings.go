package service

import (
	"context"
	"fmt"

	"github.com/blamelesshq/terraform-provider/internal/model"
)

func (s *Svc) GetIncidentTypeSettings(ctx context.Context, id string) (*model.IncidentType, error) {
	return getSettings[model.IncidentType](ctx, s, fmt.Sprintf("%s/%s", sectionIncidentType, id))
}

func (s *Svc) CreateIncidentTypeSettings(ctx context.Context, settings *model.IncidentType) (string, error) {
	resp, err := createSettings[model.IncidentType, model.IncidentType](ctx, s, sectionIncidentType, settings)
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}

func (s *Svc) UpdateIncidentTypeSettings(ctx context.Context, id string, settings *model.IncidentType) error {
	return updateSettings[model.IncidentType](ctx, s, fmt.Sprintf("%s/%s", sectionIncidentType, id), settings)
}

func (s *Svc) DeleteIncidentTypeSettings(ctx context.Context, id string) error {
	return deleteSettings[string](ctx, s, sectionIncidentType, id)
}
