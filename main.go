package main

import (
	"log"
	"net/http"
	"e-ticketing/database"
	"e-ticketing/controllers"
	"e-ticketing/middleware"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variables system")
	}

	// Initialize database
	database.InitDB()
	defer database.CloseDB()

	// Create router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")

	// Protected routes
	router.HandleFunc("/api/terminals", middleware.JWTMiddleware(controllers.CreateTerminal)).Methods("POST")

	// Start server
	log.Println("Server berjalan di port :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Server error:", err)
	}
}