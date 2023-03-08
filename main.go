package main

import (
	// 前面加點下面的AddUserRouter就不用加上前綴
	"golangAPI/config"
	. "golangAPI/src"

	"github.com/gin-gonic/gin"
)

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
	router := gin.Default()
	// 接收群組的方法UserRouter.go裡有RouterGroup
	config.Connect()
	v1 := router.Group("/v1")
	AddUserRouter(v1)

	router.Run(":8000")

}
