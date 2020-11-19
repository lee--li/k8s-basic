package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	//Grpc
	"google.golang.org/grpc"
	pb "goserver/pb"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	fmt.Println("hello world")

	port := "8080"
	grpcport := "8180"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("GRPCPORT") != "" {
		grpcport = os.Getenv("GRPCPORT")
	}
	
	r := gin.Default()
	
	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "goserver Hello world!",
		})
	})
	r.GET("/", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "goserver Hello world root!",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	go r.Run(":" + port)

	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":" + grpcport)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterGoGreeterServer(s, &server{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	fmt.Println("start Grpc server port %v", grpcport)
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}