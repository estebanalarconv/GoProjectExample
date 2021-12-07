package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"testProject/domain"
)

type UserUsecase struct {
	mock.Mock
}

func (_m *UserUsecase) Create(_a0 context.Context, _a1 domain.Users) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Users) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
