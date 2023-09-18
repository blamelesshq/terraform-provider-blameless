package incidenttypeseverity

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
	return "blameless_incident_type_severity"
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
		Description: "Incident Type Severity",
		Schema: map[string]*schema.Schema{
			"incident_type_id": {
				Type:         schema.TypeString,
				Required:     true,
				Optional:     false,
				ValidateFunc: validation.StringIsNotEmpty,
				Description:  "Id of the incident type.",
			},
			"incident_severity": {
				Type:        schema.TypeInt,
				Required:    true,
				Optional:    false,
				Description: "Severity of the incident.",
			},
			"incident_settings": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: false,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"end_of_customer_impact_status": {
							Type:         schema.TypeString,
							Required:     false,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Status that marks the end of the impact of the incident to the customer.",
						},
						"private_incident_channel": {
							Type:        schema.TypeBool,
							Required:    false,
							Optional:    true,
							Description: "Enables private channels when an incident is created.",
						},
						"channel_naming": {
							Type:     schema.TypeSet,
							Required: false,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"incident_naming_scheme": {
										Type:         schema.TypeString,
										Required:     false,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The trigger ID.",
									},
									"require_dash_separator": {
										Type:        schema.TypeBool,
										Required:    false,
										Optional:    true,
										Description: "The trigger ID.",
									},
									"custom_channel_format": {
										Type:         schema.TypeString,
										Required:     false,
										Optional:     true,
										ValidateFunc: validation.StringIsNotEmpty,
										Description:  "The trigger ID.",
									},
								},
							},
						},
						"team_notifications": {
							Type:     schema.TypeSet,
							Required: false,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_recruit_team_members": {
										Type:        schema.TypeList,
										Required:    false,
										Optional:    true,
										Description: "The trigger ID.",
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},
									"announcement_channels": {
										Type:        schema.TypeList,
										Required:    false,
										Optional:    true,
										Description: "The trigger ID.",
										Elem: &schema.Schema{
											Type:         schema.TypeString,
											ValidateFunc: validation.StringIsNotEmpty,
										},
									},
								},
							},
						},
					},
				},
			},
			"retrospective_settings": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: false,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"analysis_template": {
							Type:         schema.TypeString,
							Required:     false,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Status that marks the end of the impact of the incident to the customer.",
						},
						"questionnaire_template": {
							Type:         schema.TypeString,
							Required:     false,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Status that marks the end of the impact of the incident to the customer.",
						},
						"required": {
							Type:        schema.TypeBool,
							Required:    false,
							Optional:    true,
							Description: "Enables private channels when an incident is created.",
						},
						"daily_reminder": {
							Type:        schema.TypeBool,
							Required:    false,
							Optional:    true,
							Description: "Enables private channels when an incident is created.",
						},
						"incident_resolution_required": {
							Type:        schema.TypeBool,
							Required:    false,
							Optional:    true,
							Description: "Enables private channels when an incident is created.",
						},
					},
				},
			},
			"task_settings": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: false,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"full_permission_role": {
							Type:         schema.TypeString,
							Required:     false,
							Optional:     true,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Status that marks the end of the impact of the incident to the customer.",
						},
						"task_list": {
							Type:        schema.TypeSet,
							Required:    false,
							Optional:    true,
							Description: "Status that marks the end of the impact of the incident to the customer.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"investigating": {
										Type:        schema.TypeSet,
										Required:    false,
										Optional:    true,
										Description: "Status that marks the end of the impact of the incident to the customer.",
										Elem:        getTaskResourceSchema(),
									},
									"identified": {
										Type:        schema.TypeSet,
										Required:    false,
										Optional:    true,
										Description: "Status that marks the end of the impact of the incident to the customer.",
										Elem:        getTaskResourceSchema(),
									},
									"monitoring": {
										Type:        schema.TypeSet,
										Required:    false,
										Optional:    true,
										Description: "Status that marks the end of the impact of the incident to the customer.",
										Elem:        getTaskResourceSchema(),
									},
									"resolved": {
										Type:        schema.TypeSet,
										Required:    false,
										Optional:    true,
										Description: "Status that marks the end of the impact of the incident to the customer.",
										Elem:        getTaskResourceSchema(),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func getTaskResourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"task": {
				Type:     schema.TypeSet,
				Required: true,
				Optional: false,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							Required:     true,
							Optional:     false,
							ValidateFunc: validation.StringIsNotEmpty,
							Description:  "Status that marks the end of the impact of the incident to the customer.",
						},
						"role": {
							Type:        schema.TypeString,
							Required:    true,
							Optional:    false,
							Description: "Enables private channels when an incident is created.",
						},
						"required": {
							Type:        schema.TypeBool,
							Required:    true,
							Optional:    false,
							Description: "Enables private channels when an incident is created.",
						},
					},
				},
			},
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings := expandSettings(ctx, d.GetRawConfig())
	if err := api.UpdateIncidentTypeSeveritySettings(ctx, settings); err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%s//%d", settings.IncidentTypeId, settings.Severity))

	return read(ctx, d, m)
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings, err := api.GetIncidentTypeSeveritySettings(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	result := multierror.Append(
		d.Set("incident_type_id", settings.IncidentTypeId),
		d.Set("severity", settings.Severity),
		d.Set("incident_settings", flattenIncidentSettings(settings.IncidentSettings)),
		d.Set("retrospective_settings", flattenRetrospectiveSettings(settings.RetrospectiveSettings)),
		d.Set("task_settings", flattenTaskSettings(settings.TaskSettings)),
	)

	return diag.FromErr(result.ErrorOrNil())
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	settings := expandSettings(ctx, d.GetRawConfig())
	if err := api.UpdateIncidentTypeSeveritySettings(ctx, settings); err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	api := m.(*config.Config).GetAPI()

	if err := api.UpdateIncidentTypeSeveritySettings(ctx, &model.IncidentTypeSeverity{}); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
