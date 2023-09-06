resource "blameless_org_settings" "org_settings" {
  name     = "{{organization name}}"
  timezone = "{{organization timezone}}"
  # Example Timezones:
  #  - Asia/Jakarta
  #  - Asia/Singapore
  #  - Asia/Tokyo
  #  - Europe/London
  #  - Europe/Paris
  #  - Europe/Rome
  #  - US/Central
  #  - US/Eastern
  #  - US/Mountain
  #  - US/Pacific
  #  - UTC
}
