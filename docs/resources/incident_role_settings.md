---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "blameless_incident_role_settings Resource - terraform-provider-blameless"
subcategory: ""
description: |-
  Incident Roles
---

# blameless_incident_role_settings (Resource)

Incident Roles

## Example Usage

```terraform
resource "blameless_incident_role_settings" "role_settings" {
  roles = [
    "{{role name 1}}",
    "{{role name 2}}",
    "{{role name 3}}",
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `roles` (List of String) List of incident roles.

### Read-Only

- `id` (String) The ID of this resource.