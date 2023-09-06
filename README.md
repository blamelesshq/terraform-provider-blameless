# Blameless Terraform Provider

This provider allow use Terraform to handle Blameless resources.

## Getting Started

### Requirements

- [Terraform](https://www.terraform.io/downloads)
- A Blameless API key

### Installation

Terraform uses the [Terraform Registry](https://registry.terraform.io/) to download and install providers. To install
this provider, copy and paste the following code into your Terraform configuration. Then, run `terraform init`.

```terraform
terraform {
  required_providers {
    auth0 = {
      source  = "blameless/blameless"
      version = ">= 0.1.0" # Refer to releases for latest version
    }
  }
}

provider "blameless" {
  instance = var.blameless_instance
  key      = var.blameless_key
}
```

```shell
$ terraform init
```

#### Setup

Generate an API token in the Blamless Identity Management - Key Management section. Then please provide your API key and instance URL by either:

1. Putting the values in a `terraform.tfvars` file

```terraform
instance = "{{blameless_instance}}"
key      = "{{blameless_api_key}}"
```

2. Setting the environment variables `BLAMELESS_INSTANCE` and `BLAMELESS_KEY`

## Feedback

### Contributing

We appreciate feedback and contribution to this repo! Before you get started, please see the following:

- [Contribution Guide](./CONTRIBUTING.md)
- [Code of Conduct Guidelines](./CODE_OF_CONDUCT.md)

### Raise an issue

To provide feedback or report a bug, [please raise an issue on our issue tracker](https://github.com/blameless/terraform-provider-blameless/issues).
