##
- 示例来源
```
 article: https://wuyin.io/2018/05/02/protobuf-with-grpc-in-golang/
 github: https://github.com/wuYin/blog/tree/master/codes/protobuf-with-grpc-in-golang
```

### 目录结构
```
├── proto
│   ├── user.proto		// 定义客户端请求、服务端响应的数据格式
│   └── user.pb.go		// protoc 为 gRPC 生成的读写数据的函数
├── server.go			// 实现微服务的服务端
|-
└── client.go			// 调用微服务的客户端
```

### 步骤
- 创建proto文件夹
- 在proto文件夹下创建user.proto,定义UserInfoService微服务
```
syntax = "proto3";   //指定语法格式，注意proto3不再支持proto2的required和optional
package proto;       //指定生成的user.pb.go的包名，防止命名冲突


//service 定义开放调用的服务，即UserInfoService微服务
service UserInfoService{
    //rpc定义服务内的GetUserInfo远程调用
    rpc GetUserInfo (UserRequest) returns (UserResponse){
    }
}


//message对应生成代码的struct
//定义客户端请求的数据格式
message UserRequest{
    //[修饰符] 类型 字段名 = 标识符;
    string name = 1;
}


//定义服务端响应的数据格式
message UserResponse{
    int32 id = 1;
    string name = 2;
    int32 age = 3;
    repeated string title = 4;  //repeated修饰符表示字段是可变数组，即slice类型
}
```
- 编译 user.proto文件
```
protoc 编译器的grpc插件会处理 service字段定义的 UserInfoService
使service能编码，解码message
```
```
protoc -I . --go_out=plugins=grpc:. ./user.proto
```