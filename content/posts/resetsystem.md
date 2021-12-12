---
title: "linux系统开启root权限"
date: 2021-12-12
draft: true
author: "sjtfreaks"
tags: ["linux"]
categories: ["日常"]
series: ["日常系列"]
---
# linux系统开启root权限
1. 修改ssh服务配置文件
```sh
shdo su -
sudo vim /etc/ssh/sshd_config
```

2. 增加权限
```bash
PermitRootLogin yes
```

3. 更改root密码，重启服务
```sh
sudo passwd root
service sshd restart
```ß