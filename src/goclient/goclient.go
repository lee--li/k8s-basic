package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	pb "goclient/pb"
	"google.golang.org/grpc"
)

func main() {

	port := "8080"
	grpcport := "8180"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("GRPCPORT") != "" {
		grpcport = os.Getenv("GRPCPORT")
	}
	
	fmt.Println("hello world")

	// 连接服务器
	conn, err := grpc.Dial(":8180", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()
	grpcclient := pb.NewGoGreeterClient(conn)

	r := gin.Default()
	
	r.GET("/hello", func(c *gin.Context) {
		// 调用服务端的SayHello
		r, err := grpcclient.SayHello(context.Background(), &pb.HelloRequest{Name: "go"})
		if err != nil {
			fmt.Printf("could not greet: %v", err)
		}
		fmt.Printf("Greeting: %s !\n", r.Message)
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": r.Message,
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
 	r.Run(":" + port)
}