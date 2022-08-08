---
title: "systemd 守护命令"
date: 2022-08-08
draft: true
author: "jobcher"
tags: ["运维"]
categories: ["基础"]
series: ["基础知识系列"]
---
## 介绍
systemd 是linux中用来启动守护进程，Linux最早一直采用init进程  
  
![systemd](/static/images/systemd.png)  
(systemd 架构图)  
## systemd 命令
systemd 不是一个具体的命令，而是一组命令，用于系统管理的各个方面  
### 1.systemctl
`systemctl`是 Systemd 的主命令，用于管理系统。  
```sh
# 重启系统
$ sudo systemctl reboot

# 关闭系统，切断电源
$ sudo systemctl poweroff

# CPU停止工作
$ sudo systemctl halt

# 暂停系统
$ sudo systemctl suspend

# 让系统进入冬眠状态
$ sudo systemctl hibernate

# 让系统进入交互式休眠状态
$ sudo systemctl hybrid-sleep

# 启动进入救援状态（单用户状态）
$ sudo systemctl rescue
```
### 2.systemd-analyze
`systemd-analyze`命令用于查看启动耗时  
```sh
# 查看启动耗时
systemd-analyze                                                   

# 查看每个服务的启动耗时
$ systemd-analyze blame

# 显示瀑布状的启动过程流
$ systemd-analyze critical-chain

# 显示指定服务的启动流
$ systemd-analyze critical-chain atd.service
```
### 3.hostnamectl
`hostnamectl`命令用于查看当前主机的信息。
```sh
# 显示当前主机的信息
$ hostnamectl

# 设置主机名。
$ sudo hostnamectl set-hostname jobcher
```

### 4.localectl
`localectl`命令用于查看本地化设置
```sh
# 查看本地化设置
$ localectl

# 设置本地化参数。
$ sudo localectl set-locale LANG=en_GB.utf8
$ sudo localectl set-keymap en_GB
```
### 5.timedatectl
`timedatectl`命令用于查看当前时区设置
```sh
# 查看当前时区设置
$ timedatectl

# 显示所有可用的时区
$ timedatectl list-timezones                                                                                   

# 设置当前时区
$ sudo timedatectl set-timezone America/New_York
$ sudo timedatectl set-time YYYY-MM-DD
$ sudo timedatectl set-time HH:MM:SS
```
### 6.loginctl
`loginctl`命令用于查看当前登录的用户
```sh
# 列出当前session
$ loginctl list-sessions

# 列出当前登录用户
$ loginctl list-users

# 列出显示指定用户的信息
$ loginctl show-user ruanyf
```

## Unit
Systemd 可以管理所有系统资源。不同的资源统称为 Unit（单位）。  
|分类|资源|
|:----|:----|
|Service unit|系统服务|
|Target unit|多个 Unit 构成的一个组|
|Device Unit|硬件设备|
|Mount Unit|文件系统的挂载点|
|Automount Unit|自动挂载点|
|Path Unit|文件或路径|
|Scope Unit|不是由 Systemd 启动的外部进程|
|Slice Unit|进程组|
|Snapshot Unit|Systemd 快照，可以切回某个快照|
|Socket Unit|进程间通信的 socket|
|Swap Unit|swap 文件|
|Timer Unit|定时器|
  
### 1.systemctl list-units
`systemctl list-units`命令可以查看当前系统的所有 Unit  
```sh
# 列出正在运行的 Unit
$ systemctl list-units

# 列出所有Unit，包括没有找到配置文件的或者启动失败的
$ systemctl list-units --all

# 列出所有没有运行的 Unit
$ systemctl list-units --all --state=inactive

# 列出所有加载失败的 Unit
$ systemctl list-units --failed

# 列出所有正在运行的、类型为 service 的 Unit
$ systemctl list-units --type=service
```
### 2.Unit 的状态
`systemctl status`命令用于查看系统状态和单个 Unit 的状态。
```sh
# 显示系统状态
$ systemctl status

# 显示单个 Unit 的状态
$ sysystemctl status bluetooth.service

# 显示远程主机的某个 Unit 的状态
$ systemctl -H root@rhel7.example.com status httpd.service
```
除了`status`命令，`systemctl`还提供了三个查询状态的简单方法，主要供脚本内部的判断语句使用。
```sh
# 显示某个 Unit 是否正在运行
$ systemctl is-active application.service

# 显示某个 Unit 是否处于启动失败状态
$ systemctl is-failed application.service

# 显示某个 Unit 服务是否建立了启动链接
$ systemctl is-enabled application.service
```
### 3.Unit 管理
对于用户来说，最常用的是下面这些命令，用于启动和停止 Unit（主要是 `service`）。
```sh
# 立即启动一个服务
$ sudo systemctl start apache.service

# 立即停止一个服务
$ sudo systemctl stop apache.service

# 重启一个服务
$ sudo systemctl restart apache.service

# 杀死一个服务的所有子进程
$ sudo systemctl kill apache.service

# 重新加载一个服务的配置文件
$ sudo systemctl reload apache.service

# 重载所有修改过的配置文件
$ sudo systemctl daemon-reload

# 显示某个 Unit 的所有底层参数
$ systemctl show httpd.service

# 显示某个 Unit 的指定属性的值
$ systemctl show -p CPUShares httpd.service

# 设置某个 Unit 的指定属性
$ sudo systemctl set-property httpd.service CPUShares=500
```
### 4.依赖关系
Unit 之间存在依赖关系：A 依赖于 B，就意味着 `Systemd` 在启动 A 的时候，同时会去启动 B。  
`systemctl list-dependencies`命令列出一个 Unit 的所有依赖。
```sh
systemctl list-dependencies nginx.service
```
上面命令的输出结果之中，有些依赖是 Target 类型（详见下文），默认不会展开显示。如果要展开 Target，就需要使用--all参数。

```sh
systemctl list-dependencies --all nginx.service
```
## Unit 的配置文件
每一个 Unit 都有一个配置文件，告诉 Systemd 怎么启动这个 Unit  
Systemd 默认从目录`/etc/systemd/system/`读取配置文件。但是，里面存放的大部分文件都是符号链接，指向目录`/usr/lib/systemd/system/`，真正的配置文件存放在那个目录。  
`systemctl enable`命令用于在上面两个目录之间，建立符号链接关系。  
  
```sh
$ sudo systemctl enable jobcher.service
# 等同于
$ sudo ln -s '/usr/lib/systemd/system/jobcher.service' '/etc/systemd/system/multi-user.target.wants/jobcher.service'
```
如果配置文件里面设置了开机启动，`systemctl enable`命令相当于激活开机启动。  
  
与之对应的，`systemctl disable`命令用于在两个目录之间，撤销符号链接关系，相当于撤销开机启动。  
  
配置文件的后缀名，就是该 Unit 的种类，比如`sshd.socket`。如果省略，Systemd 默认后缀名为`.service`，所以`sshd`会被理解成`sshd.service`  

### 1.配置文件的状态
`systemctl list-unit-files`命令用于列出所有配置文件。
```sh
# 列出所有配置文件
$ systemctl list-unit-files

# 列出指定类型的配置文件
$ systemctl list-unit-files --type=service
```
这个列表显示每个配置文件的状态，一共有四种。  
|状态|连接|
|:----|:----|
|enabled|已建立启动链接|
|disabled|没建立启动链接|
|static|该配置文件没有[Install]部分（无法执行），只能作为其他配置文件的依赖|
|masked|该配置文件被禁止建立启动链接|
  
一旦修改配置文件，就要让 `SystemD` 重新加载配置文件，然后重新启动，否则修改不会生效。  
```sh
$ sudo systemctl daemon-reload
$ sudo systemctl restart httpd.service
```