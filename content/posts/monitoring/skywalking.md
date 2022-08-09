---
title: "skywalking APM 监控"
date: 2022-08-08
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
<img src="https://www.jobcher.com/images/skywalkning.png" width="100%">  
`APM `(Application Performance Management) 即应用性能管理，属于IT运维管理（ITOM)范畴.  
分为一下三个方面：  

- Logging  
服务在处理某个请求时打印的错误日志，可以将这些日志信息记录到`Elasticsearch`或是其他存储中。通过Kibana或是其他工具来分析这些日志了解服务的行为和状态，大多数情况下。日志记录的数据很分散，并且相互独立。例如错误日志，请求处理过程中关键步骤的日志等等。
- Metrics  
`Metric`是可以聚合的，例如为电商系统中每个HTTP接口添加一个计数器，计算每个接口的QPS，可以通过简单的加和计算得到系统的总负载情况。
- Tracing  
在微服务架构系统中一请求会经过很多服务处理，调用链路会非常长，要确定中间哪个服务出现异常是非常麻烦的事情，通过分布式链路追踪，运维人员就可以构建一个请求的视图。视图上战术了一个请求从进入系统开始到返回响应的整个流程。
<img src="https://www.jobcher.com/images/skywalking2.svg" width="100%">

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