package src

import (
	session "golangAPI/middlewares"
	"golangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup){
	user := r.Group("/users", session.SetSession())

	user.GET("/", service.FindAllUsers)
	user.GET("/:id", service.FindByUserId)
	user.POST("/", service.PostUser)
	// user.DELETE("/:id", service.DeleteUser)
	user.PUT("/:id", service.PutUser)
	user.POST("/more", service.CreateUserList)

	// Login
	user.POST("/login", service.LoginUser)

	user.GET("/check", service.CheckUserSession)
	user.Use(session.AuthSession())
	{
		// delete user
		user.DELETE("/:id", service.DeleteUser)
		// LogoutUser
		user.GET("/logout", service.LogoutUser)
	}

	// 以下為英文視頻的Tutorial，youtube course to CRUD to connect postgres
	user.GET("/testdb", service.TestDbGetUsers)
	user.POST("/testdb", service.TestDbPostUsers)
	user.DELETE("/testdb/:id", service.TestDbDeleteUsers)
	user.PUT("/testdb/:id", service.TestDbPutUsers)
}