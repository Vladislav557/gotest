package app

import (
	"context"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gotest/internal/resources"
	"gotest/internal/resources/postgres"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	// ENV
	if appEnv := os.Getenv("APP_ENV"); appEnv == "" {
		if err := godotenv.Load(); err != nil {
			zap.L().Fatal("error loading .env file")
		}
	}

	// DB initialization
	defer postgres.Close()
	postgres.Init(os.Getenv("DATABASE_URL"))
	postgres.AutoMigrate()

	// Logger initialization
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	// Router initialization
	router := resources.Router{}
	router.Init()

	// Server running
	srv := new(resources.Server)
	go func() {
		err := srv.Run("8000", router.Router)
		if err != nil {
			return
		}
	}()

	zap.S().Info("app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	zap.S().Info("app stopped ", sig)

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}
