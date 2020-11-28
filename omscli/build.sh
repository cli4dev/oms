#!/bin/sh

rm -rf ./out
#获取当前目录
publishenv="dev"
publishsql="mysql"
if  [ $# -eq 1 ] && [ $1 = "oracle" ] ;then
	publishsql=$1
fi

if [ $# -eq 1 ] &&  [ $1 = "prod" ] ;then
	publishenv=$1
fi

if  [ $# -eq 2 ] ;then
	if  [ $1 = "oracle" ] ;then
		publishsql=$1
	fi

	if  [ $1 = "prod" ] ;then
		publishenv=$1
	fi

	if  [ $2 = "oracle" ] ;then
		publishsql=$2
	fi

	if  [ $2 = "prod" ] ;then
		publishenv=$2
	fi
fi

publishtags="$publishenv $publishsql"
echo "-----------编译的tags:$publishtags--------"
echo "-----------编译前检查引用包是否都存在--------"
go build
if [ $? -ne 0 ]; then
	echo "项目编译出错,请检查"
	exit 1
fi
# echo "当前生成的环境为: $publishenv" 
echo "------------------------------------"
echo "---------------打包开始--------------"
echo "-----------编译apiserver-oms项目--------"
cd ../apiserver
go build -tags "$publishtags" -o "../omscli/out/apiserver/bin/apiserver-oms"
if [ $? -ne 0 ]; then
	echo "apiserver 项目编译出错,请检查"
	exit 1
fi

echo "-----------编译flowserver-oms项目--------"
cd ../flowserver
go build -tags "$publishtags" -o "../omscli/out/flowserver/bin/flowserver-oms"
if [ $? -ne 0 ]; then
	echo "flowserver 项目编译出错,请检查"
	exit 1
fi
echo "-----------打包完成-------------"
echo "-----------都放在out目录中-------------"


