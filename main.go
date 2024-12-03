package main

import (
	"golang-url-shortener/config"
	"golang-url-shortener/routes"
	"golang-url-shortener/store"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize stores
	postgresStore, err := store.NewPostgresStore(cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer postgresStore.Close()

	redisStore, err := store.NewRedisStore(cfg.RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer redisStore.Close()

	// Initialize Gin router
	r := gin.Default()

	// Setup routes using the routes package
	routes.SetupRoutes(r, postgresStore, redisStore)

	// Start server
	r.Run(":8080")
}
