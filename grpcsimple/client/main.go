package main

import (
	"context"
	"fmt"
	pb "grpcsimple/hellogrpc"
	"log"

	"google.golang.org/grpc"
)

const (
	addr string = "127.0.0.1:5656"
)

func main() {
	// 建立网络连接
	conn, err := grpc.DialContext(context.Background(), addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	cliet := pb.NewHelloSerivceClient(conn)
	reply, err := cliet.SayHello(context.Background(), &pb.HelloRequest{Name: "big wei brother"})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(reply)
}
