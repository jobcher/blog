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

                    script {
                        env.USER = sh(script:"git --no-pager show -s -n 1 --format='%cn'", returnStdout: true).trim()
                        env.BRANCH = sh(script:"git --no-pager show -s -n 1 --format='%D'", returnStdout: true).trim()
                        env.DATETIME = sh(script:"git --no-pager show -s -n 1 --format='%ci'", returnStdout: true).trim()
                        env.COMMIT_MESSAGE = sh(script:"git --no-pager show -s -n 1 --format='%B'", returnStdout: true).trim()
                        def jenkinsid = """构建:  第 ${BUILD_DISPLAY_NAME} 执行"""
                        def jenkinscommitmessage = """构建说明:  ${env.COMMIT_MESSAGE}"""
                        def jenkinsbuildid ="""${BUILD_ID}"""
                        def jenkinstime="""> 提交日期： ${env.DATETIME}  """
                        def jenkinsbranch="""> 构建分支： ${env.BRANCH}  """
                        def jenkinsuser="""> 提交者： ${env.USER}  """
                        dingtalk (
                            robot: '23bec93a-babe-486e-8f2f-f9486a6aac91',
                            type: 'MARKDOWN',
                            title: '流水线执行成功',
                            text: [
                                '# jobcher-blog-github-CI 流水线',
                                jenkinsid,
                                '',
                                '---',
                                '![success](https://i.ibb.co/hgRznG7/chenggong.png)',
                                '',
                                jenkinsbranch,
                                jenkinsuser,
                                jenkinstime,
                                '#### 更新内容',
                                '',
                                jenkinscommitmessage,
                            ],
                            at: [
                                '13250936269'
                            ]
                        )
                    }
            

                }

                failure{

                }
            }
        }
    }
}