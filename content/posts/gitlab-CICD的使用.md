---
title: "gitlab CI/CD 的使用"
date: 2021-12-03
draft: true
author: "sjtfreaks"
tags: ["gitlab"]
---

# gitlab CI/CD 的使用
我将使用gitlab的流水线自动实现hugo blog 文章的自动发布。
  
## 一、基础知识


## 二、安装过程

### 1.安装gitlab runner
首先需要安装 gitlab runner 进入服务器A  
安装方法：  
1. 容器部署
2. 手动二进制文件部署
3. 通过rpm/deb包部署

1. docker方式安装

安装文档：https://docs.gitlab.com/runne...

    docker run -dit \
    --name gitlab-runner \
    --restart always \
    -v /srv/gitlab-runner/config:/etc/gitlab-runner \
    -v /var/run/docker.sock:/var/run/docker.sock \
    gitlab/gitlab-runner
1.1 设置信息

    docker exec -it gitlab-runner gitlab-runner register
2. 非docker方式安装

2.1 安装GitLab Runner

安装环境：Linux  

其他环境参考：https://docs.gitlab.com/runne...  

下载  
  
    curl -L --output /usr/local/bin/gitlab-runner https://gitlab-runner-downloads.s3.amazonaws.com/latest/binaries/gitlab-runner-linux-amd64
添加权限  

    chmod +x /usr/local/bin/gitlab-runner  
新建gitlab-runner用户  

    sudo useradd --comment 'GitLab Runner' --create-home gitlab-runner --shell /bin/bash
安装  

安装时需要指定我们上面新建的用户  

    gitlab-runner install --user=gitlab-runner --working-directory=/home/gitlab-runner
启动  
    gitlab-runner start

### 2.配置 docker shell链接
    ssh-keygen -t rsa
    cd .ssh/
    cat id_rsa.pub >>authorized_keys
    docker cp id_rsa gitlab-runner:/root

### 3.配置.gitlab-ci.yml文件
    vim .gitlab-ci.yml
  
    stages:          # List of stages for jobs, and their order of execution
    - build
    - test
    - deploy

    build-job:       # This job runs in the build stage, which runs first.
    stage: build
    script:
        - echo "上传代码"
        - sudo ssh -i /root/id_rsa root@172.17.0.2 && cd /opt/blog && sh gitpull.sh
        - echo "上传完成."

    unit-test-job:   # This job runs in the test stage.
    stage: test    # It only starts when the job in the build stage completes successfully.
    script:
        - echo "Running unit tests... This will take about 60 seconds."
        - sleep 60
        - echo "Code coverage is 90%"

    lint-test-job:   # This job also runs in the test stage.
    stage: test    # It can run at the same time as unit-test-job (in parallel).
    script:
        - echo "Linting code... This will take about 10 seconds."
        - sleep 10
        - echo "No lint issues found."

    deploy-job:      # This job runs in the deploy stage.
    stage: deploy  # It only runs when *both* jobs in the test stage complete successfully.
    script:
        - echo "Deploying application..."
        - echo "Application successfully deployed."

