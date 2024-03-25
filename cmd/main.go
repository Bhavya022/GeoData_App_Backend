package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bhavya022/GeoData_App/Backend/internal/auth"
	"github.com/Bhavya022/GeoData_App/Backend/internal/handlers"
	"github.com/Bhavya022/GeoData_App/Backend/internal/repository"
	"github.com/Bhavya022/GeoData_App/Backend/pkg/database"

	"github.com/go-chi/chi"
)

func main() {
	// Initialize the database connection
	err := database.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Initialize repositories
	userRepository := repository.NewUserRepository(database.DBConn)

	// Initialize services
	authService := auth.NewAuthService(userRepository)

	// Initialize router
	r := chi.NewRouter()

	// Middleware
	authMiddleware := auth.NewAuthMiddleware(authService)
	r.Use(authMiddleware.Middleware)

	// Routes
	r.Post("/login", handlers.LoginHandler(authService))

	// Start the server
	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
