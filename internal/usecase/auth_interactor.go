package usecase

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/YasuhiroOsajima/go-auth-api/internal/model"
)

type AuthInteractor struct {
	repository IDatabaseRepository
	token      ITokenLogic
}

func NewAuthInteractor(dbRepository IDatabaseRepository, token ITokenLogic) *AuthInteractor {
	return &AuthInteractor{dbRepository, token}
}

func (i *AuthInteractor) GenerateToken(user *model.User) (string, error) {
	foundUser, err := i.repository.FindUserByName(user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.HashedPassword), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token, err := i.token.GenerateToken(foundUser.UserId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (i *AuthInteractor) GetUserInfoByToken(bearToken string) (*model.User, error) {
	userId, err := i.token.ExtractUserId(bearToken)
	if err != nil {
		return nil, err
	}

	foundUser, err := i.repository.FindUserByID(userId)
	if err != nil {
		return nil, err
	}

	authenticatedUser := model.NewUser(foundUser.UserId, foundUser.Username, foundUser.Password, foundUser.HashedPassword, foundUser.Disabled)
	return authenticatedUser, nil
}

func (i *AuthInteractor) GetUserInfoByName(user *model.User) (*model.User, error) {
	foundUser, err := i.repository.FindUserByName(user)
	if err != nil {
		return nil, err
	}

	authenticatedUser := model.NewUser(foundUser.UserId, foundUser.Username, foundUser.Password, foundUser.HashedPassword, foundUser.Disabled)
	return authenticatedUser, nil
}
