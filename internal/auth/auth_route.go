package auth

import (
	"encoding/json"
	"net/http"

	"github.com/shihabmridha/golang-app-template/internal/api"
)

func Handler(r *api.Router, authSvc *Service) {
	handler, render := r.GetRouterAndRenderer()

	// handler.Get("/user", func(w http.ResponseWriter, r *http.Request) {
	// 	users, _ := usrSvc.GetAll()

	// 	render.RenderJSON(w, http.StatusOK, users)
	// })

	handler.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		body := UserLogin{}
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		if err := d.Decode(&body); err != nil {
			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		token, err := authSvc.GetToken(body)

		if err != nil {
			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		render.RenderJSON(w, http.StatusOK, token)
	})

	// handler.Get("/user/activate", func(w http.ResponseWriter, r *http.Request) {
	// 	activationCode := r.URL.Query().Get("code")

	// 	if err := usrSvc.Activate(activationCode); err != nil {
	// 		render.RenderJSON(w, http.StatusBadRequest, err)
	// 		return
	// 	}
	// })
}
