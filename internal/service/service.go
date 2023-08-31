package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
)

type Service interface {
	GetOrgSettings(ctx context.Context) (*model.OrgSettings, error)
	UpdateOrgSettings(ctx context.Context, settings *model.OrgSettings) error

	GetIncidentRoleSettings(ctx context.Context) (*model.IncidentRoleSettings, error)
	UpdateIncidentRoleSettings(ctx context.Context, settings *model.IncidentRoleSettings) error

	GetIncidentSeveritySettings(ctx context.Context) (*model.IncidentSeveritySettings, error)
	UpdateIncidentSeveritySettings(ctx context.Context, settings *model.IncidentSeveritySettings) error
}

type Svc struct {
	key      string
	instance string
	client   *retryablehttp.Client
	token    *string
	mu       sync.Mutex
}

func New(key, instance string) Service {
	return &Svc{
		key:      key,
		instance: instance,
		client:   retryablehttp.NewClient(),
	}
}

func (s *Svc) Instance() string {
	return s.instance
}

func (s *Svc) Client() *retryablehttp.Client {
	return s.client
}

func getSettings[TResponse interface{}](ctx context.Context, svc *Svc, section string) (*TResponse, error) {
	return callSettings[struct{}, TResponse](ctx, svc, section, http.MethodGet, nil)
}

// TODO uncomment for incident types
// func createSettings[TRequest interface{}](ctx context.Context, svc *Svc, section string, req *TRequest) error {
// 	_, err := callSettings[TRequest, struct{}](ctx, svc, section, http.MethodPost, req)
// 	return err
// }

func updateSettings[TRequest interface{}](ctx context.Context, svc *Svc, section string, req *TRequest) error {
	_, err := callSettings[TRequest, struct{}](ctx, svc, section, http.MethodPut, req)
	return err
}

// TODO uncomment for incident types
// func deleteSettings(ctx context.Context, svc *Svc, section string) error {
// 	_, err := callSettings[struct{}, struct{}](ctx, svc, section, http.MethodDelete, nil)
// 	return err
// }

func callSettings[TRequest interface{}, TResponse interface{}](ctx context.Context, svc *Svc, section string, method string, req *TRequest) (*TResponse, error) {
	target := fmt.Sprintf("%s/api/v2/settings/%s", svc.Instance(), section)

	var payload interface{} = nil
	r := ""
	if req != nil {
		r, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		payload = bytes.NewReader(r)
	}

	request, err := retryablehttp.NewRequest(method, target, payload)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("new request error: %+v", err), map[string]interface{}{ "method": method, "target": target, "payload": fmt.Sprint(r)})
		return nil, fmt.Errorf("internal service error. code: 1")
	}
	token, err := svc.authToken()
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("auth token error: %+v", err))
		return nil, fmt.Errorf("internal service error. code: 2")
	}
	request.Header.Add("Authorization", *token)
	request.Header.Add("User-Agent", userAgent())

	resp, err := svc.Client().Do(request)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("do request error: %+v", err), map[string]interface{}{ "method": method, "target": target, "payload": fmt.Sprint(r)})
		return nil, fmt.Errorf("internal service error. code: 3")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		tflog.Debug(ctx, fmt.Sprintf("read body error: %+v", err), map[string]interface{}{ "method": method, "target": target, "payload": fmt.Sprint(r)})
		return nil, fmt.Errorf("internal service error. code: 4")
	}

	if len(body) > 0 {
		var response TResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			tflog.Debug(ctx, fmt.Sprintf("json unmarshal error: %+v", err), map[string]interface{}{"response body": string(body)})
			return nil, fmt.Errorf("internal service error. code: 5")
		}
		return &response, nil
	}

	return nil, nil
}

func userAgent() string {
	terraformSDKVersion := meta.SDKVersionString()

	return fmt.Sprintf(
		"Terraform-Provider-Blameless/dev (Terraform-SDK/%s)",
		terraformSDKVersion,
	)
}
