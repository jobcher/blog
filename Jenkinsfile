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
        }
    }
    post {
        success {
            script {
                env.DATETIME = sh(script:"date", returnStdout: true).trim()
                def job_name = "# ${JOB_NAME} 流水线 执行成功"
                def jenkinsid = """构建:  第 ${BUILD_ID} 次执行"""
                def JEN_production = "> 部署节点： k8s"
                def build_url = "> 部署详情： [详情](${BUILD_URL})"
                def jen_date = "> 执行时间： ${env.DATETIME}"

                dingtalk (
                    robot: '23bec93a-babe-486e-8f2f-f9486a6aac91',
                    type: 'MARKDOWN',
                    title: job_name,
                    text: [
                        job_name,
                        jenkinsid,
                        '',
                        '---',
                        JEN_production,
                        '',
                        build_url,
                        '',
                        jen_date
                        ],
                    at: [
                        '13250936269'
                        ]
                )            
            }
        }
        
        failure {
            script {
                env.DATETIME = sh(script:"date", returnStdout: true).trim()
                def job_name = "# ${JOB_NAME} 流水线 执行失败"
                def jenkinsid = """构建:  第 ${BUILD_ID} 次执行"""
                def JEN_production = "> 部署节点： k8s"
                def build_url = "> 部署详情： [详情](${BUILD_URL})"
                def jen_date = "> 执行时间： ${env.DATETIME}"

                dingtalk (
                    robot: '23bec93a-babe-486e-8f2f-f9486a6aac91',
                    type: 'MARKDOWN',
                    title: job_name,
                    text: [
                        job_name,
                        jenkinsid,
                        '',
                        '---',
                        JEN_production,
                        '',
                        build_url,
                        '',
                        jen_date
                        ],
                    at: [
                        '13250936269'
                        ]
                )            
            }
        }
    }
}