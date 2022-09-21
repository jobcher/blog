# 个人博客
基于hugo部署的个人博客，使用LoveIt主题

[![Tests on Linux, MacOS and Windows](https://github.com/jobcher/blog/workflows/Build/badge.svg)](https://github.com/jobcher/blog/actions?query=workflow%3ABuild)
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
docker build -t docker-hugo:latest .
docker run --name docker-hugo -d -p 4312:80 --restart=always docker-hugo:latest
```
## 代码结构
```sh
博客文章目录：content\posts
LoveIt配置目录：config.toml
```

## 个人网站地址
网站：[https://www.jobcher.com](https://www.jobcher.com)  
  
github.io：[https://jobcher.github.io](https://jobcher.github.io)