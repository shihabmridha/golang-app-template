package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/shihabmridha/golang-app-template/internal/auth"
	"github.com/shihabmridha/golang-app-template/internal/user"
	"github.com/shihabmridha/golang-app-template/pkg/render"
)

func InitRoutes(authSvc auth.Service, userSvc user.Service) *chi.Mux {
	mux := chi.NewRouter()
	render := render.NewRenderer()

	mux.Use(middleware.RedirectSlashes)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Compress(5))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.AllowContentType("application/json", "multipart/form-data"))

	cors.AllowAll().Handler(mux)

	// Auth
	mux.Post("/login", Login(render, authSvc))

	// Usesr
	mux.Post("/user", CreateUser(render, userSvc))
	mux.Get("/user/activate", ActivateUser(render, userSvc))

	// Protected routes
	mux.Group(func(r chi.Router) {
		r.Use(authSvc.Verify(render))

		r.Get("/user", GetUser(render, userSvc))
	})

	mux.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("you are lost"))
	})

	return mux
}
