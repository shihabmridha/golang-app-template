package main

import (
	"context"
	"github.com/shihabmridha/golang-app-template/internal/app"
	"github.com/shihabmridha/golang-app-template/logging"
	"github.com/shihabmridha/golang-app-template/pkg/config"
)

func main() {
	ctx := context.Background()
	cfg := config.New()

	logger := logging.NewLoggerFromEnv().With("version", cfg.App().Version())
	ctx = logging.WithLogger(ctx, logger)

	err := app.Run(ctx, cfg)

	if err != nil {
		logger.Fatal(err)
	}
}
