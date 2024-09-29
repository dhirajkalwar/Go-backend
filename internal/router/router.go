package router

import (
	"Backend-Go/internal/user"

	"github.com/gorilla/mux"
)

func NewRouter(userHandler user.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.Register).Methods("POST")
	r.HandleFunc("/login", userHandler.Login).Methods("POST")

	return r
}
