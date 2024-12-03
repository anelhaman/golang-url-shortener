package services

import (
	"golang-url-shortener/utils"
)

// DataStore defines the interface for a persistent store
type DataStore interface {
	SaveURL(shortURL, longURL string) error
	GetLongURL(shortURL string) (string, bool)
}

// CacheStore defines the interface for a cache store
type CacheStore interface {
	Set(shortURL, longURL string) error
	Get(shortURL string) (string, bool)
}

// URLService provides URL shortening and resolving
type URLService struct {
	dataStore  DataStore
	cacheStore CacheStore
}

// NewURLService creates a new URLService
func NewURLService(dataStore DataStore, cacheStore CacheStore) *URLService {
	return &URLService{
		dataStore:  dataStore,
		cacheStore: cacheStore,
	}
}

// Shorten creates a short URL and saves it
func (s *URLService) Shorten(longURL string) string {
	shortURL := utils.GenerateShortURL()
	s.dataStore.SaveURL(shortURL, longURL)
	s.cacheStore.Set(shortURL, longURL)
	return shortURL
}

// Resolve retrieves the long URL for a short URL
func (s *URLService) Resolve(shortURL string) (string, bool) {
	if longURL, found := s.cacheStore.Get(shortURL); found {
		return longURL, true
	}
	return s.dataStore.GetLongURL(shortURL)
}
