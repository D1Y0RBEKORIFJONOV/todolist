package suittest

import (
	"context"
	"github.com/google/uuid"
	"log"
	"testing"
	"time"
	"todolist/internal/config"
	entityuser "todolist/internal/entity/user"
	userpostgres "todolist/internal/infastructure/repository/postgresql/user"
	"todolist/internal/postgres"
	"todolist/logger"

	"github.com/stretchr/testify/suite"
)

type PostgresTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  *userpostgres.UserRepository
}

func (s *PostgresTest) SetupTest() {
	cfg := config.New()

	pgPool, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log1 := logger.SetupLogger("local")
	s.Repository = userpostgres.NewSolderRepository(pgPool, log1)

	s.CleanUpFunc = func() {}
}

func (s *PostgresTest) TearDownTest() {
	if s.CleanUpFunc != nil {
		s.CleanUpFunc()
	}
}
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTest))
}

func (s *PostgresTest) TestSaveAndGetUser() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	user := entityuser.User{
		Id:           uuid.NewString(),
		UserName:     "testuser",
		Email:        "testuser@gmail.com",
		PasswordHash: "password_hash",
	}

	err := s.Repository.SaveUser(ctx, &user)
	s.Require().NoError(err)

	savedUser, err := s.Repository.GetUser(ctx, "email", user.Email)
	s.Require().NoError(err)
	s.Require().NotNil(savedUser)

	s.Equal(user.Id, savedUser.Id)
	s.Equal(user.UserName, savedUser.UserName)
	s.Equal(user.Email, savedUser.Email)
	s.Equal(user.PasswordHash, savedUser.PasswordHash)

	err = s.Repository.DeleteUser(ctx, user.Id)
	s.Require().NoError(err)
}
