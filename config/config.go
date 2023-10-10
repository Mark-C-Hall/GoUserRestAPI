package config

import "os"

// Config represents the configuration structure used by the application.
// It contains fields for the server host, server port, and JWT secret key.
type Config struct {
	ServerHost     string
	ServerPort     string
	JWTSecret      string
	AllowedOrigins string
}

// C is the global configuration instance populated by the Load function.
var C Config

// Load initializes the global configuration (C) using environment variables or default values.
func Load() {
	C = Config{
		ServerHost:     getEnv("HOST", "localhost"),    // Default to localhost if HOST environment variable is not set
		ServerPort:     getEnv("PORT", "8080"),         // Default to port 8080 if PORT environment variable is not set
		JWTSecret:      getEnv("JWT_KEY", ""),          // No default for JWT secret; it should be set securely in the environment
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"), // Default to allow all origins
	}
}

// getEnv fetches the value of an environment variable or returns a default value.
// key is the name of the environment variable to fetch.
// defaultValue is the value to return if the environment variable is not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
