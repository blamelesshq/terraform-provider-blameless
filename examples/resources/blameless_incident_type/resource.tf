locals {
  analysis_template      = file("${path.module}/analysis.md")
  questionnaire_template = file("${path.module}/questionnaire.json")
}

resource "blameless_incident_type" "incident_type_security" {
  name   = "Security"
  active = true

  severity0_settings {
    end_of_customer_impact_status = "RESOLVED"
    private_incident_channel      = false

    incident_naming_scheme = "custom"
    require_dash_separator = false
    custom_channel_format  = "s0-{incident.name}"

    auto_recruit_team_members = [
      "@dillon"
    ]
    announcement_channels = [
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
        name            = "Update staus page"
        role            = "Communication Lead"
        required        = true
      }
      task {
        incident_status = "MONITORING"
        name            = "Update staus page"
        role            = "Communication Lead"
        required        = true
      }
      task {
        incident_status = "RESOLVED"
        name            = "Update staus page"
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

  severity1_settings {
    incident_naming_scheme = "custom"
    require_dash_separator = false
    custom_channel_format  = "s1-{incident.name}"
  }

  severity2_settings {
    incident_naming_scheme = "custom"
    require_dash_separator = false
    custom_channel_format  = "s2-{incident.name}"
  }

  severity3_settings {
    incident_naming_scheme = "custom"
    require_dash_separator = false
    custom_channel_format  = "s3-{incident.name}"
  }
}
