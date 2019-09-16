# Terraform Keychain Provider

A simple Terraform provider for passwords in the macOS Keychain.

WARNING: keychain passwords are copied in plaintext to your Terraform state. Consider the implications and weigh the risks before using this provider.

## Usage

```terraform
provider "keychain" {}

data "keychain_password" "example" {
  // class = "generic"
  service  = "https://data.example.com"
  username = "test@example.com"
}

resource "keychain_password" "example" {
  // class = "generic"
  // kind  = "application password"
  service  = "https://resource.example.com"
  username = "test@example.com"
  password = "Passw0rd!"
}

output "example_password" {
  value     = data.keychain_password.example.password
  sensitive = true
}
```
