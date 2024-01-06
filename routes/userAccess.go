package routes

import (
	"be_deteksi_kalimat/controllers"
	"be_deteksi_kalimat/middleware"

	"github.com/gorilla/mux"
)

func UserAccessRoutes(r *mux.Router) {
	router := r.PathPrefix("/user_access").Subrouter()

	router.Use(middleware.TokenMiddleware)
	router.Use(middleware.SuperadminMiddleware)
	router.HandleFunc("/list", controllers.ListingAccess).Methods("GET", "OPTIONS")
	router.HandleFunc("/create", controllers.CreateAccess).Methods("POST", "OPTIONS")
	router.HandleFunc("/update", controllers.UpdateAccess).Methods("PUT", "OPTIONS")
	router.HandleFunc("/delete", controllers.DeleteAccess).Methods("DELETE", "OPTIONS")
}
