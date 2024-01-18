provider "zitadeltools" {
}

terraform {
  required_version = "~> 1.6.4"
  required_providers {
    zitadeltools = {
      source  = "registry.socrate.fr/socrate-tech/zitadeltools"
      version = "~> 1.0.0"
    }
  }
}
