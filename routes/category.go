package routes

import (
	"be_deteksi_kalimat/controllers"
	"be_deteksi_kalimat/middleware"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	router := r.PathPrefix("/category").Subrouter()

	router.Use(middleware.TokenMiddleware)
	router.HandleFunc("/list", controllers.ListingCategory).Methods("GET", "OPTIONS")

	router.Use(middleware.SuperadminMiddleware)
	router.HandleFunc("/create", controllers.CreateCategory).Methods("POST", "OPTIONS")
	router.HandleFunc("/update", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/delete", controllers.DeleteCategory).Methods("DELETE")
}
