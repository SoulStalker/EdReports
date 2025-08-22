package main

import (
	"log/slog"
	"os"

	"github.com/SoulStalker/EdReports/auth-service/internal/config"
	"github.com/SoulStalker/EdReports/auth-service/internal/lib/logger/handlers/slogpretty"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	// init config
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("cfg", cfg))
	log.Debug("cyan")
	log.Error("red")
	// init logger

	// init app

	// start grpc server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
