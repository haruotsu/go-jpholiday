package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/haruotsu/go-jpholiday/model"
)

// parseFlags parses command line flags
func parseFlags() *model.Config {
	config := &model.Config{}

	flag.IntVar(&config.StartYear, "start-year", time.Now().Year(), "Start year for fetching holidays")
	flag.IntVar(&config.EndYear, "end-year", time.Now().Year()+1, "End year for fetching holidays")
	flag.StringVar(&config.CacheFile, "cache-file", "", "Path to cache file (default: data/holidays.json)")
	flag.BoolVar(&config.DryRun, "dry-run", false, "Print what would be done without making changes")
	flag.BoolVar(&config.Debug, "debug", false, "Enable debug output")
	flag.BoolVar(&config.ShowHelp, "help", false, "Show help message")
	flag.BoolVar(&config.ShowHelp, "h", false, "Show help message")
	flag.BoolVar(&config.ShowVersion, "version", false, "Show version information")
	flag.BoolVar(&config.ShowVersion, "v", false, "Show version information")

	flag.Parse()

	return config
}

// validateFlags validates the provided configuration
func validateFlags(config *model.Config) error {
	if config.StartYear > config.EndYear {
		return fmt.Errorf("start year (%d) cannot be greater than end year (%d)", config.StartYear, config.EndYear)
	}

	// Prevent fetching too many years at once (rate limiting)
	if config.EndYear-config.StartYear > 10 {
		return fmt.Errorf("year range too large (max 10 years), got %d years", config.EndYear-config.StartYear+1)
	}

	return nil
}

// printUsage prints usage information
func printUsage() {
	fmt.Printf("update-holidays - Update Japanese holidays cache\n\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("  update-holidays [options]\n\n")
	fmt.Printf("Options:\n")
	flag.PrintDefaults()
	fmt.Printf("\nEnvironment Variables:\n")
	fmt.Printf("  GOOGLE_API_KEY    Google Calendar API key (required)\n\n")
	fmt.Printf("Examples:\n")
	fmt.Printf("  # Update cache for current and next year\n")
	fmt.Printf("  GOOGLE_API_KEY=xxx update-holidays\n\n")
	fmt.Printf("  # Update cache for specific year range\n")
	fmt.Printf("  GOOGLE_API_KEY=xxx update-holidays -start-year 2024 -end-year 2025\n\n")
	fmt.Printf("  # Dry run to see what would be updated\n")
	fmt.Printf("  GOOGLE_API_KEY=xxx update-holidays -dry-run\n\n")
}