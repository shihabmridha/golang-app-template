package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shihabmridha/golang-app-template/internal/user"
	"github.com/shihabmridha/golang-app-template/pkg/render"
)

func GetUser(render *render.Renderer, usrSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, _ := usrSvc.GetAll()

		render.RenderJSON(w, http.StatusOK, users)
	}
}

func ActivateUser(render *render.Renderer, usrSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		activationCode := r.URL.Query().Get("code")

		if err := usrSvc.Activate(activationCode); err != nil {
			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}
	}
}

func CreateUser(render *render.Renderer, usrSvc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
