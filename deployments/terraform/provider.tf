terraform {
  required_providers {
    scaleway = {
      source  = "scaleway/scaleway"
      version = "~> 2.16.0"
    }
  }
  required_version = "~> 1.4.0"
}

provider "scaleway" {
  profile = "github"

  zone            = "fr-par-2"
  region          = "fr-par"
  organization_id = "15a4e529-d75e-498c-838d-1152527745fe"
  project_id      = "41896ce7-3824-4868-ba1b-141f6c48c8e8"
}
