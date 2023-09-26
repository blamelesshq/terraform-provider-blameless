package config

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/service"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
		instanceVal, ok := data.GetOk("instance")
		if !ok {
			return nil, diag.Errorf("no instance provided")
		}

		keyVal, ok := data.GetOk("key")
		if !ok {
			return nil, diag.Errorf("no key provided")
		}

		key := keyVal.(string)
		instance := instanceVal.(string)

		tflog.Debug(ctx, "configured", map[string]interface{}{"instance": instance, "key": key})

		client := service.New(key, instance)

		return New(client), nil
	}
}
