package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shihabmridha/golang-app-template/internal/api"
)

func Handler(r *api.Router, usrSvc *Service) {
	handler := r.Handler()
	render := r.Renderer()

	handler.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		users, _ := usrSvc.Get()
		fmt.Println(users)

		render.RenderJSON(w, http.StatusOK, users)
	})

	handler.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		body := &User{}
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		if err := d.Decode(body); err != nil {
			// fmt.Println(err)
			// w.WriteHeader(http.StatusBadRequest)

			render.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		fmt.Println(body)

		usrSvc.Create()
	})
}
