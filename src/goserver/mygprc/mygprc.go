package mygprc

import (
	//Grpc
	"google.golang.org/grpc"
	pb "goserver/pb"
	"net"
	"google.golang.org/grpc/reflection"
	"fmt"
	"golang.org/x/net/context"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func Run(port string){
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterGoGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	fmt.Println("start Grpc server port %v", port)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}