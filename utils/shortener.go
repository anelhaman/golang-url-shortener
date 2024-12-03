package utils

import (
	"math/rand"
	"time"
)

// GenerateShortURL generates a random short URL
func GenerateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)

	// Create a new random number generator with a seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate the random string
	for i := range b {
		b[i] = charset[rng.Intn(len(charset))]
	}
	return string(b)
}
