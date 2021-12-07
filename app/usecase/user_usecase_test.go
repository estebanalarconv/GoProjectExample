package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testProject/domain"
	"testProject/domain/mocks"
	"testing"
)

func TestUserUsecase_Create(t *testing.T) {
	mockUser := domain.Users{
		Name:     "test",
		Genre:    0,
		Birth:    "01-01-0001",
		Username: "test",
	}
	mockUserRepo := new(mocks.UserRepository)

	t.Run("success", func(t *testing.T) {
		tempMockUser := mockUser
		tempMockUser.Id = 0
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Users")).Return(nil).Once()

		u := NewUserUsecase(mockUserRepo)

		err := u.Create(context.TODO(), tempMockUser)

		assert.NoError(t, err)
		assert.Equal(t, mockUser.Name, tempMockUser.Name)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		errorWant := errors.New("test error")
		tempMockUser := mockUser
		tempMockUser.Id = 0
		mockUserRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.Users")).Return(errorWant).Once()

		u := NewUserUsecase(mockUserRepo)

		errGot := u.Create(context.TODO(), tempMockUser)

		assert.EqualError(t, errGot, errorWant.Error())
		mockUserRepo.AssertExpectations(t)
	})
}
