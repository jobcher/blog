#/!bin/bash
echo "删除旧的docker"
docker ps
docker stop nginx-hugo
docker rm nginx-hugo
docker rmi nginx:hugo
echo "生成新的docker"
docker build -t nginx:hugo .
docker run --name nginx-hugo -d -p 8080:1313 nginx:hugo
echo "显示端口"
netstat -lntp