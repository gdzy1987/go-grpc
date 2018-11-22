package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	pb "go-grpc/prepare/2-ShippingService/proto"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	ADDRESS           = "127.0.0.1:50051"
	DEFAULT_INFO_FILE = "consignment.json"
)

//读取consignment.json中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

func main() {
	//连接到grpc服务器
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect error:%v", err)
	}
	defer conn.Close()

	//初始化gprc客户端
	client := pb.NewShippingServiceClient(conn)

	//在命令行中指定新的货物信息json文件
	infoFile := DEFAULT_INFO_FILE
	if len(os.Args) > 1 {
		infoFile = os.Args[1]
	}

	consignment, err := parseFile(infoFile)
	if err != nil {
		log.Fatalf("parse info file error:%v", err)
	}

	//调用rpc
	//将货物存储到我们自己的仓库
	resp, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("create consignment error:%v", err)
	}

	//新货物是否托运成功
	fmt.Printf("created: %t", resp.Created)
	fmt.Println(resp.Consignment)
	fmt.Println(strings.Repeat("==", 30))
	//列出目前所有的货物
	resp, err = client.GetConsignments(context.Background(), &pb.GetRequest{})
	for k, c := range resp.Consignments {
		fmt.Printf("第%d次:%+v\n", k, c)
	}

}
