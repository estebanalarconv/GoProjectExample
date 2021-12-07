package repository

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testProject/database"
	"testProject/domain"
	"testing"
)

type SuiteUser struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository domain.UserRepository
	user       *domain.Users
}

func (s *SuiteUser) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{SkipDefaultTransaction: true})
	require.NoError(s.T(), err)
	s.repository = NewUserRepository(database.DBHandler{
		Conn: s.DB,
	})
}

func (s *SuiteUser) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(SuiteUser))
}

func (s *SuiteUser) Test_repository_Create() {

	user := domain.Users{
		Name:     "test",
		Genre:    0,
		Birth:    "01-01-0001",
		Username: "test",
	}
	s.mock.ExpectQuery(
		`INSERT INTO "users" ("name","genre","birth","username") 
			VALUES ($1,$2,$3,$4) RETURNING "id"`).
		WithArgs(user.Name, user.Genre, user.Birth, user.Username).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(1))

	err := s.repository.Create(context.Background(), user)

	require.NoError(s.T(), err)
}
