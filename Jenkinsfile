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
                            '# 这是消息内容的标题',
                            '消息正文：测试 markdown 类型的消息',
                            '',
                            '---',
                            '我有分割线，哈哈哈哈',
                            '> 引用内容',
                            '#### 展示列表',
                            '- 两个黄鹂鸣翠柳',
                            '- 一行白鹭上青天',
                            '- 窗含西岭千秋雪',
                            '- 门泊东吴万里船'
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