---
title: "CocoaPods 安装及碰到问题"
date: 2023-07-31
draft: true
featuredImage: "/images/cocoapods-orange-on-grey-1.jpg"
featuredImagePreview: "/images/cocoapods-orange-on-grey-1.jpg"
author: "jobcher"
tags: ["IOS 开发"]
categories: ["IOS 开发"]
series: ["IOS 开发"]
---
## 背景
CocoaPods 是OS X和IOS 下的第三类库管理工具，通过CocoaPods工具我们可以为项目添加被称为`Pods`的依赖库  

## 检查环境
```
ruby -v
gem -v
```
### 出现异常问题
> /System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin22/rbconfig.rb:21: warning: Insecure world writable dir /opt/homebrew/bin in PATH, mode 040777  
  
该警告信息表明在你的PATH环境变量中包含了一个“不安全可写”（Insecure world writable）的目录`/opt/homebrew/bin`。这可能会导致潜在的安全问题。  
为了解决这个警告，你需要修复`/opt/homebrew/bin`目录的权限，以使其不再被标记为“不安全可写”。  
#### 解决问题
```
chmod 755 /opt/homebrew/bin
chmod 755 /opt/homebrew
chmod 755 /opt/homebrew/sbin
```

## 安装cocoapods
输入安装命令
```sh
sudo gem install cocoapods
```
### 出现异常问题
>ERROR:  While executing gem ... (Gem::FilePermissionError)
    You don't have write permissions for the /System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/gems/2.6.0 directory.
        /Library/Ruby/Site/2.6.0/rubygems/installer.rb:714:in `verify_gem_home'
        /Library/Ruby/Site/2.6.0/rubygems/installer.rb:904:in `pre_install_checks'
        ……  
  
在 `macOS` 系统中，系统的Ruby目录通常是受保护的，并且普通用户没有对这些目录进行写操作的权限。为了解决这个问题，你应该避免在系统级别的Ruby目录中进行Gem的安装。相反，你应该使用用户级别的Gem安装目录。
#### 解决方案
```sh
mkdir -p ~/.gem/ruby/2.6.0
export PATH="$HOME/.gem/ruby/2.6.0/bin:$PATH"
source ~/.bash_profile
```
```sh
gem cocoapods install --user-install
```