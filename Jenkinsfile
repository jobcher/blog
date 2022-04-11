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
                sh 'docker run --name docker-hugo -d -p 4312:80 --restart=always docker-hugo:latest'
                echo "结束 end"
            }
        }
    }
}