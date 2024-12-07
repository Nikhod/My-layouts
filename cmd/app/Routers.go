package main

import (
	"Nikcase/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

// Routers for Diary
func InitRouters(h *handlers.Handlers) *mux.Router {
	router := mux.NewRouter()
	registration := router.PathPrefix("/registration").Subrouter()
	registration.HandleFunc("", h.Registration).Methods(http.MethodPost)

	getToken := router.PathPrefix("/token").Subrouter()
	getToken.HandleFunc("", h.GetToken).Methods(http.MethodGet)
	return router
}
