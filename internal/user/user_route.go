package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Handler(r *chi.Mux, usrSvc *Service) {
	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		users, _ := usrSvc.Get()
		res, _ := json.Marshal(users)

		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		usrSvc.Create()
		w.WriteHeader(http.StatusAccepted)
	})
}
