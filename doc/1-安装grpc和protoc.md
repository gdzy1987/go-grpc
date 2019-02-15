## grpc
### 按照grpc
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
## protobuf
- https://github.com/protocolbuffers/protobuf
### go protobuf
- https://github.com/golang/protobuf

## 安装protoc
```
protoc是protobuf文件（.proto）的编译器

```
- 下载对应的releases版本 https://github.com/protocolbuffers/protobuf/releases
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
- 定位
```
which protoc
```
- 文件目录
```
tree


.
├── bin
│   └── protoc
├── include
│   └── google
│       └── protobuf
│           ├── any.proto
│           ├── api.proto
│           ├── compiler
│           │   └── plugin.proto
│           ├── descriptor.proto
│           ├── duration.proto
│           ├── empty.proto
│           ├── field_mask.proto
│           ├── source_context.proto
│           ├── struct.proto
│           ├── timestamp.proto
│           ├── type.proto
│           └── wrappers.proto
└── readme.txt
```

### 使用
- protoc 的命令格式为 protoc [OPTION] PROTO_FILES   （最后是待编译的 proto文件）
- 示例1
```
$ protoc --go_out=./go/ ./proto/helloworld.proto 
拆分为3块
protoc                      protoc命令
--go_out=./go/              输出go代码的目录，指定./go/目录
./proto/helloworld.proto    指定proto文件的位置，./proto/helloworld.proto
```
- 示例2
```
protoc -I helloworld/ --go_out=plugins=grpc:helloworld helloworld/helloworld.proto

```