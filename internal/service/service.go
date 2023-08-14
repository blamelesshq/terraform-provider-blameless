package service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
)

type Service interface{}

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

func (s *Svc) authToken() (*string, error) {
	if s.token == nil {
		target := fmt.Sprintf("%s/api/v2/identity/token", s.instance)
		request, err := retryablehttp.NewRequest("POST", target, nil)
		if err != nil {
			return nil, err
		}
		request.Header.Add("Authorization", s.key)
		request.Header.Add("User-Agent", userAgent())
		resp, err := s.client.Do(request)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var response tokenResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
		s.token = &response.AccessToken
	}
	return s.token, nil
}

func userAgent() string {
	terraformSDKVersion := meta.SDKVersionString()

	return fmt.Sprintf(
		"Terraform-Provider-Blameless/dev (Terraform-SDK/%s)",
		terraformSDKVersion,
	)
}
