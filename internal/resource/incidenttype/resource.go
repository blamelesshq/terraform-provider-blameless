package incidenttype

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/config"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func GetResourceKey() string {
	return "blameless_incident_type"
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
		Description: "Incident Type",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Name of incident type.",
			},
			"active": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Active/Inactive",
			},
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings := expand(d.GetRawConfig())
	id, err := api.CreateIncidentTypeSettings(ctx, settings)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)

	return read(ctx, d, m)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings, err := api.GetIncidentTypeSettings(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	result := multierror.Append(
		d.Set("name", settings.Name),
		d.Set("active", settings.Active),
	)

	return diag.FromErr(result.ErrorOrNil())
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()
	settings := expand(d.GetRawConfig())
	if err := api.UpdateIncidentTypeSettings(ctx, d.Id(), settings); err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	if err := api.DeleteIncidentTypeSettings(ctx, d.Id()); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
