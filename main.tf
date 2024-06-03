terraform {
  required_version = ">= 1.8.4"
  
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>3.0"
    }
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
    owner = "ryouta-arisaka"
    terraform   = "true"
  }
}

resource "azurerm_resource_group" "rg" {
    name     = "rg-cft-openai-arisaka"
    location = var.location

    tags = var.tag
}

resource "azurerm_api_management" "apim" {
  name                = "apim-cft-openai-arisaka"
  location            = var.location
  resource_group_name = azurerm_resource_group.rg.name
  publisher_name      = "Ryota Arisaka"
  publisher_email     = "ryouta-arisaka@jfe-systems.com"

  sku_name = "Developer_1"
}