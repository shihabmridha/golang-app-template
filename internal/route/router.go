package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/shihabmridha/golang-app-template/pkg/render"
)

type Router struct {
	mux      *chi.Mux
	renderer *render.Renderer
}

func NewRouter() *Router {
	mux := chi.NewRouter()

	mux.Use(middleware.RedirectSlashes)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Compress(5))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(middleware.AllowContentType("application/json"))

	cors.AllowAll().Handler(mux)

	return &Router{
		mux:      mux,
		renderer: render.NewRenderer(),
	}
}

// GetRouterAndRenderer route handler and renderer
func (r *Router) GetRouterAndRenderer() (*chi.Mux, *render.Renderer) {
	return r.mux, r.renderer
}
