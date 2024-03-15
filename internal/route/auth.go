package route

import (
	"encoding/json"
	"net/http"

	"github.com/shihabmridha/golang-app-template/internal/auth"
)

func AuthHandler(r *Router, authSvc *auth.Service) {
	mux, render := r.GetRouterAndRenderer()

	mux.Post("/login", func(w http.ResponseWriter, r *http.Request) {
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
	})

}
