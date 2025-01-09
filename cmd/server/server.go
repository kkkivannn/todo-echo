package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo_echo/internal/app"
	"todo_echo/internal/config"
)

func main() {
	cfg := config.MustLoad()

	a := app.New(cfg)

	slog.Info("server starting...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go a.Srv.MustRun()

	slog.Info("server started")

	<-done

	slog.Info("stopping server")

	a.Srv.Stop(ctx)

	slog.Info("server stopped")
}
