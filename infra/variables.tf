variable "hcloud_token" {
  type      = string
  sensitive = true
}

variable "hcloud_location" {
  type = string
}

variable "ssh_keys" {
  type = list(string)
}
