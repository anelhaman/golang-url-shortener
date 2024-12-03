# Golang URL Shortener

A lightweight URL shortener service built with Go, Gin, PostgreSQL, and Redis.

## Features
- Shorten long URLs
- Redirect short URLs
- Persistent storage with PostgreSQL
- Caching with Redis

## Prerequisites
- Docker & Docker Compose
- Go 1.23+ (for local development)

## Quick Start
1. Clone the repository:
```
   git clone https://github.com/anelhaman/golang-url-shortener.git
   cd golang-url-shortener
```

Start services with Docker Compose:

```
docker-compose up --build
```
Access the service at http://localhost:8080.


API Endpoints

- `POST` /shorten: Shorten a long URL.
- `GET` /:shortURL: Redirect to the original URL.

## Build Multi-Platform Docker Image

Enable Docker Buildx:

```
docker buildx create --use
docker buildx inspect --bootstrap
```

Build and push:

```
docker buildx build --platform linux/amd64,linux/arm64 -t your-registry/golang-url-shortener:latest --push .
```

## License
#### MIT