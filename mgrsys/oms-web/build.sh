
#!/usr/bin/env bash
npm run build
go build
mv ./mgrweb ./dist
cd dist
./mgrweb install -r zk://192.168.0.101 -c mgrweb
./mgrweb run -r zk://192.168.0.101 -c mgrweb
#cp -r /root/work/img ./static/static/
#scp -r static root@[服务器地址]:[服务器路径]
