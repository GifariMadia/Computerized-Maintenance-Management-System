package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds runtime configuration for the CMMS API server.
type Config struct {
	AppEnv                string
	Port                  string
	DatabaseURL           string
	Timezone              string
	JWTSecret             string
	MaxFailedLoginAttempt int
	LockoutDuration       time.Duration
	SessionTimeout        time.Duration
	CORSAllowedOrigins    string
}

func Load() Config {
	return Config{
		AppEnv:                getEnv("APP_ENV", "development"),
		Port:                  getEnv("BACKEND_PORT", "8080"),
		DatabaseURL:           getEnv("DATABASE_URL", "postgres://cmms_app:change_me_strong_password@localhost:5432/cmms_db?sslmode=disable&TimeZone=Asia/Jakarta"),
		Timezone:              getEnv("TZ", "Asia/Jakarta"),
		JWTSecret:             getEnv("JWT_SECRET", "dev-only-secret-change-in-production"),
		MaxFailedLoginAttempt: getEnvInt("AUTH_MAX_FAILED_ATTEMPTS", 3),
		LockoutDuration:       time.Duration(getEnvInt("AUTH_LOCKOUT_MINUTES", 15)) * time.Minute,
		SessionTimeout:        time.Duration(getEnvInt("SESSION_TIMEOUT_MINUTES", 60)) * time.Minute,
		CORSAllowedOrigins:    getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			return parsed
		}
	}
	return fallback
}
