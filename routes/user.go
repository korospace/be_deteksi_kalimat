package routes

import (
	"be_deteksi_kalimat/controllers"
	"be_deteksi_kalimat/middleware"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/user").Subrouter()

	router.Use(middleware.TokenMiddleware)
	router.HandleFunc("/me", controllers.Me).Methods("GET", "OPTIONS")

	routerSuperadmin := r.PathPrefix("/user").Subrouter()

	routerSuperadmin.Use(middleware.TokenMiddleware)
	routerSuperadmin.Use(middleware.SuperadminMiddleware)
	routerSuperadmin.HandleFunc("/list", controllers.ListingUser).Methods("GET", "OPTIONS")
	routerSuperadmin.HandleFunc("/create", controllers.CreateUser).Methods("POST", "OPTIONS")
	routerSuperadmin.HandleFunc("/update", controllers.UpdateUser).Methods("PUT", "OPTIONS")
	routerSuperadmin.HandleFunc("/delete", controllers.DeleteUser).Methods("DELETE", "OPTIONS")
}
