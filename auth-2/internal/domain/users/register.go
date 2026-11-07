package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RouterRegister(r *mux.Router) {
    repo := NewRepository()
    serv := NewService(repo)
    handler := NewHandler(serv)

    // Create user
    r.HandleFunc("/users", handler.Create).
        Methods(http.MethodPost)

    // Get all users
    r.HandleFunc("/users", handler.GetAll).
        Methods(http.MethodGet)

    // Get single user
    r.HandleFunc("/users/{id}", handler.GetByID).
        Methods(http.MethodGet)

    // Update complete user
    r.HandleFunc("/users/{id}", handler.Update).
        Methods(http.MethodPut)

    // Update partial user
    r.HandleFunc("/users/{id}", handler.Patch).
        Methods(http.MethodPatch)

    // Delete user
    r.HandleFunc("/users/{id}", handler.Delete).
        Methods(http.MethodDelete)
}