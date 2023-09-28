provider "blameless" {
  instance = var.blameless_instance
  key      = var.blameless_key
}

# resource "blameless_organization" "org_settings" {
#   name     = "local"
#   timezone = "US/Pacific"
# }

# resource "blameless_incident_roles" "role_settings" {
#   roles = [
#     "Commander",
#     "Communication lead",
#     "Observer",
#   ]
# }

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
