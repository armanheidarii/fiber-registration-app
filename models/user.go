package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Age      uint   `json:"age"`
}
