pipeline {
    agent none
    stages {
        
        stage ('编译镜像'){
            agent { Dockerfile true }
            steps {
                echo "生成 docker 构建"
                sh 'docker stop docker-hugo'
                sh 'docker rm docker-hugo'
                sh 'docker run --name docker-hugo -d -p 80:80 --restart=always docker-hugo:latest'
                echo "结束 end"
            }
        }
    }
}