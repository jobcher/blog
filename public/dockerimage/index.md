# docker image镜像上传


# docker image镜像上传

登入docker hub，在https://hub.docker.com上注册你的账号。  
  
```bash
docker login
username：#输入你的用户名
password：#输入你的密码
```
## 上传镜像

```bash
docker tag nginx:hugo sjtfreaks/hogo-nginx:v1
docker push sjtfreaks/hogo-nginx:v1
```

