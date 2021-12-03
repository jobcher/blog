---
title: "hugo 命令大全"
date: 2021-12-01T17:05:35+08:00
draft: true
author: "sjtfreaks"
tags: ["hugo"]
---
# hugo命令大全
## 安装hugo
### 二进制安装  
    brew install hugo  
### 源码安装  
    export GOPATH=$HOME/go  
    go get -v github.com/spf13/hugo
    go get -u -v github.com/spf13/hugo #更新依赖库

## 生成站点
    hugo new site /opt/blog
    cd /opt/blog

## 创建文章
    hugo new about.md
    vim about.md

    hugo new post/first.md

## 安装皮肤

    cd /opt/blog/themes
    git clone https://github.com/dillonzq/LoveIt.git

## 运行hugo
    hugo server -t LoveIt -D

## 部署
你要部署在github Page上  

    hugo --theme=hyde --baseUrl="http://coderzh.github.io/"  
    
    cd public
    $ git init
    $ git remote add origin https://github.com/coderzh/coderzh.github.io.git
    $ git add -A
    $ git commit -m "first commit"
    $ git push -u origin master


## 常用命令
nohup hugo server -e production -t LoveIt -D &  