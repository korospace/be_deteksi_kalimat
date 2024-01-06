package routes

import (
	"be_deteksi_kalimat/controllers"
	"be_deteksi_kalimat/middleware"

	"github.com/gorilla/mux"
)

func DatasetRoutes(r *mux.Router) {
	router := r.PathPrefix("/dataset").Subrouter()

	router.Use(middleware.TokenMiddleware)
	router.HandleFunc("/list", controllers.ListingDataset).Methods("GET", "OPTIONS")

	routerPakar := r.PathPrefix("/dataset").Subrouter()

	routerPakar.Use(middleware.TokenMiddleware)
	routerPakar.Use(middleware.PakarMiddleware)
	routerPakar.HandleFunc("/verify", controllers.VerifyDataset).Methods("PUT", "OPTIONS")

	routerSuperadmin := r.PathPrefix("/dataset").Subrouter()

	routerSuperadmin.Use(middleware.TokenMiddleware)
	routerSuperadmin.Use(middleware.SuperadminMiddleware)
	routerSuperadmin.HandleFunc("/create", controllers.CreateDataset).Methods("POST", "OPTIONS")
	routerSuperadmin.HandleFunc("/import", controllers.ImportDataset).Methods("POST", "OPTIONS")
	routerSuperadmin.HandleFunc("/update", controllers.UpdateDataset).Methods("PUT", "OPTIONS")
	routerSuperadmin.HandleFunc("/delete", controllers.DeleteDataset).Methods("DELETE", "OPTIONS")
}
