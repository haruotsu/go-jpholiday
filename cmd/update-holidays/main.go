package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/haruotsu/go-jpholiday/fetcher"
	"github.com/haruotsu/go-jpholiday/holiday"
)

var (
	version = "dev" // Set by build flags
)

// Config holds the configuration for the CLI tool
type Config struct {
	startYear int
	endYear   int
	cacheFile string
	dryRun    bool
	debug     bool
	showHelp  bool
	showVersion bool
}

func main() {
	config := parseFlags()

	if config.showHelp {
		printUsage()
		os.Exit(0)
	}

	if config.showVersion {
		fmt.Printf("update-holidays version %s\n", version)
		os.Exit(0)
	}

	if err := validateFlags(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		printUsage()
		os.Exit(1)
	}

	// Get API key
	apiKey := getAPIKey()
	if apiKey == "" {
		fmt.Fprintf(os.Stderr, "Error: GOOGLE_API_KEY environment variable is required\n")
		os.Exit(1)
	}

	// Create fetcher
	f := fetcher.NewFetcher(apiKey)
	if f == nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to create fetcher\n")
		os.Exit(1)
	}

	if config.debug {
		log.Printf("Fetching holidays for years %d-%d", config.startYear, config.endYear)
	}

	// Fetch holidays
	holidays, err := f.FetchHolidaysRange(config.startYear, config.endYear)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching holidays: %v\n", err)
		os.Exit(1)
	}

	if config.debug {
		log.Printf("Fetched %d holidays", len(holidays))
	}

	if config.dryRun {
		fmt.Printf("Dry run: Would update cache with %d holidays\n", len(holidays))
		for _, h := range holidays {
			fmt.Printf("  %s: %s\n", h.Date.Format("2006-01-02"), h.Name)
		}
		return
	}

	// Load existing cache or create new one
	cacheFilePath := getCacheFilePath(config)
	cache, err := holiday.LoadCache(cacheFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Create new cache if file doesn't exist
			cache = &holiday.HolidayCache{
				Holidays: make(map[string]holiday.Holiday),
			}
			if config.debug {
				log.Printf("Creating new cache file: %s", cacheFilePath)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Error loading cache: %v\n", err)
			os.Exit(1)
		}
	}

	// Update cache
	holiday.UpdateCache(cache, holidays, time.Now())

	// Save cache
	if err := holiday.SaveCache(cacheFilePath, cache); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving cache: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully updated cache with %d holidays\n", len(holidays))
	fmt.Printf("Cache saved to: %s\n", cacheFilePath)
}

// parseFlags parses command line flags
func parseFlags() *Config {
	config := &Config{}

	flag.IntVar(&config.startYear, "start-year", time.Now().Year(), "Start year for fetching holidays")
	flag.IntVar(&config.endYear, "end-year", time.Now().Year()+1, "End year for fetching holidays")
	flag.StringVar(&config.cacheFile, "cache-file", "", "Path to cache file (default: data/holidays.json)")
	flag.BoolVar(&config.dryRun, "dry-run", false, "Print what would be done without making changes")
	flag.BoolVar(&config.debug, "debug", false, "Enable debug output")
	flag.BoolVar(&config.showHelp, "help", false, "Show help message")
	flag.BoolVar(&config.showHelp, "h", false, "Show help message")
	flag.BoolVar(&config.showVersion, "version", false, "Show version information")
	flag.BoolVar(&config.showVersion, "v", false, "Show version information")

	flag.Parse()

	return config
}

// getAPIKey gets the Google Calendar API key from environment variable
func getAPIKey() string {
	return os.Getenv("GOOGLE_API_KEY")
}

// validateFlags validates the provided configuration
func validateFlags(config *Config) error {
	if config.startYear > config.endYear {
		return fmt.Errorf("start year (%d) cannot be greater than end year (%d)", config.startYear, config.endYear)
	}

	// Prevent fetching too many years at once (rate limiting)
	if config.endYear-config.startYear > 5 {
		return fmt.Errorf("year range too large (max 5 years), got %d years", config.endYear-config.startYear+1)
	}

	return nil
}

// getCacheFilePath returns the cache file path
func getCacheFilePath(config *Config) string {
	if config.cacheFile != "" {
		return config.cacheFile
	}
	return "data/holidays.json"
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