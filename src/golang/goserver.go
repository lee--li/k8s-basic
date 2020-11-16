package main

import "fmt"
import "github.com/gin-gonic/gin"
import "os"

func main() {
	fmt.Println("hello world")

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
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
	r.Run(":" + port)
}