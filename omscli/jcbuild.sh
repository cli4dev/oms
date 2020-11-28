#!/bin/sh

rm -rf ./out
#获取当前目录
publishsql="mysql"
if  [ $# -eq 1 ] && [ $1 = "oracle" ] ;then
	publishsql=$1
fi

echo "-----------编译的tags:$publishsql--------"
echo "-----------编译前检查引用包是否都存在--------"
go build
if [ $? -ne 0 ]; then
	echo "项目编译出错,请检查"
	exit 1
fi
# echo "当前生成的环境为: $publishenv" 
echo "------------------------------------"
echo "---------------集成部署打包开始--------------"
echo "-----------编译flowserver-oms项目--------"
cd ../its/jcserver
go build -tags $publishsql -o "../../omscli/out/flowserver/bin/flowserver-oms"
if [ $? -ne 0 ]; then
	echo "flowserver 项目编译出错,请检查"
	exit 1
fi

echo "-----------打包完成-------------"
echo "-----------都放在out目录中-------------"


