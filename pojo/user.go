package pojo

import (
	"golangAPI/config"
	"log"
)

// TODO

// 如果Model 是 User 在DB 就是Users
type User struct {
	Id       int    `json:"UserId" binding:"required"`
	Name     string `json:"UserName" binding:"required,gt=4"`
	Password string `json:"UserPassword" binding:"min=4,max=20,userpasd"`
	Email    string `json:"UserEmail" binding:"email"`
}

type Users struct {
	UserList []User `json:"UserList" binding:"required,gt=0,lt=3"`
	UserListSize int `json:"UserListSize"`
}


func FindAllUsers() []User {
	var users []User
	config.DB.Find(&users)
	return users
}

func FindByUserId(userId string) User {
	var user User 
	config.DB.Where("id = ?",userId).First(&user)
	return user
}

func CreateUser(user User) User{
	config.DB.Create(&user)
	return user
}


func DeleteUser(userId string) bool{
	user := User{}
	result := config.DB.Where("id = ?",userId).Delete(&user)
	log.Println(result)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func UpdateUser(userId string, user User) User{
	config.DB.Model(&user).Where("id = ?",userId).Updates(user)
	return user
}


