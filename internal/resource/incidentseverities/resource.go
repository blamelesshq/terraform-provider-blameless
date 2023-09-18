package incidentseverities

import (
	"context"
	"fmt"

	"github.com/blamelesshq/terraform-provider/internal/config"
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func GetResourceKey() string {
	return "blameless_incident_severity_settings"
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
		Description: "Incident Severities",
		Schema: map[string]*schema.Schema{
			"severities": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sev0_label": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Label for Severity 0.",
						},
						"sev1_label": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Label for Severity 1.",
						},
						"sev2_label": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Label for Severity 2.",
						},
						"sev3_label": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Label for Severity 3.",
						},
					},
				},
			},
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings := expandSettings(d.GetRawConfig())
	if err := api.UpdateIncidentSeveritySettings(ctx, settings); err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings, err := api.GetIncidentSeveritySettings(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	result := multierror.Append(
		d.Set("severities", flattenIncidentSeverities(settings)),
	)
	d.SetId(fmt.Sprint(len(settings.Severities)))
	return diag.FromErr(result.ErrorOrNil())
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings := expandSettings(d.GetRawConfig())
	if err := api.UpdateIncidentSeveritySettings(ctx, settings); err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	if err := api.UpdateIncidentSeveritySettings(ctx, &model.IncidentSeveritySettings{}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
