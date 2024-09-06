package main

import (
	"todolist/internal/app"
	"todolist/internal/config"
	"todolist/logger"
)

func main() {
	cfg := config.New()
	log := logger.SetupLogger(cfg.LogLevel)
	application := app.NewApp(cfg, log)
	forever := make(chan bool)
	go application.HTTPApp.Start()
	<-forever
}
