package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Id int `json:"Id" gorm:"primary_key"`
	// Name string `json:"name"`
	// Email string `json:"email"`
	// Password string `json:"password"`
	Id int
	Name string 
	Email string
	Password string
}