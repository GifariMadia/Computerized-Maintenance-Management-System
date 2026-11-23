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
)

func main() {
	cfg := config.Load()

	// Enforce Asia/Jakarta (WIB) per URS 2.1
	if loc, err := time.LoadLocation(cfg.Timezone); err == nil {
		time.Local = loc
	}

	ctx := context.Background()
	// Note: Track 1 (Gifari) will register pgx driver (`_ "github.com/jackc/pgx/v5/stdlib"`)
	// and pass "pgx" here when wiring live PostgreSQL repositories.
	store, err := database.New(ctx, "", "")
	if err != nil {
		log.Fatalf("database init failed: %v", err)
	}
	defer func() {
		_ = store.Close()
	}()

	srv := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router.New(cfg, store),
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
