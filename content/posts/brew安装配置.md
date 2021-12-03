---
title: "brew 安装配置"
date: 2021-12-03
draft: true
author: "sjtfreaks"
tags: ["brew"]
---
# brew 安装配置
## 一.安装
### 1.在ubuntu上安装 brew
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  
### 2.在centos上安装 brew
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  
### 3.在MacOS上安装 brew
    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  
## 二、使用
### 1.安装 wget
    brew install wget

Homebrew 会将软件包安装到独立目录，并将其文件软链接至 /usr/local  
  
    $ cd /usr/local
    $ find Cellar
    Cellar/wget/1.16.1
    Cellar/wget/1.16.1/bin/wget
    Cellar/wget/1.16.1/share/man/man1/wget.1

    $ ls -l bin
    bin/wget -> ../Cellar/wget/1.16.1/bin/wget
  
### 2.创建你自己的 Homebrew 包
  
    $ brew create https://foo.com/bar-1.0.tgz
    Created /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core/Formula/bar.rb

### 3.撤销你的变更或与上游更新合并
    $ brew edit wget # 使用 $EDITOR 编辑!
