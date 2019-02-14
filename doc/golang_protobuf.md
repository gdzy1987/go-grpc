## protobuf
- https://github.com/golang/protobuf

## readme
### Go support for Protocol Buffers - Google's data interchange format
```
Go支持协议缓冲区-谷歌的数据交换格式
```
```
Google's data interchange format. Copyright 2010 The Go Authors. https://github.com/golang/protobuf
This package and the code it generates requires at least Go 1.9.
This software implements Go bindings for protocol buffers. 
For information about protocol buffers themselves, see https://developers.google.com/protocol-buffers/

谷歌的数据交换格式。Go作者版权所有。https://github.com/golang/protobuf
这个包及其生成的代码至少需要Go 1.9。
这个软件为协议缓冲区实现Go绑定。
有关协议缓冲区本身的信息，请参见https://developers.google.com/protocol-buffers/

```
### Installation 安装
```
To use this software, you must:
    Install the standard C++ implementation of protocol buffers from 
    https://developers.google.com/protocol-buffers/
    Of course, install the Go compiler and tools from https://golang.org/ 
    See https://golang.org/doc/install for details or, if you are using gccgo, 
    follow the instructions at https://golang.org/doc/install/gccgo
    Grab the code from the repository and install the proto package. 
    The simplest way is to run go get -u github.com/golang/protobuf/protoc-gen-go. 
    The compiler plugin, protoc-gen-go, will be installed in $GOPATH/bin unless $GOBIN is set. 
    It must be in your $PATH for the protocol compiler, protoc, to find it.
    If you need a particular version of protoc-gen-go (e.g., to match your proto package version), 
    one option is

    GIT_TAG="v1.2.0" # change as needed
    go get -d -u github.com/golang/protobuf/protoc-gen-go
    git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
    go install github.com/golang/protobuf/protoc-gen-go

This software has two parts: a 'protocol compiler plugin' that generates Go source files that, 
once compiled, can access and manage protocol buffers; 
and a library that implements run-time support for encoding (marshaling), 
decoding (unmarshaling), and accessing protocol buffers.
There is support for gRPC in Go using protocol buffers. 
See the note at the bottom of this file for details.
There are no insertion points in the plugin.


要使用本软件，你必须:
    从https://developers.google.com/protocol-buffers/安装协议缓冲区的标准c++实现
    当然，从https://golang.org/安装Go编译器和工具，
    详细信息请参见https://golang.org/doc/install，或者，如果您正在使用gccgo，
    请遵循https://golang.org/doc/install/gccgo的说明
    从存储库中获取代码并安装proto包。
    最简单的方法是运行
        go get -u github.com/golang/protobuf/protoc-gen-go。
    除非设置了$GOBIN，否则编译器插件protoc-gen-go将安装在$GOPATH/bin中。
    如果您需要特定版本的  protoc-gen-go (例如，匹配您的原型包版本)，一个选项是

        GIT_TAG="v1.2.0" #根据需要进行更改
        go -get -u github.com/golang/protobuf/protoc-gen-go
        git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
        go install github.com/golang/protobuf/protoc-gen-go

该软件由两部分组成:
    一个“协议编译器插件”，生成Go源文件，一旦编译完成，就可以访问和管理协议缓冲区;
    以及实现对编码(编组)、解码(解组)和访问协议缓冲区的运行时支持的库。
Go中支持使用协议缓冲区的gRPC。有关详细信息，请参见该文件底部的说明。
插件中没有插入点。
```