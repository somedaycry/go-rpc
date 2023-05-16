package main

import (
	"context"
	"errors"
	"fmt"
	pb "grpcsimple/hellogrpc"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

const (
	addr string = "127.0.0.1:5656"
)

func main() {
	// 监听服务
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("start listen 127.0.0.1:5656...")
	if err != nil {
		log.Fatalln(err)
	}
	// 创建grpc实例
	s := grpc.NewServer()

	pb.RegisterHelloSerivceServer(s, &HelloServiceServce{})
	log.Println("grpc server is running in 127.0.0.1:5656")
	err = s.Serve(listener)

	if err != nil {
		log.Println(err)
	}

}

type HelloServiceServce struct {
	pb.UnimplementedHelloSerivceServer
}

func (h *HelloServiceServce) SayHello(ctx context.Context, request *pb.HelloRequest) (response *pb.HelloResponse, err error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("failed to get peer from context")
	}
	remoteaddr := peer.Addr.String()
	log.Printf("%s,connected...\n", remoteaddr)

	return &pb.HelloResponse{
		Reply: fmt.Sprint("hello ", request.Name),
	}, nil
}
