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

	routerSuperadmin := r.PathPrefix("/category").Subrouter()

	routerSuperadmin.Use(middleware.TokenMiddleware)
	routerSuperadmin.Use(middleware.SuperadminMiddleware)
	routerSuperadmin.HandleFunc("/create", controllers.CreateCategory).Methods("POST", "OPTIONS")
	routerSuperadmin.HandleFunc("/update", controllers.UpdateCategory).Methods("PUT", "OPTIONS")
	routerSuperadmin.HandleFunc("/delete", controllers.DeleteCategory).Methods("DELETE", "OPTIONS")
}
