package routes

import (
	"golang-url-shortener/handlers"
	"golang-url-shortener/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the routes for the application
func SetupRoutes(r *gin.Engine, postgresStore services.DataStore, redisStore services.CacheStore) {
	// Initialize the URL handler
	urlHandler := handlers.NewURLHandler(postgresStore, redisStore)

	// Define routes
	r.POST("/shorten", urlHandler.ShortenURL)
	r.GET("/:shortURL", urlHandler.RedirectURL)
}
