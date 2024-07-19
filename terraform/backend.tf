terraform {
  required_version = ">= 1.8.4"

  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~>3.0"
    }
  }

  backend "azurerm" {
    resource_group_name  = "rg-techthemach-jpe-00"
    storage_account_name = "sttechthemachjpe01"
    container_name       = "tfstate"
    key                  = "terraform.tfstate"
  }
}

provider "azurerm" {
  features {}
  subscription_id = "7f174dd3-9580-404a-a527-1ae6144513ff"
}