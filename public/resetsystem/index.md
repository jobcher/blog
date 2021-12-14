# linux系统开启root权限

# linux系统开启root权限
1. 修改ssh服务配置文件
```sh
shdo su -
sudo vim /etc/ssh/sshd_config
```
2. 增加权限  
在# Authentication: 下输入  
  
```bash
PermitRootLogin yes
```

3. 更改root密码，重启服务
```sh
sudo passwd root
service sshd restart
```
