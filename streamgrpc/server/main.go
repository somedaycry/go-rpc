package main

import (

	"fmt"
	"io"
	"log"
	"net"
	"streamgrpc/pb"

	"google.golang.org/grpc"
)

func main() {

	listen, err := net.Listen("tcp4", ":5656")
	log.Println("listen localhost:5656...")
	if err != nil {
		log.Fatalln(err)
	}

	// 创建服务端
	s := grpc.NewServer()
	log.Println("grpcservice is running...")
	pb.RegisterStreamGrpcServer(s,&stremserver{})
	s.Serve(listen)

}

type stremserver struct {
	pb.UnimplementedStreamGrpcServer
}

func (s *stremserver) Channel(stream pb.StreamGrpc_ChannelServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("client closed...")
				return nil
			}
			return err
		}

		//	处理请求
		resp := &pb.Response{
			Reply: fmt.Sprint("hello ", req.GetName()),
		}
		err = stream.Send(resp)
		if err !=nil{
			if err == io.EOF{
				log.Println("client is closed...")
			}
		}


	}

}
