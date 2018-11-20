## 安装grpc
- 使用go get
```
go get google.golang.org/grpc
```
- 如果报错
```
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git   $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git  $GOPATH/src/golang.org/x/text
go  get        -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
cd $GOPATH/src/
go install google.golang.org/grpc
```

## 安装protoc
- 下载 https://github.com/protocolbuffers/protobuf/releases
- mac版本 https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-osx-x86_64.zip
- 解压 在文件夹bin目录里有protoc文件
- 使用如下命令检查是否安装成功
```
protoc --version
```
- 如果没有安装成功，可以把当前路径添加到环境变量中
```
用pwd查看路径
pwd

vi .bashrc
添加如下

export PROTOCPATH="xxxx路径"
export PATH="$PATH:$PROTOCPATH/bin"
```