# Changelog

## 0.4.0 (2021-07-10)

* Update to terraform-plugin-sdk v2, Go 1.16 and vendor dependencies.

## 0.3.0 (2019-09-14)

* Add `class` attribute to `keychain_password` data source, enabling access to 'internet' passwords.
* Add `class` and `kind` attributes to `keychain_password` resource, enabling creation of both 'generic' and 'internet' password items, and for the object kind to be overridden.

## 0.2.0 (2019-09-14)

* Implement `keychain_password` resource, allowing create, read and delete of generic macOS Keychain passwords.

## 0.1.0 (2019-09-12)

* Implement `keychain_password` data source, allowing generic macOS Keychain passwords to be read into Terraform state.
