package routes

import (
	"be_deteksi_kalimat/controllers"
	"be_deteksi_kalimat/middleware"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")

	router_me := r.PathPrefix("/auth").Subrouter()

	router_me.Use(middleware.TokenMiddleware)
	router_me.HandleFunc("/me", controllers.Me).Methods("GET", "OPTIONS")
}
