package handlers

import (
	"golang-url-shortener/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// URLHandler handles HTTP requests for URLs
type URLHandler struct {
	urlService *services.URLService
}

// NewURLHandler creates a new URLHandler
func NewURLHandler(postgresStore services.DataStore, redisStore services.CacheStore) *URLHandler {
	return &URLHandler{
		urlService: services.NewURLService(postgresStore, redisStore),
	}
}

// ShortenURL shortens a given URL
func (h *URLHandler) ShortenURL(c *gin.Context) {
	var request struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := h.urlService.Shorten(request.URL)
	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

// RedirectURL redirects to the original URL
func (h *URLHandler) RedirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	longURL, exists := h.urlService.Resolve(shortURL)

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "short URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}
