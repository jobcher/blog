---
title: "Oracle Instant Client 安装配置实现远程连接oracle"
date: 2022-11-15
draft: true
author: "jobcher"
tags: ["oracle"]
categories: ["数据库"]
series: ["oracle系列"]
---
## 背景
关于Oracle 数据库一直是许多初学者比较头疼的地方，一方面受限于线上文档比较少，令一方面在企业中不得不接触和使用Oracle数据库，这篇文章是教大家如何通过配置oracle client来远程访问Oracle数据库。本文会通过python3和cx_Oracle来实现对Oracle的访问和增删改查
## 下载oracle客户端
[官方地址下载](https://www.oracle.com/database/technologies/instant-client/linux-x86-64-downloads.html)

### 安装
下载并安装你的oracle client，因为我连接的11g oracle，所以下载11.2版本
```sh
# 下载
wget https://download.oracle.com/otn/linux/instantclient/11204/oracle-instantclient11.2-basic-11.2.0.4.0-1.x86_64.rpm
# 安装
rpm -ivh oracle-instantclient11.2-basic-11.2.0.4.0-1.x86_64.rpm
```
### 配置环境变量
```sh
# 直接运行
export ORACLE_HOME=/usr/lib/oracle/11.2/client64
export ORABIN=/usr/lib/oracle/11.2/client64/bin
```
```sh
# 编辑环境变量配置文件
vim /etc/profile
```
```sh
# 底部增加内容
export PATH USER LOGNAME MAIL HOSTNAME HISTSIZE HISTCONTROL
export ORACLE_HOME=/usr/lib/oracle/11.2/client64
export TNS_ADMIN=/usr/lib/oracle/11.2/client64
export LD_LIBRARY_PATH=/usr/lib/oracle/11.2/client64/lib
export ORABIN=/usr/lib/oracle/11.2/client64/bin
PATH=$PATH:$ORABIN
export PATH

export PATH=$ORACLE_HOME:$PATH
export PATH=$PATH:$HOME/bin:$ORACLE_HOME/bin
```
```sh
# 刷新环境变量
source /etc/profile
```

## 下载cx_Oracle
```sh
pip3 install cx_Oracle
```
创建Oracle.py文件
```py
#!/usr/bin/python3
"""
使用python 对oracle数据进行操作
提前安装好 cx_Oracle
pip3 install cx_Oracle
"""
import cx_Oracle

conn = cx_Oracle.connect('用户名','密码','IP/SN')
cursor = conn.cursor()
sql = 'SELECT * FROM test_table'
cursor.execute(sql)
res = cursor.fetchall()
print (res)

```

## 执行文件
>python3 oracle.py
```sh
#输出内容
[none,0]
```