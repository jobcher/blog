#!/bin/bash
echo "重启hugo"
killall hugo
sleep 10
nohup hugo server -e production -t LoveIt -D &
ps -ef | grep hugo
netstat -lntp
echo "运行结束"