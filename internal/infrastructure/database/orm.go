package database

import (
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/YasuhiroOsajima/go-auth-api/internal/interfaces/database"
)

type Orm struct{}

func NewOrm() *Orm {
	return &Orm{}
}

func (o Orm) SaveUser(username string, password string) (*database.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := NewUserOrm(strings.ToLower(username), string(hashedPassword), false)
	err = DB.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return database.NewUser(int(user.ID), user.Username, password, string(hashedPassword), user.Disabled), nil
}

func (o Orm) UpdateUser(username string, password string, disabled bool) (*database.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	var foundUser User
	err = DB.Where("username = ?", username).First(&foundUser).Error
	if err != nil {
		return nil, err
	}

	foundUser.Username = strings.ToLower(username)
	foundUser.Password = string(hashedPassword)
	foundUser.Disabled = disabled

	err = DB.Save(&foundUser).Error
	if err != nil {
		return nil, err
	}

	return database.NewUser(int(foundUser.ID), foundUser.Username, password, string(hashedPassword), foundUser.Disabled), nil
}

func (o Orm) FindUserByName(username string) (*database.User, error) {
	var foundUser User
	err := DB.Where("username = ?", username).First(&foundUser).Error
	if err != nil {
		return nil, err
	}

	return database.NewUser(int(foundUser.ID), foundUser.Username, "", foundUser.Password, foundUser.Disabled), nil
}

func (o Orm) FindUserByID(userId int) (*database.User, error) {
	var foundUser User
	err := DB.First(&foundUser, userId).Error
	if err != nil {
		return nil, err
	}

	return database.NewUser(int(foundUser.ID), foundUser.Username, "", foundUser.Password, foundUser.Disabled), nil
}
