## 1.6.0 (Unreleased)

### :rocket: **New Features**

* `validator/NoneOf` - Add support for formatting `NoneOf` validation description. ([GH-45](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/45))

### :bug: **Bug Fixes**

* `validator/OneOf` - Fix description after upgrade dependency `hashicorp/terraform-plugin-framework-validators` to `v0.11.0`. ([GH-45](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/45))

### :dependabot: **Dependencies**

* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.3 to 1.3.4 ([GH-42](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/42))
* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.4 to 1.3.5 ([GH-44](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/44))

## 1.5.1 (August  3, 2023)
## 1.5.0 (August  3, 2023)

### :rocket: **New Features**

* `Deprecated` - Add `deprecated` option in attribute. Now the documentation will be generated for deprecated attributes. ([GH-40](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/40))

### :dependabot: **Dependencies**

* deps: bumps github.com/hashicorp/terraform-plugin-framework from 1.3.2 to 1.3.3 ([GH-39](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/39))

## 1.4.1 (July 20, 2023)

### :bug: **Bug Fixes**

* `DataSource` - Fix `DataSource` generation for `Super`attributes. ([GH-38](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/38))

## 1.4.0 (July 19, 2023)

### :rocket: **New Features**

* `SuperTypes` - Add new attributes `Super` ([GH-37](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/37))

### :dependabot: **Dependencies**

* deps: bumps github.com/iancoleman/strcase from 0.2.0 to 0.3.0 ([GH-34](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/34))

## 1.3.3 (July  8, 2023)

### :dependabot: **Dependencies**

* deps: bumps dependabot/fetch-metadata from 1.5.1 to 1.6.0 ([GH-31](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/31))
* deps: bumps github.com/hashicorp/terraform-plugin-framework-timeouts from 0.4.0 to 0.4.1 ([GH-33](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/33))

## 1.3.2 (June 26, 2023)

### :dependabot: **Dependencies**

* deps: bumps github.com/hashicorp/terraform-plugin-framework-timeouts from 0.3.1 to 0.4.0 ([GH-25](https://github.com/orange-cloudavenue/terraform-provider-cloudavenue/issues/25))

## 1.3.1 (2023-06-19)

### Miscellaneous

* **deps:** bump github.com/hashicorp/terraform-plugin-framework ([#23](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/issues/23)) ([7662330](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/7662330ea453baedcb343d5e37a1fb929854d68b))

## 1.3.0 (2023-04-05)

### Features

* add map and mapNested ([#20](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/issues/20)) ([f84e65e](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/f84e65e9c96f989bc135e95c698ebf422ebc457c))

## 1.2.0 (2023-04-01)

### Features

* **template:** add code formatting on generated files  ([#17](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/issues/17)) ([3f43c9a](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/3f43c9a2a323a02ac4052a988d3755376e5cde0f))

### Miscellaneous

* Create CODE_OF_CONDUCT.md ([fd1e970](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/fd1e97093f453facdcc08ddc98bf11a88cae4f1b))
* create LICENCE ([f0d5d24](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/f0d5d243c784e0d7113c604fc0c9edd4f7fa8aba))

## 1.1.1 (2023-03-31)

### Bug Fixes

* datasource schema ([ec10d41](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/ec10d411498e0813b544406fb6e54118ff83db18))

## 1.1.0 (2023-03-31)

### Features

* add SingleNested ([81be9d9](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/81be9d95f84c7f66b39983c0669b6fc73e239353))
* add SingleNested ([0a466ef](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/0a466ef35da8a765ec1aaa1312f4f445e7363440))

### Bug Fixes

* change struct ([19d3d7a](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/19d3d7af3aa83b2469a1046f7d3c46e53471958f))
* schemaD attribute ([ca4a517](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/ca4a5177dbf3744f6af28f41c65bcca3d5db6a09))

## 1.0.1 (2023-03-30)

### Bug Fixes

* missing dot on validators and pm markdown desc ([#4](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/issues/4)) ([edd7a37](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/edd7a374cc73b575188a853bed51dea81e28f910))

## 1.0.0 (2023-03-30)

### Features

* super schema ([5b70824](https://github.com/FrangipaneTeam/terraform-plugin-framework-superschema/commit/5b70824b50d2a86c7589cc3f09c63bcb3809b650))
