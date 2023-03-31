terraform {
  required_providers {
    scaleway = {
      source  = "scaleway/scaleway"
      version = "~> 2.14.0"
    }
  }
  required_version = "~> 1.4.0"
}

provider "scaleway" {
  profile = "dummy"
}