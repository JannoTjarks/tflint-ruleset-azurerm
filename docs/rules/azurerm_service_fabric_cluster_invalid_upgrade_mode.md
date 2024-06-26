<!--- This file generated by `tools/apispec-rule-gen/main.go`. DO NOT EDIT --->

# azurerm_service_fabric_cluster_invalid_upgrade_mode

Warns about values that appear to be invalid based on [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs).

Allowed values are:
- Automatic
- Manual

## Example

```hcl
resource "azurerm_service_fabric_cluster" "foo" {
  upgrade_mode = ... // invalid value
}
```

```
$ tflint
1 issue(s) found:

Error: "..." is an invalid value as upgrade_mode (azurerm_service_fabric_cluster_invalid_upgrade_mode)

  on template.tf line 2:
  2:   upgrade_mode = ... // invalid value

Reference: https://github.com/terraform-linters/tflint-ruleset-azurerm/blob/v0.1.0/docs/rules/azurerm_service_fabric_cluster_invalid_upgrade_mode.md

```

## Why

Requests containing invalid values will return an error when calling the API by `terraform apply`.

## How to Fix

Replace the warned value with a valid value.

## Source

This rule is automatically generated from [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs). If you are uncertain about the warning, check the following API schema referenced by this rule.

https://github.com/Azure/azure-rest-api-specs/tree/master/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/cluster.json
