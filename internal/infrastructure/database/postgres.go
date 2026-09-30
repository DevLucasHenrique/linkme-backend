package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresConnection(dns string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dns)
	if err!=nil {
		return nil, fmt.Errorf("DSN Connection string error, %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err!=nil {
		return nil, fmt.Errorf("Can't connect with Supabase, error: %v", err)
	}

	return pool, nil
}