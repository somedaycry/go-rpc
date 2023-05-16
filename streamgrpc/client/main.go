package main

import (
	"context"
	"fmt"
	"log"
	"streamgrpc/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:5656", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewStreamGrpcClient(conn)

	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	// 发送数据
	go func(){
		for {
			err := stream.Send(&pb.Request{Name: "big wei brother"})
			if err != nil{
				// 错误可能来是服务器关闭或者是客户端发送失败
				log.Fatalln(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		resp,err := stream.Recv()
		if err!=nil{
			log.Println(err)
			break
		}
		fmt.Println(resp)
 	}

}
