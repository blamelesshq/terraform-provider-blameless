package service

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
)

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (s *Svc) authToken() (*string, error) {
	if s.token == nil {
		s.mu.Lock()
		defer s.mu.Unlock()
		if s.token != nil {
			return s.token, nil
		}

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
