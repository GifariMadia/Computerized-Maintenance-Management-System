package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Store wraps the single operational PostgreSQL connection pool (URS 2.2).
// Track 1 (Gifari): import _ "github.com/jackc/pgx/v5/stdlib" and pass driverName "pgx".
type Store struct {
	DB *sql.DB
}

func New(ctx context.Context, driverName, dsn string) (*Store, error) {
	if driverName == "" || dsn == "" {
		return &Store{DB: nil}, nil
	}
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}
	return &Store{DB: db}, nil
}

func (s *Store) Healthy(ctx context.Context) bool {
	if s == nil || s.DB == nil {
		return false
	}
	pingCtx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return s.DB.PingContext(pingCtx) == nil
}

func (s *Store) Close() error {
	if s == nil || s.DB == nil {
		return nil
	}
	return s.DB.Close()
}
