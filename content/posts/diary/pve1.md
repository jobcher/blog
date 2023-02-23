---
title: "Proxmox VE 在线扩容磁盘分区"
date: 2022-01-19
draft: true
author: "jobcher"
tags: ["扩容磁盘", "pve"]
categories: ["日常"]
series: ["日常系列"]
---

# Proxmox VE 在线扩容磁盘分区

## 添加磁盘大小

![pve1-1](/images/pve1.png)

## 在 VM 上做扩容操作

1. 安装 growpart

```sh
yum install -y epel-release
yum install -y cloud-utils
```

2. 查看系统盘 路径

```sh
fdisk -l
df -h
```

![pve1-2](/images/pve2.png)

3. 扩容设备并重启

```sh
growpart /dev/sda 2 #2代表是第二块系统分区，不是sda2,中间有空格
reboot
```

4. 重启执行命令

```sh
xfs_growfs /dev/sda2 #(xfs 文件系统)
resize2fs /dev/sda2 #(ext4 文件系统)

```

5. 更新完成

```sh
df -h
```
