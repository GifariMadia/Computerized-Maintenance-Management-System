package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/config"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/database"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/router"
	"github.com/GifariMadia/Computerized-Maintenance-Management-System/backend/internal/storage"
)

func main() {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	// Enforce Asia/Jakarta (WIB) per URS 2.1
	if loc, err := time.LoadLocation(cfg.Timezone); err == nil {
		time.Local = loc
	}

	ctx := context.Background()

	store, err := database.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database init failed: %v", err)
	}
	defer func() {
		_ = store.Close()
	}()
	log.Println("PostgreSQL connected via pgx/v5/stdlib")

	objStorage, err := storage.NewMinIO(storage.Config{
		Endpoint:  cfg.MinIOEndpoint,
		PublicURL: cfg.MinIOPublicURL,
		AccessKey: cfg.MinIOAccessKey,
		SecretKey: cfg.MinIOSecretKey,
		Bucket:    cfg.MinIOBucket,
		UseSSL:    cfg.MinIOUseSSL,
	})
	if err != nil {
		log.Fatalf("object storage init failed: %v", err)
	}
	if objStorage != nil {
		if err := objStorage.EnsureBucket(ctx); err != nil {
			log.Fatalf("object storage bucket init failed: %v", err)
		}
		log.Printf("Object storage connected (endpoint=%s, bucket=%s, public_url=%s)",
			cfg.MinIOEndpoint, cfg.MinIOBucket, cfg.MinIOPublicURL)
	}

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router.New(cfg, store, objStorage),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Printf("CMMS API listening on :%s (env=%s, tz=%s)", cfg.Port, cfg.AppEnv, cfg.Timezone)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down CMMS API gracefully...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}
}
