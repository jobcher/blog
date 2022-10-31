---
title: "brew 安装配置"
date: 2022-10-31
draft: true
author: "jobcher"
tags: ["brew"]
categories: ["Mac"]
series: ["Mac系列"]
---

# 关于macOS 13 brew软件失效
Warning: You are using macOS 13.  
We do not provide support for this pre-release version.  
You will encounter build failures with some formulae.  
Please create pull requests instead of asking for help on Homebrew's GitHub,
Twitter or any other official channels. You are responsible for resolving
any issues you experience while you are running this
pre-release version.  
> 简单来说就是macOS13版本 暂时不提供技术支持

## 解决方法
升级完macos13之后发现了比较麻烦的问题，很多软件出现了不兼容，这真的很无奈，对于我们这些做IT的人来说，这是致命的。我以git软件举例，有以下几个方法。  
### 1. 使用时间机器恢复备份
>使用前提：你之前备份了系统，并且系统正常
这个方法更加一劳永逸，因为我们并不确定还有什么软件不支持macos13

### 2. 重新安装xcode-select
```sh
xcode-select --install
```