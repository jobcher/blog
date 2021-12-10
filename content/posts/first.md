---
title: "hugo 命令大全"
date: 2021-12-01T17:05:35+08:00
draft: true
author: "sjtfreaks"
tags: ["hugo"]
categories: ["hugo"]
series: ["hugo入门系列"]
---
# hugo命令大全
## 安装hugo
### 二进制安装  
```sh
    brew install hugo  
```
### 源码安装  
```sh
    export GOPATH=$HOME/go  
    go get -v github.com/spf13/hugo
    go get -u -v github.com/spf13/hugo #更新依赖库
```
## 生成站点
```sh
    hugo new site /opt/blog
    cd /opt/blog
```
## 创建文章
```sh
    hugo new about.md
    vim about.md

    hugo new post/first.md
```
## 安装皮肤
```sh
    cd /opt/blog/themes
    git clone https://github.com/dillonzq/LoveIt.git
```
## 运行hugo
```sh
    hugo server -t LoveIt -D
```
## 部署
你要部署在github Page上  
```sh
    hugo --theme=hyde --baseUrl="http://coderzh.github.io/"  
    
    cd public
    $ git init
    $ git remote add origin https://github.com/coderzh/coderzh.github.io.git
    $ git add -A
    $ git commit -m "first commit"
    $ git push -u origin master
```

## hugo 添加搜索插件

```toml
[outputs]
  home = ["HTML", "RSS", "JSON"]
```
### 搜索配置
```toml
[params.search]
  enable = true
  # 搜索引擎的类型 ("lunr", "algolia")
  type = "lunr"
  # 文章内容最长索引长度
  contentLength = 4000
  # 搜索框的占位提示语
  placeholder = ""
  # LoveIt 新增 | 0.2.1 最大结果数目
  maxResultLength = 10
  # LoveIt 新增 | 0.2.3 结果内容片段长度
  snippetLength = 50
  # LoveIt 新增 | 0.2.1 搜索结果中高亮部分的 HTML 标签
  highlightTag = "em"
  # LoveIt 新增 | 0.2.4 是否在搜索索引中使用基于 baseURL 的绝对路径
  absoluteURL = false
  [params.search.algolia]
    index = ""
    appID = ""
    searchKey = ""
```

## 常用命令
```sh
nohup hugo server -e production -t LoveIt -D &  
```