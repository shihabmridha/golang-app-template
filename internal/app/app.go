package app

import (
	"context"
	"fmt"

	"github.com/shihabmridha/golang-app-template/internal/auth"
	"github.com/shihabmridha/golang-app-template/internal/route"
	"github.com/shihabmridha/golang-app-template/internal/user"
	"github.com/shihabmridha/golang-app-template/pkg/config"
	"github.com/shihabmridha/golang-app-template/pkg/database"
	"github.com/shihabmridha/golang-app-template/pkg/http"
)

func Run(ctx *context.Context, cfg *config.Config) error {
	db, err := database.New(ctx, cfg.Db())
	if err != nil {
		return fmt.Errorf("terminating due to db connection issue. error: %w", err)
	}

	defer db.Close()

	appCfg := cfg.App()

	// Repos
	userRepo := user.NewRepo(ctx, db)

	// Services
	authSvc := auth.NewService(appCfg, userRepo)
	userSvc := user.NewService(appCfg, userRepo)

	// Initialize chi router and register middlewares
	r := route.NewRouter()
	handler, _ := r.GetRouterAndRenderer()

	// REST handler
	route.AuthHandler(r, authSvc)
	route.UserHandler(r, userSvc, authSvc)

	httpServer := http.New(appCfg.Ip(), appCfg.Port())
	httpServer.ServeHttp(*ctx, handler)

	return nil
}
