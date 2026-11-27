package storage_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/storage"
)

func TestPresignUsesPublicURLInsteadOfInternalDockerEndpoint(t *testing.T) {
	client, err := storage.NewMinIO(storage.Config{
		Endpoint:  "minio:9000",
		PublicURL: "http://192.168.10.50:9000",
		AccessKey: "cmms_minio",
		SecretKey: "strong_secret_key",
		Bucket:    "cmms-documents",
		UseSSL:    false,
	})
	if err != nil {
		t.Fatalf("NewMinIO failed: %v", err)
	}
	if !client.Configured() {
		t.Fatal("expected storage client to report Configured() = true")
	}

	urlStr, err := client.PresignGetURL(context.Background(), "assets/AST-001/manual.pdf", 5*time.Minute)
	if err != nil {
		t.Fatalf("PresignGetURL failed: %v", err)
	}
	if !strings.HasPrefix(urlStr, "http://192.168.10.50:9000/cmms-documents/assets/AST-001/manual.pdf?") {
		t.Fatalf("expected SigV4 presigned URL for public LAN host, got %q", urlStr)
	}
	if !strings.Contains(urlStr, "X-Amz-Signature=") {
		t.Fatalf("expected SigV4 signature in URL, got %q", urlStr)
	}
}
