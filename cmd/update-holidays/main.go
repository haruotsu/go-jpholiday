package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/haruotsu/go-jpholiday/fetcher"
	"github.com/haruotsu/go-jpholiday/holiday"
	"github.com/haruotsu/go-jpholiday/model"
)

var (
	version = "dev" // Set by build flags
)

func main() {
	// for CLI
	config := parseFlags()

	if config.ShowHelp {
		printUsage()
		os.Exit(0)
	}

	if config.ShowVersion {
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

	if config.Debug {
		log.Printf("Fetching holidays for years %d-%d", config.StartYear, config.EndYear)
	}

	// Fetch holidays
	holidays, err := f.FetchHolidaysRange(config.StartYear, config.EndYear)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching holidays: %v\n", err)
		os.Exit(1)
	}

	if config.Debug {
		log.Printf("Fetched %d holidays", len(holidays))
	}

	if config.DryRun {
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
			cache = &model.HolidayCache{
				Holidays: make(map[string]model.Holiday),
			}
			if config.Debug {
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
