package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todolist/internal/app/docs"
	"todolist/internal/http/handler"
	"todolist/internal/http/middleware"
	taskusecase "todolist/internal/usecase/task"
	userusecase "todolist/internal/usecase/user"
)

// @title Artisan Connect
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host https://diyorbek.touristan-bs.uz:9000
// @BasePath        /
// @schemes         https
// @securityDefinitions.apiKey ApiKeyAuth
// @in              header
// @name            Authorization

func RegisterRouter(user userusecase.UserUseCaseImpl, tasks taskusecase.TaskUseCaseIml) *gin.Engine {
	userHandler := handler.NewUserServer(user)
	taskHandler := handler.NewTask(tasks)

	router := gin.Default()

	router.Use(middleware.Middleware)
	router.Use(middleware.TimingMiddleware)

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", userHandler.Login)
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/verify", userHandler.VerifyUser)
	}

	taskGroup := router.Group("/task")
	{
		taskGroup.POST("create", taskHandler.CreateTask)
		taskGroup.PATCH("update/:task_id/", taskHandler.UpdateTask)
		taskGroup.DELETE("delete/:task_id/", taskHandler.DeleteTask)
		taskGroup.GET("/:field/:value/", taskHandler.GetTask)
		taskGroup.GET("tasks/", taskHandler.GetTasks)
	}
	return router
}
