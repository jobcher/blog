---
title: "GNU/Linux 一键更换系统软件源脚本"
date: 2023-09-21
draft: true
featuredImage: "/images/gnulinux-logo-gnu-linux-logo-png.png"
featuredImagePreview: "/images/gnulinux-logo-gnu-linux-logo-png.png"
images: ["/images/gnulinux-logo-gnu-linux-logo-png.png"]
author: "jobcher"
tags: ["运维"]
categories: ["基础"]
series: ["基础知识系列"]
---

## 背景
有很多小伙伴，在接受公司老项目后，想要更新服务器源

## 执行命令
### 国内使用
```bash
bash <(curl -sSL https://www.jobcher.com/ChangeMirrors.sh)
```
### 教育网使用
```bash
bash <(curl -sSL https://www.jobcher.com/ChangeMirrors.sh) --edu
```
### 海外使用
```bash
bash <(curl -sSL https://www.jobcher.com/ChangeMirrors.sh) --abroad
```

## 注意事项
1. 需使用 `Root` 用户执行脚本
    切换命令为 `sudo -i` 或 `su root`，不同系统环境使用的命令不一样，因为有些系统没有在初始安装时为 `Root` 用户设置固定密码所以需要使用 `sudo` 指令来提权
2. 建议使用 `SSH` 远程工具
    如果你使用的系统终端界面不支持 UTF-8 编码那么将无法正常显示中文内容，导致无法正确选择交互内容。大部分系统都会自动开启该服务
3. 首次在新系统上执行脚本
    当前执行方式依赖 `curl` 指令来获取脚本内容并执行，所以需要先通过包管理工具来安装该软件包，否则会报错 `Command not found`，若无法安装就复制源码到本地新建.sh脚本，然后通过 bash 手动执行
