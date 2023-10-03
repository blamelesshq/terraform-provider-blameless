package incidentroles

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/config"
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func GetResourceKey() string {
	return "blameless_incident_roles"
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
		Description: "Incident Roles",
		Schema: map[string]*schema.Schema{
			"roles": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "List of incident roles.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()
	settings := expandSettings(d.GetRawConfig())
	if err := api.UpdateIncidentRoleSettings(ctx, settings); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id.UniqueId())
	return read(ctx, d, m)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()
	settings, err := api.GetIncidentRoleSettings(ctx)
	if err != nil {
		return diag.FromErr(err)
	}
	result := multierror.Append(
		d.Set("roles", settings.Roles),
	)
	return diag.FromErr(result.ErrorOrNil())
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings := expandSettings(d.GetRawConfig())
	if err := api.UpdateIncidentRoleSettings(ctx, settings); err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	if err := api.UpdateIncidentRoleSettings(ctx, &model.IncidentRoleSettings{Roles: []string{}}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
