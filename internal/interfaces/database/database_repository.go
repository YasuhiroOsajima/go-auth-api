package database

import "github.com/YasuhiroOsajima/go-auth-api/internal/model"

type DatabaseRepository struct {
	orm Orm
}

func NewDatabaseRepository(orm Orm) *DatabaseRepository {
	return &DatabaseRepository{orm}
}

func (repo *DatabaseRepository) SaveUser(user *model.User) (*model.User, error) {
	savedUser, err := repo.orm.SaveUser(user.Username, user.Password)
	if err != nil {
		return nil, err
	}

	return model.NewUser(savedUser.UserId, savedUser.Username, savedUser.Password, savedUser.HashedPassword, savedUser.Disabled), nil
}

func (repo *DatabaseRepository) UpdateUser(user *model.User) (*model.User, error) {
	updatedUser, err := repo.orm.UpdateUser(user.Username, user.Password, user.Disabled)
	if err != nil {
		return nil, err
	}

	return model.NewUser(updatedUser.UserId, updatedUser.Username, updatedUser.Password, updatedUser.HashedPassword, updatedUser.Disabled), nil
}

func (repo *DatabaseRepository) FindUserByName(user *model.User) (*model.User, error) {
	foundUser, err := repo.orm.FindUserByName(user.Username)
	if err != nil {
		return nil, err
	}

	return model.NewUser(foundUser.UserId, foundUser.Username, foundUser.Password, foundUser.HashedPassword, foundUser.Disabled), nil
}

func (repo *DatabaseRepository) FindUserByID(userId int) (*model.User, error) {
	foundUser, err := repo.orm.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	return model.NewUser(foundUser.UserId, foundUser.Username, foundUser.Password, foundUser.HashedPassword, foundUser.Disabled), nil
}
