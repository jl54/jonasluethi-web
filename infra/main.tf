terraform {
  required_providers {
    hcloud = {
      source  = "hetznercloud/hcloud"
      version = "~> 1.60.0"
    }
  }

  backend "s3" {
    key                         = "prod/terraform.tfstate"
    skip_credentials_validation = true
    skip_metadata_api_check     = true
    skip_region_validation      = true
    use_path_style              = true
  }
}

provider "hcloud" {
  token = var.hcloud_token
}

resource "hcloud_server" "jl-web-01" {
  name        = "jl-web-01"
  image       = "ubuntu-24.04"
  server_type = "cx23"
  location    = var.hcloud_location
  ssh_keys    = var.ssh_keys
  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
  }
}
