package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

type config struct {
	postgresHost     string
	postgresPort     int
	postgresDatabase string
	postgresUsername string
	postgresPass     string
}

var conf config

func LoadConfiguration() error {

	c := &config{}

	err := godotenv.Load()
	if err != nil {
		log.
			Error().
			Msg("Environment variables not found.")
	}

	errs := make([]string, 0)

	c.postgresHost = os.Getenv("DB_HOST")
	if c.postgresHost == "" {
		errs = append(errs, "Error variable database.postgres.host from .env")
	}

	c.postgresPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		errs = append(errs, "Error variable database.postgres.port from .env")
	}

	c.postgresDatabase = os.Getenv("DB_NAME")
	if c.postgresDatabase == "" {
		errs = append(errs, "Error variable database.postgres.database from .env")
	}

	c.postgresUsername = os.Getenv("DB_USER")
	if c.postgresUsername == "" {
		errs = append(errs, "Error variable database.postgres.username from .env")
	}

	c.postgresPass = os.Getenv("DB_PASSWORD")
	if c.postgresPass == "" {
		errs = append(errs, "Error variable database.postgres.pass from .env")
	}

	if len(errs) > 0 {
		log.Error().
			Interface("errors", errs).
			Msg("Required enviroments not found")
		return errors.New("errors with arguments")
	}
	conf = *c
	return nil
}

func GetPostgresHost() string {
	return conf.postgresHost
}

func GetPostgresPort() int {
	return conf.postgresPort
}
func GetPostgresDatabase() string {
	return conf.postgresDatabase
}
func GetPostgresUsername() string {
	return conf.postgresUsername
}
func GetPostgresPass() string {
	return conf.postgresPass
}
