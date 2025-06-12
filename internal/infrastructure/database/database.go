package database

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	env_config "github.com/davigomesdev/reconfile/internal/infrastructure/env-config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func GetDB() *pgxpool.Pool {
	return pool
}

func ConnectDB() *pgxpool.Pool {
	env := env_config.LoadConfig()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&timezone=%s",
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseHost,
		env.DatabasePort,
		env.DatabaseName,
		env.DatabaseSSLMode,
		env.DatabaseTimezone,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse DSN: %v", err)
	}

	config.MaxConns = 100
	config.MinConns = 10
	config.MaxConnLifetime = time.Hour

	p, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	if err := p.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}

	pool = p
	log.Println("Connected to PostgreSQL database successfully with pgxpool.")

	return pool
}

func DisconnectDB() {
	if pool != nil {
		pool.Close()
		log.Println("Database connection pool closed.")
	}
}

func RunMigrations() error {
	env := env_config.LoadConfig()

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		env.DatabaseUser,
		env.DatabasePassword,
		env.DatabaseHost,
		env.DatabasePort,
		env.DatabaseName,
		env.DatabaseSSLMode,
	)

	migrationsPath := filepath.Join("internal", "infrastructure", "database", "migrations")
	sourceURL := fmt.Sprintf("file://%s", filepath.ToSlash(migrationsPath))

	m, err := migrate.New(
		sourceURL,
		databaseURL,
	)

	if err != nil {
		log.Printf("Failed to create migrate instance: %v", err)
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Printf("Migration up error: %v", err)
		return fmt.Errorf("migration up error: %w", err)
	}

	log.Println("Database migrations applied successfully.")
	return nil
}
