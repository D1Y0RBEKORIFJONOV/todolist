package app

import (
	"log/slog"
	htppapp "todolist/internal/app/htpp"
	"todolist/internal/config"
	"todolist/internal/infastructure/repository/mongo"
	taskpostgres "todolist/internal/infastructure/repository/postgresql/task"
	"todolist/internal/infastructure/repository/postgresql/user"
	redisrepository "todolist/internal/infastructure/repository/redis"
	"todolist/internal/postgres"
	taskservice "todolist/internal/services/task"
	userservices "todolist/internal/services/user"
	taskusecase "todolist/internal/usecase/task"
	userusecase "todolist/internal/usecase/user"
)

type App struct {
	HTTPApp *htppapp.App
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	dbRedis := redisrepository.NewRedis(*cfg)
	db, err := postgres.New(cfg)
	if err != nil {
		panic(err)
	}
	dbPostgres := userpostgres.NewSolderRepository(db, logger)

	dbTaskPostgres := taskpostgres.NewTaskRepository(db, logger)
	dbTaskMongo, err := mongo.NewMongoDB(cfg)
	if err != nil {
		panic(err)
	}

	dbTaskUseCase := taskusecase.NewTaskRepoIml(dbTaskPostgres, dbTaskMongo)

	dbUseCase := userusecase.NewUserUseCaseRepo(dbRedis, dbPostgres)

	service := userservices.NewUserService(*dbUseCase, logger)

	serviceTask := taskservice.NewTask(dbTaskUseCase, logger)

	taskHandelrUseCase := taskusecase.NewTaskUseCaseIml(serviceTask)

	handlerUseCase := userusecase.NewUseCaseIml(service)
	server := htppapp.NewApp(logger, cfg.HostUrl, *handlerUseCase, *taskHandelrUseCase)
	return &App{
		HTTPApp: server,
	}
}
