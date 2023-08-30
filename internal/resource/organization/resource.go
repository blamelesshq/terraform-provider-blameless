package organization

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/config"
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func GetResourceKey() string {
	return "blameless_org_settings"
}

func NewResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: "Organization Settings",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description: "The name of the organization.",
			},
			"timezone": {
				Type:        schema.TypeString,
				Required:    true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description: "Timezone specifier",
			},
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	orgSettings := expandSettings(d.GetRawConfig())
	if err := api.UpdateOrgSettings(orgSettings); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(orgSettings.Name)

	return read(ctx, d, m)
}

func read(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	orgSettings, err := api.GetOrgSettings()
	if err != nil {
		return diag.FromErr(err)
	}

	result := multierror.Append(
		d.Set("name", orgSettings.Name),
		d.Set("timezone", orgSettings.Timezone),
	)

	return diag.FromErr(result.ErrorOrNil())
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	orgSettings := expandSettings(d.GetRawConfig())
	if err := api.UpdateOrgSettings(orgSettings); err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func delete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	if err := api.UpdateIncidentRoleSettings(&model.IncidentRoleSettings{}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
