package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"io/ioutil"

	"goserver/mygprc"
)



func main() {
	fmt.Println("hello world")

	version := "V0"
	t, err := ioutil.ReadFile("/app/setting")
	if err == nil {
		fmt.Println("read file success")
		version  = string(t)
	}
	fmt.Println("app version:", version)

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
	r.GET("/health", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	go r.Run(":" + port)

	mygprc.Run(grpcport)

}