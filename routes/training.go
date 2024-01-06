package routes

import (
	"be_deteksi_kalimat/controllers"
	"be_deteksi_kalimat/middleware"

	"github.com/gorilla/mux"
)

func TrainingRoutes(r *mux.Router) {
	router := r.PathPrefix("/training").Subrouter()

	router.Use(middleware.TokenMiddleware)
	router.HandleFunc("/single", controllers.SingleTraining).Methods("POST", "OPTIONS")
	router.HandleFunc("/bulk", controllers.BulkTraining).Methods("POST", "OPTIONS")
}
