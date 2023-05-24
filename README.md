# terraform-provider

This provider allow use Terraform to handle Blameless resources.




## Example Usage

Terraform 0.13 and later:

```
terraform {
  required_providers {
    blameless = {
      source  = "blameless"
      version = "~> 1.0"
    }
  }
}
```

## Authentication

Client id will be provide by our CS Team.

Generate an API token on the Blamless identity Management - Key Management section.

Then please provide youur client id and token.

```
provider "blameless" {
    client_id = var.blameless_client_id
    client_secret = var.blameless_client_secret
}
```