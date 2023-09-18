package incidenttypeseverity

import (
	"context"

	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expandSettings(ctx context.Context, config cty.Value) *model.IncidentTypeSeverity {
	settings := &model.IncidentTypeSeverity{
		IncidentTypeId:        value.String(config.GetAttr("incident_type_id")),
		Severity:              value.Int(config.GetAttr("incident_severity")),
		IncidentSettings:      expandIncidents(ctx, config.GetAttr("incident_settings")),
		RetrospectiveSettings: expandRetrospectives(config.GetAttr("retrospective_settings")),
		TaskSettings:          expandTasks(config.GetAttr("task_settings")),
	}
	return settings
}

func expandIncidents(ctx context.Context, incidents cty.Value) *model.IncidentTypeSeverityIncidentSettings {
	var result *model.IncidentTypeSeverityIncidentSettings
	incidents.ForEachElement(func(_, inc cty.Value) (stop bool) {
		result = &model.IncidentTypeSeverityIncidentSettings{
			EndOfCustomerImpactStatus: value.String(inc.GetAttr("end_of_customer_impact_status")),
			PrivateIncidentChannel:    value.Bool(inc.GetAttr("private_incident_channel")),
			ChannelNaming:             expandChannelNaming(inc.GetAttr("channel_naming")),
			TeamNotifications:         expandTeamNotifications(inc.GetAttr("team_notifications")),
		}
		return stop
	})
	return result
}

func expandChannelNaming(channelNaming cty.Value) *model.IncidentTypeSeverityIncidentChannelNaming {
	var result *model.IncidentTypeSeverityIncidentChannelNaming
	channelNaming.ForEachElement(func(_, cn cty.Value) (stop bool) {
		result = &model.IncidentTypeSeverityIncidentChannelNaming{
			IncidentNamingScheme: value.String(cn.GetAttr("incident_naming_scheme")),
			RequireDashSeparator: value.Bool(cn.GetAttr("require_dash_separator")),
			CustomChannelFormat:  value.String(cn.GetAttr("custom_channel_format")),
		}
		return stop
	})
	return result
}

func expandTeamNotifications(teamNotifications cty.Value) *model.IncidentTypeSeverityIncidentTeamNotification {
	var result *model.IncidentTypeSeverityIncidentTeamNotification
	teamNotifications.ForEachElement(func(_, tn cty.Value) (stop bool) {
		result = &model.IncidentTypeSeverityIncidentTeamNotification{
			AutoRecruitTeamMembers: value.StringArray(tn.GetAttr("auto_recruit_team_members")),
			AnnouncementChannels:   value.StringArray(tn.GetAttr("announcement_channels")),
		}
		return stop
	})
	return result
}

func expandRetrospectives(retrospective cty.Value) *model.IncidentTypeSeverityRetrospectiveSettings {
	var result *model.IncidentTypeSeverityRetrospectiveSettings
	retrospective.ForEachElement(func(_, r cty.Value) (stop bool) {
		result = &model.IncidentTypeSeverityRetrospectiveSettings{
			AnalysisTemplate:           value.String(r.GetAttr("analysis_template")),
			QuestionnaireTemplate:      value.String(r.GetAttr("questionnaire_template")),
			Required:                   value.Bool(r.GetAttr("required")),
			DailyReminder:              value.Bool(r.GetAttr("daily_reminder")),
			IncidentResolutionRequired: value.Bool(r.GetAttr("incident_resolution_required")),
		}
		return stop
	})
	return result
}

func expandTasks(tasks cty.Value) *model.IncidentTypeSeverityTaskSettings {
	var result *model.IncidentTypeSeverityTaskSettings
	tasks.ForEachElement(func(_, t cty.Value) (stop bool) {
		result = &model.IncidentTypeSeverityTaskSettings{
			FullPermissionRole: value.String(t.GetAttr("full_permission_role")),
			TaskList:           expandTaskList(t.GetAttr("task_list")),
		}
		return stop
	})
	return result
}

func expandTaskList(taskList cty.Value) []*model.IncidentSeverityTypeTaskList {
	results := []*model.IncidentSeverityTypeTaskList{}

	statusInvestigating := "investigating"
	statusIdentified := "identified"
	statusMonitoring := "monitoring"
	statusResolved := "resolved"

	taskList.ForEachElement(func(_, tl cty.Value) (stop bool) {
		investigating := expandTaskListItem(tl.GetAttr(statusInvestigating), statusInvestigating)
		if investigating != nil {
			results = append(results, investigating)
		}

		identified := expandTaskListItem(tl.GetAttr(statusIdentified), statusIdentified)
		if identified != nil {
			results = append(results, identified)
		}

		monitoring := expandTaskListItem(tl.GetAttr(statusMonitoring), statusMonitoring)
		if monitoring != nil {
			results = append(results, monitoring)
		}

		resolved := expandTaskListItem(tl.GetAttr(statusResolved), statusResolved)
		if resolved != nil {
			results = append(results, resolved)
		}
		return stop
	})

	return results
}

func expandTaskListItem(taskListItem cty.Value, incidentStatus string) *model.IncidentSeverityTypeTaskList {
	var result *model.IncidentSeverityTypeTaskList
	taskListItem.ForEachElement(func(_, tli cty.Value) (stop bool) {
		if tli.IsNull() {
			result = nil
		} else {
			result = &model.IncidentSeverityTypeTaskList{
				IncidentStatus: incidentStatus,
				Tasks:          expandTaskListItemTasks(tli.GetAttr("task")),
			}
		}
		return stop
	})
	return result
}

func expandTaskListItemTasks(itemTasks cty.Value) []*model.IncidentSeverityTypeTaskListTask {
	results := []*model.IncidentSeverityTypeTaskListTask{}
	itemTasks.ForEachElement(func(_, tlt cty.Value) (stop bool) {
		results = append(results, &model.IncidentSeverityTypeTaskListTask{
			Name:     value.String(tlt.GetAttr("name")),
			Role:     value.String(tlt.GetAttr("role")),
			Required: value.Bool(tlt.GetAttr("required")),
		})
		return stop
	})
	return results
}
