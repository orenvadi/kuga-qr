package main

import (
	"fmt"
	"os"

	"github.com/orenvadi/kuga-lms/internal/config"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "kuga-lms",
	Short: "kuga-lms is an open source learning management system",
	Long:  `A fast and simple lms`,

	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.MustLoad(configFile)
		fmt.Println(cfg)
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
