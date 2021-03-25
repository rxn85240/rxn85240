---
subcategory: "Base"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_resource_provider_registration"
description: |-
    Manages the Registration of a Resource Provider.
---

# azurerm_resource_provider_registration

Manages the registration of a Resource Provider - which allows access to the API's supported by this Resource Provider.

-> The Azure Provider will automatically register all of the Resource Providers which it supports on launch (unless opted-out using the `skip_provider_registration` field within the provider block).

!> **Note:** The errors returned from the Azure API when a Resource Provider is unregistered are unclear (example `API version '2019-01-01' was not found for 'Microsoft.Foo'`) - please ensure that all of the necessary Resource Providers you're using are registered - if in doubt **we strongly recommend letting Terraform register these for you**.

## Example Usage

```hcl
resource "azurerm_resource_provider_registration" "example" {
  name = "Microsoft.PolicyInsights"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The namespace of the Resource Provider which should be registered. Changing this forces a new resource to be created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when registering the Resource Provider.
* `read` - (Defaults to 5 minutes) Used when retrieving the Resource Provider.
* `delete` - (Defaults to 30 minutes) Used when unregistering the Resource Provider.

## Import

Resource Provider Registrations can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_resource_provider_registration.example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights
```
