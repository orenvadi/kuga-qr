package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/orenvadi/kuga-lms/internal/config"
	"github.com/orenvadi/kuga-lms/internal/server/handler"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "kuga-lms",
	Short: "kuga-lms is an open source learning management system",
	Long:  `A fast and simple lms`,

	Run: func(cmd *cobra.Command, args []string) {
		// ctx := context.Background()
		cfg := config.MustLoad(configFile)

		// db := postgres.New(ctx, cfg.DbUrl())
		// _ = db

		handlers := handler.New()

		srv := &http.Server{
			Addr:        cfg.Server.Port,
			IdleTimeout: cfg.Server.Timeout,
			Handler:     handlers,
		}

		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("error: %s", err)
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.yaml", "Path to YAML config file")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
