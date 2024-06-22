terraform {
  required_version = ">= 1.8.4"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>3.0"
    }
  }

  backend "azurerm" {
    resource_group_name  = "rg-cft-openai-arisaka"
    storage_account_name = "stcftopenaiarisaka"
    container_name       = "tfstate"
    key                  = "terraform.tfstate"
  }
}

provider "azurerm" {
  features {}
}

variable "location" {
  type    = string
  default = "japaneast"
}

variable "tag" {
  type = map(string)
  default = {
    owner     = "ryouta-arisaka"
    terraform = "true"
  }
}

resource "azurerm_resource_group" "rg" {
  name     = "rg-cft-openai-arisaka"
  location = var.location

  tags = var.tag
}

resource "azurerm_api_management" "apim" {
  name                = "apim-cft-openai-arisaka2"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  publisher_name      = "ryouta-arisaka"
  publisher_email     = "ryouta-arisaka@jfe-systems.com"
  sku_name            = "Consumption_0"

  identity {
    type         = "SystemAssigned"
  }

  tags = var.tag
}

resource "azurerm_cognitive_account" "openai" {
  name                = "aoai-cft-openai-arisaka"
  resource_group_name = azurerm_resource_group.rg.name
  location            = "West US"
  custom_subdomain_name = "aoai-cft-openai-arisaka"
  kind                = "OpenAI"
  sku_name            = "S0"

  tags = var.tag
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

# DB for Postgressql
resource "azurerm_postgresql_flexible_server" "pg" {
  name                = "pg-cft-openai-arisaka"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  version                       = "12"
  sku_name   = "B_Standard_B1ms"
  administrator_login           = "postgres"
  administrator_password        = "postgres"
  zone                          = "1"
  storage_mb   = 32768
  storage_tier = "P30"
  public_network_access_enabled = true
  tags = var.tag
}