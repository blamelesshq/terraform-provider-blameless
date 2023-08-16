package config

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/service"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type Config struct {
	api service.Service
}

func New(api service.Service) *Config {
	return &Config{
		api: api,
	}
}

func (c *Config) GetAPI() service.Service {
	return c.api
}

func ConfigureProvider(terraformVersion *string) schema.ConfigureContextFunc {
	return func(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
		key := data.Get("key").(string)
		instance := data.Get("instance").(string)

		client := service.New(key, instance)

		return New(client), nil
	}
}
