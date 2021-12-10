---
title: "docker 命令"
date: 2021-12-09
draft: true
author: "sjtfreaks"
tags: ["docker"]
categories: ["docker"]
series: ["docker入门系列"]
---

# 安装docker
通过docker 脚本安装  
  
    curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
    curl -sSL https://get.daocloud.io/docker | sh

# docker命令
常用docker命令  
  
    #查看容器
    docker ps
    #查看镜像
    docker images
    #停止当前所有容器
    docker stop $(docker ps -aq)
    #删除当前停止的所有容器
    docker rm $(docker ps -aq)
    #删除镜像
    docker rmi nginx
