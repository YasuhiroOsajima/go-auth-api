package usecase

import "github.com/YasuhiroOsajima/go-auth-api/internal/model"

type UserInteractor struct {
	repository IDatabaseRepository
}

func NewUserInteractor(dbRepository IDatabaseRepository) *UserInteractor {
	return &UserInteractor{dbRepository}
}

func (interactor *UserInteractor) Register(user *model.User) (*model.User, error) {
	savedUser, err := interactor.repository.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

func (interactor *UserInteractor) Enable(user *model.User) (*model.User, error) {
	user.Disabled = false
	savedUser, err := interactor.repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}

func (interactor *UserInteractor) Disable(user *model.User) (*model.User, error) {
	user.Disabled = true
	savedUser, err := interactor.repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return savedUser, nil
}
