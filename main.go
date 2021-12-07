package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"testProject/app/handler"
	"testProject/app/repository"
	"testProject/app/usecase"
	"testProject/config"
	"testProject/database"
)

func main() {
	err := config.LoadConfiguration()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.SetupConnection()
	if err != nil {
		log.Fatal("Error opening with Database")
	}

	dbHandler := database.GetConnection()

	//gb := gearbox.New()
	app := gin.Default()

	userRepository := repository.NewUserRepository(dbHandler)
	userUsecase := usecase.NewUserUsecase(userRepository)
	handler.NewUserHandler(app, userUsecase)

	//gb.Group("/testApi", []*gearbox.Route{userRoutes})
	app.Run(":3000")
}
