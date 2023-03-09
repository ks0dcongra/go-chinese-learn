package pojo

import (
	"golangAPI/config"
	"log"
)

type User struct {
	Id       int    `json:UserId`
	Name     string `json:UserName`
	Password string `json:UserPassword`
	Email    string `json:UserEmail`

	// Id   int    
	// Name string 
	// Password string 
	// Email string 
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


