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

通过docker方式安装
    docker run -dit \
    --name gitlab-runner \
    --restart always \
    -v /srv/gitlab-runner/config:/etc/gitlab-runner \
    -v /var/run/docker.sock:/var/run/docker.sock \
    gitlab/gitlab-runner

设置信息  
    docker exec -ti gitlab-runner gitlab-runner register

    Enter the GitLab instance URL (for example, https://gitlab.com/):
    https://gitlab.com/ #地址
    Enter the registration token:
    #token令牌
    Enter a description for the runner:
    [c********4]: gitlabrunner #描述
    Enter tags for the runner (comma-separated):
    hugo #tag标识
    Registering runner... succeeded                     runner=_e********
    Enter an executor: ssh, virtualbox, docker-ssh+machine, docker-ssh, parallels, shell, docker+machine, kubernetes, custom, docker:
    shell #选择你的方式
    Runner registered successfully. Feel free to start it, but if it's running already the config should be automatically reloaded!

### 2.配置