package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/orenvadi/kuga-lms/storage/sql/gen"
)

type Storage struct {
	Db         *sqlc.Queries
	connection *pgxpool.Pool
}

func New(ctx context.Context, DbUrl string) *Storage {
	// Create a connection pool config
	poolConfig, err := pgxpool.ParseConfig(DbUrl)
	if err != nil {
		log.Fatalf("could not parse db url: %v", err)
	}

	// Optional: Configure pool settings
	poolConfig.MaxConns = 20 // Adjust based on your needs

	// Create the pool
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("could not ping db: %v", err)
	}
	log.Println("database connection pool established successfully")

	db := sqlc.New(pool)

	return &Storage{
		Db:         db,
		connection: pool,
	}
}

func (s *Storage) Close() error {
	if s.connection != nil {
		s.connection.Close()
		log.Println("database connection pool closed")
	}
	return nil
}
