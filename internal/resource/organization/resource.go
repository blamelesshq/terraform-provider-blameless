package organization

import (
	"context"
	"log"
	"strings"

	"github.com/blamelesshq/terraform-provider/internal/config"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetResourceKey() string {
	return "blameless_organization"
}

func NewResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: createAction,
		ReadContext:   readAction,
		UpdateContext: updateAction,
		DeleteContext: deleteAction,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Description: "",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the action.",
			},
			"timezone": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the action.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "The name of the action.",
			},
			"incident_roles": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The name of the action.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"incident_severities": {
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

func createAction(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	orgSettings := expandSettings(d.GetRawConfig())
	if err := api.CreateOrgSettings(orgSettings); err != nil {
		log.Printf("create error: %+v", err)
		return diag.FromErr(err)
	}

	return readAction(ctx, d, m)
}

func readAction(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	orgSettings, err := api.GetOrgSettings()
	if err != nil {
		log.Printf("read error: %+v", err)
		return diag.FromErr(err)
	}

	result := multierror.Append(
		d.Set("name", orgSettings.Name),
		d.Set("timezone", orgSettings.Timezone),
		d.Set("description", orgSettings.Description),
		d.Set("incident_roles", strings.Join(orgSettings.IncidentRoles, ",")),
		d.Set("incident_severities", flattenIncidentSeverities(orgSettings.Severities)),
	)

	return diag.FromErr(result.ErrorOrNil())
}

func updateAction(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	orgSettings := expandSettings(d.GetRawConfig())
	if err := api.UpdateOrgSettings(orgSettings); err != nil {
		return diag.FromErr(err)
	}

	return readAction(ctx, d, m)
}

func deleteAction(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	if err := api.DeleteOrgSettings(); err != nil {
		return diag.FromErr(err)
	}

	return nil
}
