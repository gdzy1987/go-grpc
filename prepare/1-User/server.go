package main

import (
	"context"
	pb "go-grpc/prepare/1-User/proto"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/clientopt"
	"time"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
)

var (
	client *mongo.Client
	database *mongo.Database
	collection *mongo.Collection
	err error
	cond *FindByName
	result *mongo.DocumentResult
	user *User
)

type User struct {
	ID int32 `bson:"id" json:""`
	Name string `bson:"name" json:"name"`
	Age int32 `bson:"age" json:"age"`
	Title []string `bson:"title" json:"title"`
}

//查询过滤
type FindByName struct {
	Name string `bson:"name" json:"name"`
}
//定义服务端实现约定的接口
type UserInfoService struct{}

var u = UserInfoService{}

func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error){
	name := req.Name
	fmt.Println("收到查询 name:",name)

	//在数据库中查找用户信息
	//测试数据文件user.json
	//1,建立连接
	if client, err = mongo.Connect(context.TODO(),"mongodb://127.0.0.1:27017",clientopt.ConnectTimeout(5*time.Second)); err != nil{
		fmt.Println(err)
		return
	}

	//2,选择数据库data
	database = client.Database("data")

	//3,选择表user
	collection = database.Collection("user")

	//4,按照name过滤
	cond = &FindByName{Name:name}

	//5，查询
	result = collection.FindOne(context.TODO(),cond)

	if err = result.Decode(&user); err != nil{
		fmt.Println(err)
		return
	}

	//处理返回
	resp = &pb.UserResponse{
		Id:user.ID,
		Name:user.Name,
		Age:user.Age,
		Title:user.Title,
	}
	err = nil
	return
}

func main()  {
	port := ":2333"
	listener, err := net.Listen("tcp",port)
	if err != nil{
		log.Fatalf("listen error:%v\n",err)
	}

	fmt.Printf("listen %s\n",port)
	s := grpc.NewServer()

	//注意第二个参数 UserInfoServiceServer是接口类型的变量,需要取地址传参
	pb.RegisterUserInfoServiceServer(s,&u)
	s.Serve(listener)
}
