package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoute(r *mux.Router) {
	r.HandleFunc("/login",HealthcheckHandler)
	r.HandleFunc("/healthcheck", HealthcheckHandler)
	r.HandleFunc("/logout", LogoutHandler)
	http.Handle("/", r)
}