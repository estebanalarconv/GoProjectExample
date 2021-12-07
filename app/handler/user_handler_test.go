package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testProject/domain/mocks"
	"testing"
)

func TestUserHandler_Create(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Created", func(t *testing.T) {
		mockUCase := new(mocks.UserUsecase)

		mockUCase.On("Create", mock.Anything, mock.AnythingOfType("domain.Users")).Return(nil)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewUserHandler(router, mockUCase)
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"Name":     "test",
			"Genre":    0,
			"Birth":    "01-01-0001",
			"Username": "test",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUCase.AssertExpectations(t)
	})

	t.Run("Bad request", func(t *testing.T) {
		mockUCase := new(mocks.UserUsecase)

		mockUCase.On("Create", mock.Anything, mock.AnythingOfType("domain.Users")).Return(nil)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewUserHandler(router, mockUCase)
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"Name":     "test",
			"Genre":    "True",
			"Birth":    "01-01-0001",
			"Username": "test",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockUCase.AssertNotCalled(t, "Create")
	})

	t.Run("Error usecase", func(t *testing.T) {
		mockUCase := new(mocks.UserUsecase)
		errorWant := errors.New("error usecase")
		mockUCase.On("Create", mock.Anything, mock.AnythingOfType("domain.Users")).Return(errorWant)

		// don't need a middleware as we don't yet have authorized user
		router := gin.Default()

		NewUserHandler(router, mockUCase)
		// a response recorder for getting written http response
		rr := httptest.NewRecorder()

		// create a request body with empty email and password
		reqBody, err := json.Marshal(gin.H{
			"Name":     "test",
			"Genre":    1,
			"Birth":    "01-01-0001",
			"Username": "test",
		})
		assert.NoError(t, err)

		// use bytes.NewBuffer to create a reader
		request, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
		assert.NoError(t, err)

		request.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(rr, request)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockUCase.AssertExpectations(t)
	})
}
