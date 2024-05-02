package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/shihabmridha/golang-app-template/internal/auth"
	"github.com/shihabmridha/golang-app-template/internal/handler"
	"github.com/shihabmridha/golang-app-template/internal/user"
	"github.com/shihabmridha/golang-app-template/pkg/config"
	"github.com/shihabmridha/golang-app-template/pkg/database"
	"github.com/shihabmridha/golang-app-template/pkg/logging"
)

func main() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, os.Interrupt)

	ctx := context.Background()
	cfg := config.New()

	logger := logging.NewLoggerFromEnv().With("version", cfg.Version)
	ctx = logging.WithLogger(ctx, logger)

	db, err := database.New(ctx, cfg)
	if err != nil {
		logger.Fatalln("terminating due to db connection issue. error: %w", err)
	}

	// Repos
	userRepo := user.NewRepo(ctx, db)

	// Services
	authSvc := auth.NewService(cfg, userRepo)
	userSvc := user.NewService(cfg, userRepo)

	mux := handler.InitRoutes(authSvc, userSvc)

	addr := fmt.Sprintf(cfg.Ip + ":" + cfg.Port)
	server := &http.Server{Addr: addr, Handler: mux}

	go func() {
		logger.Infof("listening on :%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalln("server - server.ServeHttp: %w", err)
		}
	}()

	// Wait for interrupt signal
	<-exitSignal

	timeoutCtx, done := context.WithTimeout(ctx, 30*time.Second)
	defer done()

	if err := server.Shutdown(timeoutCtx); err != nil {
		logger.Errorf("HTTP server shutdown error: %s", err)
	}

	logger.Infoln("Http server shutdown complete")

	if err := db.Close(); err != nil {
		logger.Errorf("Failed to close database: %s", err)
	}
	logger.Infoln("Database connection closed")
}
