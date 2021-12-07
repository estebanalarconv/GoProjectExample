package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"testProject/config"
	domain2 "testProject/domain"
)

type DBHandler struct {
	Conn *gorm.DB
}

var dbHandler DBHandler

func SetupConnection() error {
	var err error
	dsn := "host=" + config.GetPostgresHost() +
		" user =" + config.GetPostgresUsername() +
		" password =" + config.GetPostgresPass() +
		" dbname = " + config.GetPostgresDatabase() +
		" port =" + strconv.Itoa(config.GetPostgresPort()) +
		" sslmode = disable"

	dbHandler.Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	err = dbHandler.Conn.AutoMigrate(&domain2.Users{}, &domain2.Foods{}, &domain2.FoodPerDay{})
	if err != nil {
		return err
	}
	return nil
}

func GetConnection() DBHandler {
	return dbHandler
}
