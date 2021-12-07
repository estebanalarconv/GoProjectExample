package domain

import "context"

type Users struct {
	Id       int
	Name     string
	Genre    int
	Birth    string
	Username string
}

type UserRepository interface {
	Create(ctx context.Context, user Users) error
}

type UserUsecase interface {
	Create(ctx context.Context, user Users) error
}
