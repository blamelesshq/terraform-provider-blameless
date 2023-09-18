package blameless

import (
	"github.com/blamelesshq/terraform-provider/internal/config"
	"github.com/blamelesshq/terraform-provider/internal/resource/incidentroles"
	"github.com/blamelesshq/terraform-provider/internal/resource/incidentseverities"
	"github.com/blamelesshq/terraform-provider/internal/resource/incidenttype"
	"github.com/blamelesshq/terraform-provider/internal/resource/incidenttypeseverity"
	"github.com/blamelesshq/terraform-provider/internal/resource/organization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"instance": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("BLAMELESS_INSTANCE", nil),
				Description: "Your Blameless instance URL. " +
					"It can also be sourced from the `BLAMELESS_INSTANCE` environment variable.",
			},
			"key": {
				Type:        schema.TypeString,
				Optional:    false,
				Required:    true,
				Computed:    false,
				DefaultFunc: schema.EnvDefaultFunc("BLAMELESS_KEY", nil),
				Description: "Your Blameless API key. " +
					"It can also be sourced from the `BLAMELESS_KEY` environment variable.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			organization.GetResourceKey():         organization.NewResource(),
			incidentseverities.GetResourceKey():   incidentseverities.NewResource(),
			incidentroles.GetResourceKey():        incidentroles.NewResource(),
			incidenttype.GetResourceKey():         incidenttype.NewResource(),
			incidenttypeseverity.GetResourceKey(): incidenttypeseverity.NewResource(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
	provider.ConfigureContextFunc = config.ConfigureProvider(&provider.TerraformVersion)
	return provider
}
