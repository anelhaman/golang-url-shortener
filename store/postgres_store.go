package store

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// URL represents the URL mapping
type URL struct {
	ID        uint      `gorm:"primaryKey"`
	ShortURL  string    `gorm:"uniqueIndex"`
	LongURL   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // Automatically set to current timestamp
}

// PostgresStore implements the DataStore interface for PostgreSQL
type PostgresStore struct {
	db *gorm.DB
}

// NewPostgresStore initializes a PostgresStore
func NewPostgresStore(dsn string) (*PostgresStore, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&URL{}); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, nil
}

// SaveURL saves a URL mapping
func (s *PostgresStore) SaveURL(shortURL, longURL string) error {
	return s.db.Create(&URL{ShortURL: shortURL, LongURL: longURL}).Error
}

// GetLongURL retrieves a long URL by its short URL
func (s *PostgresStore) GetLongURL(shortURL string) (string, bool) {
	var url URL
	result := s.db.Where("short_url = ?", shortURL).First(&url)
	return url.LongURL, result.RowsAffected > 0
}

// Close closes the database connection
func (s *PostgresStore) Close() error {
	db, _ := s.db.DB()
	return db.Close()
}
