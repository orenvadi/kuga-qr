package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	sqlc "github.com/orenvadi/kuga-lms/sql/gen"
)

type Storage struct {
	context    context.Context
	Db         *sqlc.Queries
	connection *pgx.Conn
}

func New(ctx context.Context, DbUrl string) *Storage {
	conn, err := pgx.Connect(ctx, DbUrl)
	if err != nil {
		panic(fmt.Errorf("could not connect to db %s", err))
	}

	db := sqlc.New(conn)

	return &Storage{context: ctx, Db: db, connection: conn}
}

func (s *Storage) Close() error {
	return s.connection.Close(s.context)
}
