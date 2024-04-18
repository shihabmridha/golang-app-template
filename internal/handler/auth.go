package handler

import (
	"encoding/json"
	"net/http"

	"github.com/shihabmridha/golang-app-template/internal/auth"
	"github.com/shihabmridha/golang-app-template/pkg/render"
)

func Login(render *render.Renderer, authSvc auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := auth.UserLogin{}
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		if err := d.Decode(&body); err != nil {
			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		token, err := authSvc.Login(body)

		if err != nil {
			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		render.RenderJSON(w, http.StatusOK, token)
	}
}
