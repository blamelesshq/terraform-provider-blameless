package incidenttype

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
)

func flattenIncidentSeverities(settings []*model.IncidentTypeSeverity) []interface{} {
	results := []interface{}{}
	for _, setting := range settings {
		results = append(results, flattenIncidentSeverity(setting))
	}
	return results
}

func flattenIncidentSeverity(settings *model.IncidentTypeSeverity) map[string]interface{} {
	if settings == nil {
		return map[string]interface{}{}
	}
	result := map[string]interface{}{}
	result["severity"] = settings.Severity
	if settings.IncidentSettings != nil {
		result["end_of_customer_impact_status"] = settings.IncidentSettings.EndOfCustomerImpactStatus
		result["private_incident_channel"] = settings.IncidentSettings.PrivateIncidentChannel

		if settings.IncidentSettings.ChannelNaming != nil {
			result["incident_naming_scheme"] = settings.IncidentSettings.ChannelNaming.IncidentNamingScheme
			result["require_dash_separator"] = settings.IncidentSettings.ChannelNaming.RequireDashSeparator
			result["custom_channel_format"] = settings.IncidentSettings.ChannelNaming.CustomChannelFormat
		}

		if settings.IncidentSettings.TeamNotifications != nil {
			result["slack_invited_users"] = settings.IncidentSettings.TeamNotifications.SlackInvitedUsers
			result["slack_announcement_channels"] = settings.IncidentSettings.TeamNotifications.SlackAnnouncementChannels
			result["teams_invited_users"] = settings.IncidentSettings.TeamNotifications.TeamsInvitedUsers
			result["teams_announcement_groups"] = settings.IncidentSettings.TeamNotifications.TeamsAnnouncementGroups
			result["teams_announcement_channels"] = settings.IncidentSettings.TeamNotifications.TeamsAnnouncementChannels
		}
	}

	if settings.RetrospectiveSettings != nil {
		result["retrospective_analysis_template"] = settings.RetrospectiveSettings.AnalysisTemplate
		result["retrospective_questionnaire_template"] = settings.RetrospectiveSettings.QuestionnaireTemplate
		result["retrospective_required"] = settings.RetrospectiveSettings.Required
		result["retrospective_daily_reminder"] = settings.RetrospectiveSettings.DailyReminder
		result["retrospective_incident_resolution_required"] = settings.RetrospectiveSettings.IncidentResolutionRequired
	}

	if settings.TaskSettings != nil {
		result["tasks_full_permission_role"] = settings.TaskSettings.FullPermissionRole
		result["task_list"] = flattenTaskList(settings.TaskSettings.TaskList)
	}
	return result
}

func flattenTaskList(settings []*model.IncidentSeverityTypeTaskList) []interface{} {
	if settings == nil {
		return []interface{}{}
	}

	result := make([]interface{}, 0)
	for _, setting := range settings {
		result = append(result, map[string]interface{}{
			"incident_status": setting.IncidentStatus,
			"name":            setting.Name,
			"role":            setting.Role,
			"required":        setting.Required,
		})
	}
	return []interface{}{
		map[string]interface{}{
			"task": result,
		},
	}
}
