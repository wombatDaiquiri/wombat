package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Seed(ctx context.Context, db *pgxpool.Pool) error {
	// TODO: insert blogposts
	return nil
}
