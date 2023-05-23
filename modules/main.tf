terraform {
  required_providers {

  }
}

provider "blameless" {}

module "settings" {
  source = "./settings"

  name = "Settings handler"

}

output "settings" {
  value = module.settings
}
