package incidentseverities

import (
	"context"
	"fmt"
	"log"

	"github.com/blamelesshq/terraform-provider/internal/config"
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		Description: "",
		Schema: map[string]*schema.Schema{
			"severities": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sev0_label": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The trigger ID.",
						},
						"sev1_label": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The trigger ID.",
						},
						"sev2_label": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The trigger ID.",
						},
						"sev3_label": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The trigger ID.",
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
	if err := api.UpdateIncidentSeveritySettings(settings); err != nil {
		log.Printf("create error: %+v", err)
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func read(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings, err := api.GetIncidentSeveritySettings()
	if err != nil {
		log.Printf("read error: %+v", err)
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
	if err := api.UpdateIncidentSeveritySettings(settings); err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func delete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	if err := api.UpdateIncidentSeveritySettings(&model.IncidentSeveritySettings{}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
