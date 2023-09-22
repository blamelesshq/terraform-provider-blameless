package incidenttype

import (
	"github.com/blamelesshq/terraform-provider/internal/model"
	"github.com/blamelesshq/terraform-provider/internal/value"
	"github.com/hashicorp/go-cty/cty"
)

func expand(config cty.Value) *model.IncidentTypeSettings {
	settings := &model.IncidentTypeSettings{
		Name:       value.String(config.GetAttr("name")),
		Active:     *value.Bool(config.GetAttr("active")),
		Severities: expandSeverities(config),
	}
	return settings
}

func expandSeverities(config cty.Value) []*model.IncidentTypeSeverity {
	sev0 := expandSeverity(0, config.GetAttr("severity0_settings"))
	sev1 := expandSeverity(1, config.GetAttr("severity1_settings"))
	sev2 := expandSeverity(2, config.GetAttr("severity2_settings"))
	sev3 := expandSeverity(3, config.GetAttr("severity3_settings"))

	results := []*model.IncidentTypeSeverity{}

	if sev0 != nil {
		results = append(results, sev0)
	}
	if sev1 != nil {
		results = append(results, sev1)
	}
	if sev2 != nil {
		results = append(results, sev2)
	}
	if sev3 != nil {
		results = append(results, sev3)
	}

	return results
}

func expandSeverity(severity int, config cty.Value) *model.IncidentTypeSeverity {
	var settings *model.IncidentTypeSeverity

	config.ForEachElement(func(key, val cty.Value) (stop bool) {
		settings = &model.IncidentTypeSeverity{
			Severity: &severity,
			IncidentSettings: &model.IncidentTypeSeverityIncidentSettings{
				EndOfCustomerImpactStatus: value.String(val.GetAttr("end_of_customer_impact_status")),
				PrivateIncidentChannel:    value.Bool(val.GetAttr("private_incident_channel")),
				ChannelNaming: &model.IncidentTypeSeverityIncidentChannelNaming{
					IncidentNamingScheme: value.String(val.GetAttr("incident_naming_scheme")),
					RequireDashSeparator: value.Bool(val.GetAttr("require_dash_separator")),
					CustomChannelFormat:  value.String(val.GetAttr("custom_channel_format")),
				},
				TeamNotifications: &model.IncidentTypeSeverityIncidentTeamNotification{
					AutoRecruitTeamMembers: value.StringArray(val.GetAttr("auto_recruit_team_members")),
					AnnouncementChannels:   value.StringArray(val.GetAttr("announcement_channels")),
				},
			},
			RetrospectiveSettings: &model.IncidentTypeSeverityRetrospectiveSettings{
				AnalysisTemplate:           value.String(val.GetAttr("retrospective_analysis_template")),
				QuestionnaireTemplate:      value.String(val.GetAttr("retrospective_questionnaire_template")),
				Required:                   value.Bool(val.GetAttr("retrospective_required")),
				DailyReminder:              value.Bool(val.GetAttr("retrospective_daily_reminder")),
				IncidentResolutionRequired: value.Bool(val.GetAttr("retrospective_incident_resolution_required")),
			},
			TaskSettings: &model.IncidentTypeSeverityTaskSettings{
				FullPermissionRole: value.String(val.GetAttr("tasks_full_permission_role")),
				TaskList:           expandTaskList(val.GetAttr("task_list")),
			},
		}
		return stop
	})

	return settings
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

// func expandTaskListItem(taskListItem cty.Value, incidentStatus string) *model.IncidentSeverityTypeTaskList {
// 	var result *model.IncidentSeverityTypeTaskList

// 	taskListItem.ForEachElement(func(_, tli cty.Value) (stop bool) {
// 		if tli.IsNull() {
// 			result = nil
// 		} else {
// 			result = &model.IncidentSeverityTypeTaskList{
// 				IncidentStatus: incidentStatus,
// 				Tasks:          expandTaskListItemTasks(tli.GetAttr("task")),
// 			}
// 		}
// 		return stop
// 	})
// 	return result
// }

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
