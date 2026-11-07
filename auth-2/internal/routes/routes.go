package routes

import (
	"auth-two/internal/auth"
	"auth-two/internal/domain/users"

	"github.com/gorilla/mux"
)

func RegistersRoutes(r *mux.Router) {
	auth.RegisterRoute(r)
	users.RouterRegister(r)
}