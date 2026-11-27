package storage

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Config defines internal S3/MinIO connectivity and the browser-accessible public URL.
type Config struct {
	Endpoint  string // Internal Docker endpoint, e.g. "minio:9000"
	PublicURL string // Browser/LAN URL, e.g. "http://localhost:9000" or "http://10.x.x.x:9000"
	AccessKey string
	SecretKey string
	Bucket    string
	UseSSL    bool
	Region    string
}

// Client wraps minio-go/v7 with separate internal (Docker) and public (browser SigV4) clients
// so presigned URLs remain cryptographically valid when accessed from localhost or LAN.
type Client struct {
	internal *minio.Client
	signer   *minio.Client
	bucket   string
}

func NewMinIO(cfg Config) (*Client, error) {
	if cfg.Endpoint == "" {
		return nil, nil
	}
	region := strings.TrimSpace(cfg.Region)
	if region == "" {
		region = "us-east-1"
	}

	internalClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
		Region: region,
	})
	if err != nil {
		return nil, fmt.Errorf("init internal minio client: %w", err)
	}

	signerClient := internalClient
	if strings.TrimSpace(cfg.PublicURL) != "" {
		parsed, err := url.Parse(strings.TrimSpace(cfg.PublicURL))
		if err != nil || parsed.Host == "" {
			return nil, fmt.Errorf("invalid MINIO_PUBLIC_URL %q", cfg.PublicURL)
		}
		signerClient, err = minio.New(parsed.Host, &minio.Options{
			Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
			Secure: parsed.Scheme == "https",
			Region: region,
		})
		if err != nil {
			return nil, fmt.Errorf("init public minio signer: %w", err)
		}
	}

	return &Client{
		internal: internalClient,
		signer:   signerClient,
		bucket:   cfg.Bucket,
	}, nil
}

// EnsureBucket verifies the configured bucket exists and creates it idempotently if missing.
func (c *Client) EnsureBucket(ctx context.Context) error {
	if c == nil || c.internal == nil {
		return nil
	}
	var lastErr error
	for attempt := 1; attempt <= 5; attempt++ {
		checkCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		exists, err := c.internal.BucketExists(checkCtx, c.bucket)
		if err == nil && !exists {
			err = c.internal.MakeBucket(checkCtx, c.bucket, minio.MakeBucketOptions{})
		}
		cancel()
		if err == nil {
			return nil
		}
		lastErr = err
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Duration(attempt) * time.Second):
		}
	}
	return fmt.Errorf("ensure minio bucket %q: %w", c.bucket, lastErr)
}

func (c *Client) Healthy(ctx context.Context) bool {
	if c == nil || c.internal == nil {
		return false
	}
	checkCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	exists, err := c.internal.BucketExists(checkCtx, c.bucket)
	return err == nil && exists
}

func (c *Client) Configured() bool {
	return c != nil && c.internal != nil
}

func (c *Client) Bucket() string {
	if c == nil {
		return ""
	}
	return c.bucket
}

// PresignGetURL generates a SigV4-valid download URL signed for MINIO_PUBLIC_URL (browser/LAN host).
func (c *Client) PresignGetURL(ctx context.Context, objectKey string, expiry time.Duration) (string, error) {
	if c == nil || c.signer == nil {
		return "", fmt.Errorf("object storage is not configured")
	}
	signed, err := c.signer.PresignedGetObject(ctx, c.bucket, objectKey, expiry, nil)
	if err != nil {
		return "", fmt.Errorf("presign get object %q: %w", objectKey, err)
	}
	return signed.String(), nil
}
