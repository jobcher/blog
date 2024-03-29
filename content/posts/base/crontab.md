---
title: "Linux crontab 命令"
date: 2022-02-22
draft: true
author: "jobcher"
tags: ["运维"]
categories: ["基础"]
series: ["基础知识系列"]
---

# Linux crontab 命令

Linux `crontab`是用来`定期执行程序`的命令。

1. `系统执行`的工作：系统周期性所要执行的工作，如`备份系统数据`、`清理缓存`
2. `个人执行`的工作：某个用户定期要做的工作，例如`每隔10分钟检查邮件服务器`是否有新信，这些工作可由每个用户自行设置

## 语法

```sh
crontab [ -u user ] file
crontab [ -u user ] { -l | -r | -e }
```

说明：  
`crontab` 是用来让使用者在`固定时间或固定间隔执行程序`之用，换句话说，也就是类似使用者的时程表。

`-u user` 是指设定指定 `user` 的时程表，这个前提是你必须要有其`权限`(比如说是 root)才能够指定他人的时程表。如果不使用 `-u user` 的话，就是表示`设定自己的时程表`。

参数说明：

`-e` : `执行文字编辑器来设定时程表`，内定的文字编辑器是 `VI`，如果你想用别的文字编辑器，则请先设定 VISUAL 环境变数来指定使用那个文字编辑器(比如说 `setenv VISUAL joe`)  
`-r` : `删除`目前的时程表  
`-l` : `列出`目前的时程表

时间格式如下：

```sh
f1  f2   f3   f4   f5    program
*    *    *    *    *
-    -    -    -    -
|    |    |    |    |
|    |    |    |    +----- 星期中星期几 (0 - 6) (星期天 为0)
|    |    |    +---------- 月份 (1 - 12)
|    |    +--------------- 一个月中的第几天 (1 - 31)
|    +-------------------- 小时 (0 - 23)
+------------------------- 分钟 (0 - 59)
```

其中 `f1` 是表示`分钟`，`f2` 表示`小时`，`f3` 表示`一个月份中的第几日`，`f4` 表示`月份`，`f5` 表示`一个星期中的第几天`。`program` 表示`要执行的程序`。
当 `f1` 为` *` 时表示`每分钟都要执行` program，`f2` 为 `*` 时表示`每小时都要执行程序`，其馀类推
当 `f1` 为 `a-b` 时表示从`第 a 分钟到第 b 分钟`这段时间内要`执行`，`f2` 为 `a-b` 时表示从`第 a 到第 b 小时`都要执行，其馀类推
当 `f1` 为 `*/n` 时表示`每 n 分钟个时间间隔执行一次`，`f2` 为 `*/n` 表示`每 n 小时个时间间隔执行一次`，其馀类推
当 `f1` 为 `a, b, c,...` 时表示`第 a, b, c,...` 分钟要`执行`，`f2` 为 `a, b, c,...` 时表示`第 a, b, c...个小时要执行`，其馀类推
