package htppapp

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"todolist/internal/http/router"
	taskusecase "todolist/internal/usecase/task"
	userusecase "todolist/internal/usecase/user"
)

type App struct {
	Logger *slog.Logger
	Port   string
	Server *gin.Engine
}

func NewApp(logger *slog.Logger, port string, handlerService userusecase.UserUseCaseImpl, taskUseCase taskusecase.TaskUseCaseIml) *App {
	sever := router.RegisterRouter(handlerService, taskUseCase)
	return &App{
		Port:   port,
		Server: sever,
		Logger: logger,
	}
}

func (app *App) Start() {
	const op = "app.Start"
	log := app.Logger.With(
		slog.String(op, "Starting server"),
		slog.String("port", app.Port))
	log.Info("Starting server")
	err := app.Server.SetTrustedProxies(nil)
	if err != nil {
		log.Error("Error setting trusted proxies", "error", err)
		return
	}
	err = app.Server.Run(app.Port)
	if err != nil {
		log.Error("Error starting server", "error", err)
		return
	}
}
