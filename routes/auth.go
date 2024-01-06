package routes

import (
	"be_deteksi_kalimat/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	router := r.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST", "OPTIONS")
}
