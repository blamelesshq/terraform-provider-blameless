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

# resource "blameless_incident_severities" "sev_settings" {
#   severities {
#     sev0_label = "SEV0"
#     sev1_label = "SEV1"
#     sev2_label = "SEV2"
#     sev3_label = "SEV3"
#   }
# }

# # TODO  Check case sensitivity in naming
# resource "blameless_incident_type" "incident_type_squirrel_attack" {
#   name   = "squirrel attack"
#   active = false
# }

# resource "blameless_incident_type_severity" "incident_type_default_sev0" {
#   incident_type_id        = blameless_incident_type.incident_type_DEFAULT.id
#   incident_severity_label = blameless_incident_severity_settings.sev_settings.severities.sev0_label

#   retro {
#     retro_template = file("")

#   }
# }

# resource "blameless_incident_type_severity" "incident_type_DEFAULT_SEV1" {
#   incident_type_id        = blameless_incident_type.incident_type_DEFAULT.id
#   incident_severity_label = blameless_incident_severity_settings.sev_settings.severities.sev1_label

#   retro {
#     retro_template = file("")

#   }
# }



