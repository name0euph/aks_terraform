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

# Cosmos DB
resource "azurerm_cosmosdb_account" "cosmosdb" {
  name                = "cosmosdb-cft-openai-arisaka2"
  resource_group_name = azurerm_resource_group.rg.name
  location            = azurerm_resource_group.rg.location
  offer_type          = "Standard"
  kind                = "GlobalDocumentDB"
  consistency_policy {
    consistency_level = "Session"
  }
  geo_location {
    location          = azurerm_resource_group.rg.location
    failover_priority = 0
  }
  capabilities {
    name = "EnableServerless"
  }
  tags = var.tag
}

resource "azurerm_cosmosdb_sql_database" "sql" {
  name                = "sqldb"
  resource_group_name = azurerm_resource_group.rg.name
  account_name        = azurerm_cosmosdb_account.cosmosdb.name
}

resource "azurerm_cosmosdb_sql_container" "users" {
  name                = "Users"
  resource_group_name = azurerm_resource_group.rg.name
  account_name        = azurerm_cosmosdb_account.cosmosdb.name
  database_name       = azurerm_cosmosdb_sql_database.sql.name
  partition_key_path  = "/email"
  partition_key_version = 1
}