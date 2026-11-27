package config_test

import (
	"testing"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/config"
)

func TestValidateRejectsMissingOrWeakSecrets(t *testing.T) {
	c := config.Config{}
	if err := c.Validate(); err == nil {
		t.Fatal("expected Validate() to fail when DATABASE_URL and JWT_SECRET are empty")
	}

	c.DatabaseURL = "postgres://u:p@localhost:5432/cmms_db"
	c.JWTSecret = "short-secret"
	if err := c.Validate(); err == nil {
		t.Fatal("expected Validate() to fail when JWT_SECRET is shorter than 32 chars")
	}

	c.JWTSecret = "0123456789abcdef0123456789abcdef"
	c.MinIOEndpoint = "minio:9000"
	if err := c.Validate(); err == nil {
		t.Fatal("expected Validate() to fail when MINIO_ENDPOINT is set without credentials")
	}

	c.MinIOAccessKey = "cmms_minio"
	c.MinIOSecretKey = "strong_minio_pass"
	c.MinIOBucket = "cmms-documents"
	if err := c.Validate(); err != nil {
		t.Fatalf("expected valid config to pass Validate(), got: %v", err)
	}
}
