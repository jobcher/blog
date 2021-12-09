#/!bin/bash
echo "删除旧的docker"
docker ps
docker stop nginx-hugo
docker rm nginx-hugo
docker rmi nginx:hugo
echo "生成新的docker"
hugo -t LoveIt -D
docker build -t nginx:hugo .
docker run --name nginx-hugo -d -p 8080:80 --restart=always nginx:hugo 
echo "显示端口"
netstat -lntp