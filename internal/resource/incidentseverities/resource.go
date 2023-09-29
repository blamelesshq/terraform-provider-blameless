package incidentseverities

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
	return "blameless_incident_severities"
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
			"severity": {
				Type:     schema.TypeList,
				MinItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(0),
							Description:  "Level of Severity.",
						},
						"label": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Label for Severity.",
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
	d.SetId(id.UniqueId())

	return read(ctx, d, m)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings, err := api.GetIncidentSeveritySettings(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	result := multierror.Append(
		d.Set("severity", flattenIncidentSeverities(settings)),
	)

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

	if err := api.UpdateIncidentSeveritySettings(ctx, &model.IncidentSeveritySettings{Severities: []*model.IncidentSeverity{}}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
