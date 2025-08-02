package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gyaan/short-urls/internal/config"
	"github.com/gyaan/short-urls/internal/handler"
	mclient "github.com/gyaan/short-urls/internal/mongo_client"
	"github.com/gyaan/short-urls/internal/repositories"
	"github.com/gyaan/short-urls/internal/router"
)

// main initializes and starts the short URL service
func main() {
	// Initialize configuration
	conf := config.New()

	// Initialize MongoDB client
	mClient, err := mclient.New(conf.MongoDbConnectionUrl, conf.MongoContextTimeout)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize repositories
	counterRepository := repositories.NewCounterRepository(mClient, conf)
	shortUrlRepository := repositories.NewShortUrlRepository(mClient, counterRepository, conf)
	userRepository := repositories.NewUserRepository(mClient, conf)

	// Initialize handlers
	authenticationHandler := handler.NewAuthenticationHandler(userRepository, conf)
	shortUrlHandler := handler.NewShortUrlHandler(shortUrlRepository)
	userHandler := handler.NewUserHandler(userRepository)

	// Create router and register routes
	r := chi.NewRouter()
	router.RegisterRoutes(shortUrlHandler, userHandler, authenticationHandler, r)

	// Start HTTP server
	log.Printf("Starting short URL application on port: %s", conf.ApplicationPort)
	if err := http.ListenAndServe(conf.ApplicationPort, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
