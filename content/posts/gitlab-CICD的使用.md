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

通过二进制文件方式安装
# Download the binary for your system
    sudo curl -L --output /usr/local/bin/gitlab-runner https://gitlab-runner-downloads.s3.amazonaws.com/latest/binaries/gitlab-runner-linux-amd64

# Give it permissions to execute
    sudo chmod +x /usr/local/bin/gitlab-runner

# Create a GitLab CI user
    sudo useradd --comment 'GitLab Runner' --create-home gitlab-runner --shell /bin/bash

# Install and run as service
    sudo gitlab-runner install --user=gitlab-runner --working-directory=/home/gitlab-runner
    sudo gitlab-runner start

注册runner的命令  
    sudo gitlab-runner register --url https://gitlab.com/ --registration-token $REGISTRATION_TOKEN

### 2.配置.gitlab-ci.yml文件
    vim .gitlab-ci.yml
  
    stages:          # List of stages for jobs, and their order of execution
    - build
    - test
    - deploy

    build-job:       # This job runs in the build stage, which runs first.
    stage: build
    script:
        - echo "上传代码"
        - cd /opt/blog && sh gitpull.sh
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

