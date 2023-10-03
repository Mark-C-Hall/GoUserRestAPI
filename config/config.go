package config

import "os"

type Config struct {
	ServerHost string
	ServerPort string
	JWTSecret  string
}

var C Config

func Load() {
	C = Config{
		ServerHost: getEnv("HOST", "localhost"),
		ServerPort: getEnv("PORT", "8080"),
		JWTSecret:  getEnv("JWT_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
