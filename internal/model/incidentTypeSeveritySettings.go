package model

type IncidentTypeSeverity struct {
	IncidentTypeId        string
	Severity              *int
	IncidentSettings      *IncidentTypeSeverityIncidentSettings      `json:"incidentSettings"`
	RetrospectiveSettings *IncidentTypeSeverityRetrospectiveSettings `json:"retrospectiveSettings"`
	TaskSettings          *IncidentTypeSeverityTaskSettings          `json:"taskSettings"`
}

type IncidentTypeSeverityIncidentSettings struct {
	EndOfCustomerImpactStatus string                                        `json:"endOfCustomerImpactStatus"`
	PrivateIncidentChannel    *bool                                         `json:"privateIncidentChannel"`
	ChannelNaming             *IncidentTypeSeverityIncidentChannelNaming    `json:"channelNaming"`
	TeamNotifications         *IncidentTypeSeverityIncidentTeamNotification `json:"teamNotifications"`
}

type IncidentTypeSeverityIncidentChannelNaming struct {
	IncidentNamingScheme string `json:"incidentNamingScheme"`
	CustomChannelFormat  string `json:"customChannelFormat"`
	RequireDashSeparator *bool  `json:"requireDashSeparator"`
}

type IncidentTypeSeverityIncidentTeamNotification struct {
	AutoRecruitTeamMembers []string `json:"autoRecruitTeamMembers"`
	AnnouncementChannels   []string `json:"announcementChannels"`
}

type IncidentTypeSeverityRetrospectiveSettings struct {
	Required                   *bool  `json:"required"`
	IncidentResolutionRequired *bool  `json:"incidentResolutionRequired"`
	DailyReminder              *bool  `json:"dailyReminder"`
	AnalysisTemplate           string `json:"analysisTemplate"`
	QuestionnaireTemplate      string `json:"questionnaireTemplate"`
}

type IncidentTypeSeverityTaskSettings struct {
	FullPermissionRole string                          `json:"fullPermissionRole"`
	TaskList           []*IncidentSeverityTypeTaskList `json:"taskList"`
}

type IncidentSeverityTypeTaskList struct {
	IncidentStatus string                              `json:"incidentStatus"`
	Tasks          []*IncidentSeverityTypeTaskListTask `json:"tasks"`
}

type IncidentSeverityTypeTaskListTask struct {
	Name     string `json:"name"`
	Required *bool  `json:"required"`
	Role     string `json:"role"`
}
