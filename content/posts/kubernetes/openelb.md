---
title: "OpenELB：让k8s私有环境对外暴露端口"
date: 2022-04-12
draft: true
author: "jobcher"
tags: ["k8s"]
categories: ["k8s"]
series: ["k8s入门系列"]
---

# OpenELB：云原生负载均衡器插件
OpenELB 是一个开源的云原生负载均衡器实现，可以在基于裸金属服务器、边缘以及虚拟化的 Kubernetes 环境中使用 LoadBalancer 类型的 Service 对外暴露服务。
## 在 Kubernetes 中安装 OpenELB
```sh
kubectl apply -f https://raw.githubusercontent.com/openelb/openelb/master/deploy/openelb.yaml
```
- 查看状态
```sh
kubectl get po -n openelb-system
```

## 使用 kubectl 删除 OpenELB
```sh
kubectl delete -f https://raw.githubusercontent.com/openelb/openelb/master/deploy/openelb.yaml
```
```sh
kubectl get ns
```

## 配置 OpenELB
