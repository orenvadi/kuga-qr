package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	sqlc "github.com/orenvadi/kuga-lms/sql/gen"
)

type Storage struct {
	Db *sqlc.Queries
}

func New(ctx context.Context, DbUrl string) *Storage {
	conn, err := pgx.Connect(ctx, DbUrl)
	if err != nil {
		panic(fmt.Errorf("could not connect to db %s", err))
	}

	db := sqlc.New(conn)

	return &Storage{Db: db}
}
