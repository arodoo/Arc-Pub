// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

// Package config provides application configuration.
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
	// Load .env file if exists (ignore error if not found)
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
