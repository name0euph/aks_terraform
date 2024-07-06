################################################################
# Resource Group
################################################################

resource "azurerm_resource_group" "rg" {
  name     = "rg-${var.workload}"
  location = var.location

  tags = var.tags
}


################################################################
# Azure OpenAI Service
################################################################

resource "azurerm_cognitive_account" "openai" {
  name                  = "aoai-${var.workload}"
  resource_group_name   = azurerm_resource_group.rg.name
  location              = "West US"
  custom_subdomain_name = "aoai-${var.workload}"
  kind                  = "OpenAI"
  sku_name              = "S0"

  tags = var.tags
}

resource "azurerm_cognitive_deployment" "gpt-35-turbo" {
  name                 = "gpt-35-turbo"
  cognitive_account_id = azurerm_cognitive_account.openai.id
  model {
    format  = "OpenAI"
    name    = "gpt-35-turbo"
    version = "1106"
  }
  scale {
    capacity = "10"
    type     = "Standard"
  }
  rai_policy_name = "Microsoft.Default"

  depends_on = [
    azurerm_cognitive_account.openai,
  ]
}

resource "azurerm_cognitive_deployment" "gpt-4o" {
  name                 = "gpt-4o"
  cognitive_account_id = azurerm_cognitive_account.openai.id
  model {
    format  = "OpenAI"
    name    = "gpt-4o"
    version = "2024-05-13"
  }
  scale {
    capacity = "10"
    type     = "Standard"
  }
  rai_policy_name = "Microsoft.Default"

  depends_on = [
    azurerm_cognitive_account.openai,
  ]
}

resource "azurerm_cognitive_deployment" "gpt-4" {
  name                 = "gpt-4"
  cognitive_account_id = azurerm_cognitive_account.openai.id
  model {
    format  = "OpenAI"
    name    = "gpt-4"
    version = "1106-Preview"
  }
  scale {
    capacity = "10"
    type     = "Standard"
  }
  rai_policy_name = "Microsoft.Default"

  depends_on = [
    azurerm_cognitive_account.openai,
  ]
}

################################################################
# PostgreSQL
################################################################

resource "azurerm_postgresql_flexible_server" "pg" {
  name                          = "pg-${var.workload}"
  resource_group_name           = azurerm_resource_group.rg.name
  location                      = azurerm_resource_group.rg.location
  version                       = "12"
  sku_name                      = "B_Standard_B1ms"
  administrator_login           = "postgres"
  administrator_password        = "postgres"
  zone                          = "1"
  storage_mb                    = 32768
  storage_tier                  = "P4"
  public_network_access_enabled = true
  tags                          = var.tags
}

################################################################
# Azure Container Registry
################################################################

resource "azurerm_container_registry" "acr" {
  name                = "acr${var.workload}"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  sku                 = "Basic"
  admin_enabled       = false
  tags                = var.tags
}