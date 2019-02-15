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
    
        go get -u github.com/golang/protobuf/protoc-gen-go
        
    除非设置了$GOBIN，否则编译器插件protoc-gen-go将安装在$GOPATH/bin中。
    如果您需要特定版本的  protoc-gen-go (例如，匹配您的原型包版本)，一个选项是

        GIT_TAG="v1.2.0" #根据需要进行更改
        go get -d -u github.com/golang/protobuf/protoc-gen-go
        git -C "$(go env GOPATH)"/src/github.com/golang/protobuf checkout $GIT_TAG
        go install github.com/golang/protobuf/protoc-gen-go

该软件由两部分组成:
    一个“协议编译器插件”，生成Go源文件，一旦编译完成，就可以访问和管理协议缓冲区;
    以及实现对编码(编组)、解码(解组)和访问协议缓冲区的运行时支持的库。
Go中支持使用协议缓冲区的gRPC。有关详细信息，请参见该文件底部的说明。
插件中没有插入点。
```

## Using protocol buffers with Go 对go使用协议缓冲区
```
Once the software is installed, there are two steps to using it. 
First you must compile the protocol buffer definitions and then import them, 
with the support library, into your program.
To compile the protocol buffer definition, 
run protoc with the --go_out parameter set to the directory you want to output the Go code to.
    
    protoc --go_out=. *.proto

The generated files will be suffixed .pb.go. See the Test code below for an example using such a file.



一旦安装了软件，使用它有两个步骤。
    首先，必须编译 protocol buffer 定义，然后使用支持库将它们导入程序。
    要编译协议缓冲区定义，请将 —go_out 参数设置为要输出Go代码的目录运行protoc。

        protoc --go_out=. *.proto

生成的文件将以.pb.go为后缀。
有关使用此类文件的示例，请参见下面的测试代码。
```
## Packages and input paths 包和输入路径
```
The protocol buffer language has a concept of "packages" which does not correspond well to the Go notion of packages. 
In generated Go code, each source .proto file is associated with a single Go package. 
The name and import path for this package is specified with the go_package proto option:

    option go_package = "github.com/golang/protobuf/ptypes/any";

The protocol buffer compiler will attempt to derive a package name and import path if a go_package option is not present, 
but it is best to always specify one explicitly.

There is a one-to-one relationship between source .proto files and generated .pb.go files, 
but any number of .pb.go files may be contained in the same Go package.

The output name of a generated file is produced by replacing the .proto suffix with .pb.go (e.g., foo.proto produces foo.pb.go). 
However, the output directory is selected in one of two ways. 
Let us say we have inputs/x.proto with a go_package option of github.com/golang/protobuf/p. 
The corresponding output file may be:

    Relative to the import path:

  protoc --go_out=. inputs/x.proto
  # writes ./github.com/golang/protobuf/p/x.pb.go

(This can work well with --go_out=$GOPATH.)

    Relative to the input file:

protoc --go_out=paths=source_relative:. inputs/x.proto
# generate ./inputs/x.pb.go



协议缓冲区语言有一个“包”的概念，这个概念与Go包的概念不太相符。
在生成的Go代码中，每个源.proto文件都与单个Go包相关联。
这个包的名称和导入路径是用go_package proto选项指定的:

    option go_package = "github.com/golang/protobuf/ptypes/any";

如果没有go_package选项，协议缓冲区编译器将尝试派生包名和导入路径，
但是最好总是显式地指定一个。

在源.proto文件和生成的.pb之间存在一对一的关系,
但任何数量的 .pb.go文件可能包含在同一个go包中。

生成的文件的输出名称是通过将.proto后缀替换为.pb.go生成的。(例如, foo.proto 生成 foo.pb.go)。
但是，选择输出目录有两种方法。
假设我们有input/x.proto的go_package选项为github.com/golang/protobuf/p。
对应的输出文件可以是:

    相对于导入路径:

    protoc --go_out=. inputs/x.proto
    #写 ./github.com/golang/protobuf/p/x.pb.go

    (这可以很好地使用   --go_out=$GOPATH.)

    相对于输入文件:

    protoc --go_out=paths=source_relative:. inputs/x.proto
    #生成 ./inputs/x.pb.go
```

## Generated code 生成代码