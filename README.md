# 个人博客
基于hugo部署的个人博客，使用LoveIt主题

[![GoDoc](https://godoc.org/github.com/gohugoio/hugo?status.svg)](https://godoc.org/github.com/jobcher/blog)
[![Tests on Linux, MacOS and Windows](https://github.com/gohugoio/hugo/workflows/Test/badge.svg)](https://github.com/jobcher/blog/actions?query=workflow%3ATest)
[![Go Report Card](https://goreportcard.com/badge/github.com/gohugoio/hugo)](https://goreportcard.com/report/github.com/jobcher/blog)
[![StandWithUkraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/badges/StandWithUkraine.svg)](https://github.com/vshymanskyy/StandWithUkraine/blob/main/docs/README.md)  
  
[![Stand With Ukraine](https://raw.githubusercontent.com/vshymanskyy/StandWithUkraine/main/banner2-direct.svg)](https://vshymanskyy.github.io/StandWithUkraine/)  
  
## 通过docker快速部署
```sh
# 安装docker
curl -sSL https://get.daocloud.io/docker | sh
# 开机docker 自启
systemctl enable docker
systemctl start docker
# docker检查
docker version
# 运行博客
sh updatedocker.sh
```
## 代码结构
```sh
博客文章目录：content\posts
LoveIt配置目录：config.toml
```

## 个人网站地址
网站：[https://www.jobcher.com](https://www.jobcher.com)  
  
github.io：[https://jobcher.github.io](https://jobcher.github.io)