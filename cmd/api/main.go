package main

import (
	"context"
	netHttp "net/http"
	"os"
	"os/signal"

	"github.com/shihabmridha/golang-app-template/pkg/config"
	"github.com/shihabmridha/golang-app-template/pkg/database"
	"github.com/shihabmridha/golang-app-template/pkg/http"
	"github.com/shihabmridha/golang-app-template/pkg/logging"
)

func main() {
	ctx, done := context.WithCancel(context.Background())
	cfg := config.New()

	logger := logging.NewLoggerFromEnv().With("version", cfg.App().Version())
	ctx = logging.WithLogger(ctx, logger)

	db, err := database.New(ctx, cfg.Db())
	if err != nil {
		logger.Errorf("terminating due to db connection issue. error: %w", err)
		done()
	}

	server := http.New(ctx, *cfg.App(), db)

	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)

	go func() {
		<-exitSignal
		logger.Infoln("received OS exit signal")

		if err := server.Shutdown(ctx); err != nil {
			logger.Errorf("HTTP server shutdown error: %s", err)
		}
		logger.Infoln("Http server shutdown complete")

		if err := db.Close(); err != nil {
			logger.Errorf("Failed to close database: %s", err)
		}
		logger.Infoln("Database connection closed")

		done()
	}()

	logger.Infof("listening on :%s", cfg.App().Port())
	if err := server.ListenAndServe(); err != nil && err != netHttp.ErrServerClosed {
		logger.Errorf("server - server.ServeHttp: %w", err)
	}

	<-ctx.Done()
}
