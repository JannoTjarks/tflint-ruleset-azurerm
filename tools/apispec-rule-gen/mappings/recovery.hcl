mapping "azurerm_recovery_services_vault" {
  import_path = "azure-rest-api-specs/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2016-06-01/vaults.json"

  name                = VaultName
  resource_group_name = ResourceGroupName
  sku                 = Sku.name
}
