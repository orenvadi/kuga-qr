package api

import (
	"context"
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

func (s Server) PostStudentScan(ctx context.Context, request PostStudentScanRequestObject) (PostStudentScanResponseObject, error) {
	return nil, nil
}

func (s Server) PostTeacherQrStream(ctx context.Context, request PostTeacherQrStreamRequestObject) (PostTeacherQrStreamResponseObject, error) {
	return nil, nil
}
