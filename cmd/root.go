package main

import (
	"ecostars-fake-api/internal/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ecostars-fake-api",
	Short: "Ecostars Fake API Server",
	Long:  `A fake API server for Ecostars to simulate hotel and measure data.`,
}

var serve = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.LoadConfig()
		BootstrapHTTPServer(config)
	},
}

var populateDB = &cobra.Command{
	Use:   "fake",
	Short: "Populate the database with sample data",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.LoadConfig()
		err := PopulateDB(config)
		if err != nil {
			fmt.Println("Error populating database:", err)
			os.Exit(1)
		}
		fmt.Println("Database populated successfully.")
	},
}

var metrics = &cobra.Command{
	Use:   "metrics",
	Short: "Start the Metrics collector",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.LoadConfig()
		BootstrapMetricsCollector(config)
	},
}

func init() {
	rootCmd.AddCommand(serve)
	rootCmd.AddCommand(populateDB)
	rootCmd.AddCommand(metrics)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
