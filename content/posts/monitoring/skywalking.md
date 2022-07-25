---
title: "skywalking 安装和配置"
date: 2022-07-25
draft: true
author: "jobcher"
tags: ["skywalking"]
categories: ["监控"]
series: ["运维监控系列"]
---
# skywalking 安装和配置
使用docker-compose安装
```yml
version: "3"
services:
  skywalking-oap-server:
    image: "apache/skywalking-oap-server:9.1.0"
    container_name: "oap-server"
    restart: "always"
    ports:
      - "10.12.12.4:12800:12800"
      - "10.12.12.4:1234:1234"
      - "10.12.12.4:11800:11800"

  skywalking-oap-ui:
    image: "apache/skywalking-ui:9.1.0"
    container_name: "oap-ui"
    restart: "always"
    environment:
      - SW_OAP_ADDRESS=http://10.12.12.4:12800
    ports:
      - "8180:8080"

```