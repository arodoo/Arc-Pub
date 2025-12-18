// File: config.go
// Purpose: Provides centralized application configuration management. Loads
// environment variables from .env file using godotenv, then parses them into
// a strongly-typed Config struct. Supports required fields (DATABASE_URL,
// JWT_SECRET) and optional fields with defaults (PORT). Single source of
// truth for all application settings following twelve-factor app principles.
// Path: server/internal/config/config.go
// All Rights Reserved. Arc-Pub.

package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// Config holds all application settings.
type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
	JWTSecret   string `env:"JWT_SECRET,required"`
	Port        string `env:"PORT" envDefault:"8080"`
}

// Load reads config from .env file and environment variables.
func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
