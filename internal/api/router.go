package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/shihabmridha/golang-app-template/pkg/render"
)

type Router struct {
	handler  *chi.Mux
	renderer *render.Renderer
}

func NewRouter() *Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Compress(5))
	r.Use(middleware.Heartbeat("/health"))
	r.Use(middleware.AllowContentType("application/json"))
	cors.AllowAll().Handler(r)

	return &Router{
		handler:  r,
		renderer: render.NewRenderer(),
	}
}

func (r *Router) Handler() *chi.Mux {
	return r.handler
}

func (r *Router) Renderer() *render.Renderer {
	return r.renderer
}
