package incidenttype

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
			"severity0_settings": {
				Type:        schema.TypeSet,
				MaxItems:    1,
				Required:    false,
				Optional:    true,
				Description: "Severity 0 configuration",
				Elem:        getIncidentSeverityResource(),
			},
			"severity1_settings": {
				Type:        schema.TypeSet,
				MaxItems:    1,
				Required:    false,
				Optional:    true,
				Description: "Severity 1 configuration",
				Elem:        getIncidentSeverityResource(),
			},
			"severity2_settings": {
				Type:        schema.TypeSet,
				MaxItems:    1,
				Required:    false,
				Optional:    true,
				Description: "Severity 2 configuration",
				Elem:        getIncidentSeverityResource(),
			},
			"severity3_settings": {
				Type:        schema.TypeSet,
				MaxItems:    1,
				Required:    false,
				Optional:    true,
				Description: "Severity 3 configuration",
				Elem:        getIncidentSeverityResource(),
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

	var sev0 *model.IncidentTypeSeverity
	var sev1 *model.IncidentTypeSeverity
	var sev2 *model.IncidentTypeSeverity
	var sev3 *model.IncidentTypeSeverity

	for _, setting := range settings.Severities {
		if setting.Severity == nil {
			continue
		}

		switch *setting.Severity {
		case 0:
			sev0 = setting
		case 1:
			sev1 = setting
		case 2:
			sev2 = setting
		case 3:
			sev3 = setting
		}
	}

	result := multierror.Append(
		d.Set("name", settings.Name),
		d.Set("active", settings.Active),
		d.Set("severity0_settings", flattenIncidentSeverity(sev0)),
		d.Set("severity1_settings", flattenIncidentSeverity(sev1)),
		d.Set("severity2_settings", flattenIncidentSeverity(sev2)),
		d.Set("severity3_settings", flattenIncidentSeverity(sev3)),
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
