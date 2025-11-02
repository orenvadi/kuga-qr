package api

import (
	"time"

	"github.com/orenvadi/kuga-lms/internal/storage/postgres"
)

type Server struct {
	db        *postgres.Storage
	jwtSecret string
	tokenTTL  time.Duration
}

func NewServer(db *postgres.Storage, jwtSecret string, tokenTTL time.Duration) Server {
	return Server{db: db, jwtSecret: jwtSecret, tokenTTL: tokenTTL}
}
