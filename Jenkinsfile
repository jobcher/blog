pipeline {
    agent any
    stages {
        stage ('编译镜像'){
            steps {
                echo "编译镜像"
                sh 'docker build -t docker-hugo:latest .'
                echo "结束 end"
            }
        }
        stage ('部署镜像'){
            steps {
                echo "部署镜像"
                sh 'docker stop docker-hugo'
                sh 'docker rm docker-hugo'
                sh 'docker run --name docker-hugo -d -p 4312:80 --restart=always docker-hugo:latest'
                echo "结束 end"
            }
        }
        stage ('上传制品库'){
            steps {
                echo "上传镜像到制品库"
                sh 'docker tag docker-hugo:latest hub.jobcher.com/blog/hugo:latest'
                sh 'docker push hub.jobcher.com/blog/hugo:latest'
                echo "结束 end"
            }
            post {
                success {
                    dingtalk (
                        robot: '23bec93a-babe-486e-8f2f-f9486a6aac91',
                        type: 'MARKDOWN',
                        title: '流水线执行成功',
                        text: [
                            '# 流水线执行成功',
                            '执行流水线：jobcher-blog-github-CI  ',
                            '![logo](https://jobcher.github.io/avatar1.png)',
                            '',
                            '---',
                            '更新内容',
                            '> 执行流水线：jobcher-blog-github-CI',
                            '#### 执行内容',
                            '- 编译镜像',
                            '- 部署镜像',
                            '- 上传制品库'
                        ],
                        at: [
                          '13250936269'
                        ]
                    )
                }
            }
        }
    }
}