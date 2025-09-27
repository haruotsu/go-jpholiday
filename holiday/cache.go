package holiday

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/haruotsu/go-jpholiday/model"
)

// LoadCache loads holiday cache from a JSON file
func LoadCache(filePath string) (*model.HolidayCache, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cache model.HolidayCache
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cache file: %w", err)
	}

	return &cache, nil
}

// SaveCache saves holiday cache to a JSON file
func SaveCache(filePath string, cache *model.HolidayCache) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Marshal cache to JSON
	data, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache: %w", err)
	}

	// Write to file
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	return nil
}

// UpdateCache updates the holiday cache with new holidays
func UpdateCache(cache *model.HolidayCache, holidays []model.Holiday, updateTime time.Time) {
	if cache.Holidays == nil {
		cache.Holidays = make(map[string]model.Holiday)
	}

	// Add new holidays to cache
	for _, holiday := range holidays {
		key := formatDateKey(holiday.Date)
		cache.Holidays[key] = holiday
	}

	// Update the last updated time
	cache.LastUpdated = updateTime
}

// IsStale checks if the cache is older than the specified duration
func IsStale(cache *model.HolidayCache, maxAge time.Duration) bool {
	if cache == nil {
		return true
	}

	return time.Since(cache.LastUpdated) > maxAge
}
