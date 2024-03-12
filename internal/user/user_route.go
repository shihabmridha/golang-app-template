package user

import (
	"encoding/json"
	"net/http"

	"github.com/shihabmridha/golang-app-template/internal/api"
)

func Handler(r *api.Router, usrSvc *Service) {
	handler, render := r.GetRouterAndRenderer()

	handler.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		users, _ := usrSvc.GetAll()

		render.RenderJSON(w, http.StatusOK, users)
	})

	handler.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		body := User{}
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
