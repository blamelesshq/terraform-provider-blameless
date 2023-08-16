# terraform-provider

This provider allow use Terraform to handle Blameless resources.

## Setup

Generate an API token in the Blamless Identity Management - Key Management section. Then please provide your API key and instance URL by:

1. Putting the values in a `.tfvars` file

```
instance = "https://{{instance}}.blameless-dev.io"
key = "123qwe"
```

2. Setting the environment variables `BLAMELESS_INSTANCE` and `BLAMELESS_KEY`

## Local Testing

1. Run `make install` in the root directory
2. Go to the `modules` directory
3. Run `terraform init`
4. Run `terraform plan`, `terraform validate`, or `terraform apply`