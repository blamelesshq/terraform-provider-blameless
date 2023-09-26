package model

type IncidentType struct {
	Id         string                  `json:"id"`
	Name       string                  `json:"name"`
	Active     bool                    `json:"active"`
	Severities []*IncidentTypeSeverity `json:"severities,omitempty"`
}

type IncidentTypeSeverity struct {
	Severity              *int                               `json:"severity,omitempty"`
	IncidentSettings      *IncidentTypeSeverityIncident      `json:"incidentSettings,omitempty"`
	RetrospectiveSettings *IncidentTypeSeverityRetrospective `json:"retrospectiveSettings,omitempty"`
	TaskSettings          *IncidentTypeSeverityTask          `json:"taskSettings,omitempty"`
}

type IncidentTypeSeverityIncident struct {
	EndOfCustomerImpactStatus *string                                       `json:"endOfCustomerImpactStatus,omitempty"`
	PrivateIncidentChannel    *bool                                         `json:"privateIncidentChannel,omitempty"`
	ChannelNaming             *IncidentTypeSeverityIncidentChannelNaming    `json:"channelNaming,omitempty"`
	TeamNotifications         *IncidentTypeSeverityIncidentTeamNotification `json:"teamNotifications,omitempty"`
}

type IncidentTypeSeverityIncidentChannelNaming struct {
	IncidentNamingScheme *string `json:"incidentNamingScheme,omitempty"`
	CustomChannelFormat  *string `json:"customChannelFormat,omitempty"`
	RequireDashSeparator *bool   `json:"requireDashSeparator,omitempty"`
}

type IncidentTypeSeverityIncidentTeamNotification struct {
	AutoRecruitTeamMembers []string `json:"autoRecruitTeamMembers,omitempty"`
	AnnouncementChannels   []string `json:"announcementChannels,omitempty"`
}

type IncidentTypeSeverityRetrospective struct {
	Required                   *bool   `json:"required,omitempty"`
	IncidentResolutionRequired *bool   `json:"incidentResolutionRequired,omitempty"`
	DailyReminder              *bool   `json:"dailyReminder,omitempty"`
	AnalysisTemplate           *string `json:"analysisTemplate,omitempty"`
	QuestionnaireTemplate      *string `json:"questionnaireTemplate,omitempty"`
}

type IncidentTypeSeverityTask struct {
	FullPermissionRole *string                         `json:"fullPermissionRole,omitempty"`
	TaskList           []*IncidentSeverityTypeTaskList `json:"taskList,omitempty"`
}

type IncidentSeverityTypeTaskList struct {
	IncidentStatus *string `json:"incidentStatus,omitempty"`
	Name           *string `json:"name,omitempty"`
	Required       *bool   `json:"required,omitempty"`
	Role           *string `json:"role,omitempty"`
}
