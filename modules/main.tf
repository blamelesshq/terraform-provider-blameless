provider "blameless" {
  instance = var.blameless_instance
  key = var.blameless_key
}

resource "blameless_organization" "org_settings" {
  name = "swat-2"
  timezone = "US/Pacific"
  description = "swat-2 desc"
  incident_roles = ["Commander", "Communication lead"]
  incident_severities {
    sev0_label = "SEV0"
    sev1_label = "SEV1"
    sev2_label = "SEV2"
    sev3_label = "SEV3"
  }
}