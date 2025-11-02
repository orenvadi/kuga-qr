package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/orenvadi/kuga-lms/internal/app"
	"github.com/orenvadi/kuga-lms/internal/config"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "kuga-lms",
	Short: "kuga-lms is an open source learning management system",
	Long:  `A fast and simple lms`,

	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cfg := config.MustLoad(configFile)
		fmt.Printf("\nServer config: \n%+v\n\n", cfg)

		log.Println("...STARTING APPLICATION...")

		application := app.New(ctx, cfg)

		application.Run()

		log.Println("...APPLICATION RUNNING...")

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

		// reading from chan is blocking operaiton
		sgnl := <-stop

		log.Printf("stopping application, signal: %v \n", sgnl.String())

		// gracefuly stop application
		application.Stop()

		log.Println("application stopped")
	},
}

func main() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", "config.yaml", "Path to YAML config file")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
