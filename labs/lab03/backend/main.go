package main

import (
	"log"
	"net/http"
	"time"

	"lab03-backend/api"
	"lab03-backend/storage"
)

func main() {
	// Create a new memory storage instance
	memStorage := storage.NewMemoryStorage()

	// Create a new API handler with the storage
	handler := api.NewHandler(memStorage)

	// Setup routes using the handler
	router := handler.SetupRoutes()

	// Configure HTTP server
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Log that the server is starting
	log.Println("Starting server on http://localhost:8080")

	// Start the server and handle errors
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
