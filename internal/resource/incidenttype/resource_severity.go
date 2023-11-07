package incidenttype

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func getIncidentSeverityResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"severity": {
				Type:         schema.TypeInt,
				Required:     true,
				Optional:     false,
				ValidateFunc: validation.IntAtLeast(0),
				Description:  "Severity level.",
			},
			"end_of_customer_impact_status": {
				Type:         schema.TypeString,
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "Status that marks the end of the impact of the incident to the customer.",
			},
			"private_incident_channel": {
				Type:        schema.TypeBool,
				Required:    false,
				Optional:    true,
				Description: "Enables private channels when an incident is created.",
			},
			"incident_naming_scheme": {
				Type:         schema.TypeString,
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "Naming scheme for incidents.",
			},
			"require_dash_separator": {
				Type:        schema.TypeBool,
				Required:    false,
				Optional:    true,
				Description: "Require a dash separator in custom incident names.",
			},
			"custom_channel_format": {
				Type:         schema.TypeString,
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "Custom format for incident channel names.",
			},
			"slack_invited_users": {
				Type:        schema.TypeList,
				Required:    false,
				Optional:    true,
				Description: "Slack team members to automatically recruit into the incident channel.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
			"slack_announcement_channels": {
				Type:        schema.TypeList,
				Required:    false,
				Optional:    true,
				Description: "Slack channels to notify when an incident is created.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
			"teams_invited_users": {
				Type:        schema.TypeList,
				Required:    false,
				Optional:    true,
				Description: "MS Teams team members to automatically recruit into the incident channel.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
			"teams_announcement_group": {
				Type:        schema.TypeString,
				Required:    false,
				Optional:    true,
				Description: "MS Teams group to notify when an incident is created.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
			"teams_announcement_channels": {
				Type:        schema.TypeList,
				Required:    false,
				Optional:    true,
				Description: "MS Teams channels to notify when an incident is created.",
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotWhiteSpace,
				},
			},
			"retrospective_analysis_template": {
				Type:         schema.TypeString,
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "Markdown template for the retrospective analysis.",
			},
			"retrospective_questionnaire_template": {
				Type:         schema.TypeString,
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "JSON schema for the retrospective questionnaire.",
			},
			"retrospective_required": {
				Type:        schema.TypeBool,
				Required:    false,
				Optional:    true,
				Description: "Requires retrospectives.",
			},
			"retrospective_daily_reminder": {
				Type:        schema.TypeBool,
				Required:    false,
				Optional:    true,
				Description: "Enables daily reminder to complete the retrospective.",
			},
			"retrospective_incident_resolution_required": {
				Type:        schema.TypeBool,
				Required:    false,
				Optional:    true,
				Description: "Requires incident resolution for the retrospective to be completed.",
			},
			"tasks_full_permission_role": {
				Type:         schema.TypeString,
				Required:     false,
				Optional:     true,
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Description:  "Role with full permissions to tasks.",
			},
			"task_list": {
				Type:        schema.TypeSet,
				Required:    false,
				Optional:    true,
				MaxItems:    1,
				Description: "Tasks to complete at each status of an incident.",
				Elem:        getTaskResourceSchema(),
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
						"incident_status": {
							Type:         schema.TypeString,
							Required:     true,
							Optional:     false,
							ValidateFunc: validation.StringIsNotWhiteSpace,
							Description:  "Status to present the task.",
						},
						"name": {
							Type:         schema.TypeString,
							Required:     true,
							Optional:     false,
							ValidateFunc: validation.StringIsNotWhiteSpace,
							Description:  "Name of the task.",
						},
						"role": {
							Type:        schema.TypeString,
							Required:    true,
							Optional:    false,
							Description: "Role assigned to the task.",
						},
						"required": {
							Type:        schema.TypeBool,
							Required:    true,
							Optional:    false,
							Description: "Requires task to be completed.",
						},
					},
				},
			},
		},
	}
}
