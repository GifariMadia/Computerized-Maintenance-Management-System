package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
	CORSAllowedOrigins    []string

	// MinIO / S3-compatible document storage (URS Bab 3 Interfaces)
	MinIOEndpoint  string // Internal Docker/service endpoint, e.g. minio:9000
	MinIOPublicURL string // Browser-accessible base URL, e.g. http://localhost:9000
	MinIOAccessKey string
	MinIOSecretKey string
	MinIOBucket    string
	MinIOUseSSL    bool
}

func Load() Config {
	accessKey := firstNonEmpty(os.Getenv("MINIO_ACCESS_KEY"), os.Getenv("MINIO_ROOT_USER"))
	secretKey := firstNonEmpty(os.Getenv("MINIO_SECRET_KEY"), os.Getenv("MINIO_ROOT_PASSWORD"))

	return Config{
		AppEnv:                getEnv("APP_ENV", "development"),
		Port:                  getEnv("BACKEND_PORT", "8080"),
		DatabaseURL:           strings.TrimSpace(os.Getenv("DATABASE_URL")),
		Timezone:              getEnv("TZ", "Asia/Jakarta"),
		JWTSecret:             strings.TrimSpace(os.Getenv("JWT_SECRET")),
		MaxFailedLoginAttempt: getEnvInt("AUTH_MAX_FAILED_ATTEMPTS", 3),
		LockoutDuration:       time.Duration(getEnvInt("AUTH_LOCKOUT_MINUTES", 15)) * time.Minute,
		SessionTimeout:        time.Duration(getEnvInt("SESSION_TIMEOUT_MINUTES", 60)) * time.Minute,
		CORSAllowedOrigins:    parseCSV(getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000")),

		MinIOEndpoint:  strings.TrimSpace(os.Getenv("MINIO_ENDPOINT")),
		MinIOPublicURL: strings.TrimRight(strings.TrimSpace(os.Getenv("MINIO_PUBLIC_URL")), "/"),
		MinIOAccessKey: strings.TrimSpace(accessKey),
		MinIOSecretKey: strings.TrimSpace(secretKey),
		MinIOBucket:    getEnv("MINIO_BUCKET", "cmms-documents"),
		MinIOUseSSL:    getEnvBool("MINIO_USE_SSL", false),
	}
}

// Validate enforces fail-fast checks so weak or missing secrets are caught at startup.
func (c Config) Validate() error {
	var missing []string
	if c.DatabaseURL == "" {
		missing = append(missing, "DATABASE_URL")
	}
	if c.JWTSecret == "" {
		missing = append(missing, "JWT_SECRET")
	} else if len(c.JWTSecret) < 32 {
		return fmt.Errorf("JWT_SECRET must be at least 32 characters long")
	}
	if c.MinIOEndpoint != "" {
		if c.MinIOAccessKey == "" {
			missing = append(missing, "MINIO_ACCESS_KEY (or MINIO_ROOT_USER)")
		}
		if c.MinIOSecretKey == "" {
			missing = append(missing, "MINIO_SECRET_KEY (or MINIO_ROOT_PASSWORD)")
		}
		if c.MinIOBucket == "" {
			missing = append(missing, "MINIO_BUCKET")
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}
	return nil
}

func getEnv(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func firstNonEmpty(values ...string) string {
	for _, v := range values {
		if strings.TrimSpace(v) != "" {
			return strings.TrimSpace(v)
		}
	}
	return ""
}

func getEnvInt(key string, fallback int) int {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			return parsed
		}
	}
	return fallback
}

func getEnvBool(key string, fallback bool) bool {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		if parsed, err := strconv.ParseBool(v); err == nil {
			return parsed
		}
	}
	return fallback
}

func parseCSV(raw string) []string {
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if trimmed := strings.TrimRight(strings.TrimSpace(p), "/"); trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}
