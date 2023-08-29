package incidentroles

import (
	"context"
	"log"
	"sort"
	"strings"

	"github.com/blamelesshq/terraform-provider/internal/config"
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetResourceKey() string {
	return "blameless_incident_role_settings"
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
			"roles": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of incident roles.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()
	settings := expandSettings(d.GetRawConfig())
	if err := api.UpdateIncidentRoleSettings(settings); err != nil {
		log.Printf("create error: %+v\n", err)
		return diag.FromErr(err)
	}
	sort.Strings(settings.Roles)
	d.SetId(strings.Join(settings.Roles, ","))
	return read(ctx, d, m)
}

func read(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()
	settings, err := api.GetIncidentRoleSettings()
	if err != nil {
		log.Printf("read error: %+v\n", err)
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
	if err := api.UpdateIncidentRoleSettings(settings); err != nil {
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
