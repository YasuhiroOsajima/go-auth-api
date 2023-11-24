package database

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Disabled bool   `gorm:"default:false" json:"disabled"`
}

func NewUserOrm(username string, password string, disabled bool) User {
	return User{Username: username, Password: password, Disabled: disabled}
}
