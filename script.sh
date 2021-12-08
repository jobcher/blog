#!/bin/bash
echo "git操作："
echo " "
echo "1.拉取main代码"
echo "2.拉取devlop代码"
echo " "
read -p "请输入你想执行的命令（1-2）：" var


case ${var} in
1)
echo "拉取main代码"
git checkout main
git pull origin main
echo "结束"
;;
2)
echo "拉取devlop代码"
git checkout devlop
git pull origin develop
echo "结束"
;;
*)
echo "输入错误，请重新输入"
;;
esac