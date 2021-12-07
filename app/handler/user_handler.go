package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/domain"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(app *gin.Engine, us domain.UserUsecase) {
	handler := &UserHandler{
		userUsecase: us,
	}
	app.POST("/users", handler.Create)
}

func (u *UserHandler) Create(c *gin.Context) {
	var user domain.Users
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.userUsecase.Create(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
