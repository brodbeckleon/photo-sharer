package config

import "os"

type AppConfig struct {
	SecretCode string
	JWTSecret  []byte
}

func Load() *AppConfig {
	// In a real app, handle missing variables gracefully
	return &AppConfig{
		SecretCode: os.Getenv("SECRET_CODE"),
		JWTSecret:  []byte(os.Getenv("JWT_SECRET")),
	}
}
