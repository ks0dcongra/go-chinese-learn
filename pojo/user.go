package pojo

import (
	"golangAPI/config"
)

type User struct {
	// Id   int    `json:UserId`
	// Name string `json:UserName`
	// Password string `json:UserPassword`
	// Email string `json:UserEmail`

	Id   int    
	Name string 
	Password string 
	Email string 
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

