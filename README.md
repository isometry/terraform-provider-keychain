# Terraform Keychain Provider

A trivial Terraform provider exposing passwords from the macOS Keychain.

## Usage

```terraform
provider "keychain" {}

data "keychain_password" "example" {
  service  = "Example Service"
  username = "test@example.com"
}

output "example_password" {
  value     = data.keychain_password.example.password
  sensitive = true
}
```
