terraform {
  backend "s3" {
    bucket                      = "dummy-terraform-aa0f8df3-3b62-4d27-900a-1f03fad68f66"
    key                         = "terraform.tfstate"
    region                      = "fr-par"
    endpoint                    = "https://s3.fr-par.scw.cloud"
    profile                     = "dummy-terraform"
    skip_credentials_validation = true
    skip_region_validation      = true
  }
}
