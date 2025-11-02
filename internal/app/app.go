package app

import (
	"context"
	"log"
	"net/http"

	"github.com/orenvadi/kuga-lms/internal/config"
	"github.com/orenvadi/kuga-lms/internal/storage/postgres"
	"github.com/orenvadi/kuga-lms/openapi/pkg/api"
)

type App struct {
	context context.Context
	Server  *http.Server
	db      *postgres.Storage
}

func New(ctx context.Context, cfg config.Config) *App {
	db := postgres.New(ctx, cfg.DbUrl())

	// handlers := handlers.New(db, cfg.Jwt.Secret)

	server := api.NewServer(db, cfg.Jwt.Secret, cfg.Jwt.TokenTTL)

	strictHandler := api.NewStrictHandler(server, []api.StrictMiddlewareFunc{api.StrictJWTMiddlewareWithSecretKey(cfg.Jwt.Secret)})

	mux := http.NewServeMux()

	handlers := api.HandlerFromMux(strictHandler, mux)

	return &App{
		context: ctx,
		Server: &http.Server{
			Addr:        cfg.Server.Port,
			IdleTimeout: cfg.Server.Timeout,
			Handler:     handlers,
		},
		db: db,
	}
}

func (a *App) Run() {
	go func() {
		a.Server.ListenAndServe()
	}()
}

func (a *App) Stop() {
	if err := a.Server.Shutdown(a.context); err != nil {
		log.Printf("could not stop application server: %v\n", err)
	}
	if err := a.db.Close(); err != nil {
		log.Printf("could not stop application database: %v\n", err)
	}
}
