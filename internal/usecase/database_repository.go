package usecase

import "github.com/YasuhiroOsajima/go-auth-api/internal/model"

type IDatabaseRepository interface {
	SaveUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	FindUserByName(user *model.User) (*model.User, error)
	FindUserByID(userId int) (*model.User, error)
}
