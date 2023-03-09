package pojo

import (
	"golangAPI/config"
)

// TODO

// 如果Model 是 User 在DB 就是Users
type User struct {
	Id int `binding:"required"`
	Name string `binding:"required,gt=4"`
	Password string `binding:"min=4,max=20,userpasd"`
	Email    string `binding:"email"`
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
	return result.RowsAffected > 0
}

func UpdateUser(userId string, user User) User{
	config.DB.Model(&user).Where("id = ?",userId).Updates(user)
	return user
}

func CheckUserPassword(name string, password string) User {
	user := User{}
	config.DB.Where("name = ? and password = ?", name, password).First(&user)
	return user
}


