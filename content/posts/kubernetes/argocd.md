---
title: "Argo cd 安装和部署"
date: 2021-12-24
draft: true
author: "jobcher"
tags: ["gitlab"]
categories: ["gitlab"]
series: ["git入门系列"]
---
# Argo cd 安装和部署
Argo CD 是一个为 Kubernetes 而生的，遵循声明式 GitOps 理念的持续部署（CD）工具。Argo CD 可在 Git 存储库更改时自动同步和部署应用程序
![argocd](/images/argocd-1.png)
## 安装
k8s快速安装
```sh
k3s kubectl create namespace argocd
k3s kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```
## 使用
![argocd-2](/images/argocd_2.png)