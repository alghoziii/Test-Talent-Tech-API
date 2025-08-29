package main

import (
    "log"
    "net/http"
    "github.com/joho/godotenv"
    "e-ticketing/database"
    "e-ticketing/controllers"
    "e-ticketing/middleware"
    "github.com/gorilla/mux"
)

func main() {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    // Initialize database
    database.InitDB()
    defer database.CloseDB()

    router := mux.NewRouter()

    // Public routes
    router.HandleFunc("/api/login", controllers.Login).Methods("POST")

    // Protected routes
    router.HandleFunc("/api/terminals", middleware.JWTMiddleware(controllers.CreateTerminal)).Methods("POST")

    log.Println("Server starting on :8080")
    http.ListenAndServe(":8080", router)
}