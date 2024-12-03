package store

import (
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

// RedisStore implements the CacheStore interface for Redis
type RedisStore struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisStore initializes a RedisStore
func NewRedisStore(addr string) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{Addr: addr})
	ctx := context.Background()

	// Test the connection
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisStore{client: client, ctx: ctx}, nil
}

// Set caches a URL mapping
func (s *RedisStore) Set(shortURL, longURL string) error {
	return s.client.Set(s.ctx, shortURL, longURL, 0).Err()
}

// Get retrieves a long URL by its short URL
func (s *RedisStore) Get(shortURL string) (string, bool) {
	val, err := s.client.Get(s.ctx, shortURL).Result()
	return val, err == nil
}

// Close closes the Redis connection
func (s *RedisStore) Close() error {
	return s.client.Close()
}
