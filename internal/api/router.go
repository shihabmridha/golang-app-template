package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func Init() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/health"))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	cors.AllowAll().Handler(r)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	return r
}
