locals {
  analysis_template      = file("${path.module}/analysis.md")
  questionnaire_template = file("${path.module}/questionnaire.json")
}

resource "blameless_incident_type" "incident_type_security" {
  name   = "Security"
  active = true

  severity_settings {
    severity                      = 0
    end_of_customer_impact_status = "RESOLVED"
    private_incident_channel      = false

    incident_naming_scheme = "custom"
    require_dash_separator = false
    custom_channel_format  = "s0-{incident.name}"

    # Slack Notifications
    slack_invited_users = [
      "@dillon"
    ]
    slack_announcement_channels = [
      "#general"
    ]

    # MS Teams Notifications
    teams_invited_users = [
      "@dillon"
    ]
    teams_announcement_groups = [
      "@responders"
    ]
    teams_announcement_channels = [
      "#general"
    ]

    retrospective_analysis_template            = local.analysis_template
    retrospective_questionnaire_template       = local.questionnaire_template
    retrospective_required                     = true
    retrospective_daily_reminder               = true
    retrospective_incident_resolution_required = true

    tasks_full_permission_role = "Communication Lead"

    task_list {
      task {
        incident_status = "INVESTIGATING"
        name            = "Assign Commander role"
        role            = "Creator"
        required        = true
      }
      task {
        incident_status = "INVESTIGATING"
        name            = "Page on-call responders"
        role            = "Commander"
        required        = true
      }
      task {
        incident_status = "IDENTIFIED"
        name            = "Update status page"
        role            = "Communication Lead"
        required        = true
      }
      task {
        incident_status = "MONITORING"
        name            = "Update status page"
        role            = "Communication Lead"
        required        = true
      }
      task {
        incident_status = "RESOLVED"
        name            = "Update status page"
        role            = "Communication Lead"
        required        = true
      }
      task {
        incident_status = "RESOLVED"
        name            = "Schedule retrospective"
        role            = "Commander"
        required        = true
      }
    }
  }

  severity_settings {
    severity               = 1
    incident_naming_scheme = "custom"
    require_dash_separator = false
    custom_channel_format  = "s1-{incident.name}"
  }

  severity_settings {
    severity               = 2
    incident_naming_scheme = "custom"
    require_dash_separator = false
    custom_channel_format  = "s2-{incident.name}"
  }

  severity_settings {
    severity               = 3
    incident_naming_scheme = "custom"
    require_dash_separator = false
    custom_channel_format  = "s3-{incident.name}"
  }
}
