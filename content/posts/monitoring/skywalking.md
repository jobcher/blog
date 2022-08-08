---
title: "skywalking APM 监控"
date: 2022-07-25
draft: true
author: "jobcher"
tags: ["skywalking"]
categories: ["监控"]
series: ["运维监控系列"]
---
# skywalking 
基于OpenTracing规范，专门为微服务架构以及云原生服务。
## APM 监控
一个基于微服务架构的电商系统  
![shop](/images/skywalkning.png)  
`APM `(Application Performance Management) 即应用性能管理，属于IT运维管理（ITOM)范畴.  
分为一下三个方面：  

- Logging
- Metrics
- Tracing

## 使用docker-compose安装
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