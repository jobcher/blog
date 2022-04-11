---
title: "RocketMQ 安装和启动"
date: 2022-01-07
draft: true
author: "jobcher"
tags: ["RocketMQ"]
categories: ["日常"]
series: ["日常系列"]
---
# RocketMQ 安装和部署
部署RocketMQ

## 单机安装构建
1. 安装JDK 1.8.0
```sh
yum install java-1.8.0-openjdk*
```

2. 安装Maven
```sh
wget http://dlcdn.apache.org/maven/maven-3/3.8.4/binaries/apache-maven-3.8.4-bin.tar.gz
tar -zxvf apache-maven-3.8.4-bin.tar.gz
mv -f apache-maven-3.8.4 /usr/local/
vim /etc/profile
# 末尾添加
export MAVEN_HOME=/usr/local/apache-maven-3.8.4
export PATH=${PATH}:${MAVEN_HOME}/bin
# 保存
source /etc/profile
# 查看maven是否正常
mvn -v
```

## 快速部署
```sh
#构建 DLedger
git clone https://github.com/openmessaging/openmessaging-storage-dledger.git
cd openmessaging-storage-dledger
mvn clean install -DskipTests
# 构建 RocketMQ
git clone https://github.com/apache/rocketmq.git
cd rocketmq
git checkout -b store_with_dledger origin/store_with_dledger
mvn -Prelease-all -DskipTests clean install -U
# 部署
cd rocketmq/distribution/target/apache-rocketmq
sh bin/dledger/fast-try.sh start
# 通过 mqadmin 运维命令查看集群状态
sh bin/mqadmin clusterList -n 127.0.0.1:9876
```

启动单节点  
```sh
cd distribution/target/rocketmq-4.9.3-SNAPSHOT/rocketmq-4.9.3-SNAPSHOT
nohup sh bin/mqnamesrv &
# 查看 Namesrv 日志
tail -f ~/logs/rocketmqlogs/namesrv.log
2022-01-07 14:59:29 INFO main - The Name Server boot success. serializeType=JSON
# 启动 Broker
nohup sh bin/mqbroker -c conf/broker.conf  -n 127.0.0.1:9876 &
# 查看 Broker 日志
tail -f ~/logs/rocketmqlogs/broker.log
```


如果提示找不到上面的日志文件，应该是没启动成功。  
应该是内存不够，RocketMQ默认用8g内存，如果你服务器的内存比较小，可以修改下bin/runbroker.sh脚本，将 Broker JVM 内存调小。如：JAVA_OPT="${JAVA_OPT} -server -Xms2g -Xmx2g -Xmn1g"。  
再次启动broker，可以正常启动。  
  
默认情况下，Broker 日志文件所在地址为~/logs/rocketmqlogs/broker.log。如果想要自定义，可以通过conf/logback_broker.xml配置文件来进行修改。  