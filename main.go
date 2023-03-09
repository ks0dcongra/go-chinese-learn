package main

import (
	// 前面加點下面的AddUserRouter就不用加上前綴
	"golangAPI/config"
	"golangAPI/middlewares"
	"golangAPI/src"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// 創建Gin Log 日誌
func setupLogging(){
	f, _  := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {
	//最簡單環境配置
	// router := gin.New()
	// config.Connect()
	// routes.UserRoute(router)
	// router.Run(":8000")
	// 最基礎表達GET && POST 的code
	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "ping",
	// 	})
	// })
	// router.POST("/ping/:id", func(c *gin.Context) {
	// 	id := c.Param("id")
	// 	c.JSON(200, gin.H{
	// 		"id": id,
	// 	})
	// })
	
	// 產生Log
	setupLogging()

	router := gin.Default()

	// BasicAuth可以驗證帳號
	// 設置middlewares，可以將滑鼠停在func上面會注意到他要回傳的東西是handler，所以可以在middlewares對應要回傳的參數。 
	router.Use(gin.BasicAuth(gin.Accounts{"Tom": "123456"}), middlewares.Logger())
	
	// 接收群組的方法UserRouter.go裡有RouterGroup
	config.Connect()
	v1 := router.Group("/v1")
	src.AddUserRouter(v1)

	router.Run(":8000")

}
