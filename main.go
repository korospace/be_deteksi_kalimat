package main

import (
	"be_deteksi_kalimat/database"
	"be_deteksi_kalimat/middleware"
	"be_deteksi_kalimat/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	// Global Middleware
	router.Use(middleware.CorsMiddleware)

	// Routes Registration
	routes.AuthRoutes(router)
	routes.CategoryRoutes(router)
	routes.UserAccessRoutes(router)
	routes.DatasetRoutes(router)
	routes.TrainingRoutes(router)

	log.Println("server running at http://localhost:8080")

	http.ListenAndServe(":8080", router)
}
