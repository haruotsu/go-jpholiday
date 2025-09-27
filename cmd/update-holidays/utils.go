package main

import (
	"os"

	"github.com/haruotsu/go-jpholiday/holiday"
	"github.com/haruotsu/go-jpholiday/model"
)

// getAPIKey gets the Google Calendar API key from environment variable
func getAPIKey() string {
	return os.Getenv("GOOGLE_API_KEY")
}

// getCacheFilePath returns the cache file path
func getCacheFilePath(config *model.Config) string {
	if config.CacheFile != "" {
		return config.CacheFile
	}
	return holiday.DefaultCacheFile
}
