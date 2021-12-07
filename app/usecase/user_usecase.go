package usecase

import (
	"context"
	"testProject/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (u *UserUsecase) Create(c context.Context, user domain.Users) error {
	err := u.userRepo.Create(c, user)
	if err != nil {
		return err
	}
	return nil
}
