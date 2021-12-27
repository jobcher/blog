---
title: "安装配置 Terraform"
date: 2021-12-27
draft: true
author: "sjtfreaks"
tags: ["Terraform"]
categories: ["日常"]
series: ["日常系列"]
---

# 安装配置 Terraform
## 安装
1. macOS 苹果系统安装
```sh
#安装
brew tap hashicorp/tap
brew install hashicorp/tap/terraform
# 更新
brew update
brew upgrade hashicorp/tap/terraform
#验证安装
terraform -help
```
2. windows 系统安装
```sh
#安装
choco install terraform
#直接到这个url里下载64位系统
https://www.terraform.io/downloads
#验证安装
terraform -help
```

3. Linux 安装
```sh
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
sudo apt-get update && sudo apt-get install terraform
#验证安装
terraform -help
```

## 配置
配置这个terraform我们这个需要持续更新，首先我们先配置Azure吧
### Azure 配置
1. 安装azurecli
```sh
# linux
curl -L https://aka.ms/InstallAzureCli | bash
apt install azure-cli
# macOS
brew update && brew install azure-cli
```
2. 登录azure
```sh
# 中国区azure
az cloud set --name AzureCloud
az login -u <账户> -p <密码>

#海外azure
az cloud set --name AzureChinaCloud
az login -u <账户> -p <密码>
```

3. 创建terrafrom代码
创建main.tf
```tf
# 正在使用的 Azure 提供程序源和版本
terraform {
    required_version = ">=0.12"

    required_providers {
        azurerm = {
            source = "hashicorp/azurerm"
            version = "~>2.0"
        }
    }
}
# 配置 Microsoft Azure 提供程序
provider "azurerm" {
  features {}
}

# 资源组前缀
resource "random_pet" "rg-name" {
  prefix    = var.resource_group_name_prefix
}

# 创建资源组
resource "azurerm_resource_group" "rg" {
  name      = random_pet.rg-name.id
  location  = var.resource_group_location
}
```
创建variable.tf
```tf
variable "resource_group_name_prefix" {
  default       = "rg"
  description   = "Prefix of the resource group name that's combined with a random ID so name is unique in your Azure subscription."
}

variable "resource_group_location" {
  default = "eastus"
  description   = "Location of the resource group."
}
```