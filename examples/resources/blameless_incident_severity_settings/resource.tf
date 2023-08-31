resource "blameless_incident_severity_settings" "sev_settings" {
  severities {
    sev0_label = "{{severity 0 label}}"
    sev1_label = "{{severity 1 label}}"
    sev2_label = "{{severity 2 label}}"
    sev3_label = "{{severity 3 label}}"
  }
}
