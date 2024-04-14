package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shihabmridha/golang-app-template/internal/auth"
	"github.com/shihabmridha/golang-app-template/internal/route"
	"github.com/shihabmridha/golang-app-template/internal/user"
	"github.com/shihabmridha/golang-app-template/pkg/config"
	"github.com/shihabmridha/golang-app-template/pkg/database"
)

func New(ctx context.Context, appCfg config.App, db *database.Sql) *http.Server {
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
	route.UserHandler(r, authSvc, userSvc)

	addr := fmt.Sprintf(appCfg.Ip() + ":" + appCfg.Port())
	server := &http.Server{Addr: addr, Handler: handler}

	return server
}
