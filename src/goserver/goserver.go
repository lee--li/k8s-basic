package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"goserver/mygprc"
)



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

	mygprc.Run(grpcport)

}