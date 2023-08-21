package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
)

type Service interface {
	GetOrgSettings() (*model.OrgSettings, error)
	UpdateOrgSettings(settings *model.OrgSettings) error

	GetIncidentRoleSettings() (*model.IncidentRoleSettings, error)
	UpdateIncidentRoleSettings(settings *model.IncidentRoleSettings) error

	GetIncidentSeveritySettings() (*model.IncidentSeveritySettings, error)
	UpdateIncidentSeveritySettings(settings *model.IncidentSeveritySettings) error
}

type Svc struct {
	key      string
	instance string
	client   *retryablehttp.Client
	token    *string
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

func getSettings[TResponse interface{}](svc *Svc, section string) (*TResponse, error) {
	return callSettings[struct{}, TResponse](svc, section, http.MethodGet, nil)
}

// TODO uncomment for incident types
// func createSettings[TRequest interface{}](svc *Svc, section string, req *TRequest) error {
// 	_, err := callSettings[TRequest, struct{}](svc, section, http.MethodPost, req)
// 	return err
// }

func updateSettings[TRequest interface{}](svc *Svc, section string, req *TRequest) error {
	_, err := callSettings[TRequest, struct{}](svc, section, http.MethodPut, req)
	return err
}

// TODO uncomment for incident types
// func deleteSettings(svc *Svc, section string) error {
// 	_, err := callSettings[struct{}, struct{}](svc, section, http.MethodDelete, nil)
// 	return err
// }

func callSettings[TRequest interface{}, TResponse interface{}](svc *Svc, section string, method string, req *TRequest) (*TResponse, error) {
	target := fmt.Sprintf("%s/api/v2/settings/%s", svc.Instance(), section)
	var payload interface{} = nil
	if req != nil {
		r, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		payload = bytes.NewReader(r)
	}

	request, err := retryablehttp.NewRequest(method, target, payload)
	if err != nil {
		log.Printf("new request error: %+v", err)
		return nil, err
	}
	token, err := svc.authToken()
	if err != nil {
		log.Printf("auth token error: %+v", err)
		return nil, err
	}
	request.Header.Add("Authorization", *token)
	request.Header.Add("User-Agent", userAgent())

	resp, err := svc.Client().Do(request)
	if err != nil {
		log.Printf("do request error: %+v", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read body error: %+v", err)
		return nil, err
	}

	if len(body) > 0 {
		var response TResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			log.Printf("json unmarshal error: %+v", err)
			return nil, err
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
