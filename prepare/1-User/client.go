package main

import (
	"google.golang.org/grpc"
	"log"
	pb "go-grpc/prepare/1-User/proto"
	"context"
	"fmt"
)

func main()  {
	conn, err := grpc.Dial(":2333",grpc.WithInsecure())
	if err != nil{
		log.Fatalf("dial error: %v\n",err)
	}
	defer conn.Close()

	client := pb.NewUserInfoServiceClient(conn)

	//调用服务
	req := new(pb.UserRequest)
	req.Name = "张三"

	resp, err := client.GetUserInfo(context.Background(),req)
	if err != nil{
		log.Fatalf("resp error:%v\n",err)
	}

	fmt.Println("id:",resp.Id)
	fmt.Println("name:",resp.Name)
	fmt.Println("age:",resp.Age)
	fmt.Println("title:",resp.Title)
}