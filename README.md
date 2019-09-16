# Terraform Keychain Provider

A simple Terraform provider for passwords in the macOS Keychain.

## Provider

```terraform
provider "keychain" {}
```

## Data Sources

### keychain_password

The `keychain_password` data source can be used to retrieve the password associated with an existing macOS Keychain item.

> WARNING: use of this data source will result in a Keychain password being copied into your Terraform state *in plaintext*. Please consider the security implications and weigh the risks before use!

#### Example Usage

```terraform
data "keychain_password" "example" {
  service  = "https://data.example.com"
  username = "test@example.com"
}
```

#### Argument Reference

The following arguments are supported:

* `class` – (Optional) The class of item to filter on. Allowed options are `generic` for Generic Passwords and `internet` for Internet Passwords. Default: `generic`.
* `service` – (Required) The service (typically the host or website) to filter on.
* `username` – (Required) The username to filter on.

#### Attribute Reference

The only exported attribute is `password`, which is the password of the matching item, or `null` if no matching item was found.

## Resources

### keychain_password

The `keychain_password` resource can be used to create and manage macOS Keychain password items.

> WARNING: use of this resource will leave a copy of the password in your Terraform state *in plaintext*. Please consider the security implications and weigh the risks before use!

#### Example Usage

```terraform
resource "keychain_password" "example" {
  class    = "internet"
  kind     = "Internet password"
  service  = "https://resource.example.com"
  username = "test@example.com"
  password = "Passw0rd!"
}
```

#### Argument Reference

The following arguments are supported:

* `class` – (Optional) The class of item. Allowed options are `generic` for Generic Passwords and `internet` for Internet Passwords. Default: `generic`.
* `kind` – (Optional) The kind of item. Default: `terraform password`.
* `service` – (Required) The service (typically the host or website).
* `username` – (Required) The username or account name.
* `password` – (Required) The password.

#### Attribute Reference

No useful attributes are exported.
