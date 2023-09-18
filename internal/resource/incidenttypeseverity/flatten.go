package incidenttypeseverity

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
)

func flattenIncidentSettings(settings *model.IncidentTypeSeverityIncidentSettings) []interface{} {
	result := map[string]interface{}{
		"end_of_customer_impact_status": settings.EndOfCustomerImpactStatus,
		"private_incident_channel":      settings.PrivateIncidentChannel,
		"channel_naming":                flattenChannelNaming(settings.ChannelNaming),
		"team_notifications":            flattenTeamNotifications(settings.TeamNotifications),
	}
	return []interface{}{result}
}

func flattenChannelNaming(settings *model.IncidentTypeSeverityIncidentChannelNaming) map[string]interface{} {
	return map[string]interface{}{
		"incident_naming_scheme": settings.IncidentNamingScheme,
		"require_dash_separator": settings.RequireDashSeparator,
		"custom_channel_format":  settings.CustomChannelFormat,
	}
}

func flattenTeamNotifications(settings *model.IncidentTypeSeverityIncidentTeamNotification) map[string]interface{} {
	return map[string]interface{}{
		"auto_recruit_team_members": settings.AutoRecruitTeamMembers,
		"announcement_channels":     settings.AnnouncementChannels,
	}
}

func flattenRetrospectiveSettings(settings *model.IncidentTypeSeverityRetrospectiveSettings) []interface{} {
	result := map[string]interface{}{
		"analysis_template":            settings.AnalysisTemplate,
		"questionnaire_template":       settings.QuestionnaireTemplate,
		"required":                     settings.Required,
		"daily_reminder":               settings.DailyReminder,
		"incident_resolution_required": settings.IncidentResolutionRequired,
	}
	return []interface{}{result}
}

func flattenTaskSettings(settings *model.IncidentTypeSeverityTaskSettings) []interface{} {
	result := map[string]interface{}{
		"full_permission_role": settings.FullPermissionRole,
		"task_list":            flattenTaskList(settings.TaskList),
	}
	return []interface{}{result}
}

func flattenTaskList(settings []*model.IncidentSeverityTypeTaskList) map[string]interface{} {
	result := make(map[string]interface{})
	for _, setting := range settings {
		result[setting.IncidentStatus] = flattenTaskListTasks(setting.Tasks)
	}
	return result
}

func flattenTaskListTasks(settings []*model.IncidentSeverityTypeTaskListTask) []interface{} {
	result := make([]interface{}, 0)

	for _, setting := range settings {
		result = append(result, map[string]interface{}{
			"name":     setting.Name,
			"role":     setting.Role,
			"required": setting.Required,
		})
	}

	return result
}
