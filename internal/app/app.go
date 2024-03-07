package app

import (
	"context"
	"fmt"

	"github.com/shihabmridha/golang-app-template/internal/api"
	"github.com/shihabmridha/golang-app-template/internal/user"
	"github.com/shihabmridha/golang-app-template/pkg/config"
	"github.com/shihabmridha/golang-app-template/pkg/database"
	"github.com/shihabmridha/golang-app-template/pkg/http"
)

func Run(ctx context.Context, cfg *config.Config) error {
	db, err := database.New(ctx, cfg.Db())
	if err != nil {
		return fmt.Errorf("app - Run - database.New: %w", err)
	}

	defer db.Close()

	// Repos
	userRepo := user.NewRepo(db)

	// Services
	userSvc := user.NewSvc(*userRepo)

	// Initialize chi router and register middlewares
	r := api.Init()

	// REST handler
	user.Handler(r, userSvc)

	appCfg := cfg.App()

	httpServer := http.New(appCfg.Ip(), appCfg.Port())
	httpServer.ServeHttp(ctx, r)

	return nil
}
