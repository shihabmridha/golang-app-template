package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shihabmridha/golang-app-template/internal/auth"
	"github.com/shihabmridha/golang-app-template/internal/user"
)

func UserHandler(r *Router, usrSvc *user.Service, authSvc *auth.Service) {
	handler, render := r.GetRouterAndRenderer()

	handler.Group(func(r chi.Router) {
		r.Use(authSvc.Verify(render))

		r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
			users, _ := usrSvc.GetAll()

			render.RenderJSON(w, http.StatusOK, users)
		})
	})

	handler.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		body := user.User{}
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		if err := d.Decode(&body); err != nil {
			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		if err := usrSvc.Create(body); err != nil {
			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}
	})

	handler.Get("/user/activate", func(w http.ResponseWriter, r *http.Request) {
		activationCode := r.URL.Query().Get("code")

		if err := usrSvc.Activate(activationCode); err != nil {
			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}
	})
}
