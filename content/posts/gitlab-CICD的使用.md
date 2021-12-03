---
title: "gitlab CI/CD 的使用"
date: 2021-12-03
draft: true
author: "sjtfreaks"
tags: ["gitlab"]
---

# gitlab CI/CD 的使用
我将使用gitlab的流水线自动实现hugo blog 文章的自动发布。
  
## 一、基础知识


## 二、安装过程

### 1.安装gitlab runner
首先需要安装 gitlab runner 进入服务器A  
安装方法：  
1. 容器部署
2. 手动二进制文件部署
3. 通过rpm/deb包部署

我这边选择docker部署  
1. 拉取镜像  
    docker pull gitlab/gitlab-runner:latest

