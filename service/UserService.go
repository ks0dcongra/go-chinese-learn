package service

import (
	"golangAPI/config"
	"golangAPI/middlewares"
	"golangAPI/models"
	"golangAPI/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET
func FindAllUsers(c *gin.Context) {
	// c.JSON(http.StatusOK, userList)
	users := pojo.FindAllUsers()
	c.JSON(http.StatusOK,users)
}

func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 {
		c.JSON(http.StatusNotFound,"Error")
		return
	}
	log.Println("User ->", user)
	c.JSON(http.StatusOK,user)
}

// POST
func PostUser(c *gin.Context){
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable,"Error:" + err.Error())
		return
	}
	newUser := pojo.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

// DELETE
func DeleteUser(c *gin.Context){
	user := pojo.DeleteUser(c.Param("id"))
	if  !user {
		c.JSON(http.StatusNotFound,"Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

// PUT
func PutUser(c *gin.Context){
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest,"Error")
		return
	}
	user = pojo.UpdateUser(c.Param("id"),user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound,"Error")
		return
	}
	c.JSON(http.StatusOK,user)
}

//CreateUserList 
func CreateUserList(c *gin.Context){
	users := pojo.Users{}
	err := c.BindJSON(&users)
	if err != nil {
		c.String(400, "Error%s", err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func LoginUser(c *gin.Context){
	name := c.PostForm("name")
	password := c.PostForm("password")
	user := pojo.CheckUserPassword(name, password)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	middlewares.SaveSession(c, user.Id)
	c.JSON(http.StatusOK, gin.H{
		"message" : "Login Successfully",
		"User" : user,
		"Sessions": middlewares.GetSession(c),
	})
}

func LogoutUser(c *gin.Context){
	middlewares.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{
		"message" : "Logout Successfully",
	})
}

func CheckUserSession(c *gin.Context){
	sessionId := middlewares.GetSession(c)
	if sessionId == 0 {
		c.JSON(http.StatusUnauthorized, "Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message" : "Check Session Successfully",
		"User" : middlewares.GetSession(c),
	})
}


// 以下為英文視頻的Tutorial
func TestDbGetUsers(c *gin.Context){
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func TestDbPostUsers(c *gin.Context){
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func TestDbDeleteUsers(c *gin.Context){
	var user models.User
	config.DB.Where("id = ?",c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func TestDbPutUsers(c *gin.Context){
	var user models.User
	config.DB.Where("id = ?",c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}