resource "blameless_incident_severities" "sev_settings" {
  severity {
    level = 0
    label = "SEV0"
  }
  severity {
    level = 1
    label = "SEV1"
  }
  severity {
    level = 2
    label = "SEV2"
  }
  severity {
    level = 3
    label = "SEV3"
  }
}
