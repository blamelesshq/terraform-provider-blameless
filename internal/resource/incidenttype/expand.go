package incidenttype

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expand(config cty.Value) *model.IncidentTypeSettings {
	settings := &model.IncidentTypeSettings{
		Name:       *value.String(config.GetAttr("name")),
		Active:     *value.Bool(config.GetAttr("active")),
		Severities: expandSeverities(config.GetAttr("severity_settings")),
	}
	return settings
}

func expandSeverities(config cty.Value) []*model.IncidentTypeSeverity {
	results := []*model.IncidentTypeSeverity{}
	config.ForEachElement(func(_, val cty.Value) (stop bool) {
		results = append(results, expandSeverity(val))
		return stop
	})
	return results
}

func expandSeverity(config cty.Value) *model.IncidentTypeSeverity {
	return &model.IncidentTypeSeverity{
		Severity: value.Int(config.GetAttr("severity")),
		IncidentSettings: &model.IncidentTypeSeverityIncidentSettings{
			EndOfCustomerImpactStatus: value.String(config.GetAttr("end_of_customer_impact_status")),
			PrivateIncidentChannel:    value.Bool(config.GetAttr("private_incident_channel")),
			ChannelNaming: &model.IncidentTypeSeverityIncidentChannelNaming{
				IncidentNamingScheme: value.String(config.GetAttr("incident_naming_scheme")),
				RequireDashSeparator: value.Bool(config.GetAttr("require_dash_separator")),
				CustomChannelFormat:  value.String(config.GetAttr("custom_channel_format")),
			},
			TeamNotifications: &model.IncidentTypeSeverityIncidentTeamNotification{
				SlackInvitedUsers:         value.StringArray(config.GetAttr("slack_invited_users")),
				SlackAnnouncementChannels: value.StringArray(config.GetAttr("slack_announcement_channels")),
				TeamsInvitedUsers:         value.StringArray(config.GetAttr("teams_invited_users")),
				TeamsAnnouncementGroups:   value.StringArray(config.GetAttr("teams_announcement_groups")),
				TeamsAnnouncementChannels: value.StringArray(config.GetAttr("teams_announcement_channels")),
			},
		},
		RetrospectiveSettings: &model.IncidentTypeSeverityRetrospectiveSettings{
			AnalysisTemplate:           value.String(config.GetAttr("retrospective_analysis_template")),
			QuestionnaireTemplate:      value.String(config.GetAttr("retrospective_questionnaire_template")),
			Required:                   value.Bool(config.GetAttr("retrospective_required")),
			DailyReminder:              value.Bool(config.GetAttr("retrospective_daily_reminder")),
			IncidentResolutionRequired: value.Bool(config.GetAttr("retrospective_incident_resolution_required")),
		},
		TaskSettings: &model.IncidentTypeSeverityTaskSettings{
			FullPermissionRole: value.String(config.GetAttr("tasks_full_permission_role")),
			TaskList:           expandTaskList(config.GetAttr("task_list")),
		},
	}
}

func expandTaskList(taskList cty.Value) []*model.IncidentSeverityTypeTaskList {
	results := []*model.IncidentSeverityTypeTaskList{}
	taskList.ForEachElement(func(_, val cty.Value) (stop bool) {
		tasks := expandTaskListItems(val.GetAttr("task"))
		results = append(results, tasks...)
		return stop
	})

	return results
}

func expandTaskListItems(itemTasks cty.Value) []*model.IncidentSeverityTypeTaskList {
	results := []*model.IncidentSeverityTypeTaskList{}
	itemTasks.ForEachElement(func(_, val cty.Value) (stop bool) {
		results = append(results, &model.IncidentSeverityTypeTaskList{
			IncidentStatus: value.String(val.GetAttr("incident_status")),
			Name:           value.String(val.GetAttr("name")),
			Role:           value.String(val.GetAttr("role")),
			Required:       value.Bool(val.GetAttr("required")),
		})
		return stop
	})
	return results
}
