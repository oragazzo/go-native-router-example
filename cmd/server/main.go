package main

import (
	"log"
	"net/http"

	"github.com/oragazzo/go-native-router-example/internal/handlers"
	"github.com/oragazzo/go-native-router-example/internal/storage"
)

func main() {
	// Initialize storage
	memStorage := storage.NewMemoryStorage()

	// Initialize handlers
	userHandler := handlers.NewUserHandler(memStorage)

	// Setup router
	serverMux := http.NewServeMux()

	// Register routes
	serverMux.HandleFunc("GET /users", userHandler.GetUsers)
	serverMux.HandleFunc("POST /users", userHandler.CreateUser)
	serverMux.HandleFunc("GET /users/{id}", userHandler.GetUser)
	serverMux.HandleFunc("PUT /users/{id}", userHandler.UpdateUser)
	serverMux.HandleFunc("DELETE /users/{id}", userHandler.DeleteUser)

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", serverMux); err != nil {
		log.Fatal(err)
	}
}
