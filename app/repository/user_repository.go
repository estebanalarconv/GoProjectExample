package repository

import (
	"context"
	"testProject/database"
	"testProject/domain"
)

type UserRepository struct {
	dbHandler database.DBHandler
}

func NewUserRepository(dbHandler database.DBHandler) domain.UserRepository {
	return &UserRepository{dbHandler}
}

func (m UserRepository) Create(ctx context.Context, user domain.Users) error {
	err := m.dbHandler.Conn.Create(&user).Error
	return err
}
