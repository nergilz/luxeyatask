package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/nergilz/luxeyatask/internal/repository"
	"github.com/nergilz/luxeyatask/internal/server"
	"github.com/nergilz/luxeyatask/internal/service"
)

func main() {
	logger := setupLogger("local")
	logger.Info("init logger")

	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	url := "postgres://admin:mysecret@localhost:5432/tourdb"

	storage, err := repository.New(ctx, url, logger)
	if err != nil {
		logger.ErrorContext(ctx, "error", slog.String("new store", err.Error()))
	}

	usecase := service.New(logger, storage)

	handler := server.NewHandler(logger, usecase)

	server.Run(ctx, handler.ServeMux)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "info":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
