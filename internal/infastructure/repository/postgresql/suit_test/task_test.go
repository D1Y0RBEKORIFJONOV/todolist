package suittest

import (
	"context"
	"github.com/google/uuid"
	"log"
	"testing"
	"time"
	"todolist/internal/config"
	tasksentity "todolist/internal/entity/tasks"
	entityuser "todolist/internal/entity/user"
	taskpostgres "todolist/internal/infastructure/repository/postgresql/task"
	userpostgres "todolist/internal/infastructure/repository/postgresql/user"
	"todolist/internal/postgres"
	"todolist/logger"

	"github.com/stretchr/testify/suite"
)

type PostgresTestTask struct {
	suite.Suite
	CleanUpFunc    func()
	Repository     *taskpostgres.TaskRepository
	RepositoryUser *userpostgres.UserRepository
}

func (s *PostgresTestTask) SetupTest() {
	cfg := config.New()

	pgPool, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log1 := logger.SetupLogger("local")
	s.Repository = taskpostgres.NewTaskRepository(pgPool, log1)
	s.RepositoryUser = userpostgres.NewSolderRepository(pgPool, log1)
	s.CleanUpFunc = func() {}
}

func (s *PostgresTestTask) TaskDownTest() {
	if s.CleanUpFunc != nil {
		s.CleanUpFunc()
	}
}
func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestTask))
}

func (s *PostgresTestTask) TestTask() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	user := entityuser.User{
		Id:           uuid.NewString(),
		UserName:     "testuser",
		Email:        "testuser@gmail.com",
		PasswordHash: "password_hash",
	}

	err := s.RepositoryUser.SaveUser(ctx, &user)
	s.Require().NoError(err)

	taskSave := tasksentity.TaskPostgres{
		UserId:   user.Id,
		Id:       uuid.NewString(),
		Title:    "test1",
		CreateAt: "tes1_time",
	}

	err = s.Repository.CreateTask(ctx, &taskSave)
	s.Require().NoError(err)

	task, err := s.Repository.GetTask(ctx, taskSave.UserId, "id", taskSave.Id)
	s.Require().NoError(err)
	s.Require().Equal(taskSave.Title, task.Title)
	s.Require().Equal(taskSave.CreateAt, taskSave.CreateAt)
	s.Require().Equal(taskSave.Id, task.Id)
	s.Require().Equal(taskSave.UserId, task.UserId)

	tasks, err := s.Repository.GetAllTasks(ctx, &tasksentity.GetAllTaskReq{
		UserId: taskSave.UserId,
	})
	s.Require().NoError(err)
	s.Require().Len(tasks, 1)
	s.Require().Equal(taskSave.Title, tasks[0].Title)
	s.Require().Equal(taskSave.CreateAt, tasks[0].CreateAt)
	s.Require().Equal(taskSave.Id, tasks[0].Id)
	s.Require().Equal(taskSave.UserId, tasks[0].UserId)

	err = s.Repository.UpdateTask(ctx, task.UserId, "QWERT")
	s.Require().NoError(err)

	err = s.Repository.DeleteTask(ctx, task.UserId, task.Id)
	s.Require().NoError(err)

	err = s.RepositoryUser.DeleteUser(ctx, taskSave.UserId)
}
