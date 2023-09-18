locals {
  analysis_template      = file("${path.module}/analysis.md")
  questionnaire_template = file("${path.module}/questionnaire.json")
}

provider "blameless" {
  instance = var.blameless_instance
  key      = var.blameless_key
}

resource "blameless_org_settings" "org_settings" {
  name     = "local"
  timezone = "US/Pacific"
}

resource "blameless_incident_role_settings" "role_settings" {
  roles = [
    "Commander",
    "Communication lead",
    "Observer",
  ]
}

resource "blameless_incident_severity_settings" "sev_settings" {
  severities {
    sev0_label = "SEV0"
    sev1_label = "SEV1"
    sev2_label = "SEV2"
    sev3_label = "SEV3"
  }
}

resource "blameless_incident_type" "incident_type_security" {
  name   = "Security"
  active = true
}

resource "blameless_incident_type_severity" "sev" {
  incident_type_id  = "afdssa"
  incident_severity = 0

  incident_settings {
    end_of_customer_impact_status = "RESOLVED"
    private_incident_channel      = false
    channel_naming {
      incident_naming_scheme = "custom"
      require_dash_separator = false
      custom_channel_format  = "{incident.name}-1123"
    }
    team_notifications {
      auto_recruit_team_members = [
        "@dillon"
      ]
      announcement_channels = [
        "#general"
      ]
    }
  }

  retrospective_settings {
    analysis_template            = local.analysis_template
    questionnaire_template       = local.questionnaire_template
    required                     = true
    daily_reminder               = true
    incident_resolution_required = true
  }


  task_settings {
    full_permission_role = "Commander"
    task_list {
      investigating {
        task {
          name     = "Assign Commander role"
          role     = "Creator"
          required = true
        }
        task {
          name     = "Assign Commander role2"
          role     = "Creator"
          required = true
        }
      }
      identified {
        task {
          name     = "Assign Commander role"
          role     = "Creator"
          required = true
        }
        task {
          name     = "Assign Commander role2"
          role     = "Creator"
          required = true
        }
      }
      monitoring {
        task {
          name     = "Assign Commander role"
          role     = "Creator"
          required = true
        }
      }
      resolved {
        task {
          name     = "Assign Commander role"
          role     = "Creator"
          required = true
        }
      }
    }
  }
}
