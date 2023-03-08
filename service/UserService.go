package service

import (
	"golangAPI/pojo"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

// GET
func FindAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}
// POST
func PostUser(c *gin.Context){
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable,"Error:" + err.Error())
		return
	}
	userList = append(userList,user)
	c.JSON(http.StatusOK, "Successfully posted")
}

// DELETE
func DeleteUser(c *gin.Context){
	userId,_ := strconv.Atoi(c.Param("id"))
	for _, user := range userList {
		log.Println(user)
		if user.Id == userId {
			userList = append(userList[:userId],userList[userId+1:]...)
			c.JSON(http.StatusOK, "Successfully deleted")
			return
		}
	}
	c.JSON(http.StatusNotFound,"Error")
}

// PUT
func PutUser(c *gin.Context){
	beforeUser := pojo.User{}
	err := c.BindJSON(&beforeUser)
	if err != nil {
		c.JSON(http.StatusBadRequest,"Error")
	}
	userId, _ := strconv.Atoi(c.Param("id"))
	for key, user := range userList{
		if userId == user.Id {
			userList[key] = beforeUser
			log.Println(userList[key])
			c.JSON(http.StatusOK, "Successfully updated")
			return
		}
	}
	c.JSON(http.StatusNotFound,"Error")

}